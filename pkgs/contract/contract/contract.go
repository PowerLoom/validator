// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PowerloomProtocolStateRequest is an auto generated low-level Go binding around an user-defined struct.
type PowerloomProtocolStateRequest struct {
	Deadline    *big.Int
	SnapshotCid string
	EpochId     *big.Int
	ProjectId   string
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"epochSize\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainBlockTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useBlockNumberAsEpochId\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"snapshotterAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DelayedSnapshotSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EpochReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"enableEpochId\",\"type\":\"uint256\"}],\"name\":\"ProjectsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SnapshotBatchFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SnapshotBatchSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"ValidatorsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"snapshotterAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"allSnapshottersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"snapshotterAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"masterSnapshottersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"enableEpochId\",\"type\":\"uint256\"}],\"name\":\"pretestProjectsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DeploymentBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EPOCH_SIZE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_BLOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USE_BLOCK_NUMBER_AS_EPOCH_ID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allSnapshotters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"name\":\"attestationFinalizedStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"attestationsReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"attestationsReceivedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"batchIdAttestationStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batchIdToProjects\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"projectids\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochEnd\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"forceSkipEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMasterSnapshotters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSnapshotters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalMasterSnapshotterCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSnapshotterCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"name\":\"lastFinalizedSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"masterSnapshotters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"maxAttestationFinalizedRootHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"maxAttestationsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"maxSnapshotsCid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSubmissionsForConsensus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"name\":\"projectFirstEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recoverAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"releaseEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"snapshotStatus\",\"outputs\":[{\"internalType\":\"enumPowerloomProtocolState.SnapshotStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshotSubmissionWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"finalizedCidsRootHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchAttestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"projectIds\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"snapshotCids\",\"type\":\"string[]\"}],\"name\":\"submitSubmissionBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_snapshotters\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_status\",\"type\":\"bool[]\"}],\"name\":\"updateAllSnapshotters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_snapshotters\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_status\",\"type\":\"bool[]\"}],\"name\":\"updateMasterSnapshotters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSubmissionsForConsensus\",\"type\":\"uint256\"}],\"name\":\"updateMinSnapshottersForConsensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newsnapshotSubmissionWindow\",\"type\":\"uint256\"}],\"name\":\"updateSnapshotSubmissionWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_status\",\"type\":\"bool[]\"}],\"name\":\"updateValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"internalType\":\"structPowerloomProtocolState.Request\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// DeploymentBlockNumber is a free data retrieval call binding the contract method 0xb1288f71.
//
// Solidity: function DeploymentBlockNumber() view returns(uint256)
func (_Contract *ContractCaller) DeploymentBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DeploymentBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeploymentBlockNumber is a free data retrieval call binding the contract method 0xb1288f71.
//
// Solidity: function DeploymentBlockNumber() view returns(uint256)
func (_Contract *ContractSession) DeploymentBlockNumber() (*big.Int, error) {
	return _Contract.Contract.DeploymentBlockNumber(&_Contract.CallOpts)
}

// DeploymentBlockNumber is a free data retrieval call binding the contract method 0xb1288f71.
//
// Solidity: function DeploymentBlockNumber() view returns(uint256)
func (_Contract *ContractCallerSession) DeploymentBlockNumber() (*big.Int, error) {
	return _Contract.Contract.DeploymentBlockNumber(&_Contract.CallOpts)
}

// EPOCHSIZE is a free data retrieval call binding the contract method 0x62656003.
//
// Solidity: function EPOCH_SIZE() view returns(uint8)
func (_Contract *ContractCaller) EPOCHSIZE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "EPOCH_SIZE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// EPOCHSIZE is a free data retrieval call binding the contract method 0x62656003.
//
// Solidity: function EPOCH_SIZE() view returns(uint8)
func (_Contract *ContractSession) EPOCHSIZE() (uint8, error) {
	return _Contract.Contract.EPOCHSIZE(&_Contract.CallOpts)
}

// EPOCHSIZE is a free data retrieval call binding the contract method 0x62656003.
//
// Solidity: function EPOCH_SIZE() view returns(uint8)
func (_Contract *ContractCallerSession) EPOCHSIZE() (uint8, error) {
	return _Contract.Contract.EPOCHSIZE(&_Contract.CallOpts)
}

// SOURCECHAINBLOCKTIME is a free data retrieval call binding the contract method 0x351b6155.
//
// Solidity: function SOURCE_CHAIN_BLOCK_TIME() view returns(uint256)
func (_Contract *ContractCaller) SOURCECHAINBLOCKTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "SOURCE_CHAIN_BLOCK_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SOURCECHAINBLOCKTIME is a free data retrieval call binding the contract method 0x351b6155.
//
// Solidity: function SOURCE_CHAIN_BLOCK_TIME() view returns(uint256)
func (_Contract *ContractSession) SOURCECHAINBLOCKTIME() (*big.Int, error) {
	return _Contract.Contract.SOURCECHAINBLOCKTIME(&_Contract.CallOpts)
}

// SOURCECHAINBLOCKTIME is a free data retrieval call binding the contract method 0x351b6155.
//
// Solidity: function SOURCE_CHAIN_BLOCK_TIME() view returns(uint256)
func (_Contract *ContractCallerSession) SOURCECHAINBLOCKTIME() (*big.Int, error) {
	return _Contract.Contract.SOURCECHAINBLOCKTIME(&_Contract.CallOpts)
}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(uint256)
func (_Contract *ContractCaller) SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(uint256)
func (_Contract *ContractSession) SOURCECHAINID() (*big.Int, error) {
	return _Contract.Contract.SOURCECHAINID(&_Contract.CallOpts)
}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(uint256)
func (_Contract *ContractCallerSession) SOURCECHAINID() (*big.Int, error) {
	return _Contract.Contract.SOURCECHAINID(&_Contract.CallOpts)
}

// USEBLOCKNUMBERASEPOCHID is a free data retrieval call binding the contract method 0x2d46247b.
//
// Solidity: function USE_BLOCK_NUMBER_AS_EPOCH_ID() view returns(bool)
func (_Contract *ContractCaller) USEBLOCKNUMBERASEPOCHID(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "USE_BLOCK_NUMBER_AS_EPOCH_ID")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// USEBLOCKNUMBERASEPOCHID is a free data retrieval call binding the contract method 0x2d46247b.
//
// Solidity: function USE_BLOCK_NUMBER_AS_EPOCH_ID() view returns(bool)
func (_Contract *ContractSession) USEBLOCKNUMBERASEPOCHID() (bool, error) {
	return _Contract.Contract.USEBLOCKNUMBERASEPOCHID(&_Contract.CallOpts)
}

// USEBLOCKNUMBERASEPOCHID is a free data retrieval call binding the contract method 0x2d46247b.
//
// Solidity: function USE_BLOCK_NUMBER_AS_EPOCH_ID() view returns(bool)
func (_Contract *ContractCallerSession) USEBLOCKNUMBERASEPOCHID() (bool, error) {
	return _Contract.Contract.USEBLOCKNUMBERASEPOCHID(&_Contract.CallOpts)
}

// AllSnapshotters is a free data retrieval call binding the contract method 0x3d15d0f4.
//
// Solidity: function allSnapshotters(address ) view returns(bool)
func (_Contract *ContractCaller) AllSnapshotters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allSnapshotters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllSnapshotters is a free data retrieval call binding the contract method 0x3d15d0f4.
//
// Solidity: function allSnapshotters(address ) view returns(bool)
func (_Contract *ContractSession) AllSnapshotters(arg0 common.Address) (bool, error) {
	return _Contract.Contract.AllSnapshotters(&_Contract.CallOpts, arg0)
}

// AllSnapshotters is a free data retrieval call binding the contract method 0x3d15d0f4.
//
// Solidity: function allSnapshotters(address ) view returns(bool)
func (_Contract *ContractCallerSession) AllSnapshotters(arg0 common.Address) (bool, error) {
	return _Contract.Contract.AllSnapshotters(&_Contract.CallOpts, arg0)
}

// AttestationFinalizedStatus is a free data retrieval call binding the contract method 0xbfefe303.
//
// Solidity: function attestationFinalizedStatus(uint256 epochId, string projectId) view returns(bool finalized)
func (_Contract *ContractCaller) AttestationFinalizedStatus(opts *bind.CallOpts, epochId *big.Int, projectId string) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestationFinalizedStatus", epochId, projectId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AttestationFinalizedStatus is a free data retrieval call binding the contract method 0xbfefe303.
//
// Solidity: function attestationFinalizedStatus(uint256 epochId, string projectId) view returns(bool finalized)
func (_Contract *ContractSession) AttestationFinalizedStatus(epochId *big.Int, projectId string) (bool, error) {
	return _Contract.Contract.AttestationFinalizedStatus(&_Contract.CallOpts, epochId, projectId)
}

// AttestationFinalizedStatus is a free data retrieval call binding the contract method 0xbfefe303.
//
// Solidity: function attestationFinalizedStatus(uint256 epochId, string projectId) view returns(bool finalized)
func (_Contract *ContractCallerSession) AttestationFinalizedStatus(epochId *big.Int, projectId string) (bool, error) {
	return _Contract.Contract.AttestationFinalizedStatus(&_Contract.CallOpts, epochId, projectId)
}

// AttestationsReceived is a free data retrieval call binding the contract method 0x33420a0c.
//
// Solidity: function attestationsReceived(uint256 batchId, address validatorAddr) view returns(bool)
func (_Contract *ContractCaller) AttestationsReceived(opts *bind.CallOpts, batchId *big.Int, validatorAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestationsReceived", batchId, validatorAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AttestationsReceived is a free data retrieval call binding the contract method 0x33420a0c.
//
// Solidity: function attestationsReceived(uint256 batchId, address validatorAddr) view returns(bool)
func (_Contract *ContractSession) AttestationsReceived(batchId *big.Int, validatorAddr common.Address) (bool, error) {
	return _Contract.Contract.AttestationsReceived(&_Contract.CallOpts, batchId, validatorAddr)
}

// AttestationsReceived is a free data retrieval call binding the contract method 0x33420a0c.
//
// Solidity: function attestationsReceived(uint256 batchId, address validatorAddr) view returns(bool)
func (_Contract *ContractCallerSession) AttestationsReceived(batchId *big.Int, validatorAddr common.Address) (bool, error) {
	return _Contract.Contract.AttestationsReceived(&_Contract.CallOpts, batchId, validatorAddr)
}

// AttestationsReceivedCount is a free data retrieval call binding the contract method 0x66f94346.
//
// Solidity: function attestationsReceivedCount(uint256 batchId, bytes32 rootHash) view returns(uint256 count)
func (_Contract *ContractCaller) AttestationsReceivedCount(opts *bind.CallOpts, batchId *big.Int, rootHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestationsReceivedCount", batchId, rootHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationsReceivedCount is a free data retrieval call binding the contract method 0x66f94346.
//
// Solidity: function attestationsReceivedCount(uint256 batchId, bytes32 rootHash) view returns(uint256 count)
func (_Contract *ContractSession) AttestationsReceivedCount(batchId *big.Int, rootHash [32]byte) (*big.Int, error) {
	return _Contract.Contract.AttestationsReceivedCount(&_Contract.CallOpts, batchId, rootHash)
}

// AttestationsReceivedCount is a free data retrieval call binding the contract method 0x66f94346.
//
// Solidity: function attestationsReceivedCount(uint256 batchId, bytes32 rootHash) view returns(uint256 count)
func (_Contract *ContractCallerSession) AttestationsReceivedCount(batchId *big.Int, rootHash [32]byte) (*big.Int, error) {
	return _Contract.Contract.AttestationsReceivedCount(&_Contract.CallOpts, batchId, rootHash)
}

// BatchIdAttestationStatus is a free data retrieval call binding the contract method 0xa047977d.
//
// Solidity: function batchIdAttestationStatus(uint256 batchId) view returns(bool status)
func (_Contract *ContractCaller) BatchIdAttestationStatus(opts *bind.CallOpts, batchId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "batchIdAttestationStatus", batchId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchIdAttestationStatus is a free data retrieval call binding the contract method 0xa047977d.
//
// Solidity: function batchIdAttestationStatus(uint256 batchId) view returns(bool status)
func (_Contract *ContractSession) BatchIdAttestationStatus(batchId *big.Int) (bool, error) {
	return _Contract.Contract.BatchIdAttestationStatus(&_Contract.CallOpts, batchId)
}

// BatchIdAttestationStatus is a free data retrieval call binding the contract method 0xa047977d.
//
// Solidity: function batchIdAttestationStatus(uint256 batchId) view returns(bool status)
func (_Contract *ContractCallerSession) BatchIdAttestationStatus(batchId *big.Int) (bool, error) {
	return _Contract.Contract.BatchIdAttestationStatus(&_Contract.CallOpts, batchId)
}

// BatchIdToProjects is a free data retrieval call binding the contract method 0x510c9f4c.
//
// Solidity: function batchIdToProjects(uint256 batchId, uint256 ) view returns(string projectids)
func (_Contract *ContractCaller) BatchIdToProjects(opts *bind.CallOpts, batchId *big.Int, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "batchIdToProjects", batchId, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BatchIdToProjects is a free data retrieval call binding the contract method 0x510c9f4c.
//
// Solidity: function batchIdToProjects(uint256 batchId, uint256 ) view returns(string projectids)
func (_Contract *ContractSession) BatchIdToProjects(batchId *big.Int, arg1 *big.Int) (string, error) {
	return _Contract.Contract.BatchIdToProjects(&_Contract.CallOpts, batchId, arg1)
}

// BatchIdToProjects is a free data retrieval call binding the contract method 0x510c9f4c.
//
// Solidity: function batchIdToProjects(uint256 batchId, uint256 ) view returns(string projectids)
func (_Contract *ContractCallerSession) BatchIdToProjects(batchId *big.Int, arg1 *big.Int) (string, error) {
	return _Contract.Contract.BatchIdToProjects(&_Contract.CallOpts, batchId, arg1)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256 begin, uint256 end, uint256 epochId)
func (_Contract *ContractCaller) CurrentEpoch(opts *bind.CallOpts) (struct {
	Begin   *big.Int
	End     *big.Int
	EpochId *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentEpoch")

	outstruct := new(struct {
		Begin   *big.Int
		End     *big.Int
		EpochId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Begin = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256 begin, uint256 end, uint256 epochId)
func (_Contract *ContractSession) CurrentEpoch() (struct {
	Begin   *big.Int
	End     *big.Int
	EpochId *big.Int
}, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256 begin, uint256 end, uint256 epochId)
func (_Contract *ContractCallerSession) CurrentEpoch() (struct {
	Begin   *big.Int
	End     *big.Int
	EpochId *big.Int
}, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) view returns(uint256 timestamp, uint256 blocknumber, uint256 epochEnd)
func (_Contract *ContractCaller) EpochInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	EpochEnd    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "epochInfo", arg0)

	outstruct := new(struct {
		Timestamp   *big.Int
		Blocknumber *big.Int
		EpochEnd    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Blocknumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochEnd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) view returns(uint256 timestamp, uint256 blocknumber, uint256 epochEnd)
func (_Contract *ContractSession) EpochInfo(arg0 *big.Int) (struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	EpochEnd    *big.Int
}, error) {
	return _Contract.Contract.EpochInfo(&_Contract.CallOpts, arg0)
}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) view returns(uint256 timestamp, uint256 blocknumber, uint256 epochEnd)
func (_Contract *ContractCallerSession) EpochInfo(arg0 *big.Int) (struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	EpochEnd    *big.Int
}, error) {
	return _Contract.Contract.EpochInfo(&_Contract.CallOpts, arg0)
}

// GetMasterSnapshotters is a free data retrieval call binding the contract method 0x90110313.
//
// Solidity: function getMasterSnapshotters() view returns(address[])
func (_Contract *ContractCaller) GetMasterSnapshotters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getMasterSnapshotters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMasterSnapshotters is a free data retrieval call binding the contract method 0x90110313.
//
// Solidity: function getMasterSnapshotters() view returns(address[])
func (_Contract *ContractSession) GetMasterSnapshotters() ([]common.Address, error) {
	return _Contract.Contract.GetMasterSnapshotters(&_Contract.CallOpts)
}

// GetMasterSnapshotters is a free data retrieval call binding the contract method 0x90110313.
//
// Solidity: function getMasterSnapshotters() view returns(address[])
func (_Contract *ContractCallerSession) GetMasterSnapshotters() ([]common.Address, error) {
	return _Contract.Contract.GetMasterSnapshotters(&_Contract.CallOpts)
}

// GetSnapshotters is a free data retrieval call binding the contract method 0x8deed336.
//
// Solidity: function getSnapshotters() view returns(address[])
func (_Contract *ContractCaller) GetSnapshotters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSnapshotters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSnapshotters is a free data retrieval call binding the contract method 0x8deed336.
//
// Solidity: function getSnapshotters() view returns(address[])
func (_Contract *ContractSession) GetSnapshotters() ([]common.Address, error) {
	return _Contract.Contract.GetSnapshotters(&_Contract.CallOpts)
}

// GetSnapshotters is a free data retrieval call binding the contract method 0x8deed336.
//
// Solidity: function getSnapshotters() view returns(address[])
func (_Contract *ContractCallerSession) GetSnapshotters() ([]common.Address, error) {
	return _Contract.Contract.GetSnapshotters(&_Contract.CallOpts)
}

// GetTotalMasterSnapshotterCount is a free data retrieval call binding the contract method 0xd551672b.
//
// Solidity: function getTotalMasterSnapshotterCount() view returns(uint256)
func (_Contract *ContractCaller) GetTotalMasterSnapshotterCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalMasterSnapshotterCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalMasterSnapshotterCount is a free data retrieval call binding the contract method 0xd551672b.
//
// Solidity: function getTotalMasterSnapshotterCount() view returns(uint256)
func (_Contract *ContractSession) GetTotalMasterSnapshotterCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalMasterSnapshotterCount(&_Contract.CallOpts)
}

// GetTotalMasterSnapshotterCount is a free data retrieval call binding the contract method 0xd551672b.
//
// Solidity: function getTotalMasterSnapshotterCount() view returns(uint256)
func (_Contract *ContractCallerSession) GetTotalMasterSnapshotterCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalMasterSnapshotterCount(&_Contract.CallOpts)
}

// GetTotalSnapshotterCount is a free data retrieval call binding the contract method 0x92ae6f66.
//
// Solidity: function getTotalSnapshotterCount() view returns(uint256)
func (_Contract *ContractCaller) GetTotalSnapshotterCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalSnapshotterCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSnapshotterCount is a free data retrieval call binding the contract method 0x92ae6f66.
//
// Solidity: function getTotalSnapshotterCount() view returns(uint256)
func (_Contract *ContractSession) GetTotalSnapshotterCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalSnapshotterCount(&_Contract.CallOpts)
}

// GetTotalSnapshotterCount is a free data retrieval call binding the contract method 0x92ae6f66.
//
// Solidity: function getTotalSnapshotterCount() view returns(uint256)
func (_Contract *ContractCallerSession) GetTotalSnapshotterCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalSnapshotterCount(&_Contract.CallOpts)
}

// GetTotalValidatorsCount is a free data retrieval call binding the contract method 0x983d52e7.
//
// Solidity: function getTotalValidatorsCount() view returns(uint256)
func (_Contract *ContractCaller) GetTotalValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalValidatorsCount is a free data retrieval call binding the contract method 0x983d52e7.
//
// Solidity: function getTotalValidatorsCount() view returns(uint256)
func (_Contract *ContractSession) GetTotalValidatorsCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalValidatorsCount(&_Contract.CallOpts)
}

// GetTotalValidatorsCount is a free data retrieval call binding the contract method 0x983d52e7.
//
// Solidity: function getTotalValidatorsCount() view returns(uint256)
func (_Contract *ContractCallerSession) GetTotalValidatorsCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalValidatorsCount(&_Contract.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Contract *ContractCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Contract *ContractSession) GetValidators() ([]common.Address, error) {
	return _Contract.Contract.GetValidators(&_Contract.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Contract *ContractCallerSession) GetValidators() ([]common.Address, error) {
	return _Contract.Contract.GetValidators(&_Contract.CallOpts)
}

// LastFinalizedSnapshot is a free data retrieval call binding the contract method 0x4ea16b0a.
//
// Solidity: function lastFinalizedSnapshot(string projectId) view returns(uint256 epochId)
func (_Contract *ContractCaller) LastFinalizedSnapshot(opts *bind.CallOpts, projectId string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastFinalizedSnapshot", projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFinalizedSnapshot is a free data retrieval call binding the contract method 0x4ea16b0a.
//
// Solidity: function lastFinalizedSnapshot(string projectId) view returns(uint256 epochId)
func (_Contract *ContractSession) LastFinalizedSnapshot(projectId string) (*big.Int, error) {
	return _Contract.Contract.LastFinalizedSnapshot(&_Contract.CallOpts, projectId)
}

// LastFinalizedSnapshot is a free data retrieval call binding the contract method 0x4ea16b0a.
//
// Solidity: function lastFinalizedSnapshot(string projectId) view returns(uint256 epochId)
func (_Contract *ContractCallerSession) LastFinalizedSnapshot(projectId string) (*big.Int, error) {
	return _Contract.Contract.LastFinalizedSnapshot(&_Contract.CallOpts, projectId)
}

// MasterSnapshotters is a free data retrieval call binding the contract method 0x34b739d9.
//
// Solidity: function masterSnapshotters(address ) view returns(bool)
func (_Contract *ContractCaller) MasterSnapshotters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "masterSnapshotters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MasterSnapshotters is a free data retrieval call binding the contract method 0x34b739d9.
//
// Solidity: function masterSnapshotters(address ) view returns(bool)
func (_Contract *ContractSession) MasterSnapshotters(arg0 common.Address) (bool, error) {
	return _Contract.Contract.MasterSnapshotters(&_Contract.CallOpts, arg0)
}

// MasterSnapshotters is a free data retrieval call binding the contract method 0x34b739d9.
//
// Solidity: function masterSnapshotters(address ) view returns(bool)
func (_Contract *ContractCallerSession) MasterSnapshotters(arg0 common.Address) (bool, error) {
	return _Contract.Contract.MasterSnapshotters(&_Contract.CallOpts, arg0)
}

// MaxAttestationFinalizedRootHash is a free data retrieval call binding the contract method 0xcc542228.
//
// Solidity: function maxAttestationFinalizedRootHash(uint256 batchId) view returns(bytes32 rootHash)
func (_Contract *ContractCaller) MaxAttestationFinalizedRootHash(opts *bind.CallOpts, batchId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxAttestationFinalizedRootHash", batchId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MaxAttestationFinalizedRootHash is a free data retrieval call binding the contract method 0xcc542228.
//
// Solidity: function maxAttestationFinalizedRootHash(uint256 batchId) view returns(bytes32 rootHash)
func (_Contract *ContractSession) MaxAttestationFinalizedRootHash(batchId *big.Int) ([32]byte, error) {
	return _Contract.Contract.MaxAttestationFinalizedRootHash(&_Contract.CallOpts, batchId)
}

// MaxAttestationFinalizedRootHash is a free data retrieval call binding the contract method 0xcc542228.
//
// Solidity: function maxAttestationFinalizedRootHash(uint256 batchId) view returns(bytes32 rootHash)
func (_Contract *ContractCallerSession) MaxAttestationFinalizedRootHash(batchId *big.Int) ([32]byte, error) {
	return _Contract.Contract.MaxAttestationFinalizedRootHash(&_Contract.CallOpts, batchId)
}

// MaxAttestationsCount is a free data retrieval call binding the contract method 0x30c8ecb2.
//
// Solidity: function maxAttestationsCount(uint256 batchId) view returns(uint256 count)
func (_Contract *ContractCaller) MaxAttestationsCount(opts *bind.CallOpts, batchId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxAttestationsCount", batchId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAttestationsCount is a free data retrieval call binding the contract method 0x30c8ecb2.
//
// Solidity: function maxAttestationsCount(uint256 batchId) view returns(uint256 count)
func (_Contract *ContractSession) MaxAttestationsCount(batchId *big.Int) (*big.Int, error) {
	return _Contract.Contract.MaxAttestationsCount(&_Contract.CallOpts, batchId)
}

// MaxAttestationsCount is a free data retrieval call binding the contract method 0x30c8ecb2.
//
// Solidity: function maxAttestationsCount(uint256 batchId) view returns(uint256 count)
func (_Contract *ContractCallerSession) MaxAttestationsCount(batchId *big.Int) (*big.Int, error) {
	return _Contract.Contract.MaxAttestationsCount(&_Contract.CallOpts, batchId)
}

// MaxSnapshotsCid is a free data retrieval call binding the contract method 0xc2b97d4c.
//
// Solidity: function maxSnapshotsCid(string , uint256 ) view returns(string)
func (_Contract *ContractCaller) MaxSnapshotsCid(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxSnapshotsCid", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MaxSnapshotsCid is a free data retrieval call binding the contract method 0xc2b97d4c.
//
// Solidity: function maxSnapshotsCid(string , uint256 ) view returns(string)
func (_Contract *ContractSession) MaxSnapshotsCid(arg0 string, arg1 *big.Int) (string, error) {
	return _Contract.Contract.MaxSnapshotsCid(&_Contract.CallOpts, arg0, arg1)
}

// MaxSnapshotsCid is a free data retrieval call binding the contract method 0xc2b97d4c.
//
// Solidity: function maxSnapshotsCid(string , uint256 ) view returns(string)
func (_Contract *ContractCallerSession) MaxSnapshotsCid(arg0 string, arg1 *big.Int) (string, error) {
	return _Contract.Contract.MaxSnapshotsCid(&_Contract.CallOpts, arg0, arg1)
}

// MinSubmissionsForConsensus is a free data retrieval call binding the contract method 0x66752e1e.
//
// Solidity: function minSubmissionsForConsensus() view returns(uint256)
func (_Contract *ContractCaller) MinSubmissionsForConsensus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minSubmissionsForConsensus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSubmissionsForConsensus is a free data retrieval call binding the contract method 0x66752e1e.
//
// Solidity: function minSubmissionsForConsensus() view returns(uint256)
func (_Contract *ContractSession) MinSubmissionsForConsensus() (*big.Int, error) {
	return _Contract.Contract.MinSubmissionsForConsensus(&_Contract.CallOpts)
}

// MinSubmissionsForConsensus is a free data retrieval call binding the contract method 0x66752e1e.
//
// Solidity: function minSubmissionsForConsensus() view returns(uint256)
func (_Contract *ContractCallerSession) MinSubmissionsForConsensus() (*big.Int, error) {
	return _Contract.Contract.MinSubmissionsForConsensus(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// ProjectFirstEpochId is a free data retrieval call binding the contract method 0xfa30dbe0.
//
// Solidity: function projectFirstEpochId(string projectId) view returns(uint256 epochId)
func (_Contract *ContractCaller) ProjectFirstEpochId(opts *bind.CallOpts, projectId string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "projectFirstEpochId", projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectFirstEpochId is a free data retrieval call binding the contract method 0xfa30dbe0.
//
// Solidity: function projectFirstEpochId(string projectId) view returns(uint256 epochId)
func (_Contract *ContractSession) ProjectFirstEpochId(projectId string) (*big.Int, error) {
	return _Contract.Contract.ProjectFirstEpochId(&_Contract.CallOpts, projectId)
}

// ProjectFirstEpochId is a free data retrieval call binding the contract method 0xfa30dbe0.
//
// Solidity: function projectFirstEpochId(string projectId) view returns(uint256 epochId)
func (_Contract *ContractCallerSession) ProjectFirstEpochId(projectId string) (*big.Int, error) {
	return _Contract.Contract.ProjectFirstEpochId(&_Contract.CallOpts, projectId)
}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 messageHash, bytes signature) pure returns(address)
func (_Contract *ContractCaller) RecoverAddress(opts *bind.CallOpts, messageHash [32]byte, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "recoverAddress", messageHash, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 messageHash, bytes signature) pure returns(address)
func (_Contract *ContractSession) RecoverAddress(messageHash [32]byte, signature []byte) (common.Address, error) {
	return _Contract.Contract.RecoverAddress(&_Contract.CallOpts, messageHash, signature)
}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 messageHash, bytes signature) pure returns(address)
func (_Contract *ContractCallerSession) RecoverAddress(messageHash [32]byte, signature []byte) (common.Address, error) {
	return _Contract.Contract.RecoverAddress(&_Contract.CallOpts, messageHash, signature)
}

// SnapshotStatus is a free data retrieval call binding the contract method 0x3aaf384d.
//
// Solidity: function snapshotStatus(string , uint256 ) view returns(uint8 status, string snapshotCid, uint256 timestamp)
func (_Contract *ContractCaller) SnapshotStatus(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Status      uint8
	SnapshotCid string
	Timestamp   *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "snapshotStatus", arg0, arg1)

	outstruct := new(struct {
		Status      uint8
		SnapshotCid string
		Timestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.SnapshotCid = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SnapshotStatus is a free data retrieval call binding the contract method 0x3aaf384d.
//
// Solidity: function snapshotStatus(string , uint256 ) view returns(uint8 status, string snapshotCid, uint256 timestamp)
func (_Contract *ContractSession) SnapshotStatus(arg0 string, arg1 *big.Int) (struct {
	Status      uint8
	SnapshotCid string
	Timestamp   *big.Int
}, error) {
	return _Contract.Contract.SnapshotStatus(&_Contract.CallOpts, arg0, arg1)
}

// SnapshotStatus is a free data retrieval call binding the contract method 0x3aaf384d.
//
// Solidity: function snapshotStatus(string , uint256 ) view returns(uint8 status, string snapshotCid, uint256 timestamp)
func (_Contract *ContractCallerSession) SnapshotStatus(arg0 string, arg1 *big.Int) (struct {
	Status      uint8
	SnapshotCid string
	Timestamp   *big.Int
}, error) {
	return _Contract.Contract.SnapshotStatus(&_Contract.CallOpts, arg0, arg1)
}

// SnapshotSubmissionWindow is a free data retrieval call binding the contract method 0x059080f6.
//
// Solidity: function snapshotSubmissionWindow() view returns(uint256)
func (_Contract *ContractCaller) SnapshotSubmissionWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "snapshotSubmissionWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SnapshotSubmissionWindow is a free data retrieval call binding the contract method 0x059080f6.
//
// Solidity: function snapshotSubmissionWindow() view returns(uint256)
func (_Contract *ContractSession) SnapshotSubmissionWindow() (*big.Int, error) {
	return _Contract.Contract.SnapshotSubmissionWindow(&_Contract.CallOpts)
}

// SnapshotSubmissionWindow is a free data retrieval call binding the contract method 0x059080f6.
//
// Solidity: function snapshotSubmissionWindow() view returns(uint256)
func (_Contract *ContractCallerSession) SnapshotSubmissionWindow() (*big.Int, error) {
	return _Contract.Contract.SnapshotSubmissionWindow(&_Contract.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xc22e1a53.
//
// Solidity: function verify(string snapshotCid, uint256 epochId, string projectId, (uint256,string,uint256,string) request, bytes signature, address signer) view returns(bool)
func (_Contract *ContractCaller) Verify(opts *bind.CallOpts, snapshotCid string, epochId *big.Int, projectId string, request PowerloomProtocolStateRequest, signature []byte, signer common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verify", snapshotCid, epochId, projectId, request, signature, signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xc22e1a53.
//
// Solidity: function verify(string snapshotCid, uint256 epochId, string projectId, (uint256,string,uint256,string) request, bytes signature, address signer) view returns(bool)
func (_Contract *ContractSession) Verify(snapshotCid string, epochId *big.Int, projectId string, request PowerloomProtocolStateRequest, signature []byte, signer common.Address) (bool, error) {
	return _Contract.Contract.Verify(&_Contract.CallOpts, snapshotCid, epochId, projectId, request, signature, signer)
}

// Verify is a free data retrieval call binding the contract method 0xc22e1a53.
//
// Solidity: function verify(string snapshotCid, uint256 epochId, string projectId, (uint256,string,uint256,string) request, bytes signature, address signer) view returns(bool)
func (_Contract *ContractCallerSession) Verify(snapshotCid string, epochId *big.Int, projectId string, request PowerloomProtocolStateRequest, signature []byte, signer common.Address) (bool, error) {
	return _Contract.Contract.Verify(&_Contract.CallOpts, snapshotCid, epochId, projectId, request, signature, signer)
}

// ForceSkipEpoch is a paid mutator transaction binding the contract method 0xf537a3e2.
//
// Solidity: function forceSkipEpoch(uint256 begin, uint256 end) returns()
func (_Contract *ContractTransactor) ForceSkipEpoch(opts *bind.TransactOpts, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "forceSkipEpoch", begin, end)
}

// ForceSkipEpoch is a paid mutator transaction binding the contract method 0xf537a3e2.
//
// Solidity: function forceSkipEpoch(uint256 begin, uint256 end) returns()
func (_Contract *ContractSession) ForceSkipEpoch(begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ForceSkipEpoch(&_Contract.TransactOpts, begin, end)
}

// ForceSkipEpoch is a paid mutator transaction binding the contract method 0xf537a3e2.
//
// Solidity: function forceSkipEpoch(uint256 begin, uint256 end) returns()
func (_Contract *ContractTransactorSession) ForceSkipEpoch(begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ForceSkipEpoch(&_Contract.TransactOpts, begin, end)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0x132c290f.
//
// Solidity: function releaseEpoch(uint256 begin, uint256 end) returns()
func (_Contract *ContractTransactor) ReleaseEpoch(opts *bind.TransactOpts, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "releaseEpoch", begin, end)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0x132c290f.
//
// Solidity: function releaseEpoch(uint256 begin, uint256 end) returns()
func (_Contract *ContractSession) ReleaseEpoch(begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReleaseEpoch(&_Contract.TransactOpts, begin, end)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0x132c290f.
//
// Solidity: function releaseEpoch(uint256 begin, uint256 end) returns()
func (_Contract *ContractTransactorSession) ReleaseEpoch(begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReleaseEpoch(&_Contract.TransactOpts, begin, end)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SubmitBatchAttestation is a paid mutator transaction binding the contract method 0x70dfe736.
//
// Solidity: function submitBatchAttestation(uint256 batchId, uint256 epochId, bytes32 finalizedCidsRootHash) returns()
func (_Contract *ContractTransactor) SubmitBatchAttestation(opts *bind.TransactOpts, batchId *big.Int, epochId *big.Int, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submitBatchAttestation", batchId, epochId, finalizedCidsRootHash)
}

// SubmitBatchAttestation is a paid mutator transaction binding the contract method 0x70dfe736.
//
// Solidity: function submitBatchAttestation(uint256 batchId, uint256 epochId, bytes32 finalizedCidsRootHash) returns()
func (_Contract *ContractSession) SubmitBatchAttestation(batchId *big.Int, epochId *big.Int, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SubmitBatchAttestation(&_Contract.TransactOpts, batchId, epochId, finalizedCidsRootHash)
}

// SubmitBatchAttestation is a paid mutator transaction binding the contract method 0x70dfe736.
//
// Solidity: function submitBatchAttestation(uint256 batchId, uint256 epochId, bytes32 finalizedCidsRootHash) returns()
func (_Contract *ContractTransactorSession) SubmitBatchAttestation(batchId *big.Int, epochId *big.Int, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SubmitBatchAttestation(&_Contract.TransactOpts, batchId, epochId, finalizedCidsRootHash)
}

// SubmitSubmissionBatch is a paid mutator transaction binding the contract method 0xca467e62.
//
// Solidity: function submitSubmissionBatch(string batchCid, uint256 batchId, uint256 epochId, string[] projectIds, string[] snapshotCids) returns()
func (_Contract *ContractTransactor) SubmitSubmissionBatch(opts *bind.TransactOpts, batchCid string, batchId *big.Int, epochId *big.Int, projectIds []string, snapshotCids []string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submitSubmissionBatch", batchCid, batchId, epochId, projectIds, snapshotCids)
}

// SubmitSubmissionBatch is a paid mutator transaction binding the contract method 0xca467e62.
//
// Solidity: function submitSubmissionBatch(string batchCid, uint256 batchId, uint256 epochId, string[] projectIds, string[] snapshotCids) returns()
func (_Contract *ContractSession) SubmitSubmissionBatch(batchCid string, batchId *big.Int, epochId *big.Int, projectIds []string, snapshotCids []string) (*types.Transaction, error) {
	return _Contract.Contract.SubmitSubmissionBatch(&_Contract.TransactOpts, batchCid, batchId, epochId, projectIds, snapshotCids)
}

// SubmitSubmissionBatch is a paid mutator transaction binding the contract method 0xca467e62.
//
// Solidity: function submitSubmissionBatch(string batchCid, uint256 batchId, uint256 epochId, string[] projectIds, string[] snapshotCids) returns()
func (_Contract *ContractTransactorSession) SubmitSubmissionBatch(batchCid string, batchId *big.Int, epochId *big.Int, projectIds []string, snapshotCids []string) (*types.Transaction, error) {
	return _Contract.Contract.SubmitSubmissionBatch(&_Contract.TransactOpts, batchCid, batchId, epochId, projectIds, snapshotCids)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// UpdateAllSnapshotters is a paid mutator transaction binding the contract method 0xa88d5a46.
//
// Solidity: function updateAllSnapshotters(address[] _snapshotters, bool[] _status) returns()
func (_Contract *ContractTransactor) UpdateAllSnapshotters(opts *bind.TransactOpts, _snapshotters []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateAllSnapshotters", _snapshotters, _status)
}

// UpdateAllSnapshotters is a paid mutator transaction binding the contract method 0xa88d5a46.
//
// Solidity: function updateAllSnapshotters(address[] _snapshotters, bool[] _status) returns()
func (_Contract *ContractSession) UpdateAllSnapshotters(_snapshotters []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAllSnapshotters(&_Contract.TransactOpts, _snapshotters, _status)
}

// UpdateAllSnapshotters is a paid mutator transaction binding the contract method 0xa88d5a46.
//
// Solidity: function updateAllSnapshotters(address[] _snapshotters, bool[] _status) returns()
func (_Contract *ContractTransactorSession) UpdateAllSnapshotters(_snapshotters []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAllSnapshotters(&_Contract.TransactOpts, _snapshotters, _status)
}

// UpdateMasterSnapshotters is a paid mutator transaction binding the contract method 0x8f9eacf3.
//
// Solidity: function updateMasterSnapshotters(address[] _snapshotters, bool[] _status) returns()
func (_Contract *ContractTransactor) UpdateMasterSnapshotters(opts *bind.TransactOpts, _snapshotters []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateMasterSnapshotters", _snapshotters, _status)
}

// UpdateMasterSnapshotters is a paid mutator transaction binding the contract method 0x8f9eacf3.
//
// Solidity: function updateMasterSnapshotters(address[] _snapshotters, bool[] _status) returns()
func (_Contract *ContractSession) UpdateMasterSnapshotters(_snapshotters []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMasterSnapshotters(&_Contract.TransactOpts, _snapshotters, _status)
}

// UpdateMasterSnapshotters is a paid mutator transaction binding the contract method 0x8f9eacf3.
//
// Solidity: function updateMasterSnapshotters(address[] _snapshotters, bool[] _status) returns()
func (_Contract *ContractTransactorSession) UpdateMasterSnapshotters(_snapshotters []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMasterSnapshotters(&_Contract.TransactOpts, _snapshotters, _status)
}

// UpdateMinSnapshottersForConsensus is a paid mutator transaction binding the contract method 0x38deacd3.
//
// Solidity: function updateMinSnapshottersForConsensus(uint256 _minSubmissionsForConsensus) returns()
func (_Contract *ContractTransactor) UpdateMinSnapshottersForConsensus(opts *bind.TransactOpts, _minSubmissionsForConsensus *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateMinSnapshottersForConsensus", _minSubmissionsForConsensus)
}

// UpdateMinSnapshottersForConsensus is a paid mutator transaction binding the contract method 0x38deacd3.
//
// Solidity: function updateMinSnapshottersForConsensus(uint256 _minSubmissionsForConsensus) returns()
func (_Contract *ContractSession) UpdateMinSnapshottersForConsensus(_minSubmissionsForConsensus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMinSnapshottersForConsensus(&_Contract.TransactOpts, _minSubmissionsForConsensus)
}

// UpdateMinSnapshottersForConsensus is a paid mutator transaction binding the contract method 0x38deacd3.
//
// Solidity: function updateMinSnapshottersForConsensus(uint256 _minSubmissionsForConsensus) returns()
func (_Contract *ContractTransactorSession) UpdateMinSnapshottersForConsensus(_minSubmissionsForConsensus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMinSnapshottersForConsensus(&_Contract.TransactOpts, _minSubmissionsForConsensus)
}

// UpdateSnapshotSubmissionWindow is a paid mutator transaction binding the contract method 0x9b2f89ce.
//
// Solidity: function updateSnapshotSubmissionWindow(uint256 newsnapshotSubmissionWindow) returns()
func (_Contract *ContractTransactor) UpdateSnapshotSubmissionWindow(opts *bind.TransactOpts, newsnapshotSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateSnapshotSubmissionWindow", newsnapshotSubmissionWindow)
}

// UpdateSnapshotSubmissionWindow is a paid mutator transaction binding the contract method 0x9b2f89ce.
//
// Solidity: function updateSnapshotSubmissionWindow(uint256 newsnapshotSubmissionWindow) returns()
func (_Contract *ContractSession) UpdateSnapshotSubmissionWindow(newsnapshotSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSnapshotSubmissionWindow(&_Contract.TransactOpts, newsnapshotSubmissionWindow)
}

// UpdateSnapshotSubmissionWindow is a paid mutator transaction binding the contract method 0x9b2f89ce.
//
// Solidity: function updateSnapshotSubmissionWindow(uint256 newsnapshotSubmissionWindow) returns()
func (_Contract *ContractTransactorSession) UpdateSnapshotSubmissionWindow(newsnapshotSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSnapshotSubmissionWindow(&_Contract.TransactOpts, newsnapshotSubmissionWindow)
}

// UpdateValidators is a paid mutator transaction binding the contract method 0x1b0b3ae3.
//
// Solidity: function updateValidators(address[] _validators, bool[] _status) returns()
func (_Contract *ContractTransactor) UpdateValidators(opts *bind.TransactOpts, _validators []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateValidators", _validators, _status)
}

// UpdateValidators is a paid mutator transaction binding the contract method 0x1b0b3ae3.
//
// Solidity: function updateValidators(address[] _validators, bool[] _status) returns()
func (_Contract *ContractSession) UpdateValidators(_validators []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidators(&_Contract.TransactOpts, _validators, _status)
}

// UpdateValidators is a paid mutator transaction binding the contract method 0x1b0b3ae3.
//
// Solidity: function updateValidators(address[] _validators, bool[] _status) returns()
func (_Contract *ContractTransactorSession) UpdateValidators(_validators []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidators(&_Contract.TransactOpts, _validators, _status)
}

// ContractDelayedSnapshotSubmittedIterator is returned from FilterDelayedSnapshotSubmitted and is used to iterate over the raw logs and unpacked data for DelayedSnapshotSubmitted events raised by the Contract contract.
type ContractDelayedSnapshotSubmittedIterator struct {
	Event *ContractDelayedSnapshotSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelayedSnapshotSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelayedSnapshotSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelayedSnapshotSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelayedSnapshotSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelayedSnapshotSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelayedSnapshotSubmitted represents a DelayedSnapshotSubmitted event raised by the Contract contract.
type ContractDelayedSnapshotSubmitted struct {
	SnapshotterAddr common.Address
	SnapshotCid     string
	EpochId         *big.Int
	ProjectId       string
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelayedSnapshotSubmitted is a free log retrieval operation binding the contract event 0xa4b1762053ee970f50692b6936d4e58a9a01291449e4da16bdf758891c8de752.
//
// Solidity: event DelayedSnapshotSubmitted(address indexed snapshotterAddr, string snapshotCid, uint256 indexed epochId, string projectId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterDelayedSnapshotSubmitted(opts *bind.FilterOpts, snapshotterAddr []common.Address, epochId []*big.Int) (*ContractDelayedSnapshotSubmittedIterator, error) {

	var snapshotterAddrRule []interface{}
	for _, snapshotterAddrItem := range snapshotterAddr {
		snapshotterAddrRule = append(snapshotterAddrRule, snapshotterAddrItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DelayedSnapshotSubmitted", snapshotterAddrRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelayedSnapshotSubmittedIterator{contract: _Contract.contract, event: "DelayedSnapshotSubmitted", logs: logs, sub: sub}, nil
}

// WatchDelayedSnapshotSubmitted is a free log subscription operation binding the contract event 0xa4b1762053ee970f50692b6936d4e58a9a01291449e4da16bdf758891c8de752.
//
// Solidity: event DelayedSnapshotSubmitted(address indexed snapshotterAddr, string snapshotCid, uint256 indexed epochId, string projectId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchDelayedSnapshotSubmitted(opts *bind.WatchOpts, sink chan<- *ContractDelayedSnapshotSubmitted, snapshotterAddr []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var snapshotterAddrRule []interface{}
	for _, snapshotterAddrItem := range snapshotterAddr {
		snapshotterAddrRule = append(snapshotterAddrRule, snapshotterAddrItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DelayedSnapshotSubmitted", snapshotterAddrRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelayedSnapshotSubmitted)
				if err := _Contract.contract.UnpackLog(event, "DelayedSnapshotSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedSnapshotSubmitted is a log parse operation binding the contract event 0xa4b1762053ee970f50692b6936d4e58a9a01291449e4da16bdf758891c8de752.
//
// Solidity: event DelayedSnapshotSubmitted(address indexed snapshotterAddr, string snapshotCid, uint256 indexed epochId, string projectId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseDelayedSnapshotSubmitted(log types.Log) (*ContractDelayedSnapshotSubmitted, error) {
	event := new(ContractDelayedSnapshotSubmitted)
	if err := _Contract.contract.UnpackLog(event, "DelayedSnapshotSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEpochReleasedIterator is returned from FilterEpochReleased and is used to iterate over the raw logs and unpacked data for EpochReleased events raised by the Contract contract.
type ContractEpochReleasedIterator struct {
	Event *ContractEpochReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEpochReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEpochReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEpochReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEpochReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEpochReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEpochReleased represents a EpochReleased event raised by the Contract contract.
type ContractEpochReleased struct {
	EpochId   *big.Int
	Begin     *big.Int
	End       *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEpochReleased is a free log retrieval operation binding the contract event 0x108f87075a74f81fa2271fdf9fc0883a1811431182601fc65d24513970336640.
//
// Solidity: event EpochReleased(uint256 indexed epochId, uint256 begin, uint256 end, uint256 timestamp)
func (_Contract *ContractFilterer) FilterEpochReleased(opts *bind.FilterOpts, epochId []*big.Int) (*ContractEpochReleasedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EpochReleased", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractEpochReleasedIterator{contract: _Contract.contract, event: "EpochReleased", logs: logs, sub: sub}, nil
}

// WatchEpochReleased is a free log subscription operation binding the contract event 0x108f87075a74f81fa2271fdf9fc0883a1811431182601fc65d24513970336640.
//
// Solidity: event EpochReleased(uint256 indexed epochId, uint256 begin, uint256 end, uint256 timestamp)
func (_Contract *ContractFilterer) WatchEpochReleased(opts *bind.WatchOpts, sink chan<- *ContractEpochReleased, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EpochReleased", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEpochReleased)
				if err := _Contract.contract.UnpackLog(event, "EpochReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEpochReleased is a log parse operation binding the contract event 0x108f87075a74f81fa2271fdf9fc0883a1811431182601fc65d24513970336640.
//
// Solidity: event EpochReleased(uint256 indexed epochId, uint256 begin, uint256 end, uint256 timestamp)
func (_Contract *ContractFilterer) ParseEpochReleased(log types.Log) (*ContractEpochReleased, error) {
	event := new(ContractEpochReleased)
	if err := _Contract.contract.UnpackLog(event, "EpochReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractProjectsUpdatedIterator is returned from FilterProjectsUpdated and is used to iterate over the raw logs and unpacked data for ProjectsUpdated events raised by the Contract contract.
type ContractProjectsUpdatedIterator struct {
	Event *ContractProjectsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractProjectsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProjectsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractProjectsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractProjectsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProjectsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProjectsUpdated represents a ProjectsUpdated event raised by the Contract contract.
type ContractProjectsUpdated struct {
	ProjectId     string
	Allowed       bool
	EnableEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProjectsUpdated is a free log retrieval operation binding the contract event 0x96103183e0788ea5e37ef5b430d1e822a85c862531186baeea043a82dcbcb924.
//
// Solidity: event ProjectsUpdated(string projectId, bool allowed, uint256 enableEpochId)
func (_Contract *ContractFilterer) FilterProjectsUpdated(opts *bind.FilterOpts) (*ContractProjectsUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProjectsUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractProjectsUpdatedIterator{contract: _Contract.contract, event: "ProjectsUpdated", logs: logs, sub: sub}, nil
}

// WatchProjectsUpdated is a free log subscription operation binding the contract event 0x96103183e0788ea5e37ef5b430d1e822a85c862531186baeea043a82dcbcb924.
//
// Solidity: event ProjectsUpdated(string projectId, bool allowed, uint256 enableEpochId)
func (_Contract *ContractFilterer) WatchProjectsUpdated(opts *bind.WatchOpts, sink chan<- *ContractProjectsUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProjectsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProjectsUpdated)
				if err := _Contract.contract.UnpackLog(event, "ProjectsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProjectsUpdated is a log parse operation binding the contract event 0x96103183e0788ea5e37ef5b430d1e822a85c862531186baeea043a82dcbcb924.
//
// Solidity: event ProjectsUpdated(string projectId, bool allowed, uint256 enableEpochId)
func (_Contract *ContractFilterer) ParseProjectsUpdated(log types.Log) (*ContractProjectsUpdated, error) {
	event := new(ContractProjectsUpdated)
	if err := _Contract.contract.UnpackLog(event, "ProjectsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSnapshotBatchFinalizedIterator is returned from FilterSnapshotBatchFinalized and is used to iterate over the raw logs and unpacked data for SnapshotBatchFinalized events raised by the Contract contract.
type ContractSnapshotBatchFinalizedIterator struct {
	Event *ContractSnapshotBatchFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSnapshotBatchFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSnapshotBatchFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSnapshotBatchFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSnapshotBatchFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSnapshotBatchFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSnapshotBatchFinalized represents a SnapshotBatchFinalized event raised by the Contract contract.
type ContractSnapshotBatchFinalized struct {
	EpochId   *big.Int
	BatchId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSnapshotBatchFinalized is a free log retrieval operation binding the contract event 0x5234f54aa94bd8e44799076dbbe59404d669eb94833cfbff03360715b9820e40.
//
// Solidity: event SnapshotBatchFinalized(uint256 indexed epochId, uint256 indexed batchId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterSnapshotBatchFinalized(opts *bind.FilterOpts, epochId []*big.Int, batchId []*big.Int) (*ContractSnapshotBatchFinalizedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SnapshotBatchFinalized", epochIdRule, batchIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSnapshotBatchFinalizedIterator{contract: _Contract.contract, event: "SnapshotBatchFinalized", logs: logs, sub: sub}, nil
}

// WatchSnapshotBatchFinalized is a free log subscription operation binding the contract event 0x5234f54aa94bd8e44799076dbbe59404d669eb94833cfbff03360715b9820e40.
//
// Solidity: event SnapshotBatchFinalized(uint256 indexed epochId, uint256 indexed batchId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchSnapshotBatchFinalized(opts *bind.WatchOpts, sink chan<- *ContractSnapshotBatchFinalized, epochId []*big.Int, batchId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SnapshotBatchFinalized", epochIdRule, batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSnapshotBatchFinalized)
				if err := _Contract.contract.UnpackLog(event, "SnapshotBatchFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotBatchFinalized is a log parse operation binding the contract event 0x5234f54aa94bd8e44799076dbbe59404d669eb94833cfbff03360715b9820e40.
//
// Solidity: event SnapshotBatchFinalized(uint256 indexed epochId, uint256 indexed batchId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseSnapshotBatchFinalized(log types.Log) (*ContractSnapshotBatchFinalized, error) {
	event := new(ContractSnapshotBatchFinalized)
	if err := _Contract.contract.UnpackLog(event, "SnapshotBatchFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSnapshotBatchSubmittedIterator is returned from FilterSnapshotBatchSubmitted and is used to iterate over the raw logs and unpacked data for SnapshotBatchSubmitted events raised by the Contract contract.
type ContractSnapshotBatchSubmittedIterator struct {
	Event *ContractSnapshotBatchSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSnapshotBatchSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSnapshotBatchSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSnapshotBatchSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSnapshotBatchSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSnapshotBatchSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSnapshotBatchSubmitted represents a SnapshotBatchSubmitted event raised by the Contract contract.
type ContractSnapshotBatchSubmitted struct {
	BatchId   *big.Int
	BatchCid  string
	EpochId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSnapshotBatchSubmitted is a free log retrieval operation binding the contract event 0x71f7e37d775c89547d94f23aae240bec29afd96657dc34975abbdd326e8e6d5d.
//
// Solidity: event SnapshotBatchSubmitted(uint256 batchId, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterSnapshotBatchSubmitted(opts *bind.FilterOpts, epochId []*big.Int) (*ContractSnapshotBatchSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SnapshotBatchSubmitted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSnapshotBatchSubmittedIterator{contract: _Contract.contract, event: "SnapshotBatchSubmitted", logs: logs, sub: sub}, nil
}

// WatchSnapshotBatchSubmitted is a free log subscription operation binding the contract event 0x71f7e37d775c89547d94f23aae240bec29afd96657dc34975abbdd326e8e6d5d.
//
// Solidity: event SnapshotBatchSubmitted(uint256 batchId, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchSnapshotBatchSubmitted(opts *bind.WatchOpts, sink chan<- *ContractSnapshotBatchSubmitted, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SnapshotBatchSubmitted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSnapshotBatchSubmitted)
				if err := _Contract.contract.UnpackLog(event, "SnapshotBatchSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotBatchSubmitted is a log parse operation binding the contract event 0x71f7e37d775c89547d94f23aae240bec29afd96657dc34975abbdd326e8e6d5d.
//
// Solidity: event SnapshotBatchSubmitted(uint256 batchId, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseSnapshotBatchSubmitted(log types.Log) (*ContractSnapshotBatchSubmitted, error) {
	event := new(ContractSnapshotBatchSubmitted)
	if err := _Contract.contract.UnpackLog(event, "SnapshotBatchSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValidatorsUpdatedIterator is returned from FilterValidatorsUpdated and is used to iterate over the raw logs and unpacked data for ValidatorsUpdated events raised by the Contract contract.
type ContractValidatorsUpdatedIterator struct {
	Event *ContractValidatorsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorsUpdated represents a ValidatorsUpdated event raised by the Contract contract.
type ContractValidatorsUpdated struct {
	ValidatorAddress common.Address
	Allowed          bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorsUpdated is a free log retrieval operation binding the contract event 0x7f3079c058f3e3dee87048158309898b46e9741ff53b6c7a3afac7c370649afc.
//
// Solidity: event ValidatorsUpdated(address validatorAddress, bool allowed)
func (_Contract *ContractFilterer) FilterValidatorsUpdated(opts *bind.FilterOpts) (*ContractValidatorsUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractValidatorsUpdatedIterator{contract: _Contract.contract, event: "ValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorsUpdated is a free log subscription operation binding the contract event 0x7f3079c058f3e3dee87048158309898b46e9741ff53b6c7a3afac7c370649afc.
//
// Solidity: event ValidatorsUpdated(address validatorAddress, bool allowed)
func (_Contract *ContractFilterer) WatchValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *ContractValidatorsUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorsUpdated)
				if err := _Contract.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorsUpdated is a log parse operation binding the contract event 0x7f3079c058f3e3dee87048158309898b46e9741ff53b6c7a3afac7c370649afc.
//
// Solidity: event ValidatorsUpdated(address validatorAddress, bool allowed)
func (_Contract *ContractFilterer) ParseValidatorsUpdated(log types.Log) (*ContractValidatorsUpdated, error) {
	event := new(ContractValidatorsUpdated)
	if err := _Contract.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllSnapshottersUpdatedIterator is returned from FilterAllSnapshottersUpdated and is used to iterate over the raw logs and unpacked data for AllSnapshottersUpdated events raised by the Contract contract.
type ContractAllSnapshottersUpdatedIterator struct {
	Event *ContractAllSnapshottersUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllSnapshottersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllSnapshottersUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllSnapshottersUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllSnapshottersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllSnapshottersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllSnapshottersUpdated represents a AllSnapshottersUpdated event raised by the Contract contract.
type ContractAllSnapshottersUpdated struct {
	SnapshotterAddress common.Address
	Allowed            bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAllSnapshottersUpdated is a free log retrieval operation binding the contract event 0x743e47fbcd2e3a64a2ab8f5dcdeb4c17f892c17c9f6ab58d3a5c235953d60058.
//
// Solidity: event allSnapshottersUpdated(address snapshotterAddress, bool allowed)
func (_Contract *ContractFilterer) FilterAllSnapshottersUpdated(opts *bind.FilterOpts) (*ContractAllSnapshottersUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "allSnapshottersUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractAllSnapshottersUpdatedIterator{contract: _Contract.contract, event: "allSnapshottersUpdated", logs: logs, sub: sub}, nil
}

// WatchAllSnapshottersUpdated is a free log subscription operation binding the contract event 0x743e47fbcd2e3a64a2ab8f5dcdeb4c17f892c17c9f6ab58d3a5c235953d60058.
//
// Solidity: event allSnapshottersUpdated(address snapshotterAddress, bool allowed)
func (_Contract *ContractFilterer) WatchAllSnapshottersUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllSnapshottersUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "allSnapshottersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllSnapshottersUpdated)
				if err := _Contract.contract.UnpackLog(event, "allSnapshottersUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllSnapshottersUpdated is a log parse operation binding the contract event 0x743e47fbcd2e3a64a2ab8f5dcdeb4c17f892c17c9f6ab58d3a5c235953d60058.
//
// Solidity: event allSnapshottersUpdated(address snapshotterAddress, bool allowed)
func (_Contract *ContractFilterer) ParseAllSnapshottersUpdated(log types.Log) (*ContractAllSnapshottersUpdated, error) {
	event := new(ContractAllSnapshottersUpdated)
	if err := _Contract.contract.UnpackLog(event, "allSnapshottersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMasterSnapshottersUpdatedIterator is returned from FilterMasterSnapshottersUpdated and is used to iterate over the raw logs and unpacked data for MasterSnapshottersUpdated events raised by the Contract contract.
type ContractMasterSnapshottersUpdatedIterator struct {
	Event *ContractMasterSnapshottersUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMasterSnapshottersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMasterSnapshottersUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMasterSnapshottersUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMasterSnapshottersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMasterSnapshottersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMasterSnapshottersUpdated represents a MasterSnapshottersUpdated event raised by the Contract contract.
type ContractMasterSnapshottersUpdated struct {
	SnapshotterAddress common.Address
	Allowed            bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMasterSnapshottersUpdated is a free log retrieval operation binding the contract event 0x9c2d4c2b4cf1ca90e31b1448712dae908d3568785b68a411f6617479c2b9913b.
//
// Solidity: event masterSnapshottersUpdated(address snapshotterAddress, bool allowed)
func (_Contract *ContractFilterer) FilterMasterSnapshottersUpdated(opts *bind.FilterOpts) (*ContractMasterSnapshottersUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "masterSnapshottersUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractMasterSnapshottersUpdatedIterator{contract: _Contract.contract, event: "masterSnapshottersUpdated", logs: logs, sub: sub}, nil
}

// WatchMasterSnapshottersUpdated is a free log subscription operation binding the contract event 0x9c2d4c2b4cf1ca90e31b1448712dae908d3568785b68a411f6617479c2b9913b.
//
// Solidity: event masterSnapshottersUpdated(address snapshotterAddress, bool allowed)
func (_Contract *ContractFilterer) WatchMasterSnapshottersUpdated(opts *bind.WatchOpts, sink chan<- *ContractMasterSnapshottersUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "masterSnapshottersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMasterSnapshottersUpdated)
				if err := _Contract.contract.UnpackLog(event, "masterSnapshottersUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMasterSnapshottersUpdated is a log parse operation binding the contract event 0x9c2d4c2b4cf1ca90e31b1448712dae908d3568785b68a411f6617479c2b9913b.
//
// Solidity: event masterSnapshottersUpdated(address snapshotterAddress, bool allowed)
func (_Contract *ContractFilterer) ParseMasterSnapshottersUpdated(log types.Log) (*ContractMasterSnapshottersUpdated, error) {
	event := new(ContractMasterSnapshottersUpdated)
	if err := _Contract.contract.UnpackLog(event, "masterSnapshottersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPretestProjectsUpdatedIterator is returned from FilterPretestProjectsUpdated and is used to iterate over the raw logs and unpacked data for PretestProjectsUpdated events raised by the Contract contract.
type ContractPretestProjectsUpdatedIterator struct {
	Event *ContractPretestProjectsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPretestProjectsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPretestProjectsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPretestProjectsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPretestProjectsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPretestProjectsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPretestProjectsUpdated represents a PretestProjectsUpdated event raised by the Contract contract.
type ContractPretestProjectsUpdated struct {
	ProjectId     string
	Allowed       bool
	EnableEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPretestProjectsUpdated is a free log retrieval operation binding the contract event 0x9a3ed6f278d4fc2537eb88005da8f892d6a8838b3cda9d15d2d30f5639d8d861.
//
// Solidity: event pretestProjectsUpdated(string projectId, bool allowed, uint256 enableEpochId)
func (_Contract *ContractFilterer) FilterPretestProjectsUpdated(opts *bind.FilterOpts) (*ContractPretestProjectsUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "pretestProjectsUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractPretestProjectsUpdatedIterator{contract: _Contract.contract, event: "pretestProjectsUpdated", logs: logs, sub: sub}, nil
}

// WatchPretestProjectsUpdated is a free log subscription operation binding the contract event 0x9a3ed6f278d4fc2537eb88005da8f892d6a8838b3cda9d15d2d30f5639d8d861.
//
// Solidity: event pretestProjectsUpdated(string projectId, bool allowed, uint256 enableEpochId)
func (_Contract *ContractFilterer) WatchPretestProjectsUpdated(opts *bind.WatchOpts, sink chan<- *ContractPretestProjectsUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "pretestProjectsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPretestProjectsUpdated)
				if err := _Contract.contract.UnpackLog(event, "pretestProjectsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePretestProjectsUpdated is a log parse operation binding the contract event 0x9a3ed6f278d4fc2537eb88005da8f892d6a8838b3cda9d15d2d30f5639d8d861.
//
// Solidity: event pretestProjectsUpdated(string projectId, bool allowed, uint256 enableEpochId)
func (_Contract *ContractFilterer) ParsePretestProjectsUpdated(log types.Log) (*ContractPretestProjectsUpdated, error) {
	event := new(ContractPretestProjectsUpdated)
	if err := _Contract.contract.UnpackLog(event, "pretestProjectsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
