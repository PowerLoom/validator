package main

import (
	"sync"
	"time"
	"validator/config"
	"validator/pkgs/clients"
	"validator/pkgs/helpers"
)

func main() {
	var wg sync.WaitGroup

	helpers.InitLogger()
	config.LoadConfig()
	clients.InitializeReportingClient(config.SettingsObj.SlackReportingUrl, 60*time.Second)
	helpers.ConfigureClient()
	helpers.ConfigureContractInstance()
	helpers.RedisClient = helpers.NewRedisClient()
	helpers.ConnectIPFSNode()
	helpers.PopulateStateVars()

	wg.Add(1)
	go helpers.StartFetchingBlocks()
	wg.Wait()
}
