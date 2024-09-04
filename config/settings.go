package config

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var SettingsObj *Settings

type Settings struct {
	ClientUrl            string
	ContractAddress      string
	DataMarketAddress    common.Address
	RedisHost            string
	RedisPort            string
	IPFSUrl              string
	SignerAccountAddress common.Address
	PrivateKey           *ecdsa.PrivateKey
	ChainID              int
	BlockTime            int
	BatchSubmissionLimit int
	SlackReportingUrl    string
	RedisDB              int
}

func LoadConfig() {
	var err error

	missingEnvVars := []string{}

	requiredEnvVars := []string{
		"PROST_RPC_URL",
		"PROTOCOL_STATE_CONTRACT",
		"REDIS_HOST",
		"REDIS_DB",
		"REDIS_PORT",
		"IPFS_URL",
		"DATA_MARKET_CONTRACT",
		"SIGNER_ACCOUNT_PRIVATE_KEY",
		"BATCH_SUBMISSION_LIMIT",
		"PROST_CHAIN_ID",
		"BLOCK_TIME",
		"SLACK_REPORTING_URL",
	}

	for envVar := range requiredEnvVars {
		if getEnv(requiredEnvVars[envVar], "") == "" {
			missingEnvVars = append(missingEnvVars, requiredEnvVars[envVar])
		}
	}

	if len(missingEnvVars) > 0 {
		log.Fatalf("Missing required environment variables: %v", missingEnvVars)
	}

	config := Settings{
		ClientUrl:         getEnv("PROST_RPC_URL", ""),
		ContractAddress:   getEnv("PROTOCOL_STATE_CONTRACT", ""),
		RedisHost:         getEnv("REDIS_HOST", ""),
		RedisPort:         getEnv("REDIS_PORT", ""),
		IPFSUrl:           getEnv("IPFS_URL", ""),
		SlackReportingUrl: getEnv("SLACK_REPORTING_URL", ""),
	}

	config.ChainID, err = strconv.Atoi(getEnv("PROST_CHAIN_ID", ""))
	if err != nil {
		log.Fatalf("PROST_CHAIN_ID is not a valid integer")
	}

	config.BlockTime, err = strconv.Atoi(getEnv("BLOCK_TIME", ""))
	if err != nil {
		log.Fatalf("BLOCK_TIME is not a valid integer")
	}

	config.BatchSubmissionLimit, err = strconv.Atoi(getEnv("BATCH_SUBMISSION_LIMIT", ""))
	if err != nil {
		log.Fatalf("BATCH_SUBMISSION_LIMIT is not a valid integer")
	}

	config.PrivateKey, err = crypto.HexToECDSA(getEnv("SIGNER_ACCOUNT_PRIVATE_KEY", ""))
	if err != nil {
		log.Fatalf("SIGNER_ACCOUNT_PRIVATE_KEY is not a valid private key")
	}

	config.RedisDB, err = strconv.Atoi(getEnv("REDIS_DB", ""))
	if err != nil {
		log.Fatalf("REDIS_DB is not a valid integer")
	}

	// get signer address from private key
	config.SignerAccountAddress = crypto.PubkeyToAddress(config.PrivateKey.PublicKey)

	config.DataMarketAddress = common.HexToAddress(getEnv("DATA_MARKET_CONTRACT", ""))

	log.Infoln("Configuration loaded successfully")
	log.Infoln("Client URL: ", config.ClientUrl)
	log.Infoln("Contract Address: ", config.ContractAddress)
	log.Infoln("Redis Host: ", config.RedisHost)
	log.Infoln("Redis Port: ", config.RedisPort)
	log.Infoln("Redis DB: ", config.RedisDB)
	log.Infoln("IPFS URL: ", config.IPFSUrl)
	log.Infoln("Chain ID: ", config.ChainID)
	log.Infoln("Block Time: ", config.BlockTime)
	log.Infoln("Batch Submission Limit: ", config.BatchSubmissionLimit)
	log.Infoln("Signer Account Address: ", config.SignerAccountAddress.Hex())
	log.Infoln("Data Market Address: ", config.DataMarketAddress.Hex())
	log.Infoln("Slack Reporting URL: ", config.SlackReportingUrl)

	SettingsObj = &config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func checkOptionalEnvVar(value, key string) {
	if value == "" {
		log.Warnf("Optional environment variable %s is not set", key)
	}
}
