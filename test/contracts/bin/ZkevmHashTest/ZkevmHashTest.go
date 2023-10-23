// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ZkevmHashTest

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

// ZkevmHashTestMetaData contains all meta data concerning the ZkevmHashTest contract.
var ZkevmHashTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"testKeccak256Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"testSha256Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610402806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063897e31ee1461003b578063c1171f231461006b575b600080fd5b610055600480360381019061005091906101ca565b61009b565b6040516100629190610210565b60405180910390f35b610085600480360381019061008091906101ca565b610139565b6040516100929190610210565b60405180910390f35b60008060405180606001604052806040815260200161038d604091399050600080600090505b8481101561012e576002836040516100d991906102a5565b602060405180830381855afa1580156100f6573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061011991906102e8565b9150808061012690610344565b9150506100c1565b508092505050919050565b60008060405180606001604052806040815260200161038d604091399050600080600090505b848110156101845782805190602001209150808061017c90610344565b91505061015f565b508092505050919050565b600080fd5b6000819050919050565b6101a781610194565b81146101b257600080fd5b50565b6000813590506101c48161019e565b92915050565b6000602082840312156101e0576101df61018f565b5b60006101ee848285016101b5565b91505092915050565b6000819050919050565b61020a816101f7565b82525050565b60006020820190506102256000830184610201565b92915050565b600081519050919050565b600081905092915050565b60005b8381101561025f578082015181840152602081019050610244565b8381111561026e576000848401525b50505050565b600061027f8261022b565b6102898185610236565b9350610299818560208601610241565b80840191505092915050565b60006102b18284610274565b915081905092915050565b6102c5816101f7565b81146102d057600080fd5b50565b6000815190506102e2816102bc565b92915050565b6000602082840312156102fe576102fd61018f565b5b600061030c848285016102d3565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061034f82610194565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361038157610380610315565b5b60018201905091905056fe68656c6c6f776f726c6468656c6c6f776f726c6468656c6c6f776f726c6468656c6c6f776f726c6468656c6c6f776f726c6468656c6c6f776f726c6431323334a2646970667358221220d63f1d7e613ebf42eba863bbdae8bc4a35aea6afae801cea16737c694f0066ca64736f6c634300080f0033",
}

// ZkevmHashTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkevmHashTestMetaData.ABI instead.
var ZkevmHashTestABI = ZkevmHashTestMetaData.ABI

// ZkevmHashTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZkevmHashTestMetaData.Bin instead.
var ZkevmHashTestBin = ZkevmHashTestMetaData.Bin

// DeployZkevmHashTest deploys a new Ethereum contract, binding an instance of ZkevmHashTest to it.
func DeployZkevmHashTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZkevmHashTest, error) {
	parsed, err := ZkevmHashTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZkevmHashTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZkevmHashTest{ZkevmHashTestCaller: ZkevmHashTestCaller{contract: contract}, ZkevmHashTestTransactor: ZkevmHashTestTransactor{contract: contract}, ZkevmHashTestFilterer: ZkevmHashTestFilterer{contract: contract}}, nil
}

// ZkevmHashTest is an auto generated Go binding around an Ethereum contract.
type ZkevmHashTest struct {
	ZkevmHashTestCaller     // Read-only binding to the contract
	ZkevmHashTestTransactor // Write-only binding to the contract
	ZkevmHashTestFilterer   // Log filterer for contract events
}

// ZkevmHashTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkevmHashTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkevmHashTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkevmHashTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkevmHashTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkevmHashTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkevmHashTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkevmHashTestSession struct {
	Contract     *ZkevmHashTest    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkevmHashTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkevmHashTestCallerSession struct {
	Contract *ZkevmHashTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ZkevmHashTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkevmHashTestTransactorSession struct {
	Contract     *ZkevmHashTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ZkevmHashTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkevmHashTestRaw struct {
	Contract *ZkevmHashTest // Generic contract binding to access the raw methods on
}

// ZkevmHashTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkevmHashTestCallerRaw struct {
	Contract *ZkevmHashTestCaller // Generic read-only contract binding to access the raw methods on
}

// ZkevmHashTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkevmHashTestTransactorRaw struct {
	Contract *ZkevmHashTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkevmHashTest creates a new instance of ZkevmHashTest, bound to a specific deployed contract.
func NewZkevmHashTest(address common.Address, backend bind.ContractBackend) (*ZkevmHashTest, error) {
	contract, err := bindZkevmHashTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkevmHashTest{ZkevmHashTestCaller: ZkevmHashTestCaller{contract: contract}, ZkevmHashTestTransactor: ZkevmHashTestTransactor{contract: contract}, ZkevmHashTestFilterer: ZkevmHashTestFilterer{contract: contract}}, nil
}

// NewZkevmHashTestCaller creates a new read-only instance of ZkevmHashTest, bound to a specific deployed contract.
func NewZkevmHashTestCaller(address common.Address, caller bind.ContractCaller) (*ZkevmHashTestCaller, error) {
	contract, err := bindZkevmHashTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkevmHashTestCaller{contract: contract}, nil
}

// NewZkevmHashTestTransactor creates a new write-only instance of ZkevmHashTest, bound to a specific deployed contract.
func NewZkevmHashTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkevmHashTestTransactor, error) {
	contract, err := bindZkevmHashTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkevmHashTestTransactor{contract: contract}, nil
}

// NewZkevmHashTestFilterer creates a new log filterer instance of ZkevmHashTest, bound to a specific deployed contract.
func NewZkevmHashTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkevmHashTestFilterer, error) {
	contract, err := bindZkevmHashTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkevmHashTestFilterer{contract: contract}, nil
}

// bindZkevmHashTest binds a generic wrapper to an already deployed contract.
func bindZkevmHashTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZkevmHashTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkevmHashTest *ZkevmHashTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkevmHashTest.Contract.ZkevmHashTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkevmHashTest *ZkevmHashTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkevmHashTest.Contract.ZkevmHashTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkevmHashTest *ZkevmHashTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkevmHashTest.Contract.ZkevmHashTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkevmHashTest *ZkevmHashTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkevmHashTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkevmHashTest *ZkevmHashTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkevmHashTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkevmHashTest *ZkevmHashTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkevmHashTest.Contract.contract.Transact(opts, method, params...)
}

// TestKeccak256Hash is a paid mutator transaction binding the contract method 0xc1171f23.
//
// Solidity: function testKeccak256Hash(uint256 count) returns(bytes32)
func (_ZkevmHashTest *ZkevmHashTestTransactor) TestKeccak256Hash(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _ZkevmHashTest.contract.Transact(opts, "testKeccak256Hash", count)
}

// TestKeccak256Hash is a paid mutator transaction binding the contract method 0xc1171f23.
//
// Solidity: function testKeccak256Hash(uint256 count) returns(bytes32)
func (_ZkevmHashTest *ZkevmHashTestSession) TestKeccak256Hash(count *big.Int) (*types.Transaction, error) {
	return _ZkevmHashTest.Contract.TestKeccak256Hash(&_ZkevmHashTest.TransactOpts, count)
}

// TestKeccak256Hash is a paid mutator transaction binding the contract method 0xc1171f23.
//
// Solidity: function testKeccak256Hash(uint256 count) returns(bytes32)
func (_ZkevmHashTest *ZkevmHashTestTransactorSession) TestKeccak256Hash(count *big.Int) (*types.Transaction, error) {
	return _ZkevmHashTest.Contract.TestKeccak256Hash(&_ZkevmHashTest.TransactOpts, count)
}

// TestSha256Hash is a paid mutator transaction binding the contract method 0x897e31ee.
//
// Solidity: function testSha256Hash(uint256 count) returns(bytes32)
func (_ZkevmHashTest *ZkevmHashTestTransactor) TestSha256Hash(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _ZkevmHashTest.contract.Transact(opts, "testSha256Hash", count)
}

// TestSha256Hash is a paid mutator transaction binding the contract method 0x897e31ee.
//
// Solidity: function testSha256Hash(uint256 count) returns(bytes32)
func (_ZkevmHashTest *ZkevmHashTestSession) TestSha256Hash(count *big.Int) (*types.Transaction, error) {
	return _ZkevmHashTest.Contract.TestSha256Hash(&_ZkevmHashTest.TransactOpts, count)
}

// TestSha256Hash is a paid mutator transaction binding the contract method 0x897e31ee.
//
// Solidity: function testSha256Hash(uint256 count) returns(bytes32)
func (_ZkevmHashTest *ZkevmHashTestTransactorSession) TestSha256Hash(count *big.Int) (*types.Transaction, error) {
	return _ZkevmHashTest.Contract.TestSha256Hash(&_ZkevmHashTest.TransactOpts, count)
}
