package clients

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"
	"validator/config"
)

type TxRelayerClient struct {
	url    string
	client *http.Client
}

type BatchAttestationRequest struct {
	DataMarketAddress string   `json:"dataMarketAddress"`
	BatchCID          string   `json:"batchCID"`
	EpochID           *big.Int `json:"epochID"`
	RootHash          string   `json:"rootHash"`
	AuthToken         string   `json:"authToken"`
}

var txRelayerClient *TxRelayerClient

// InitializeTxClient initializes the TxRelayerClient with the provided URL and timeout
func InitializeTxClient(url string, timeout time.Duration) {
	txRelayerClient = &TxRelayerClient{
		url: url,
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

// SubmitBatchAttestationRequest submits details for batch attestattion to the transaction relayer service
func SubmitBatchAttestationRequest(dataMarketAddress, batchCID, rootHash string, epochID *big.Int) error {
	request := BatchAttestationRequest{
		DataMarketAddress: dataMarketAddress,
		BatchCID:          batchCID,
		EpochID:           epochID,
		RootHash:          rootHash,
		AuthToken:         config.SettingsObj.TxRelayerAuthWriteToken,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("unable to marshal batch attestation request: %w", err)
	}

	url := fmt.Sprintf("%s/submitBatchAttestation", txRelayerClient.url)

	resp, err := txRelayerClient.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("unable to send batch attestation request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send batch attestation request, status code: %d", resp.StatusCode)
	}

	return nil
}
