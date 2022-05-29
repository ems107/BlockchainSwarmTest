// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
)

// RobotSwarmMetaData contains all meta data concerning the RobotSwarm contract.
var RobotSwarmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d36dedd2": "getOrder()",
		"775a25e3": "getTotal()",
		"7cf5dab0": "increment(uint256)",
		"2ddbd13a": "total()",
	},
	Bin: "0x608060405234801561001057600080fd5b506101fb806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80632ddbd13a14610051578063775a25e31461006d5780637cf5dab014610075578063d36dedd21461008a575b600080fd5b61005a60005481565b6040519081526020015b60405180910390f35b60005461005a565b61008861008336600461010f565b61009f565b005b6100926100b8565b6040516100649190610128565b806000808282546100b0919061017d565b909155505050565b606060026000546100c991906101a3565b6000036100f05750604080518082019091526005815264149a59da1d60da1b602082015290565b506040805180820190915260048152631319599d60e21b602082015290565b60006020828403121561012157600080fd5b5035919050565b600060208083528351808285015260005b8181101561015557858101830151858201604001528201610139565b81811115610167576000604083870101525b50601f01601f1916929092016040019392505050565b6000821982111561019e57634e487b7160e01b600052601160045260246000fd5b500190565b6000826101c057634e487b7160e01b600052601260045260246000fd5b50069056fea2646970667358221220c4b21286a3c7b1903d5d440c144226e2eeac3aafd0b0428be97723803f73aa2964736f6c634300080e0033",
}

// RobotSwarmABI is the input ABI used to generate the binding from.
// Deprecated: Use RobotSwarmMetaData.ABI instead.
var RobotSwarmABI = RobotSwarmMetaData.ABI

// Deprecated: Use RobotSwarmMetaData.Sigs instead.
// RobotSwarmFuncSigs maps the 4-byte function signature to its string representation.
var RobotSwarmFuncSigs = RobotSwarmMetaData.Sigs

// RobotSwarmBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RobotSwarmMetaData.Bin instead.
var RobotSwarmBin = RobotSwarmMetaData.Bin

// DeployRobotSwarm deploys a new Ethereum contract, binding an instance of RobotSwarm to it.
func DeployRobotSwarm(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RobotSwarm, error) {
	parsed, err := RobotSwarmMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RobotSwarmBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RobotSwarm{RobotSwarmCaller: RobotSwarmCaller{contract: contract}, RobotSwarmTransactor: RobotSwarmTransactor{contract: contract}, RobotSwarmFilterer: RobotSwarmFilterer{contract: contract}}, nil
}

// RobotSwarm is an auto generated Go binding around an Ethereum contract.
type RobotSwarm struct {
	RobotSwarmCaller     // Read-only binding to the contract
	RobotSwarmTransactor // Write-only binding to the contract
	RobotSwarmFilterer   // Log filterer for contract events
}

// RobotSwarmCaller is an auto generated read-only Go binding around an Ethereum contract.
type RobotSwarmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RobotSwarmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RobotSwarmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RobotSwarmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RobotSwarmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RobotSwarmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RobotSwarmSession struct {
	Contract     *RobotSwarm       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RobotSwarmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RobotSwarmCallerSession struct {
	Contract *RobotSwarmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RobotSwarmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RobotSwarmTransactorSession struct {
	Contract     *RobotSwarmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RobotSwarmRaw is an auto generated low-level Go binding around an Ethereum contract.
type RobotSwarmRaw struct {
	Contract *RobotSwarm // Generic contract binding to access the raw methods on
}

// RobotSwarmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RobotSwarmCallerRaw struct {
	Contract *RobotSwarmCaller // Generic read-only contract binding to access the raw methods on
}

// RobotSwarmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RobotSwarmTransactorRaw struct {
	Contract *RobotSwarmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRobotSwarm creates a new instance of RobotSwarm, bound to a specific deployed contract.
func NewRobotSwarm(address common.Address, backend bind.ContractBackend) (*RobotSwarm, error) {
	contract, err := bindRobotSwarm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RobotSwarm{RobotSwarmCaller: RobotSwarmCaller{contract: contract}, RobotSwarmTransactor: RobotSwarmTransactor{contract: contract}, RobotSwarmFilterer: RobotSwarmFilterer{contract: contract}}, nil
}

// NewRobotSwarmCaller creates a new read-only instance of RobotSwarm, bound to a specific deployed contract.
func NewRobotSwarmCaller(address common.Address, caller bind.ContractCaller) (*RobotSwarmCaller, error) {
	contract, err := bindRobotSwarm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RobotSwarmCaller{contract: contract}, nil
}

// NewRobotSwarmTransactor creates a new write-only instance of RobotSwarm, bound to a specific deployed contract.
func NewRobotSwarmTransactor(address common.Address, transactor bind.ContractTransactor) (*RobotSwarmTransactor, error) {
	contract, err := bindRobotSwarm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RobotSwarmTransactor{contract: contract}, nil
}

// NewRobotSwarmFilterer creates a new log filterer instance of RobotSwarm, bound to a specific deployed contract.
func NewRobotSwarmFilterer(address common.Address, filterer bind.ContractFilterer) (*RobotSwarmFilterer, error) {
	contract, err := bindRobotSwarm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RobotSwarmFilterer{contract: contract}, nil
}

// bindRobotSwarm binds a generic wrapper to an already deployed contract.
func bindRobotSwarm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RobotSwarmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RobotSwarm *RobotSwarmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RobotSwarm.Contract.RobotSwarmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RobotSwarm *RobotSwarmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RobotSwarm.Contract.RobotSwarmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RobotSwarm *RobotSwarmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RobotSwarm.Contract.RobotSwarmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RobotSwarm *RobotSwarmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RobotSwarm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RobotSwarm *RobotSwarmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RobotSwarm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RobotSwarm *RobotSwarmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RobotSwarm.Contract.contract.Transact(opts, method, params...)
}

// GetOrder is a free data retrieval call binding the contract method 0xd36dedd2.
//
// Solidity: function getOrder() view returns(string)
func (_RobotSwarm *RobotSwarmCaller) GetOrder(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RobotSwarm.contract.Call(opts, &out, "getOrder")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0xd36dedd2.
//
// Solidity: function getOrder() view returns(string)
func (_RobotSwarm *RobotSwarmSession) GetOrder() (string, error) {
	return _RobotSwarm.Contract.GetOrder(&_RobotSwarm.CallOpts)
}

// GetOrder is a free data retrieval call binding the contract method 0xd36dedd2.
//
// Solidity: function getOrder() view returns(string)
func (_RobotSwarm *RobotSwarmCallerSession) GetOrder() (string, error) {
	return _RobotSwarm.Contract.GetOrder(&_RobotSwarm.CallOpts)
}

// GetTotal is a free data retrieval call binding the contract method 0x775a25e3.
//
// Solidity: function getTotal() view returns(uint256)
func (_RobotSwarm *RobotSwarmCaller) GetTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotSwarm.contract.Call(opts, &out, "getTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotal is a free data retrieval call binding the contract method 0x775a25e3.
//
// Solidity: function getTotal() view returns(uint256)
func (_RobotSwarm *RobotSwarmSession) GetTotal() (*big.Int, error) {
	return _RobotSwarm.Contract.GetTotal(&_RobotSwarm.CallOpts)
}

// GetTotal is a free data retrieval call binding the contract method 0x775a25e3.
//
// Solidity: function getTotal() view returns(uint256)
func (_RobotSwarm *RobotSwarmCallerSession) GetTotal() (*big.Int, error) {
	return _RobotSwarm.Contract.GetTotal(&_RobotSwarm.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_RobotSwarm *RobotSwarmCaller) Total(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotSwarm.contract.Call(opts, &out, "total")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_RobotSwarm *RobotSwarmSession) Total() (*big.Int, error) {
	return _RobotSwarm.Contract.Total(&_RobotSwarm.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_RobotSwarm *RobotSwarmCallerSession) Total() (*big.Int, error) {
	return _RobotSwarm.Contract.Total(&_RobotSwarm.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0x7cf5dab0.
//
// Solidity: function increment(uint256 number) returns()
func (_RobotSwarm *RobotSwarmTransactor) Increment(opts *bind.TransactOpts, number *big.Int) (*types.Transaction, error) {
	return _RobotSwarm.contract.Transact(opts, "increment", number)
}

// Increment is a paid mutator transaction binding the contract method 0x7cf5dab0.
//
// Solidity: function increment(uint256 number) returns()
func (_RobotSwarm *RobotSwarmSession) Increment(number *big.Int) (*types.Transaction, error) {
	return _RobotSwarm.Contract.Increment(&_RobotSwarm.TransactOpts, number)
}

// Increment is a paid mutator transaction binding the contract method 0x7cf5dab0.
//
// Solidity: function increment(uint256 number) returns()
func (_RobotSwarm *RobotSwarmTransactorSession) Increment(number *big.Int) (*types.Transaction, error) {
	return _RobotSwarm.Contract.Increment(&_RobotSwarm.TransactOpts, number)
}
