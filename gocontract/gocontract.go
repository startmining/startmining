// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gocontract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// OraclizeAddrResolverIABI is the input ABI used to generate the binding from.
const OraclizeAddrResolverIABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OraclizeAddrResolverIBin is the compiled bytecode used for deploying new contracts.
const OraclizeAddrResolverIBin = `0x`

// DeployOraclizeAddrResolverI deploys a new Ethereum contract, binding an instance of OraclizeAddrResolverI to it.
func DeployOraclizeAddrResolverI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OraclizeAddrResolverI, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeAddrResolverIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OraclizeAddrResolverIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OraclizeAddrResolverI{OraclizeAddrResolverICaller: OraclizeAddrResolverICaller{contract: contract}, OraclizeAddrResolverITransactor: OraclizeAddrResolverITransactor{contract: contract}}, nil
}

// OraclizeAddrResolverI is an auto generated Go binding around an Ethereum contract.
type OraclizeAddrResolverI struct {
	OraclizeAddrResolverICaller     // Read-only binding to the contract
	OraclizeAddrResolverITransactor // Write-only binding to the contract
}

// OraclizeAddrResolverICaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclizeAddrResolverICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverITransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclizeAddrResolverITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclizeAddrResolverISession struct {
	Contract     *OraclizeAddrResolverI // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OraclizeAddrResolverICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclizeAddrResolverICallerSession struct {
	Contract *OraclizeAddrResolverICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// OraclizeAddrResolverITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclizeAddrResolverITransactorSession struct {
	Contract     *OraclizeAddrResolverITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// OraclizeAddrResolverIRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclizeAddrResolverIRaw struct {
	Contract *OraclizeAddrResolverI // Generic contract binding to access the raw methods on
}

// OraclizeAddrResolverICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclizeAddrResolverICallerRaw struct {
	Contract *OraclizeAddrResolverICaller // Generic read-only contract binding to access the raw methods on
}

// OraclizeAddrResolverITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclizeAddrResolverITransactorRaw struct {
	Contract *OraclizeAddrResolverITransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclizeAddrResolverI creates a new instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverI(address common.Address, backend bind.ContractBackend) (*OraclizeAddrResolverI, error) {
	contract, err := bindOraclizeAddrResolverI(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverI{OraclizeAddrResolverICaller: OraclizeAddrResolverICaller{contract: contract}, OraclizeAddrResolverITransactor: OraclizeAddrResolverITransactor{contract: contract}}, nil
}

// NewOraclizeAddrResolverICaller creates a new read-only instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverICaller(address common.Address, caller bind.ContractCaller) (*OraclizeAddrResolverICaller, error) {
	contract, err := bindOraclizeAddrResolverI(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverICaller{contract: contract}, nil
}

// NewOraclizeAddrResolverITransactor creates a new write-only instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverITransactor(address common.Address, transactor bind.ContractTransactor) (*OraclizeAddrResolverITransactor, error) {
	contract, err := bindOraclizeAddrResolverI(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverITransactor{contract: contract}, nil
}

// bindOraclizeAddrResolverI binds a generic wrapper to an already deployed contract.
func bindOraclizeAddrResolverI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeAddrResolverIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeAddrResolverI *OraclizeAddrResolverIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeAddrResolverI.Contract.OraclizeAddrResolverICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeAddrResolverI *OraclizeAddrResolverIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.OraclizeAddrResolverITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeAddrResolverI *OraclizeAddrResolverIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.OraclizeAddrResolverITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeAddrResolverI *OraclizeAddrResolverICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeAddrResolverI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_addr address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactor) GetAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.contract.Transact(opts, "getAddress")
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_addr address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverISession) GetAddress() (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.GetAddress(&_OraclizeAddrResolverI.TransactOpts)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_addr address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactorSession) GetAddress() (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.GetAddress(&_OraclizeAddrResolverI.TransactOpts)
}

// OraclizeIABI is the input ABI used to generate the binding from.
const OraclizeIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"_dsprice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_datasource\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"_dsprice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_coupon\",\"type\":\"string\"}],\"name\":\"useCoupon\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proofType\",\"type\":\"bytes1\"}],\"name\":\"setProofType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg1\",\"type\":\"string\"},{\"name\":\"_arg2\",\"type\":\"string\"}],\"name\":\"query2\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_argN\",\"type\":\"bytes\"}],\"name\":\"queryN\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg1\",\"type\":\"string\"},{\"name\":\"_arg2\",\"type\":\"string\"},{\"name\":\"_gaslimit\",\"type\":\"uint256\"}],\"name\":\"query2_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"randomDS_getSessionPubKeyHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cbAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg\",\"type\":\"string\"},{\"name\":\"_gaslimit\",\"type\":\"uint256\"}],\"name\":\"query_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_argN\",\"type\":\"bytes\"},{\"name\":\"_gaslimit\",\"type\":\"uint256\"}],\"name\":\"queryN_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"setCustomGasPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_config\",\"type\":\"bytes32\"}],\"name\":\"setConfig\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OraclizeIBin is the compiled bytecode used for deploying new contracts.
const OraclizeIBin = `0x`

// DeployOraclizeI deploys a new Ethereum contract, binding an instance of OraclizeI to it.
func DeployOraclizeI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OraclizeI, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OraclizeIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OraclizeI{OraclizeICaller: OraclizeICaller{contract: contract}, OraclizeITransactor: OraclizeITransactor{contract: contract}}, nil
}

// OraclizeI is an auto generated Go binding around an Ethereum contract.
type OraclizeI struct {
	OraclizeICaller     // Read-only binding to the contract
	OraclizeITransactor // Write-only binding to the contract
}

// OraclizeICaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclizeICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeITransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclizeITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclizeISession struct {
	Contract     *OraclizeI        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OraclizeICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclizeICallerSession struct {
	Contract *OraclizeICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OraclizeITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclizeITransactorSession struct {
	Contract     *OraclizeITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OraclizeIRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclizeIRaw struct {
	Contract *OraclizeI // Generic contract binding to access the raw methods on
}

// OraclizeICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclizeICallerRaw struct {
	Contract *OraclizeICaller // Generic read-only contract binding to access the raw methods on
}

// OraclizeITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclizeITransactorRaw struct {
	Contract *OraclizeITransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclizeI creates a new instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeI(address common.Address, backend bind.ContractBackend) (*OraclizeI, error) {
	contract, err := bindOraclizeI(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OraclizeI{OraclizeICaller: OraclizeICaller{contract: contract}, OraclizeITransactor: OraclizeITransactor{contract: contract}}, nil
}

// NewOraclizeICaller creates a new read-only instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeICaller(address common.Address, caller bind.ContractCaller) (*OraclizeICaller, error) {
	contract, err := bindOraclizeI(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeICaller{contract: contract}, nil
}

// NewOraclizeITransactor creates a new write-only instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeITransactor(address common.Address, transactor bind.ContractTransactor) (*OraclizeITransactor, error) {
	contract, err := bindOraclizeI(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OraclizeITransactor{contract: contract}, nil
}

// bindOraclizeI binds a generic wrapper to an already deployed contract.
func bindOraclizeI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeI *OraclizeIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeI.Contract.OraclizeICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeI *OraclizeIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeI.Contract.OraclizeITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeI *OraclizeIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeI.Contract.OraclizeITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeI *OraclizeICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeI *OraclizeITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeI *OraclizeITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeI.Contract.contract.Transact(opts, method, params...)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeI *OraclizeICaller) CbAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OraclizeI.contract.Call(opts, out, "cbAddress")
	return *ret0, err
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeI *OraclizeISession) CbAddress() (common.Address, error) {
	return _OraclizeI.Contract.CbAddress(&_OraclizeI.CallOpts)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeI *OraclizeICallerSession) CbAddress() (common.Address, error) {
	return _OraclizeI.Contract.CbAddress(&_OraclizeI.CallOpts)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(_datasource string) returns(_dsprice uint256)
func (_OraclizeI *OraclizeITransactor) GetPrice(opts *bind.TransactOpts, _datasource string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "getPrice", _datasource)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(_datasource string) returns(_dsprice uint256)
func (_OraclizeI *OraclizeISession) GetPrice(_datasource string) (*types.Transaction, error) {
	return _OraclizeI.Contract.GetPrice(&_OraclizeI.TransactOpts, _datasource)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(_datasource string) returns(_dsprice uint256)
func (_OraclizeI *OraclizeITransactorSession) GetPrice(_datasource string) (*types.Transaction, error) {
	return _OraclizeI.Contract.GetPrice(&_OraclizeI.TransactOpts, _datasource)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(_timestamp uint256, _datasource string, _arg string) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) Query(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query", _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(_timestamp uint256, _datasource string, _arg string) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(_timestamp uint256, _datasource string, _arg string) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(_timestamp uint256, _datasource string, _arg1 string, _arg2 string) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) Query2(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query2", _timestamp, _datasource, _arg1, _arg2)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(_timestamp uint256, _datasource string, _arg1 string, _arg2 string) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) Query2(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(_timestamp uint256, _datasource string, _arg1 string, _arg2 string) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) Query2(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2)
}

// Query2_withGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(_timestamp uint256, _datasource string, _arg1 string, _arg2 string, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) Query2_withGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query2_withGasLimit", _timestamp, _datasource, _arg1, _arg2, _gaslimit)
}

// Query2_withGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(_timestamp uint256, _datasource string, _arg1 string, _arg2 string, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) Query2_withGasLimit(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2_withGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2, _gaslimit)
}

// Query2_withGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(_timestamp uint256, _datasource string, _arg1 string, _arg2 string, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) Query2_withGasLimit(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2_withGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2, _gaslimit)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(_timestamp uint256, _datasource string, _argN bytes) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) QueryN(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "queryN", _timestamp, _datasource, _argN)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(_timestamp uint256, _datasource string, _argN bytes) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) QueryN(_timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryN(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(_timestamp uint256, _datasource string, _argN bytes) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) QueryN(_timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryN(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN)
}

// QueryN_withGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(_timestamp uint256, _datasource string, _argN bytes, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) QueryN_withGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _argN []byte, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "queryN_withGasLimit", _timestamp, _datasource, _argN, _gaslimit)
}

// QueryN_withGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(_timestamp uint256, _datasource string, _argN bytes, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) QueryN_withGasLimit(_timestamp *big.Int, _datasource string, _argN []byte, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryN_withGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN, _gaslimit)
}

// QueryN_withGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(_timestamp uint256, _datasource string, _argN bytes, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) QueryN_withGasLimit(_timestamp *big.Int, _datasource string, _argN []byte, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryN_withGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN, _gaslimit)
}

// Query_withGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(_timestamp uint256, _datasource string, _arg string, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) Query_withGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query_withGasLimit", _timestamp, _datasource, _arg, _gaslimit)
}

// Query_withGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(_timestamp uint256, _datasource string, _arg string, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) Query_withGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query_withGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg, _gaslimit)
}

// Query_withGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(_timestamp uint256, _datasource string, _arg string, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) Query_withGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query_withGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg, _gaslimit)
}

// RandomDS_getSessionPubKeyHash is a paid mutator transaction binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() returns(bytes32)
func (_OraclizeI *OraclizeITransactor) RandomDS_getSessionPubKeyHash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "randomDS_getSessionPubKeyHash")
}

// RandomDS_getSessionPubKeyHash is a paid mutator transaction binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() returns(bytes32)
func (_OraclizeI *OraclizeISession) RandomDS_getSessionPubKeyHash() (*types.Transaction, error) {
	return _OraclizeI.Contract.RandomDS_getSessionPubKeyHash(&_OraclizeI.TransactOpts)
}

// RandomDS_getSessionPubKeyHash is a paid mutator transaction binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() returns(bytes32)
func (_OraclizeI *OraclizeITransactorSession) RandomDS_getSessionPubKeyHash() (*types.Transaction, error) {
	return _OraclizeI.Contract.RandomDS_getSessionPubKeyHash(&_OraclizeI.TransactOpts)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe8a5282d.
//
// Solidity: function setConfig(_config bytes32) returns()
func (_OraclizeI *OraclizeITransactor) SetConfig(opts *bind.TransactOpts, _config [32]byte) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "setConfig", _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe8a5282d.
//
// Solidity: function setConfig(_config bytes32) returns()
func (_OraclizeI *OraclizeISession) SetConfig(_config [32]byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetConfig(&_OraclizeI.TransactOpts, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe8a5282d.
//
// Solidity: function setConfig(_config bytes32) returns()
func (_OraclizeI *OraclizeITransactorSession) SetConfig(_config [32]byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetConfig(&_OraclizeI.TransactOpts, _config)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(_gasPrice uint256) returns()
func (_OraclizeI *OraclizeITransactor) SetCustomGasPrice(opts *bind.TransactOpts, _gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "setCustomGasPrice", _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(_gasPrice uint256) returns()
func (_OraclizeI *OraclizeISession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetCustomGasPrice(&_OraclizeI.TransactOpts, _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(_gasPrice uint256) returns()
func (_OraclizeI *OraclizeITransactorSession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetCustomGasPrice(&_OraclizeI.TransactOpts, _gasPrice)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(_proofType bytes1) returns()
func (_OraclizeI *OraclizeITransactor) SetProofType(opts *bind.TransactOpts, _proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "setProofType", _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(_proofType bytes1) returns()
func (_OraclizeI *OraclizeISession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetProofType(&_OraclizeI.TransactOpts, _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(_proofType bytes1) returns()
func (_OraclizeI *OraclizeITransactorSession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetProofType(&_OraclizeI.TransactOpts, _proofType)
}

// UseCoupon is a paid mutator transaction binding the contract method 0x60f66701.
//
// Solidity: function useCoupon(_coupon string) returns()
func (_OraclizeI *OraclizeITransactor) UseCoupon(opts *bind.TransactOpts, _coupon string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "useCoupon", _coupon)
}

// UseCoupon is a paid mutator transaction binding the contract method 0x60f66701.
//
// Solidity: function useCoupon(_coupon string) returns()
func (_OraclizeI *OraclizeISession) UseCoupon(_coupon string) (*types.Transaction, error) {
	return _OraclizeI.Contract.UseCoupon(&_OraclizeI.TransactOpts, _coupon)
}

// UseCoupon is a paid mutator transaction binding the contract method 0x60f66701.
//
// Solidity: function useCoupon(_coupon string) returns()
func (_OraclizeI *OraclizeITransactorSession) UseCoupon(_coupon string) (*types.Transaction, error) {
	return _OraclizeI.Contract.UseCoupon(&_OraclizeI.TransactOpts, _coupon)
}

// WrapperOraclizeABI is the input ABI used to generate the binding from.
const WrapperOraclizeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"myid\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"string\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"myid\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"string\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWrapperData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWrapperBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"datasource\",\"type\":\"string\"},{\"name\":\"arg\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"datasource\",\"type\":\"string\"},{\"name\":\"arg\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"info\",\"type\":\"string\"}],\"name\":\"ShowOraclizeInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"description\",\"type\":\"string\"}],\"name\":\"newOraclizeQuery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"data\",\"type\":\"string\"}],\"name\":\"newOraclizeData\",\"type\":\"event\"}]"

// WrapperOraclizeBin is the compiled bytecode used for deploying new contracts.
const WrapperOraclizeBin = `0x606060405234156200001057600080fd5b5b7f80bda1fc37f774cb87cb9cf0c74b1239a6b7b14a7e24f24d6a843d53124357dd6040516020808252600f908201527f494e495449414c495a452e2e2e2e2e00000000000000000000000000000000006040808301919091526060909101905180910390a16200009060a1640100000000620012686200009882021704565b505b62000561565b600080620000c8731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed640100000000620012606200049582021704565b11156200014a5760008054600160a060020a031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed1790556200014160408051908101604052600b81527f6574685f6d61696e6e65740000000000000000000000000000000000000000006020820152640100000000620015a76200049d82021704565b50600162000490565b60006200017973c03a2615d5efaf5f49f60b7bb6583eaec212fdf1640100000000620012606200049582021704565b1115620001fb5760008054600160a060020a03191673c03a2615d5efaf5f49f60b7bb6583eaec212fdf11790556200014160408051908101604052600c81527f6574685f726f707374656e3300000000000000000000000000000000000000006020820152640100000000620015a76200049d82021704565b50600162000490565b60006200022a73b7a07bcf2ba2f2703b24c0691b5278999c59ac7e640100000000620012606200049582021704565b1115620002ac5760008054600160a060020a03191673b7a07bcf2ba2f2703b24c0691b5278999c59ac7e1790556200014160408051908101604052600981527f6574685f6b6f76616e00000000000000000000000000000000000000000000006020820152640100000000620015a76200049d82021704565b50600162000490565b6000620002db73146500cfd35b22e4a392fe0adc06de1a1368ed48640100000000620012606200049582021704565b11156200035d5760008054600160a060020a03191673146500cfd35b22e4a392fe0adc06de1a1368ed481790556200014160408051908101604052600b81527f6574685f72696e6b6562790000000000000000000000000000000000000000006020820152640100000000620015a76200049d82021704565b50600162000490565b60006200038c736f485c8bf6fc43ea212e93bbf8ce046c7f1cb475640100000000620012606200049582021704565b1115620003c2575060008054600160a060020a031916736f485c8bf6fc43ea212e93bbf8ce046c7f1cb475179055600162000490565b6000620003f17320e12a1f859b3feae5fb2a0a32c18f5a65555bbf640100000000620012606200049582021704565b111562000427575060008054600160a060020a0319167320e12a1f859b3feae5fb2a0a32c18f5a65555bbf179055600162000490565b6000620004567351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa640100000000620012606200049582021704565b11156200048c575060008054600160a060020a0319167351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa179055600162000490565b5060005b919050565b803b5b919050565b6002818051620004b2929160200190620004b7565b505b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004fa57805160ff19168380011785556200052a565b828001600101855582156200052a579182015b828111156200052a5782518255916020019190600101906200050d565b5b50620005399291506200053d565b5090565b6200055e91905b8082111562000539576000815560010162000544565b5090565b90565b6116bd80620005716000396000f300606060405236156100675763ffffffff60e060020a60003504166327dc297e811461006c57806338bbfa50146100c45780633d7569731461015e57806373d4a13a146101835780638747a6741461020e578063f4c84d1914610233578063f95e0a54146102bd575b600080fd5b341561007757600080fd5b6100c2600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061034c95505050505050565b005b34156100cf57600080fd5b6100c2600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061050095505050505050565b005b341561016957600080fd5b610171610506565b60405190815260200160405180910390f35b341561018e57600080fd5b6101966105b1565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101d35780820151818401525b6020016101ba565b50505050905090810190601f1680156102005780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561021957600080fd5b61017161064f565b60405190815260200160405180910390f35b6100c260046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061065e95505050505050565b005b6100c2600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506107d995505050505050565b005b7f80bda1fc37f774cb87cb9cf0c74b1239a6b7b14a7e24f24d6a843d53124357dd60405160208082526017908201527f455845435554494e472043414c4c4241434b2e2e2e2e2e0000000000000000006040808301919091526060909101905180910390a16103b9610956565b600160a060020a031633600160a060020a03161415156103d857600080fd5b7f80bda1fc37f774cb87cb9cf0c74b1239a6b7b14a7e24f24d6a843d53124357dd60405160208082526013908201527f57524954494e4720524553554c542e2e2e2e2e000000000000000000000000006040808301919091526060909101905180910390a160058180516104509291602001906115bf565b507f6378a25beb362c1c8d44a1d61ea44ec4b6253f3cf3c6bb187e0801346043d8b26005604051602080825282546002600019610100600184161502019091160490820181905281906040820190849080156104ed5780601f106104c2576101008083540402835291602001916104ed565b820191906000526020600020905b8154815290600101906020018083116104d057829003601f168201915b50509250505060405180910390a15b5050565b5b505050565b60006105ab60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105a15780601f10610576576101008083540402835291602001916105a1565b820191906000526020600020905b81548152906001019060200180831161058457829003601f168201915b5050505050610a72565b90505b90565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106475780601f1061061c57610100808354040283529160200191610647565b820191906000526020600020905b81548152906001019060200180831161062a57829003601f168201915b505050505081565b600160a060020a033016315b90565b30600160a060020a0316316106a560408051908101604052600381527f55524c00000000000000000000000000000000000000000000000000000000006020820152610ae7565b111561074f576000805160206116728339815191526040516020808252604b908201527f4f7261636c697a6520717565727920776173204e4f542073656e742c20706c656040808301919091527f6173652061646420736f6d652045544820746f20636f76657220666f7220746860608301527f6520717565727920666565000000000000000000000000000000000000000000608083015260a0909101905180910390a16104fc565b60008051602061167283398151915260405160208082526035908201527f4f7261636c697a65207175657279207761732073656e742c207374616e64696e6040808301919091527f6720627920666f722074686520616e737765722e2e000000000000000000000060608301526080909101905180910390a16105008282610c73565b505b5b5050565b30600160a060020a03163161082060408051908101604052600381527f55524c00000000000000000000000000000000000000000000000000000000006020820152610ae7565b11156108ca576000805160206116728339815191526040516020808252604b908201527f4f7261636c697a6520717565727920776173204e4f542073656e742c20706c656040808301919091527f6173652061646420736f6d652045544820746f20636f76657220666f7220746860608301527f6520717565727920666565000000000000000000000000000000000000000000608083015260a0909101905180910390a1610500565b60008051602061167283398151915260405160208082526035908201527f4f7261636c697a65207175657279207761732073656e742c207374616e64696e6040808301919091527f6720627920666f722074686520616e737765722e2e000000000000000000000060608301526080909101905180910390a161094e838383610f69565b505b5b505050565b60008054600160a060020a03161580610981575060005461097f90600160a060020a0316611260565b155b15610992576109906000611268565b505b60008054600160a060020a0316906338cc483190604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156109db57600080fd5b6102c65a03f115156109ec57600080fd5b505050604051805160018054600160a060020a031916600160a060020a03928316179081905516905063c281d19e6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610a5157600080fd5b6102c65a03f11515610a6257600080fd5b50505060405180519150505b5b90565b6000610a7c61163e565b50816000805b6020811015610ad8578251811015610abe57828181518110610aa057fe5b016020015160f860020a900460f860020a0260f860020a9004821791505b601f811015610acf57816101000291505b5b600101610a82565b8160010293505b505050919050565b60008054600160a060020a03161580610b125750600054610b1090600160a060020a0316611260565b155b15610b2357610b216000611268565b505b60008054600160a060020a0316906338cc483190604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610b6c57600080fd5b6102c65a03f11515610b7d57600080fd5b505050604051805160018054600160a060020a031916600160a060020a03928316179081905516905063524f3889836000604051602001526040518263ffffffff1660e060020a0281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610c055780820151818401525b602001610bec565b50505050905090810190601f168015610c325780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1515610c5057600080fd5b6102c65a03f11515610c6157600080fd5b50505060405180519150505b5b919050565b600080548190600160a060020a03161580610ca05750600054610c9e90600160a060020a0316611260565b155b15610cb157610caf6000611268565b505b60008054600160a060020a0316906338cc483190604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610cfa57600080fd5b6102c65a03f11515610d0b57600080fd5b505050604051805160018054600160a060020a031916600160a060020a03928316179081905516905063524f3889856000604051602001526040518263ffffffff1660e060020a0281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610d935780820151818401525b602001610d7a565b50505050905090810190601f168015610dc05780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1515610dde57600080fd5b6102c65a03f11515610def57600080fd5b5050506040518051915050670de0b6b3a764000062030d403a0201811115610e1a5760009150610f61565b600154600160a060020a031663adf59f99826000878782604051602001526040518563ffffffff1660e060020a028152600401808481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015610e925780820151818401525b602001610e79565b50505050905090810190601f168015610ebf5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610ef65780820151818401525b602001610edd565b50505050905090810190601f168015610f235780820380516001836020036101000a031916815260200191505b50955050505050506020604051808303818588803b1515610f4357600080fd5b6125ee5a03f11515610f5457600080fd5b5050505060405180519250505b5b5092915050565b600080548190600160a060020a03161580610f965750600054610f9490600160a060020a0316611260565b155b15610fa757610fa56000611268565b505b60008054600160a060020a0316906338cc483190604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610ff057600080fd5b6102c65a03f1151561100157600080fd5b505050604051805160018054600160a060020a031916600160a060020a03928316179081905516905063524f3889856000604051602001526040518263ffffffff1660e060020a0281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156110895780820151818401525b602001611070565b50505050905090810190601f1680156110b65780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15156110d457600080fd5b6102c65a03f115156110e557600080fd5b5050506040518051915050670de0b6b3a764000062030d403a02018111156111105760009150611257565b600154600160a060020a031663adf59f99828787876000604051602001526040518563ffffffff1660e060020a028152600401808481526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156111885780820151818401525b60200161116f565b50505050905090810190601f1680156111b55780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156111ec5780820151818401525b6020016111d3565b50505050905090810190601f1680156112195780820380516001836020036101000a031916815260200191505b50955050505050506020604051808303818588803b151561123957600080fd5b6125ee5a03f1151561124a57600080fd5b5050505060405180519250505b5b509392505050565b803b5b919050565b600080611288731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed611260565b11156112f85760008054600160a060020a031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed1790556112f060408051908101604052600b81527f6574685f6d61696e6e657400000000000000000000000000000000000000000060208201526115a7565b506001610c6d565b600061131773c03a2615d5efaf5f49f60b7bb6583eaec212fdf1611260565b11156113875760008054600160a060020a03191673c03a2615d5efaf5f49f60b7bb6583eaec212fdf11790556112f060408051908101604052600c81527f6574685f726f707374656e33000000000000000000000000000000000000000060208201526115a7565b506001610c6d565b60006113a673b7a07bcf2ba2f2703b24c0691b5278999c59ac7e611260565b11156114165760008054600160a060020a03191673b7a07bcf2ba2f2703b24c0691b5278999c59ac7e1790556112f060408051908101604052600981527f6574685f6b6f76616e000000000000000000000000000000000000000000000060208201526115a7565b506001610c6d565b600061143573146500cfd35b22e4a392fe0adc06de1a1368ed48611260565b11156114a55760008054600160a060020a03191673146500cfd35b22e4a392fe0adc06de1a1368ed481790556112f060408051908101604052600b81527f6574685f72696e6b65627900000000000000000000000000000000000000000060208201526115a7565b506001610c6d565b60006114c4736f485c8bf6fc43ea212e93bbf8ce046c7f1cb475611260565b11156114f8575060008054600160a060020a031916736f485c8bf6fc43ea212e93bbf8ce046c7f1cb4751790556001610c6d565b60006115177320e12a1f859b3feae5fb2a0a32c18f5a65555bbf611260565b111561154b575060008054600160a060020a0319167320e12a1f859b3feae5fb2a0a32c18f5a65555bbf1790556001610c6d565b600061156a7351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa611260565b111561159e575060008054600160a060020a0319167351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa1790556001610c6d565b5060005b919050565b60028180516104fc9291602001906115bf565b505b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061160057805160ff191683800117855561162d565b8280016001018555821561162d579182015b8281111561162d578251825591602001919060010190611612565b5b5061163a929150611650565b5090565b60206040519081016040526000815290565b6105ae91905b8082111561163a5760008155600101611656565b5090565b90560046cb989ef9cef13e930e3b7f286225a086e716a90d63e0b7da85d310a9db0c9aa165627a7a723058201bbc73fc93fac26ac1d7e1fd5ef3565c7a103f242197d327a3f5b28e110f48d90029`

// DeployWrapperOraclize deploys a new Ethereum contract, binding an instance of WrapperOraclize to it.
func DeployWrapperOraclize(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WrapperOraclize, error) {
	parsed, err := abi.JSON(strings.NewReader(WrapperOraclizeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WrapperOraclizeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WrapperOraclize{WrapperOraclizeCaller: WrapperOraclizeCaller{contract: contract}, WrapperOraclizeTransactor: WrapperOraclizeTransactor{contract: contract}}, nil
}

// WrapperOraclize is an auto generated Go binding around an Ethereum contract.
type WrapperOraclize struct {
	WrapperOraclizeCaller     // Read-only binding to the contract
	WrapperOraclizeTransactor // Write-only binding to the contract
}

// WrapperOraclizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrapperOraclizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapperOraclizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrapperOraclizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapperOraclizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrapperOraclizeSession struct {
	Contract     *WrapperOraclize  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WrapperOraclizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrapperOraclizeCallerSession struct {
	Contract *WrapperOraclizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WrapperOraclizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrapperOraclizeTransactorSession struct {
	Contract     *WrapperOraclizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WrapperOraclizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrapperOraclizeRaw struct {
	Contract *WrapperOraclize // Generic contract binding to access the raw methods on
}

// WrapperOraclizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrapperOraclizeCallerRaw struct {
	Contract *WrapperOraclizeCaller // Generic read-only contract binding to access the raw methods on
}

// WrapperOraclizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrapperOraclizeTransactorRaw struct {
	Contract *WrapperOraclizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrapperOraclize creates a new instance of WrapperOraclize, bound to a specific deployed contract.
func NewWrapperOraclize(address common.Address, backend bind.ContractBackend) (*WrapperOraclize, error) {
	contract, err := bindWrapperOraclize(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrapperOraclize{WrapperOraclizeCaller: WrapperOraclizeCaller{contract: contract}, WrapperOraclizeTransactor: WrapperOraclizeTransactor{contract: contract}}, nil
}

// NewWrapperOraclizeCaller creates a new read-only instance of WrapperOraclize, bound to a specific deployed contract.
func NewWrapperOraclizeCaller(address common.Address, caller bind.ContractCaller) (*WrapperOraclizeCaller, error) {
	contract, err := bindWrapperOraclize(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &WrapperOraclizeCaller{contract: contract}, nil
}

// NewWrapperOraclizeTransactor creates a new write-only instance of WrapperOraclize, bound to a specific deployed contract.
func NewWrapperOraclizeTransactor(address common.Address, transactor bind.ContractTransactor) (*WrapperOraclizeTransactor, error) {
	contract, err := bindWrapperOraclize(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &WrapperOraclizeTransactor{contract: contract}, nil
}

// bindWrapperOraclize binds a generic wrapper to an already deployed contract.
func bindWrapperOraclize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WrapperOraclizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrapperOraclize *WrapperOraclizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WrapperOraclize.Contract.WrapperOraclizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrapperOraclize *WrapperOraclizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapperOraclize.Contract.WrapperOraclizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrapperOraclize *WrapperOraclizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrapperOraclize.Contract.WrapperOraclizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrapperOraclize *WrapperOraclizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WrapperOraclize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrapperOraclize *WrapperOraclizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapperOraclize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrapperOraclize *WrapperOraclizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrapperOraclize.Contract.contract.Transact(opts, method, params...)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_WrapperOraclize *WrapperOraclizeCaller) Data(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WrapperOraclize.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_WrapperOraclize *WrapperOraclizeSession) Data() (string, error) {
	return _WrapperOraclize.Contract.Data(&_WrapperOraclize.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_WrapperOraclize *WrapperOraclizeCallerSession) Data() (string, error) {
	return _WrapperOraclize.Contract.Data(&_WrapperOraclize.CallOpts)
}

// GetWrapperBalance is a free data retrieval call binding the contract method 0x8747a674.
//
// Solidity: function getWrapperBalance() constant returns(uint256)
func (_WrapperOraclize *WrapperOraclizeCaller) GetWrapperBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WrapperOraclize.contract.Call(opts, out, "getWrapperBalance")
	return *ret0, err
}

// GetWrapperBalance is a free data retrieval call binding the contract method 0x8747a674.
//
// Solidity: function getWrapperBalance() constant returns(uint256)
func (_WrapperOraclize *WrapperOraclizeSession) GetWrapperBalance() (*big.Int, error) {
	return _WrapperOraclize.Contract.GetWrapperBalance(&_WrapperOraclize.CallOpts)
}

// GetWrapperBalance is a free data retrieval call binding the contract method 0x8747a674.
//
// Solidity: function getWrapperBalance() constant returns(uint256)
func (_WrapperOraclize *WrapperOraclizeCallerSession) GetWrapperBalance() (*big.Int, error) {
	return _WrapperOraclize.Contract.GetWrapperBalance(&_WrapperOraclize.CallOpts)
}

// GetWrapperData is a free data retrieval call binding the contract method 0x3d756973.
//
// Solidity: function getWrapperData() constant returns(bytes32)
func (_WrapperOraclize *WrapperOraclizeCaller) GetWrapperData(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WrapperOraclize.contract.Call(opts, out, "getWrapperData")
	return *ret0, err
}

// GetWrapperData is a free data retrieval call binding the contract method 0x3d756973.
//
// Solidity: function getWrapperData() constant returns(bytes32)
func (_WrapperOraclize *WrapperOraclizeSession) GetWrapperData() ([32]byte, error) {
	return _WrapperOraclize.Contract.GetWrapperData(&_WrapperOraclize.CallOpts)
}

// GetWrapperData is a free data retrieval call binding the contract method 0x3d756973.
//
// Solidity: function getWrapperData() constant returns(bytes32)
func (_WrapperOraclize *WrapperOraclizeCallerSession) GetWrapperData() ([32]byte, error) {
	return _WrapperOraclize.Contract.GetWrapperData(&_WrapperOraclize.CallOpts)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_WrapperOraclize *WrapperOraclizeTransactor) __callback(opts *bind.TransactOpts, myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _WrapperOraclize.contract.Transact(opts, "__callback", myid, result, proof)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_WrapperOraclize *WrapperOraclizeSession) __callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _WrapperOraclize.Contract.__callback(&_WrapperOraclize.TransactOpts, myid, result, proof)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_WrapperOraclize *WrapperOraclizeTransactorSession) __callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _WrapperOraclize.Contract.__callback(&_WrapperOraclize.TransactOpts, myid, result, proof)
}

// Update is a paid mutator transaction binding the contract method 0xf95e0a54.
//
// Solidity: function update(timestamp uint256, datasource string, arg string) returns()
func (_WrapperOraclize *WrapperOraclizeTransactor) Update(opts *bind.TransactOpts, timestamp *big.Int, datasource string, arg string) (*types.Transaction, error) {
	return _WrapperOraclize.contract.Transact(opts, "update", timestamp, datasource, arg)
}

// Update is a paid mutator transaction binding the contract method 0xf95e0a54.
//
// Solidity: function update(timestamp uint256, datasource string, arg string) returns()
func (_WrapperOraclize *WrapperOraclizeSession) Update(timestamp *big.Int, datasource string, arg string) (*types.Transaction, error) {
	return _WrapperOraclize.Contract.Update(&_WrapperOraclize.TransactOpts, timestamp, datasource, arg)
}

// Update is a paid mutator transaction binding the contract method 0xf95e0a54.
//
// Solidity: function update(timestamp uint256, datasource string, arg string) returns()
func (_WrapperOraclize *WrapperOraclizeTransactorSession) Update(timestamp *big.Int, datasource string, arg string) (*types.Transaction, error) {
	return _WrapperOraclize.Contract.Update(&_WrapperOraclize.TransactOpts, timestamp, datasource, arg)
}

// UsingOraclizeABI is the input ABI used to generate the binding from.
const UsingOraclizeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"myid\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"string\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"myid\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"string\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UsingOraclizeBin is the compiled bytecode used for deploying new contracts.
const UsingOraclizeBin = `0x6060604052341561000f57600080fd5b5b61019c8061001f6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327dc297e811461004857806338bbfa50146100a0575b600080fd5b341561005357600080fd5b61009e600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061013a95505050505050565b005b34156100ab57600080fd5b61009e600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061016a95505050505050565b005b6101658282600060405180591061014e5750595b908082528060200260200182016040525b5061016a565b5b5050565b5b5050505600a165627a7a72305820d06e136aca38046280712076b232fae1fe156528fab62273098519cc82cb581e0029`

// DeployUsingOraclize deploys a new Ethereum contract, binding an instance of UsingOraclize to it.
func DeployUsingOraclize(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UsingOraclize, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingOraclizeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsingOraclizeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UsingOraclize{UsingOraclizeCaller: UsingOraclizeCaller{contract: contract}, UsingOraclizeTransactor: UsingOraclizeTransactor{contract: contract}}, nil
}

// UsingOraclize is an auto generated Go binding around an Ethereum contract.
type UsingOraclize struct {
	UsingOraclizeCaller     // Read-only binding to the contract
	UsingOraclizeTransactor // Write-only binding to the contract
}

// UsingOraclizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsingOraclizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingOraclizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsingOraclizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingOraclizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsingOraclizeSession struct {
	Contract     *UsingOraclize    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsingOraclizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsingOraclizeCallerSession struct {
	Contract *UsingOraclizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UsingOraclizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsingOraclizeTransactorSession struct {
	Contract     *UsingOraclizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UsingOraclizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsingOraclizeRaw struct {
	Contract *UsingOraclize // Generic contract binding to access the raw methods on
}

// UsingOraclizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsingOraclizeCallerRaw struct {
	Contract *UsingOraclizeCaller // Generic read-only contract binding to access the raw methods on
}

// UsingOraclizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsingOraclizeTransactorRaw struct {
	Contract *UsingOraclizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsingOraclize creates a new instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclize(address common.Address, backend bind.ContractBackend) (*UsingOraclize, error) {
	contract, err := bindUsingOraclize(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsingOraclize{UsingOraclizeCaller: UsingOraclizeCaller{contract: contract}, UsingOraclizeTransactor: UsingOraclizeTransactor{contract: contract}}, nil
}

// NewUsingOraclizeCaller creates a new read-only instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclizeCaller(address common.Address, caller bind.ContractCaller) (*UsingOraclizeCaller, error) {
	contract, err := bindUsingOraclize(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &UsingOraclizeCaller{contract: contract}, nil
}

// NewUsingOraclizeTransactor creates a new write-only instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclizeTransactor(address common.Address, transactor bind.ContractTransactor) (*UsingOraclizeTransactor, error) {
	contract, err := bindUsingOraclize(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &UsingOraclizeTransactor{contract: contract}, nil
}

// bindUsingOraclize binds a generic wrapper to an already deployed contract.
func bindUsingOraclize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingOraclizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingOraclize *UsingOraclizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsingOraclize.Contract.UsingOraclizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingOraclize *UsingOraclizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingOraclize.Contract.UsingOraclizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingOraclize *UsingOraclizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingOraclize.Contract.UsingOraclizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingOraclize *UsingOraclizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsingOraclize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingOraclize *UsingOraclizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingOraclize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingOraclize *UsingOraclizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingOraclize.Contract.contract.Transact(opts, method, params...)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_UsingOraclize *UsingOraclizeTransactor) __callback(opts *bind.TransactOpts, myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _UsingOraclize.contract.Transact(opts, "__callback", myid, result, proof)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_UsingOraclize *UsingOraclizeSession) __callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _UsingOraclize.Contract.__callback(&_UsingOraclize.TransactOpts, myid, result, proof)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_UsingOraclize *UsingOraclizeTransactorSession) __callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _UsingOraclize.Contract.__callback(&_UsingOraclize.TransactOpts, myid, result, proof)
}
