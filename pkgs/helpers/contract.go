package helpers

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strings"
	"time"
	"validator/config"
	"validator/pkgs/contract/contract"
)

var (
	Auth     *bind.TransactOpts
	Instance *contract.Contract
)

func ConfigureContractInstance() {
	Instance, _ = contract.NewContract(common.HexToAddress(config.SettingsObj.ContractAddress), Client)
}

func SubmitAttestation(key string, cid []byte) {
	vals := strings.Split(key, ".")
	epochId, _ := new(big.Int).SetString(vals[1], 10)
	batchId, _ := new(big.Int).SetString(vals[2], 10)

	multiplier := 1
	UpdateGasPrice(multiplier)
	var tx *types.Transaction
	var err error
	nonce := Auth.Nonce.String()
	for tx, err = Instance.SubmitBatchAttestation(Auth, batchId, epochId, [32]byte(cid)); err != nil; {
		time.Sleep(time.Duration(config.SettingsObj.BlockTime) * time.Second)
		nonce = Auth.Nonce.String()
		multiplier = HandleAttestationSubmissionError(err, multiplier, batchId.String())
	}
	RedisClient.Set(context.Background(), fmt.Sprintf("%s.%s.%s.%s", TxsKey, epochId.String(), batchId.String(), nonce), fmt.Sprintf("%s.%s.%s.%s", tx.Hash().Hex(), epochId.String(), batchId.String(), common.Bytes2Hex(cid)), time.Hour)
	log.Debugf("Successfully submitted attestation for batch %s with roothash %s and nonce %s\n ", batchId.String(), common.Bytes2Hex(cid), nonce)
	UpdateAuth(1)
}
