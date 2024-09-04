package helpers

import (
	"bytes"
	"encoding/json"
	shell "github.com/ipfs/go-ipfs-api"
	log "github.com/sirupsen/logrus"
	"math/big"
	"validator/config"
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
		return "", err
	}
	return cid, nil
}

func FetchSubmission(sh *shell.Shell, cid string) *Batch {
	data, err := sh.Cat(cid)
	if err != nil {
		return nil
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(data)
	if err != nil {
		return nil
	}

	batch := &Batch{}
	err = json.Unmarshal(buf.Bytes(), batch) // Unmarshal takes a byte slice directly
	if err != nil {
		return nil
	}

	return batch
}
