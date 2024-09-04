package config

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

var SettingsObj *Settings

type Settings struct {
	ClientUrl               string `json:"ClientUrl"`
	ContractAddress         string `json:"ContractAddress"`
	RedisHost               string `json:"RedisHost"`
	RedisPort               string `json:"RedisPort"`
	IPFSUrl                 string `json:"IPFSUrl"`
	SignerAccountAddressStr string `json:"SignerAccountAddress"`
	SignerAccountAddress    common.Address
	PrivateKeyStr           string `json:"PrivateKey"`
	PrivateKey              *ecdsa.PrivateKey
	ChainID                 int `json:"ChainID"`
	BlockTime               int `json:"BlockTime"`
	BatchSubmissionLimit    int `json:"BatchSubmissionLimit"`
}

func LoadConfig() {
	//file, err := os.Open("/Users/mukundrawat/power2/validator-alpha/config/settings.json")
	file, err := os.Open(strings.TrimSuffix(os.Getenv("CONFIG_PATH"), "/") + "/config/settings.json")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Errorf("Unable to close file: %s", err.Error())
		}
	}(file)

	decoder := json.NewDecoder(file)
	config := Settings{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}

	config.SignerAccountAddress = common.HexToAddress(config.SignerAccountAddressStr)
	config.PrivateKey, _ = crypto.HexToECDSA(config.PrivateKeyStr)

	SettingsObj = &config
}
