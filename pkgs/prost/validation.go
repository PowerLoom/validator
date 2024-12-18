package prost

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"time"
	"validator/pkgs"
	"validator/pkgs/clients"
	"validator/pkgs/ipfs"
	"validator/pkgs/merkle"
	"validator/pkgs/redis"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sergerad/incremental-merkle-tree/imt"
	log "github.com/sirupsen/logrus"
)

type BatchDetails struct {
	DataMarketAddress string
	BatchCID          string
	EpochID           *big.Int
}

func StartBatchAttestation() {
	log.Info("🚀 Batch Attestation has started...")

	for {
		// Retrieve batch details from the Redis attestation queue
		result, err := redis.RedisClient.BRPop(context.Background(), 0, "attestorQueue").Result()
		if err != nil {
			log.Errorf("Error fetching data from Redis queue: %v", err)
			continue
		}

		// Ensure valid data was retrieved from the queue
		if len(result) < 2 {
			log.Println("Invalid data received from Redis queue, skipping this entry")
			continue
		}

		// Parse the queue data into the BatchDetails structure
		var batchDetails BatchDetails
		err = json.Unmarshal([]byte(result[1]), &batchDetails)
		if err != nil {
			log.Errorf("Failed to parse batch details from Redis queue: %v", err)
			continue
		}

		// Log the details of the batch being processed
		log.Infof("🔄 Starting attestation for batchCID %s in epoch %s within data market %s", batchDetails.BatchCID, batchDetails.EpochID.String(), batchDetails.DataMarketAddress)

		// Perform batch attestation
		if err := batchDetails.Attest(); err != nil {
			log.Errorf("Failed to attest batchCID %s in epoch %s within data market %s: %v", batchDetails.BatchCID, batchDetails.EpochID.String(), batchDetails.DataMarketAddress, err)
			continue
		}

		// Log the successful attestation of the batch
		log.Infof("✅ Successfully attested batchCID %s for epoch %s in data market %s", batchDetails.BatchCID, batchDetails.EpochID.String(), batchDetails.DataMarketAddress)
	}
}

func (batchDetails *BatchDetails) Attest() error {
	// Extract the epoch ID and the data market address from batch details
	epochID := batchDetails.EpochID.String()
	dataMarketAddress := batchDetails.DataMarketAddress

	// Fetch the batch from IPFS
	batch, err := ipfs.FetchSubmission(batchDetails.BatchCID)
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to fetch batch with CID %s, epoch %s, data market %s from IPFS: %s", batchDetails.BatchCID, epochID, dataMarketAddress, err.Error())
		clients.SendFailureNotification(pkgs.FetchBatchSubmission, errorMsg, time.Now().String(), "High")
		log.Error(errorMsg)
		return err
	}

	log.Infof("📦 Fetched batch for CID %s, epoch %s, data market %s from IPFS", batchDetails.BatchCID, epochID, dataMarketAddress)

	// Marshal the retreived batch to JSON
	batchJSON, err := json.Marshal(batch)
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to marshal batch with CID %s in epoch %s, data market %s: %s", batchDetails.BatchCID, epochID, dataMarketAddress, err.Error())
		clients.SendFailureNotification(pkgs.StoreBatchSubmission, errorMsg, time.Now().String(), "High")
		log.Error(errorMsg)
		return err
	}

	// Store the details associated with the batch in Redis
	if err := redis.StoreValidatorDetails(context.Background(), dataMarketAddress, epochID, batch.RootHash, string(batchJSON)); err != nil {
		errorMsg := fmt.Sprintf("Failed to store batch details for epoch %s, data market %s in Redis: %s", epochID, dataMarketAddress, err.Error())
		clients.SendFailureNotification(pkgs.StoreBatchSubmission, errorMsg, time.Now().String(), "High")
		log.Error(errorMsg)
		return err
	}

	log.Infof("✅ Successfully stored batch details for CID %s, epoch %s, data market %s in Redis", batchDetails.BatchCID, epochID, dataMarketAddress)

	// Create a new Merkle tree
	merkleTree, err := imt.New()
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to initialize Merkle tree for epoch %s, data market %s: %s", epochID, dataMarketAddress, err.Error())
		clients.SendFailureNotification(pkgs.BuildMerkleTree, errorMsg, time.Now().String(), "High")
		log.Error(errorMsg)
		return err
	}

	log.Infof("🌲 Successfully updated Merkle tree for epoch %s, data market %s", epochID, dataMarketAddress)

	// Update the Merkle tree with finalized CIDs from the batch
	if _, err = merkle.UpdateMerkleTree(batch.CIDs, merkleTree); err != nil {
		errorMsg := fmt.Sprintf("Error updating Merkle tree with finalized CIDs for epoch %s, data market %s: %v", epochID, dataMarketAddress, err.Error())
		clients.SendFailureNotification(pkgs.BuildMerkleTree, errorMsg, time.Now().String(), "High")
		log.Error(errorMsg)
		return err
	}

	// Calculate the root hash of the Merkle tree
	finalizedCIDsRootHash := GetRootHash(merkleTree)

	// Submit batch attestation to the external tx Relayer service with retry mechanism
	if err := SubmitBatchAttestationToRelayer(dataMarketAddress, batchDetails.BatchCID, hex.EncodeToString(finalizedCIDsRootHash[:]), batchDetails.EpochID); err != nil {
		errorMsg := fmt.Sprintf("🚨 Relayer submission failed: CID %s, epoch %s, data market %s: %v", batchDetails.BatchCID, batchDetails.EpochID, dataMarketAddress, err)
		clients.SendFailureNotification(pkgs.SendBatchAttestationToRelayer, errorMsg, time.Now().String(), "High")
		log.Error(errorMsg)
		return err
	}

	log.Infof("📤 Successfully submitted batch attestation to relayer with CID %s, epoch %s, data market %s", batchDetails.BatchCID, epochID, dataMarketAddress)

	return nil
}

func SubmitBatchAttestationToRelayer(dataMarketAddress, batchCID, rootHash string, epochID *big.Int) error {
	// Define the operation that will be retried
	operation := func() error {
		// Attempt batch attestation submission
		err := clients.SubmitBatchAttestationRequest(dataMarketAddress, batchCID, rootHash, epochID)
		if err != nil {
			log.Errorf("Failed to send batch attestation for epoch %s, data market %s: %v", epochID, dataMarketAddress, err)
			return err // Return error to trigger retry
		}

		log.Infof("Successfully submitted batch attestation for epoch %s, data market %s", epochID, dataMarketAddress)
		return nil // Successful submission, no need for further retries
	}

	// Customize the backoff configuration
	backoffConfig := backoff.NewExponentialBackOff()
	backoffConfig.InitialInterval = 1 * time.Second // Start with a 1-second delay
	backoffConfig.Multiplier = 1.5                  // Increase interval by 1.5x after each retry
	backoffConfig.MaxInterval = 4 * time.Second     // Set max interval between retries
	backoffConfig.MaxElapsedTime = 10 * time.Second // Retry for a maximum of 10 seconds

	// Limit retries to 3 times within 10 seconds
	if err := backoff.Retry(operation, backoff.WithMaxRetries(backoffConfig, 3)); err != nil {
		log.Errorf("Failed to submit batch attestation for epoch %s, data market %s after multiple retries: %v", epochID, dataMarketAddress, err)
		return err
	}

	return nil
}

// GetRootHash returns hash representation of the root digest of the Merkle tree
func GetRootHash(tree *imt.IncrementalMerkleTree) [32]byte {
	return crypto.Keccak256Hash(tree.RootDigest())
}
