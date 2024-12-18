package main

import (
	"sync"
	"time"
	"validator/config"
	"validator/pkgs/clients"
	"validator/pkgs/ipfs"
	"validator/pkgs/prost"
	"validator/pkgs/redis"
	"validator/pkgs/utils"
)

func main() {
	// Initiate logger
	utils.InitLogger()

	// Load the config object
	config.LoadConfig()

	// Initialize reporting service
	clients.InitializeReportingClient(config.SettingsObj.SlackReportingUrl, 5*time.Second)

	// Initialize tx relayer service
	clients.InitializeTxClient(config.SettingsObj.TxRelayerUrl, time.Duration(config.SettingsObj.HttpTimeout)*time.Second)

	// Setup redis
	redis.RedisClient = redis.NewRedisClient()

	// Connect to IPFS node
	ipfs.ConnectIPFSNode()

	// Set up the RPC client, contract, and ABI instance
	prost.ConfigureClient()
	prost.ConfigureContractInstance()
	prost.ConfigureABI()

	var wg sync.WaitGroup

	wg.Add(1)
	go prost.StartBatchAttestation() // Start batch attestation process
	wg.Wait()
}
