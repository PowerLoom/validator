package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sergerad/incremental-merkle-tree/imt"
	log "github.com/sirupsen/logrus"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"
	"validator/config"
	"validator/pkgs/contract/contract"
)

var (
	Client             *ethclient.Client
	CurrentBlockNumber = new(big.Int)
	CurrentEpochID     = new(big.Int)
)

func ConfigureClient() {
	var err error
	Client, err = ethclient.Dial(config.SettingsObj.ClientUrl)
	if err != nil {
		log.Fatal(err)
	}
}

func SetupAuth() {
	nonce, err := Client.PendingNonceAt(context.Background(), config.SettingsObj.SignerAccountAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	Auth, err = bind.NewKeyedTransactorWithChainID(config.SettingsObj.PrivateKey, big.NewInt(int64(config.SettingsObj.ChainID)))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	Auth.Nonce = big.NewInt(int64(nonce))
	Auth.Value = big.NewInt(0)      // in wei
	Auth.GasLimit = uint64(3000000) // in units
	Auth.From = config.SettingsObj.SignerAccountAddress
}

func UpdateGasPrice(multiplier int) {
	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("Failed to get gas price: %v", err)
	}
	Auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(int64(multiplier)))
}

func UpdateAuth(num int64) {
	Auth.Nonce = new(big.Int).Add(Auth.Nonce, big.NewInt(num))
	UpdateGasPrice(1)
}

func StartFetchingBlocks() {
	contractABI, err := abi.JSON(strings.NewReader(contract.ContractMetaData.ABI)) // Replace with your contract ABI

	if err != nil {
		log.Fatal(err)
	}

	for {
		var block *types.Block
		block, err = Client.BlockByNumber(context.Background(), nil)
		if err != nil || block == nil {
			log.Errorf("Failed to fetch latest block: %s", err.Error())
			continue
		}

		if CurrentBlockNumber.Cmp(block.Header().Number) < 0 {
			CurrentBlockNumber.Set(block.Header().Number)
			log.Debugln("Current block: ", CurrentBlockNumber.String())

			// iterate all transactions in parallel and search for events
			go func() {
				for _, tx := range block.Transactions() {
					receipt, err := Client.TransactionReceipt(context.Background(), tx.Hash())
					if err != nil {
						log.Errorln(err.Error())
						continue
					}
					for _, vLog := range receipt.Logs {
						if vLog.Address.Hex() != config.SettingsObj.ContractAddress {
							continue
						}
						switch vLog.Topics[0].Hex() {
						case contractABI.Events["SnapshotBatchSubmitted"].ID.Hex():
							event, err := Instance.ParseSnapshotBatchSubmitted(*vLog)
							if err != nil {
								log.Debugln("Error unpacking SnapshotBatchSubmitted event:", err)
								continue
							}
							// begin building merkle tree
							go storeBatchSubmission(event)
						case contractABI.Events["EpochReleased"].ID.Hex():
							event, err := Instance.ParseEpochReleased(*vLog)
							if err != nil {
								log.Debugln("Error unpacking epochReleased event:", err)
								continue
							}
							event.EpochId = new(big.Int).SetBytes(vLog.Topics[1][:])
							if CurrentEpochID.Cmp(event.EpochId) < 0 {
								CurrentEpochID.Set(event.EpochId)
								go triggerValidationFlow(new(big.Int).Set(CurrentEpochID))
								log.Debugln("Epoch Released: ", event.EpochId.String())
							}
						}
					}
				}
			}()
			time.Sleep(time.Duration(config.SettingsObj.BlockTime*500) * time.Millisecond)
		} else {
			time.Sleep(100 * time.Millisecond)
			continue
		}
	}
}

func PopulateStateVars() {
	for {
		if num, err := Client.BlockNumber(context.Background()); err == nil {
			CurrentBlockNumber.SetUint64(num)
			break
		} else {
			log.Debugln("Encountered error while fetching current block: ", err.Error())
		}
	}
	CurrentEpochID.Set(big.NewInt(0))
}

func storeBatchSubmission(event *contract.ContractSnapshotBatchSubmitted) {
	batch := FetchSubmission(IPFSCon, event.BatchCid)
	log.Debugf("Fetched batch %s for epoch %s with roothash %s from IPFS: ", batch.ID.String(), event.EpochId, batch.RootHash)
	//submissions, err := json.Marshal(batch.Submissions)
	//if err != nil {
	//	log.Errorf("Unable to unmarshal submissions for batch %d epochId %s: %s\n", batch.ID, event.EpochId.String(), err.Error())
	//}
	submissionIds, err := json.Marshal(batch.SubmissionIds)
	if err != nil {
		log.Errorf("Unable to unmarshal submissionIds for batch %d epochId %s: %s\n", batch.ID, event.EpochId.String(), err.Error())
	}
	err = Set(context.Background(), RedisClient, fmt.Sprintf("%s.%s.%s", ValidatorKey, event.EpochId.String(), batch.ID.String()), string(submissionIds), time.Hour)
	if err != nil {
		log.Errorf("Unable to store submissions for batch %d epochId %s: %s\n", batch.ID, event.EpochId.String(), err.Error())
	}
}

func triggerValidationFlow(epochId *big.Int) {
	batchSubmissionLimit := new(big.Int).Add(CurrentBlockNumber, big.NewInt(int64(config.SettingsObj.BatchSubmissionLimit)))

	for CurrentBlockNumber.Cmp(batchSubmissionLimit) < 0 {
		time.Sleep(time.Duration(config.SettingsObj.BlockTime) * time.Second)
	}

	pattern := fmt.Sprintf("%s.%s.*", ValidatorKey, epochId)
	keys, err := FetchKeysForPattern(context.Background(), RedisClient, pattern)

	if err != nil {
		log.Errorf("Unable to fetch keys for pattern %s: %s\n", pattern, err.Error())
	}

	sort.Slice(keys, func(i, j int) bool {
		numI, _ := strconv.Atoi(strings.Split(keys[i], ".")[2])
		numJ, _ := strconv.Atoi(strings.Split(keys[j], ".")[2])
		return numI < numJ
	})

	tree, _ := imt.New()

	SetupAuth()
	for _, key := range keys {
		value, err := Get(context.Background(), RedisClient, key)
		if err != nil {
			log.Errorln("Error fetching data from redis: ", err.Error())
		}
		log.Debugf("Fetched submissions for key %s\n", key)
		var batchSubmissionIds []string
		err = json.Unmarshal([]byte(value), &batchSubmissionIds)
		if err != nil {
			log.Errorf("Unable to unmarshal batch submissionIds for key: %s\n", key)
		}
		_, err = UpdateMerkleTree(batchSubmissionIds, tree)
		if err != nil {
			log.Errorf("Unable to build Merkel tree: %s\n", err.Error())
		}
		SubmitAttestation(key, tree.RootDigest())
	}

	ResetValidatorDBSubmissions(context.Background(), RedisClient, epochId)

	time.Sleep(time.Second * time.Duration(config.SettingsObj.BlockTime))

	EnsureTxSuccess(epochId)
}

func EnsureTxSuccess(epochID *big.Int) {
	for {
		keys, err := FetchKeysForPattern(context.Background(), RedisClient, fmt.Sprintf("%s.%s.*", TxsKey, epochID.String()))
		SetupAuth()
		if err != nil {
			log.Debugf("Could not fetch submitted transactions: %s\n", err.Error())
			return
		} else {
			if keys == nil {
				log.Debugln("No unsuccessful transactions remaining for epochId: ", epochID.String())
				return
			}
			log.Debugf("Fetched %d transactions for epoch %d", len(keys), epochID)
			for _, key := range keys {
				if value, err := Get(context.Background(), RedisClient, key); err != nil {
					log.Errorf("Unable to fetch value for key: %s\n", key)
				} else {
					log.Debugf("Fetched value %s for key %s\n", value, key)
					vals := strings.Split(value, ".")

					tx := vals[0]
					cid := vals[3]
					batchID := new(big.Int)
					_, ok := batchID.SetString(vals[2], 10)
					if !ok {
						log.Errorf("Unable to convert bigInt string to bigInt: %s\n", vals[2])
					}

					nonce := strings.Split(key, ".")[3]
					multiplier := 1
					if _, err := Client.TransactionReceipt(context.Background(), common.HexToHash(tx)); err != nil {
						log.Errorf("Found unsuccessful transaction: %s, batchID: %d, nonce: %s", tx, batchID, nonce)
						updatedNonce := Auth.Nonce.String()
						UpdateGasPrice(1)
						var reTx *types.Transaction
						for reTx, err = Instance.SubmitBatchAttestation(Auth, batchID, epochID, [32]byte(common.Hex2Bytes(cid))); err != nil; {
							updatedNonce = Auth.Nonce.String()
							multiplier = HandleAttestationSubmissionError(err, multiplier, batchID.String())
						}
						UpdateAuth(1)
						RedisClient.Set(context.Background(), fmt.Sprintf("%s.%s.%d.%s", TxsKey, epochID.String(), batchID, updatedNonce), fmt.Sprintf("%s.%s.%s.%s", reTx.Hash().Hex(), epochID.String(), batchID.String(), cid), time.Hour)
					}
					if _, err := RedisClient.Del(context.Background(), fmt.Sprintf("%s.%s.%s.%s", TxsKey, epochID.String(), batchID.String(), nonce)).Result(); err != nil {
						log.Errorf("Unable to delete transaction from redis: %s\n", err.Error())
					}
				}
			}
		}
		time.Sleep(time.Second * time.Duration(config.SettingsObj.BlockTime))
	}
}
