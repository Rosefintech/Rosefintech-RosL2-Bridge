// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FundABI is the input ABI used to generate the binding from.
const FundABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wbtcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETHAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_fetchAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_returnAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fetchNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_returnNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"ExchangeLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_billNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenBrokerageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenChargeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"OutOfPositionLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_ethNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_usdcNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"PositionLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPATCH_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXCHANGE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_deviation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_exchangeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addFundERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"brokerageAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chargeAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_return_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_queryId\",\"type\":\"uint256\"}],\"name\":\"fetchETH2Token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fetch_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_queryId\",\"type\":\"uint256\"}],\"name\":\"fetchToken2ETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fetch_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_return_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_queryId\",\"type\":\"uint256\"}],\"name\":\"fetchToken2Token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssert\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_eth\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"_erc20Symbol\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_balance\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundERC20Address\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addr\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenNum\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"billTokenNum\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenBrokerageNum\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenChargeNum\",\"type\":\"uint256[]\"}],\"name\":\"outOfPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenRes\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"removeFundERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_brokerageAddr\",\"type\":\"address\"}],\"name\":\"setBrokerageAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chargeAddress\",\"type\":\"address\"}],\"name\":\"setChargeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dev\",\"type\":\"uint256\"}],\"name\":\"setDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exchangeContract\",\"type\":\"address\"}],\"name\":\"setExchangeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_USDCAddress\",\"type\":\"address\"}],\"name\":\"setUSDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETHAddress\",\"type\":\"address\"}],\"name\":\"setWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdcNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"takePosition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Fund is an auto generated Go binding around an Ethereum contract.
type Fund struct {
	FundCaller     // Read-only binding to the contract
	FundTransactor // Write-only binding to the contract
	FundFilterer   // Log filterer for contract events
}

// FundCaller is an auto generated read-only Go binding around an Ethereum contract.
type FundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FundSession struct {
	Contract     *Fund             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FundCallerSession struct {
	Contract *FundCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FundTransactorSession struct {
	Contract     *FundTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FundRaw is an auto generated low-level Go binding around an Ethereum contract.
type FundRaw struct {
	Contract *Fund // Generic contract binding to access the raw methods on
}

// FundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FundCallerRaw struct {
	Contract *FundCaller // Generic read-only contract binding to access the raw methods on
}

// FundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FundTransactorRaw struct {
	Contract *FundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFund creates a new instance of Fund, bound to a specific deployed contract.
func NewFund(address common.Address, backend bind.ContractBackend) (*Fund, error) {
	contract, err := bindFund(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fund{FundCaller: FundCaller{contract: contract}, FundTransactor: FundTransactor{contract: contract}, FundFilterer: FundFilterer{contract: contract}}, nil
}

// NewFundCaller creates a new read-only instance of Fund, bound to a specific deployed contract.
func NewFundCaller(address common.Address, caller bind.ContractCaller) (*FundCaller, error) {
	contract, err := bindFund(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FundCaller{contract: contract}, nil
}

// NewFundTransactor creates a new write-only instance of Fund, bound to a specific deployed contract.
func NewFundTransactor(address common.Address, transactor bind.ContractTransactor) (*FundTransactor, error) {
	contract, err := bindFund(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FundTransactor{contract: contract}, nil
}

// NewFundFilterer creates a new log filterer instance of Fund, bound to a specific deployed contract.
func NewFundFilterer(address common.Address, filterer bind.ContractFilterer) (*FundFilterer, error) {
	contract, err := bindFund(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FundFilterer{contract: contract}, nil
}

// bindFund binds a generic wrapper to an already deployed contract.
func bindFund(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FundABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fund *FundRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fund.Contract.FundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fund *FundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fund.Contract.FundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fund *FundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fund.Contract.FundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fund *FundCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fund.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fund *FundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fund.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fund *FundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fund.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Fund *FundCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Fund *FundSession) ADMINROLE() ([32]byte, error) {
	return _Fund.Contract.ADMINROLE(&_Fund.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Fund *FundCallerSession) ADMINROLE() ([32]byte, error) {
	return _Fund.Contract.ADMINROLE(&_Fund.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Fund *FundCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Fund *FundSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Fund.Contract.DEFAULTADMINROLE(&_Fund.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Fund *FundCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Fund.Contract.DEFAULTADMINROLE(&_Fund.CallOpts)
}

// DISPATCHROLE is a free data retrieval call binding the contract method 0x501bcb31.
//
// Solidity: function DISPATCH_ROLE() view returns(bytes32)
func (_Fund *FundCaller) DISPATCHROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "DISPATCH_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DISPATCHROLE is a free data retrieval call binding the contract method 0x501bcb31.
//
// Solidity: function DISPATCH_ROLE() view returns(bytes32)
func (_Fund *FundSession) DISPATCHROLE() ([32]byte, error) {
	return _Fund.Contract.DISPATCHROLE(&_Fund.CallOpts)
}

// DISPATCHROLE is a free data retrieval call binding the contract method 0x501bcb31.
//
// Solidity: function DISPATCH_ROLE() view returns(bytes32)
func (_Fund *FundCallerSession) DISPATCHROLE() ([32]byte, error) {
	return _Fund.Contract.DISPATCHROLE(&_Fund.CallOpts)
}

// EXCHANGEROLE is a free data retrieval call binding the contract method 0x3b734501.
//
// Solidity: function EXCHANGE_ROLE() view returns(bytes32)
func (_Fund *FundCaller) EXCHANGEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "EXCHANGE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXCHANGEROLE is a free data retrieval call binding the contract method 0x3b734501.
//
// Solidity: function EXCHANGE_ROLE() view returns(bytes32)
func (_Fund *FundSession) EXCHANGEROLE() ([32]byte, error) {
	return _Fund.Contract.EXCHANGEROLE(&_Fund.CallOpts)
}

// EXCHANGEROLE is a free data retrieval call binding the contract method 0x3b734501.
//
// Solidity: function EXCHANGE_ROLE() view returns(bytes32)
func (_Fund *FundCallerSession) EXCHANGEROLE() ([32]byte, error) {
	return _Fund.Contract.EXCHANGEROLE(&_Fund.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Fund *FundCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Fund *FundSession) USDC() (common.Address, error) {
	return _Fund.Contract.USDC(&_Fund.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Fund *FundCallerSession) USDC() (common.Address, error) {
	return _Fund.Contract.USDC(&_Fund.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Fund *FundCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Fund *FundSession) WETH() (common.Address, error) {
	return _Fund.Contract.WETH(&_Fund.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Fund *FundCallerSession) WETH() (common.Address, error) {
	return _Fund.Contract.WETH(&_Fund.CallOpts)
}

// Deviation is a free data retrieval call binding the contract method 0xb7778175.
//
// Solidity: function _deviation() view returns(uint256)
func (_Fund *FundCaller) Deviation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "_deviation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deviation is a free data retrieval call binding the contract method 0xb7778175.
//
// Solidity: function _deviation() view returns(uint256)
func (_Fund *FundSession) Deviation() (*big.Int, error) {
	return _Fund.Contract.Deviation(&_Fund.CallOpts)
}

// Deviation is a free data retrieval call binding the contract method 0xb7778175.
//
// Solidity: function _deviation() view returns(uint256)
func (_Fund *FundCallerSession) Deviation() (*big.Int, error) {
	return _Fund.Contract.Deviation(&_Fund.CallOpts)
}

// ExchangeAddress is a free data retrieval call binding the contract method 0x91b9cd82.
//
// Solidity: function _exchangeAddress() view returns(address)
func (_Fund *FundCaller) ExchangeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "_exchangeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExchangeAddress is a free data retrieval call binding the contract method 0x91b9cd82.
//
// Solidity: function _exchangeAddress() view returns(address)
func (_Fund *FundSession) ExchangeAddress() (common.Address, error) {
	return _Fund.Contract.ExchangeAddress(&_Fund.CallOpts)
}

// ExchangeAddress is a free data retrieval call binding the contract method 0x91b9cd82.
//
// Solidity: function _exchangeAddress() view returns(address)
func (_Fund *FundCallerSession) ExchangeAddress() (common.Address, error) {
	return _Fund.Contract.ExchangeAddress(&_Fund.CallOpts)
}

// BrokerageAddr is a free data retrieval call binding the contract method 0xcac413d8.
//
// Solidity: function brokerageAddr() view returns(address)
func (_Fund *FundCaller) BrokerageAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "brokerageAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BrokerageAddr is a free data retrieval call binding the contract method 0xcac413d8.
//
// Solidity: function brokerageAddr() view returns(address)
func (_Fund *FundSession) BrokerageAddr() (common.Address, error) {
	return _Fund.Contract.BrokerageAddr(&_Fund.CallOpts)
}

// BrokerageAddr is a free data retrieval call binding the contract method 0xcac413d8.
//
// Solidity: function brokerageAddr() view returns(address)
func (_Fund *FundCallerSession) BrokerageAddr() (common.Address, error) {
	return _Fund.Contract.BrokerageAddr(&_Fund.CallOpts)
}

// ChargeAddr is a free data retrieval call binding the contract method 0x8f04ec16.
//
// Solidity: function chargeAddr() view returns(address)
func (_Fund *FundCaller) ChargeAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "chargeAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChargeAddr is a free data retrieval call binding the contract method 0x8f04ec16.
//
// Solidity: function chargeAddr() view returns(address)
func (_Fund *FundSession) ChargeAddr() (common.Address, error) {
	return _Fund.Contract.ChargeAddr(&_Fund.CallOpts)
}

// ChargeAddr is a free data retrieval call binding the contract method 0x8f04ec16.
//
// Solidity: function chargeAddr() view returns(address)
func (_Fund *FundCallerSession) ChargeAddr() (common.Address, error) {
	return _Fund.Contract.ChargeAddr(&_Fund.CallOpts)
}

// GetAssert is a free data retrieval call binding the contract method 0x7cc465f0.
//
// Solidity: function getAssert() view returns(uint256 _eth, string[] _erc20Symbol, uint256[] _balance)
func (_Fund *FundCaller) GetAssert(opts *bind.CallOpts) (struct {
	Eth         *big.Int
	Erc20Symbol []string
	Balance     []*big.Int
}, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "getAssert")

	outstruct := new(struct {
		Eth         *big.Int
		Erc20Symbol []string
		Balance     []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Eth = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Erc20Symbol = *abi.ConvertType(out[1], new([]string)).(*[]string)
	outstruct.Balance = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAssert is a free data retrieval call binding the contract method 0x7cc465f0.
//
// Solidity: function getAssert() view returns(uint256 _eth, string[] _erc20Symbol, uint256[] _balance)
func (_Fund *FundSession) GetAssert() (struct {
	Eth         *big.Int
	Erc20Symbol []string
	Balance     []*big.Int
}, error) {
	return _Fund.Contract.GetAssert(&_Fund.CallOpts)
}

// GetAssert is a free data retrieval call binding the contract method 0x7cc465f0.
//
// Solidity: function getAssert() view returns(uint256 _eth, string[] _erc20Symbol, uint256[] _balance)
func (_Fund *FundCallerSession) GetAssert() (struct {
	Eth         *big.Int
	Erc20Symbol []string
	Balance     []*big.Int
}, error) {
	return _Fund.Contract.GetAssert(&_Fund.CallOpts)
}

// GetFundERC20Address is a free data retrieval call binding the contract method 0x79404cb9.
//
// Solidity: function getFundERC20Address() view returns(address[])
func (_Fund *FundCaller) GetFundERC20Address(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "getFundERC20Address")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFundERC20Address is a free data retrieval call binding the contract method 0x79404cb9.
//
// Solidity: function getFundERC20Address() view returns(address[])
func (_Fund *FundSession) GetFundERC20Address() ([]common.Address, error) {
	return _Fund.Contract.GetFundERC20Address(&_Fund.CallOpts)
}

// GetFundERC20Address is a free data retrieval call binding the contract method 0x79404cb9.
//
// Solidity: function getFundERC20Address() view returns(address[])
func (_Fund *FundCallerSession) GetFundERC20Address() ([]common.Address, error) {
	return _Fund.Contract.GetFundERC20Address(&_Fund.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Fund *FundCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Fund *FundSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Fund.Contract.GetRoleAdmin(&_Fund.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Fund *FundCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Fund.Contract.GetRoleAdmin(&_Fund.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Fund *FundCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Fund *FundSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Fund.Contract.GetRoleMember(&_Fund.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Fund *FundCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Fund.Contract.GetRoleMember(&_Fund.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Fund *FundCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Fund *FundSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Fund.Contract.GetRoleMemberCount(&_Fund.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Fund *FundCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Fund.Contract.GetRoleMemberCount(&_Fund.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Fund *FundCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Fund.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Fund *FundSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Fund.Contract.HasRole(&_Fund.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Fund *FundCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Fund.Contract.HasRole(&_Fund.CallOpts, role, account)
}

// AddFundERC20 is a paid mutator transaction binding the contract method 0xcbdd2fed.
//
// Solidity: function addFundERC20(address _addr) returns(bool)
func (_Fund *FundTransactor) AddFundERC20(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "addFundERC20", _addr)
}

// AddFundERC20 is a paid mutator transaction binding the contract method 0xcbdd2fed.
//
// Solidity: function addFundERC20(address _addr) returns(bool)
func (_Fund *FundSession) AddFundERC20(_addr common.Address) (*types.Transaction, error) {
	return _Fund.Contract.AddFundERC20(&_Fund.TransactOpts, _addr)
}

// AddFundERC20 is a paid mutator transaction binding the contract method 0xcbdd2fed.
//
// Solidity: function addFundERC20(address _addr) returns(bool)
func (_Fund *FundTransactorSession) AddFundERC20(_addr common.Address) (*types.Transaction, error) {
	return _Fund.Contract.AddFundERC20(&_Fund.TransactOpts, _addr)
}

// FetchETH2Token is a paid mutator transaction binding the contract method 0xdc1d2bf4.
//
// Solidity: function fetchETH2Token(address _return_address, uint256 _tokenNum, uint256 _queryId) returns(uint256)
func (_Fund *FundTransactor) FetchETH2Token(opts *bind.TransactOpts, _return_address common.Address, _tokenNum *big.Int, _queryId *big.Int) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "fetchETH2Token", _return_address, _tokenNum, _queryId)
}

// FetchETH2Token is a paid mutator transaction binding the contract method 0xdc1d2bf4.
//
// Solidity: function fetchETH2Token(address _return_address, uint256 _tokenNum, uint256 _queryId) returns(uint256)
func (_Fund *FundSession) FetchETH2Token(_return_address common.Address, _tokenNum *big.Int, _queryId *big.Int) (*types.Transaction, error) {
	return _Fund.Contract.FetchETH2Token(&_Fund.TransactOpts, _return_address, _tokenNum, _queryId)
}

// FetchETH2Token is a paid mutator transaction binding the contract method 0xdc1d2bf4.
//
// Solidity: function fetchETH2Token(address _return_address, uint256 _tokenNum, uint256 _queryId) returns(uint256)
func (_Fund *FundTransactorSession) FetchETH2Token(_return_address common.Address, _tokenNum *big.Int, _queryId *big.Int) (*types.Transaction, error) {
	return _Fund.Contract.FetchETH2Token(&_Fund.TransactOpts, _return_address, _tokenNum, _queryId)
}

// FetchToken2ETH is a paid mutator transaction binding the contract method 0xc5e9314f.
//
// Solidity: function fetchToken2ETH(address _fetch_address, uint256 _tokenNum, uint256 _queryId) returns(uint256)
func (_Fund *FundTransactor) FetchToken2ETH(opts *bind.TransactOpts, _fetch_address common.Address, _tokenNum *big.Int, _queryId *big.Int) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "fetchToken2ETH", _fetch_address, _tokenNum, _queryId)
}

// FetchToken2ETH is a paid mutator transaction binding the contract method 0xc5e9314f.
//
// Solidity: function fetchToken2ETH(address _fetch_address, uint256 _tokenNum, uint256 _queryId) returns(uint256)
func (_Fund *FundSession) FetchToken2ETH(_fetch_address common.Address, _tokenNum *big.Int, _queryId *big.Int) (*types.Transaction, error) {
	return _Fund.Contract.FetchToken2ETH(&_Fund.TransactOpts, _fetch_address, _tokenNum, _queryId)
}

// FetchToken2ETH is a paid mutator transaction binding the contract method 0xc5e9314f.
//
// Solidity: function fetchToken2ETH(address _fetch_address, uint256 _tokenNum, uint256 _queryId) returns(uint256)
func (_Fund *FundTransactorSession) FetchToken2ETH(_fetch_address common.Address, _tokenNum *big.Int, _queryId *big.Int) (*types.Transaction, error) {
	return _Fund.Contract.FetchToken2ETH(&_Fund.TransactOpts, _fetch_address, _tokenNum, _queryId)
}

// FetchToken2Token is a paid mutator transaction binding the contract method 0x552a793c.
//
// Solidity: function fetchToken2Token(address _fetch_address, address _return_address, uint256 _tokenNum, uint256 _queryId) returns(uint256)
func (_Fund *FundTransactor) FetchToken2Token(opts *bind.TransactOpts, _fetch_address common.Address, _return_address common.Address, _tokenNum *big.Int, _queryId *big.Int) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "fetchToken2Token", _fetch_address, _return_address, _tokenNum, _queryId)
}

// FetchToken2Token is a paid mutator transaction binding the contract method 0x552a793c.
//
// Solidity: function fetchToken2Token(address _fetch_address, address _return_address, uint256 _tokenNum, uint256 _queryId) returns(uint256)
func (_Fund *FundSession) FetchToken2Token(_fetch_address common.Address, _return_address common.Address, _tokenNum *big.Int, _queryId *big.Int) (*types.Transaction, error) {
	return _Fund.Contract.FetchToken2Token(&_Fund.TransactOpts, _fetch_address, _return_address, _tokenNum, _queryId)
}

// FetchToken2Token is a paid mutator transaction binding the contract method 0x552a793c.
//
// Solidity: function fetchToken2Token(address _fetch_address, address _return_address, uint256 _tokenNum, uint256 _queryId) returns(uint256)
func (_Fund *FundTransactorSession) FetchToken2Token(_fetch_address common.Address, _return_address common.Address, _tokenNum *big.Int, _queryId *big.Int) (*types.Transaction, error) {
	return _Fund.Contract.FetchToken2Token(&_Fund.TransactOpts, _fetch_address, _return_address, _tokenNum, _queryId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Fund *FundTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Fund *FundSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fund.Contract.GrantRole(&_Fund.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Fund *FundTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fund.Contract.GrantRole(&_Fund.TransactOpts, role, account)
}

// OutOfPosition is a paid mutator transaction binding the contract method 0x3a99a509.
//
// Solidity: function outOfPosition(address[] _addr, address tokenAddr, uint256[] tokenNum, uint256[] billTokenNum, uint256[] tokenBrokerageNum, uint256[] tokenChargeNum) returns(uint256 _tokenRes)
func (_Fund *FundTransactor) OutOfPosition(opts *bind.TransactOpts, _addr []common.Address, tokenAddr common.Address, tokenNum []*big.Int, billTokenNum []*big.Int, tokenBrokerageNum []*big.Int, tokenChargeNum []*big.Int) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "outOfPosition", _addr, tokenAddr, tokenNum, billTokenNum, tokenBrokerageNum, tokenChargeNum)
}

// OutOfPosition is a paid mutator transaction binding the contract method 0x3a99a509.
//
// Solidity: function outOfPosition(address[] _addr, address tokenAddr, uint256[] tokenNum, uint256[] billTokenNum, uint256[] tokenBrokerageNum, uint256[] tokenChargeNum) returns(uint256 _tokenRes)
func (_Fund *FundSession) OutOfPosition(_addr []common.Address, tokenAddr common.Address, tokenNum []*big.Int, billTokenNum []*big.Int, tokenBrokerageNum []*big.Int, tokenChargeNum []*big.Int) (*types.Transaction, error) {
	return _Fund.Contract.OutOfPosition(&_Fund.TransactOpts, _addr, tokenAddr, tokenNum, billTokenNum, tokenBrokerageNum, tokenChargeNum)
}

// OutOfPosition is a paid mutator transaction binding the contract method 0x3a99a509.
//
// Solidity: function outOfPosition(address[] _addr, address tokenAddr, uint256[] tokenNum, uint256[] billTokenNum, uint256[] tokenBrokerageNum, uint256[] tokenChargeNum) returns(uint256 _tokenRes)
func (_Fund *FundTransactorSession) OutOfPosition(_addr []common.Address, tokenAddr common.Address, tokenNum []*big.Int, billTokenNum []*big.Int, tokenBrokerageNum []*big.Int, tokenChargeNum []*big.Int) (*types.Transaction, error) {
	return _Fund.Contract.OutOfPosition(&_Fund.TransactOpts, _addr, tokenAddr, tokenNum, billTokenNum, tokenBrokerageNum, tokenChargeNum)
}

// RemoveFundERC20 is a paid mutator transaction binding the contract method 0xeccd2247.
//
// Solidity: function removeFundERC20(address _addr) returns(bool)
func (_Fund *FundTransactor) RemoveFundERC20(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "removeFundERC20", _addr)
}

// RemoveFundERC20 is a paid mutator transaction binding the contract method 0xeccd2247.
//
// Solidity: function removeFundERC20(address _addr) returns(bool)
func (_Fund *FundSession) RemoveFundERC20(_addr common.Address) (*types.Transaction, error) {
	return _Fund.Contract.RemoveFundERC20(&_Fund.TransactOpts, _addr)
}

// RemoveFundERC20 is a paid mutator transaction binding the contract method 0xeccd2247.
//
// Solidity: function removeFundERC20(address _addr) returns(bool)
func (_Fund *FundTransactorSession) RemoveFundERC20(_addr common.Address) (*types.Transaction, error) {
	return _Fund.Contract.RemoveFundERC20(&_Fund.TransactOpts, _addr)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Fund *FundTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Fund *FundSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fund.Contract.RenounceRole(&_Fund.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Fund *FundTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fund.Contract.RenounceRole(&_Fund.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Fund *FundTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Fund *FundSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fund.Contract.RevokeRole(&_Fund.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Fund *FundTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fund.Contract.RevokeRole(&_Fund.TransactOpts, role, account)
}

// SetBrokerageAddress is a paid mutator transaction binding the contract method 0xfc449038.
//
// Solidity: function setBrokerageAddress(address _brokerageAddr) returns()
func (_Fund *FundTransactor) SetBrokerageAddress(opts *bind.TransactOpts, _brokerageAddr common.Address) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "setBrokerageAddress", _brokerageAddr)
}

// SetBrokerageAddress is a paid mutator transaction binding the contract method 0xfc449038.
//
// Solidity: function setBrokerageAddress(address _brokerageAddr) returns()
func (_Fund *FundSession) SetBrokerageAddress(_brokerageAddr common.Address) (*types.Transaction, error) {
	return _Fund.Contract.SetBrokerageAddress(&_Fund.TransactOpts, _brokerageAddr)
}

// SetBrokerageAddress is a paid mutator transaction binding the contract method 0xfc449038.
//
// Solidity: function setBrokerageAddress(address _brokerageAddr) returns()
func (_Fund *FundTransactorSession) SetBrokerageAddress(_brokerageAddr common.Address) (*types.Transaction, error) {
	return _Fund.Contract.SetBrokerageAddress(&_Fund.TransactOpts, _brokerageAddr)
}

// SetChargeAddress is a paid mutator transaction binding the contract method 0x10e45e4c.
//
// Solidity: function setChargeAddress(address _chargeAddress) returns()
func (_Fund *FundTransactor) SetChargeAddress(opts *bind.TransactOpts, _chargeAddress common.Address) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "setChargeAddress", _chargeAddress)
}

// SetChargeAddress is a paid mutator transaction binding the contract method 0x10e45e4c.
//
// Solidity: function setChargeAddress(address _chargeAddress) returns()
func (_Fund *FundSession) SetChargeAddress(_chargeAddress common.Address) (*types.Transaction, error) {
	return _Fund.Contract.SetChargeAddress(&_Fund.TransactOpts, _chargeAddress)
}

// SetChargeAddress is a paid mutator transaction binding the contract method 0x10e45e4c.
//
// Solidity: function setChargeAddress(address _chargeAddress) returns()
func (_Fund *FundTransactorSession) SetChargeAddress(_chargeAddress common.Address) (*types.Transaction, error) {
	return _Fund.Contract.SetChargeAddress(&_Fund.TransactOpts, _chargeAddress)
}

// SetDeviation is a paid mutator transaction binding the contract method 0xe2fb8cbe.
//
// Solidity: function setDeviation(uint256 _dev) returns()
func (_Fund *FundTransactor) SetDeviation(opts *bind.TransactOpts, _dev *big.Int) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "setDeviation", _dev)
}

// SetDeviation is a paid mutator transaction binding the contract method 0xe2fb8cbe.
//
// Solidity: function setDeviation(uint256 _dev) returns()
func (_Fund *FundSession) SetDeviation(_dev *big.Int) (*types.Transaction, error) {
	return _Fund.Contract.SetDeviation(&_Fund.TransactOpts, _dev)
}

// SetDeviation is a paid mutator transaction binding the contract method 0xe2fb8cbe.
//
// Solidity: function setDeviation(uint256 _dev) returns()
func (_Fund *FundTransactorSession) SetDeviation(_dev *big.Int) (*types.Transaction, error) {
	return _Fund.Contract.SetDeviation(&_Fund.TransactOpts, _dev)
}

// SetExchangeAddress is a paid mutator transaction binding the contract method 0xd082ea8c.
//
// Solidity: function setExchangeAddress(address _exchangeContract) returns()
func (_Fund *FundTransactor) SetExchangeAddress(opts *bind.TransactOpts, _exchangeContract common.Address) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "setExchangeAddress", _exchangeContract)
}

// SetExchangeAddress is a paid mutator transaction binding the contract method 0xd082ea8c.
//
// Solidity: function setExchangeAddress(address _exchangeContract) returns()
func (_Fund *FundSession) SetExchangeAddress(_exchangeContract common.Address) (*types.Transaction, error) {
	return _Fund.Contract.SetExchangeAddress(&_Fund.TransactOpts, _exchangeContract)
}

// SetExchangeAddress is a paid mutator transaction binding the contract method 0xd082ea8c.
//
// Solidity: function setExchangeAddress(address _exchangeContract) returns()
func (_Fund *FundTransactorSession) SetExchangeAddress(_exchangeContract common.Address) (*types.Transaction, error) {
	return _Fund.Contract.SetExchangeAddress(&_Fund.TransactOpts, _exchangeContract)
}

// SetUSDC is a paid mutator transaction binding the contract method 0xb3e089a2.
//
// Solidity: function setUSDC(address _USDCAddress) returns()
func (_Fund *FundTransactor) SetUSDC(opts *bind.TransactOpts, _USDCAddress common.Address) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "setUSDC", _USDCAddress)
}

// SetUSDC is a paid mutator transaction binding the contract method 0xb3e089a2.
//
// Solidity: function setUSDC(address _USDCAddress) returns()
func (_Fund *FundSession) SetUSDC(_USDCAddress common.Address) (*types.Transaction, error) {
	return _Fund.Contract.SetUSDC(&_Fund.TransactOpts, _USDCAddress)
}

// SetUSDC is a paid mutator transaction binding the contract method 0xb3e089a2.
//
// Solidity: function setUSDC(address _USDCAddress) returns()
func (_Fund *FundTransactorSession) SetUSDC(_USDCAddress common.Address) (*types.Transaction, error) {
	return _Fund.Contract.SetUSDC(&_Fund.TransactOpts, _USDCAddress)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address _WETHAddress) returns()
func (_Fund *FundTransactor) SetWETH(opts *bind.TransactOpts, _WETHAddress common.Address) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "setWETH", _WETHAddress)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address _WETHAddress) returns()
func (_Fund *FundSession) SetWETH(_WETHAddress common.Address) (*types.Transaction, error) {
	return _Fund.Contract.SetWETH(&_Fund.TransactOpts, _WETHAddress)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address _WETHAddress) returns()
func (_Fund *FundTransactorSession) SetWETH(_WETHAddress common.Address) (*types.Transaction, error) {
	return _Fund.Contract.SetWETH(&_Fund.TransactOpts, _WETHAddress)
}

// TakePosition is a paid mutator transaction binding the contract method 0x8f6b95d3.
//
// Solidity: function takePosition(address _account, uint256 _usdcNum, bytes32 _id) payable returns(bool)
func (_Fund *FundTransactor) TakePosition(opts *bind.TransactOpts, _account common.Address, _usdcNum *big.Int, _id [32]byte) (*types.Transaction, error) {
	return _Fund.contract.Transact(opts, "takePosition", _account, _usdcNum, _id)
}

// TakePosition is a paid mutator transaction binding the contract method 0x8f6b95d3.
//
// Solidity: function takePosition(address _account, uint256 _usdcNum, bytes32 _id) payable returns(bool)
func (_Fund *FundSession) TakePosition(_account common.Address, _usdcNum *big.Int, _id [32]byte) (*types.Transaction, error) {
	return _Fund.Contract.TakePosition(&_Fund.TransactOpts, _account, _usdcNum, _id)
}

// TakePosition is a paid mutator transaction binding the contract method 0x8f6b95d3.
//
// Solidity: function takePosition(address _account, uint256 _usdcNum, bytes32 _id) payable returns(bool)
func (_Fund *FundTransactorSession) TakePosition(_account common.Address, _usdcNum *big.Int, _id [32]byte) (*types.Transaction, error) {
	return _Fund.Contract.TakePosition(&_Fund.TransactOpts, _account, _usdcNum, _id)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fund *FundTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fund.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fund *FundSession) Receive() (*types.Transaction, error) {
	return _Fund.Contract.Receive(&_Fund.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fund *FundTransactorSession) Receive() (*types.Transaction, error) {
	return _Fund.Contract.Receive(&_Fund.TransactOpts)
}

// FundExchangeLogIterator is returned from FilterExchangeLog and is used to iterate over the raw logs and unpacked data for ExchangeLog events raised by the Fund contract.
type FundExchangeLogIterator struct {
	Event *FundExchangeLog // Event containing the contract specifics and raw log

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
func (it *FundExchangeLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundExchangeLog)
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
		it.Event = new(FundExchangeLog)
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
func (it *FundExchangeLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundExchangeLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundExchangeLog represents a ExchangeLog event raised by the Fund contract.
type FundExchangeLog struct {
	FetchAddress  common.Address
	ReturnAddress common.Address
	FetchNum      *big.Int
	ReturnNum     *big.Int
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExchangeLog is a free log retrieval operation binding the contract event 0x54d1da88a25f3be7e7f0dafca34f087cfd940ad5a386cd40dc4aa8458d5b73ea.
//
// Solidity: event ExchangeLog(address _fetchAddress, address _returnAddress, uint256 _fetchNum, uint256 _returnNum, uint256 _timestamp)
func (_Fund *FundFilterer) FilterExchangeLog(opts *bind.FilterOpts) (*FundExchangeLogIterator, error) {

	logs, sub, err := _Fund.contract.FilterLogs(opts, "ExchangeLog")
	if err != nil {
		return nil, err
	}
	return &FundExchangeLogIterator{contract: _Fund.contract, event: "ExchangeLog", logs: logs, sub: sub}, nil
}

// WatchExchangeLog is a free log subscription operation binding the contract event 0x54d1da88a25f3be7e7f0dafca34f087cfd940ad5a386cd40dc4aa8458d5b73ea.
//
// Solidity: event ExchangeLog(address _fetchAddress, address _returnAddress, uint256 _fetchNum, uint256 _returnNum, uint256 _timestamp)
func (_Fund *FundFilterer) WatchExchangeLog(opts *bind.WatchOpts, sink chan<- *FundExchangeLog) (event.Subscription, error) {

	logs, sub, err := _Fund.contract.WatchLogs(opts, "ExchangeLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundExchangeLog)
				if err := _Fund.contract.UnpackLog(event, "ExchangeLog", log); err != nil {
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

// ParseExchangeLog is a log parse operation binding the contract event 0x54d1da88a25f3be7e7f0dafca34f087cfd940ad5a386cd40dc4aa8458d5b73ea.
//
// Solidity: event ExchangeLog(address _fetchAddress, address _returnAddress, uint256 _fetchNum, uint256 _returnNum, uint256 _timestamp)
func (_Fund *FundFilterer) ParseExchangeLog(log types.Log) (*FundExchangeLog, error) {
	event := new(FundExchangeLog)
	if err := _Fund.contract.UnpackLog(event, "ExchangeLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FundOutOfPositionLogIterator is returned from FilterOutOfPositionLog and is used to iterate over the raw logs and unpacked data for OutOfPositionLog events raised by the Fund contract.
type FundOutOfPositionLogIterator struct {
	Event *FundOutOfPositionLog // Event containing the contract specifics and raw log

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
func (it *FundOutOfPositionLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundOutOfPositionLog)
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
		it.Event = new(FundOutOfPositionLog)
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
func (it *FundOutOfPositionLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundOutOfPositionLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundOutOfPositionLog represents a OutOfPositionLog event raised by the Fund contract.
type FundOutOfPositionLog struct {
	Account           common.Address
	TokenAddr         common.Address
	BillNum           *big.Int
	TokenNum          *big.Int
	TokenBrokerageNum *big.Int
	TokenChargeNum    *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOutOfPositionLog is a free log retrieval operation binding the contract event 0xd062038d03821c9b9536916a4b617a4f50c19e298ab93209790adfac0ec047c3.
//
// Solidity: event OutOfPositionLog(address _account, address _tokenAddr, uint256 _billNum, uint256 _tokenNum, uint256 tokenBrokerageNum, uint256 _tokenChargeNum, uint256 _timestamp)
func (_Fund *FundFilterer) FilterOutOfPositionLog(opts *bind.FilterOpts) (*FundOutOfPositionLogIterator, error) {

	logs, sub, err := _Fund.contract.FilterLogs(opts, "OutOfPositionLog")
	if err != nil {
		return nil, err
	}
	return &FundOutOfPositionLogIterator{contract: _Fund.contract, event: "OutOfPositionLog", logs: logs, sub: sub}, nil
}

// WatchOutOfPositionLog is a free log subscription operation binding the contract event 0xd062038d03821c9b9536916a4b617a4f50c19e298ab93209790adfac0ec047c3.
//
// Solidity: event OutOfPositionLog(address _account, address _tokenAddr, uint256 _billNum, uint256 _tokenNum, uint256 tokenBrokerageNum, uint256 _tokenChargeNum, uint256 _timestamp)
func (_Fund *FundFilterer) WatchOutOfPositionLog(opts *bind.WatchOpts, sink chan<- *FundOutOfPositionLog) (event.Subscription, error) {

	logs, sub, err := _Fund.contract.WatchLogs(opts, "OutOfPositionLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundOutOfPositionLog)
				if err := _Fund.contract.UnpackLog(event, "OutOfPositionLog", log); err != nil {
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

// ParseOutOfPositionLog is a log parse operation binding the contract event 0xd062038d03821c9b9536916a4b617a4f50c19e298ab93209790adfac0ec047c3.
//
// Solidity: event OutOfPositionLog(address _account, address _tokenAddr, uint256 _billNum, uint256 _tokenNum, uint256 tokenBrokerageNum, uint256 _tokenChargeNum, uint256 _timestamp)
func (_Fund *FundFilterer) ParseOutOfPositionLog(log types.Log) (*FundOutOfPositionLog, error) {
	event := new(FundOutOfPositionLog)
	if err := _Fund.contract.UnpackLog(event, "OutOfPositionLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FundPositionLogIterator is returned from FilterPositionLog and is used to iterate over the raw logs and unpacked data for PositionLog events raised by the Fund contract.
type FundPositionLogIterator struct {
	Event *FundPositionLog // Event containing the contract specifics and raw log

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
func (it *FundPositionLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundPositionLog)
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
		it.Event = new(FundPositionLog)
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
func (it *FundPositionLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundPositionLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundPositionLog represents a PositionLog event raised by the Fund contract.
type FundPositionLog struct {
	Id        [32]byte
	Addr      common.Address
	EthNum    *big.Int
	UsdcNum   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPositionLog is a free log retrieval operation binding the contract event 0xa14109c4360e407ec70417975c35107b2e1a042b28881702f3e67f95f51e249e.
//
// Solidity: event PositionLog(bytes32 _id, address _addr, uint256 _ethNum, uint256 _usdcNum, uint256 _timestamp)
func (_Fund *FundFilterer) FilterPositionLog(opts *bind.FilterOpts) (*FundPositionLogIterator, error) {

	logs, sub, err := _Fund.contract.FilterLogs(opts, "PositionLog")
	if err != nil {
		return nil, err
	}
	return &FundPositionLogIterator{contract: _Fund.contract, event: "PositionLog", logs: logs, sub: sub}, nil
}

// WatchPositionLog is a free log subscription operation binding the contract event 0xa14109c4360e407ec70417975c35107b2e1a042b28881702f3e67f95f51e249e.
//
// Solidity: event PositionLog(bytes32 _id, address _addr, uint256 _ethNum, uint256 _usdcNum, uint256 _timestamp)
func (_Fund *FundFilterer) WatchPositionLog(opts *bind.WatchOpts, sink chan<- *FundPositionLog) (event.Subscription, error) {

	logs, sub, err := _Fund.contract.WatchLogs(opts, "PositionLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundPositionLog)
				if err := _Fund.contract.UnpackLog(event, "PositionLog", log); err != nil {
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

// ParsePositionLog is a log parse operation binding the contract event 0xa14109c4360e407ec70417975c35107b2e1a042b28881702f3e67f95f51e249e.
//
// Solidity: event PositionLog(bytes32 _id, address _addr, uint256 _ethNum, uint256 _usdcNum, uint256 _timestamp)
func (_Fund *FundFilterer) ParsePositionLog(log types.Log) (*FundPositionLog, error) {
	event := new(FundPositionLog)
	if err := _Fund.contract.UnpackLog(event, "PositionLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FundRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Fund contract.
type FundRoleAdminChangedIterator struct {
	Event *FundRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FundRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundRoleAdminChanged)
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
		it.Event = new(FundRoleAdminChanged)
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
func (it *FundRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundRoleAdminChanged represents a RoleAdminChanged event raised by the Fund contract.
type FundRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Fund *FundFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FundRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Fund.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FundRoleAdminChangedIterator{contract: _Fund.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Fund *FundFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FundRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Fund.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundRoleAdminChanged)
				if err := _Fund.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Fund *FundFilterer) ParseRoleAdminChanged(log types.Log) (*FundRoleAdminChanged, error) {
	event := new(FundRoleAdminChanged)
	if err := _Fund.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FundRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Fund contract.
type FundRoleGrantedIterator struct {
	Event *FundRoleGranted // Event containing the contract specifics and raw log

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
func (it *FundRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundRoleGranted)
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
		it.Event = new(FundRoleGranted)
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
func (it *FundRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundRoleGranted represents a RoleGranted event raised by the Fund contract.
type FundRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fund *FundFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FundRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fund.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FundRoleGrantedIterator{contract: _Fund.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fund *FundFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FundRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fund.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundRoleGranted)
				if err := _Fund.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fund *FundFilterer) ParseRoleGranted(log types.Log) (*FundRoleGranted, error) {
	event := new(FundRoleGranted)
	if err := _Fund.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FundRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Fund contract.
type FundRoleRevokedIterator struct {
	Event *FundRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FundRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundRoleRevoked)
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
		it.Event = new(FundRoleRevoked)
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
func (it *FundRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundRoleRevoked represents a RoleRevoked event raised by the Fund contract.
type FundRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fund *FundFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FundRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fund.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FundRoleRevokedIterator{contract: _Fund.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fund *FundFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FundRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fund.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundRoleRevoked)
				if err := _Fund.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fund *FundFilterer) ParseRoleRevoked(log types.Log) (*FundRoleRevoked, error) {
	event := new(FundRoleRevoked)
	if err := _Fund.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
