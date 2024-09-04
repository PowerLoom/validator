package main

import (
	"sync"
	"validator/config"
	"validator/pkgs/helpers"
)

func main() {
	var wg sync.WaitGroup

	helpers.InitLogger()
	config.LoadConfig()
	helpers.ConfigureClient()
	helpers.ConfigureContractInstance()
	helpers.RedisClient = helpers.NewRedisClient()
	helpers.ConnectIPFSNode()
	helpers.PopulateStateVars()

	wg.Add(1)
	go helpers.StartFetchingBlocks()
	wg.Wait()
}
