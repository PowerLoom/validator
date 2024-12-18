package ipfs

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"time"
	"validator/config"

	shell "github.com/ipfs/go-ipfs-api"
	log "github.com/sirupsen/logrus"
)

var IPFSClient *shell.Shell

// Batch represents your data structure
type Batch struct {
	SubmissionIDs []string `json:"submissionIDs"`
	Submissions   []string `json:"submissions"`
	RootHash      string   `json:"roothash"`
	PIDs          []string `json:"pids"`
	CIDs          []string `json:"cids"`
}

// ConnectIPFSNode connects to the IPFS node using the provided configuration
func ConnectIPFSNode() {
	log.Debugf("Connecting to IPFS node at: %s", config.SettingsObj.IPFSUrl)

	IPFSClient = shell.NewShellWithClient(
		config.SettingsObj.IPFSUrl,
		&http.Client{
			Timeout: time.Duration(config.SettingsObj.HttpTimeout) * time.Second,
			Transport: &http.Transport{
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
				MaxIdleConns:      10,
				IdleConnTimeout:   5 * time.Second,
				DisableKeepAlives: true,
			},
		},
	)
}

func FetchSubmission(batchCID string) (*Batch, error) {
	// Fetch data from IPFS
	data, err := IPFSClient.Cat(batchCID)
	if err != nil {
		log.Errorf("Error fetching data from IPFS: %s", err.Error())
		return nil, err
	}

	// Read the data directly from the reader
	fetchedData, err := io.ReadAll(data)
	if err != nil {
		log.Errorf("Error reading data from IPFS: %v", err)
		return nil, err
	}

	// Unmarshal JSON into the Batch struct
	var batch Batch
	if err := json.Unmarshal(fetchedData, &batch); err != nil {
		log.Errorf("Error unmarshalling fetched data from IPFS: %s", err.Error())
		return nil, err
	}

	return &batch, nil
}
