package helpers

import (
	"bytes"
	"encoding/json"
	shell "github.com/ipfs/go-ipfs-api"
	log "github.com/sirupsen/logrus"
	"math/big"
	"time"
	"validator/config"
	"validator/pkgs/clients"
)

var IPFSCon *shell.Shell

// Batch represents your data structure
type Batch struct {
	ID            *big.Int `json:"id"`
	SubmissionIds []string `json:"submissionIds"`
	Submissions   []string `json:"submissions"`
	RootHash      string   `json:"roothash"`
}

// Connect to the local IPFS node
func ConnectIPFSNode() {
	log.Debugf("Connecting to IPFS host: %s", config.SettingsObj.IPFSUrl)
	IPFSCon = shell.NewShell(config.SettingsObj.IPFSUrl)
}

func StoreOnIPFS(sh *shell.Shell, data *Batch) (string, error) {
	jsonData, err := json.Marshal(data)
	cid, err := sh.Add(bytes.NewReader(jsonData))
	if err != nil {
		clients.SendFailureNotification("IPFS", "Error Storing on IPFS: "+err.Error(), time.Now().String(), "High")
		return "", err
	}
	return cid, nil
}

func FetchSubmission(sh *shell.Shell, cid string) *Batch {
	data, err := sh.Cat(cid)
	if err != nil {
		clients.SendFailureNotification("IPFS", "Error Fetching from IPFS: "+err.Error(), time.Now().String(), "High")
		return nil
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(data)
	if err != nil {
		clients.SendFailureNotification("IPFS", "Error Fetching from IPFS: "+err.Error(), time.Now().String(), "High")
		return nil
	}

	batch := &Batch{}
	err = json.Unmarshal(buf.Bytes(), batch) // Unmarshal takes a byte slice directly
	if err != nil {
		clients.SendFailureNotification("IPFS", "Error unmarshalling fetched data from IPFS: "+err.Error(), time.Now().String(), "High")
		return nil
	}

	return batch
}
