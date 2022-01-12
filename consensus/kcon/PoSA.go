// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kcon

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

// PoSAMetaData contains all meta data concerning the PoSA contract.
var PoSAMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MINIMUM_STAKE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNER_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"initialSigners\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signerList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newSigners\",\"type\":\"address[]\"}],\"name\":\"updateSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PoSAABI is the input ABI used to generate the binding from.
// Deprecated: Use PoSAMetaData.ABI instead.
var PoSAABI = PoSAMetaData.ABI

// PoSA is an auto generated Go binding around an Ethereum contract.
type PoSA struct {
	PoSACaller     // Read-only binding to the contract
	PoSATransactor // Write-only binding to the contract
	PoSAFilterer   // Log filterer for contract events
}

// PoSACaller is an auto generated read-only Go binding around an Ethereum contract.
type PoSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoSASession struct {
	Contract     *PoSA             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoSACallerSession struct {
	Contract *PoSACaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoSATransactorSession struct {
	Contract     *PoSATransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoSARaw is an auto generated low-level Go binding around an Ethereum contract.
type PoSARaw struct {
	Contract *PoSA // Generic contract binding to access the raw methods on
}

// PoSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoSACallerRaw struct {
	Contract *PoSACaller // Generic read-only contract binding to access the raw methods on
}

// PoSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoSATransactorRaw struct {
	Contract *PoSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoSA creates a new instance of PoSA, bound to a specific deployed contract.
func NewPoSA(address common.Address, backend bind.ContractBackend) (*PoSA, error) {
	contract, err := bindPoSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoSA{PoSACaller: PoSACaller{contract: contract}, PoSATransactor: PoSATransactor{contract: contract}, PoSAFilterer: PoSAFilterer{contract: contract}}, nil
}

// NewPoSACaller creates a new read-only instance of PoSA, bound to a specific deployed contract.
func NewPoSACaller(address common.Address, caller bind.ContractCaller) (*PoSACaller, error) {
	contract, err := bindPoSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoSACaller{contract: contract}, nil
}

// NewPoSATransactor creates a new write-only instance of PoSA, bound to a specific deployed contract.
func NewPoSATransactor(address common.Address, transactor bind.ContractTransactor) (*PoSATransactor, error) {
	contract, err := bindPoSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoSATransactor{contract: contract}, nil
}

// NewPoSAFilterer creates a new log filterer instance of PoSA, bound to a specific deployed contract.
func NewPoSAFilterer(address common.Address, filterer bind.ContractFilterer) (*PoSAFilterer, error) {
	contract, err := bindPoSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoSAFilterer{contract: contract}, nil
}

// bindPoSA binds a generic wrapper to an already deployed contract.
func bindPoSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoSA *PoSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoSA.Contract.PoSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoSA *PoSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoSA.Contract.PoSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoSA *PoSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoSA.Contract.PoSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoSA *PoSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoSA *PoSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoSA *PoSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoSA.Contract.contract.Transact(opts, method, params...)
}

// MINIMUMSTAKE is a free data retrieval call binding the contract method 0x08dbbb03.
//
// Solidity: function MINIMUM_STAKE() view returns(uint256)
func (_PoSA *PoSACaller) MINIMUMSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoSA.contract.Call(opts, &out, "MINIMUM_STAKE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMSTAKE is a free data retrieval call binding the contract method 0x08dbbb03.
//
// Solidity: function MINIMUM_STAKE() view returns(uint256)
func (_PoSA *PoSASession) MINIMUMSTAKE() (*big.Int, error) {
	return _PoSA.Contract.MINIMUMSTAKE(&_PoSA.CallOpts)
}

// MINIMUMSTAKE is a free data retrieval call binding the contract method 0x08dbbb03.
//
// Solidity: function MINIMUM_STAKE() view returns(uint256)
func (_PoSA *PoSACallerSession) MINIMUMSTAKE() (*big.Int, error) {
	return _PoSA.Contract.MINIMUMSTAKE(&_PoSA.CallOpts)
}

// SIGNERAMOUNT is a free data retrieval call binding the contract method 0x09c4fe38.
//
// Solidity: function SIGNER_AMOUNT() view returns(uint256)
func (_PoSA *PoSACaller) SIGNERAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoSA.contract.Call(opts, &out, "SIGNER_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGNERAMOUNT is a free data retrieval call binding the contract method 0x09c4fe38.
//
// Solidity: function SIGNER_AMOUNT() view returns(uint256)
func (_PoSA *PoSASession) SIGNERAMOUNT() (*big.Int, error) {
	return _PoSA.Contract.SIGNERAMOUNT(&_PoSA.CallOpts)
}

// SIGNERAMOUNT is a free data retrieval call binding the contract method 0x09c4fe38.
//
// Solidity: function SIGNER_AMOUNT() view returns(uint256)
func (_PoSA *PoSACallerSession) SIGNERAMOUNT() (*big.Int, error) {
	return _PoSA.Contract.SIGNERAMOUNT(&_PoSA.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_PoSA *PoSACaller) GetSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PoSA.contract.Call(opts, &out, "getSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_PoSA *PoSASession) GetSigners() ([]common.Address, error) {
	return _PoSA.Contract.GetSigners(&_PoSA.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_PoSA *PoSACallerSession) GetSigners() ([]common.Address, error) {
	return _PoSA.Contract.GetSigners(&_PoSA.CallOpts)
}

// SignerList is a free data retrieval call binding the contract method 0x577ed7e1.
//
// Solidity: function signerList(uint256 ) view returns(address)
func (_PoSA *PoSACaller) SignerList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PoSA.contract.Call(opts, &out, "signerList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerList is a free data retrieval call binding the contract method 0x577ed7e1.
//
// Solidity: function signerList(uint256 ) view returns(address)
func (_PoSA *PoSASession) SignerList(arg0 *big.Int) (common.Address, error) {
	return _PoSA.Contract.SignerList(&_PoSA.CallOpts, arg0)
}

// SignerList is a free data retrieval call binding the contract method 0x577ed7e1.
//
// Solidity: function signerList(uint256 ) view returns(address)
func (_PoSA *PoSACallerSession) SignerList(arg0 *big.Int) (common.Address, error) {
	return _PoSA.Contract.SignerList(&_PoSA.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] initialSigners) returns()
func (_PoSA *PoSATransactor) Initialize(opts *bind.TransactOpts, initialSigners []common.Address) (*types.Transaction, error) {
	return _PoSA.contract.Transact(opts, "initialize", initialSigners)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] initialSigners) returns()
func (_PoSA *PoSASession) Initialize(initialSigners []common.Address) (*types.Transaction, error) {
	return _PoSA.Contract.Initialize(&_PoSA.TransactOpts, initialSigners)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] initialSigners) returns()
func (_PoSA *PoSATransactorSession) Initialize(initialSigners []common.Address) (*types.Transaction, error) {
	return _PoSA.Contract.Initialize(&_PoSA.TransactOpts, initialSigners)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0x96751ae9.
//
// Solidity: function updateSigners(address[] newSigners) returns()
func (_PoSA *PoSATransactor) UpdateSigners(opts *bind.TransactOpts, newSigners []common.Address) (*types.Transaction, error) {
	return _PoSA.contract.Transact(opts, "updateSigners", newSigners)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0x96751ae9.
//
// Solidity: function updateSigners(address[] newSigners) returns()
func (_PoSA *PoSASession) UpdateSigners(newSigners []common.Address) (*types.Transaction, error) {
	return _PoSA.Contract.UpdateSigners(&_PoSA.TransactOpts, newSigners)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0x96751ae9.
//
// Solidity: function updateSigners(address[] newSigners) returns()
func (_PoSA *PoSATransactorSession) UpdateSigners(newSigners []common.Address) (*types.Transaction, error) {
	return _PoSA.Contract.UpdateSigners(&_PoSA.TransactOpts, newSigners)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PoSA *PoSATransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoSA.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PoSA *PoSASession) Receive() (*types.Transaction, error) {
	return _PoSA.Contract.Receive(&_PoSA.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PoSA *PoSATransactorSession) Receive() (*types.Transaction, error) {
	return _PoSA.Contract.Receive(&_PoSA.TransactOpts)
}
