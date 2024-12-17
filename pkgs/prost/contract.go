package prost

import (
	"context"
	"crypto/tls"
	"net/http"
	"strings"
	"time"
	"validator/config"
	"validator/pkgs/clients"
	"validator/pkgs/contract"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	log "github.com/sirupsen/logrus"
)

var (
	Client      *ethclient.Client
	Instance    *contract.Contract
	ContractABI abi.ABI
	Auth        *bind.TransactOpts
)

func ConfigureClient() {
	rpcClient, err := rpc.DialOptions(context.Background(), config.SettingsObj.ClientUrl, rpc.WithHTTPClient(&http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}))
	if err != nil {
		log.Errorf("Failed to connect to client: %s", err)
		log.Fatal(err)
	}

	Client = ethclient.NewClient(rpcClient)
}

func ConfigureContractInstance() {
	Instance, _ = contract.NewContract(common.HexToAddress(config.SettingsObj.ContractAddress), Client)
}

func ConfigureABI() {
	contractABI, err := abi.JSON(strings.NewReader(contract.ContractMetaData.ABI))
	if err != nil {
		log.Errorf("Failed to configure contract ABI: %s", err)
		log.Fatal(err)
	}

	ContractABI = contractABI
}

func MustQuery[K any](ctx context.Context, call func() (val K, err error)) (K, error) {
	expBackOff := backoff.NewConstantBackOff(1 * time.Second)

	var val K
	operation := func() error {
		var err error
		val, err = call()
		return err
	}
	// Use the retry package to execute the operation with backoff
	err := backoff.Retry(operation, backoff.WithMaxRetries(expBackOff, 3))
	if err != nil {
		clients.SendFailureNotification("Contract query error [MustQuery]", err.Error(), time.Now().String(), "High")
		return *new(K), err
	}
	return val, err
}
