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

// ScheduleABI is the input ABI used to generate the binding from.
const ScheduleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_Rose\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_storeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_USDC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_RosePledge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bill\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wbtc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_VerifierAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"}],\"name\":\"WithdrawLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BillAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IPledge\",\"outputs\":[{\"internalType\":\"contractIRosPledge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPS\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POSITION_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PositionMemberMinValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Rose\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StorageAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VerifierAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WBTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isUsedTag\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"builder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"periodManager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"rosPay\",\"type\":\"bool\"}],\"name\":\"openPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"builder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"periodManager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"rosPay\",\"type\":\"bool\"}],\"name\":\"openPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bill\",\"type\":\"address\"}],\"name\":\"setBillAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fundManageContract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setFundManageAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_ops\",\"type\":\"address\"}],\"name\":\"setOPSAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setPositionMemberMinValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rosPledge\",\"type\":\"address\"}],\"name\":\"setRosPledgeAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ros\",\"type\":\"address\"}],\"name\":\"setRose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setSignAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_Fund2MemAddr\",\"type\":\"address\"}],\"name\":\"setStorageAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdc\",\"type\":\"address\"}],\"name\":\"setUSDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setUniswapAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setVerifierAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"name\":\"setWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"address[2]\",\"name\":\"_contractAddr\",\"type\":\"address[2]\"},{\"internalType\":\"address[]\",\"name\":\"_addr\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenNum\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewardNum\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"billTokenNum\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenChargeNum\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenBrokerageNum\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"surplusNum\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2][2][2]\",\"name\":\"a\",\"type\":\"uint256[2][2][2]\"},{\"internalType\":\"uint256[21]\",\"name\":\"input\",\"type\":\"uint256[21]\"},{\"internalType\":\"bytes32[4]\",\"name\":\"_vrsHash\",\"type\":\"bytes32[4]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Schedule is an auto generated Go binding around an Ethereum contract.
type Schedule struct {
	ScheduleCaller     // Read-only binding to the contract
	ScheduleTransactor // Write-only binding to the contract
	ScheduleFilterer   // Log filterer for contract events
}

// ScheduleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScheduleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScheduleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScheduleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScheduleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScheduleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScheduleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScheduleSession struct {
	Contract     *Schedule         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScheduleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScheduleCallerSession struct {
	Contract *ScheduleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ScheduleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScheduleTransactorSession struct {
	Contract     *ScheduleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ScheduleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScheduleRaw struct {
	Contract *Schedule // Generic contract binding to access the raw methods on
}

// ScheduleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScheduleCallerRaw struct {
	Contract *ScheduleCaller // Generic read-only contract binding to access the raw methods on
}

// ScheduleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScheduleTransactorRaw struct {
	Contract *ScheduleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSchedule creates a new instance of Schedule, bound to a specific deployed contract.
func NewSchedule(address common.Address, backend bind.ContractBackend) (*Schedule, error) {
	contract, err := bindSchedule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Schedule{ScheduleCaller: ScheduleCaller{contract: contract}, ScheduleTransactor: ScheduleTransactor{contract: contract}, ScheduleFilterer: ScheduleFilterer{contract: contract}}, nil
}

// NewScheduleCaller creates a new read-only instance of Schedule, bound to a specific deployed contract.
func NewScheduleCaller(address common.Address, caller bind.ContractCaller) (*ScheduleCaller, error) {
	contract, err := bindSchedule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScheduleCaller{contract: contract}, nil
}

// NewScheduleTransactor creates a new write-only instance of Schedule, bound to a specific deployed contract.
func NewScheduleTransactor(address common.Address, transactor bind.ContractTransactor) (*ScheduleTransactor, error) {
	contract, err := bindSchedule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScheduleTransactor{contract: contract}, nil
}

// NewScheduleFilterer creates a new log filterer instance of Schedule, bound to a specific deployed contract.
func NewScheduleFilterer(address common.Address, filterer bind.ContractFilterer) (*ScheduleFilterer, error) {
	contract, err := bindSchedule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScheduleFilterer{contract: contract}, nil
}

// bindSchedule binds a generic wrapper to an already deployed contract.
func bindSchedule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScheduleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Schedule *ScheduleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Schedule.Contract.ScheduleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Schedule *ScheduleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Schedule.Contract.ScheduleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Schedule *ScheduleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Schedule.Contract.ScheduleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Schedule *ScheduleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Schedule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Schedule *ScheduleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Schedule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Schedule *ScheduleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Schedule.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Schedule *ScheduleCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Schedule *ScheduleSession) ADMINROLE() ([32]byte, error) {
	return _Schedule.Contract.ADMINROLE(&_Schedule.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Schedule *ScheduleCallerSession) ADMINROLE() ([32]byte, error) {
	return _Schedule.Contract.ADMINROLE(&_Schedule.CallOpts)
}

// BillAddress is a free data retrieval call binding the contract method 0x7d266319.
//
// Solidity: function BillAddress() view returns(address)
func (_Schedule *ScheduleCaller) BillAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "BillAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BillAddress is a free data retrieval call binding the contract method 0x7d266319.
//
// Solidity: function BillAddress() view returns(address)
func (_Schedule *ScheduleSession) BillAddress() (common.Address, error) {
	return _Schedule.Contract.BillAddress(&_Schedule.CallOpts)
}

// BillAddress is a free data retrieval call binding the contract method 0x7d266319.
//
// Solidity: function BillAddress() view returns(address)
func (_Schedule *ScheduleCallerSession) BillAddress() (common.Address, error) {
	return _Schedule.Contract.BillAddress(&_Schedule.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Schedule *ScheduleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Schedule *ScheduleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Schedule.Contract.DEFAULTADMINROLE(&_Schedule.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Schedule *ScheduleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Schedule.Contract.DEFAULTADMINROLE(&_Schedule.CallOpts)
}

// IPledge is a free data retrieval call binding the contract method 0x605ebc07.
//
// Solidity: function IPledge() view returns(address)
func (_Schedule *ScheduleCaller) IPledge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "IPledge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IPledge is a free data retrieval call binding the contract method 0x605ebc07.
//
// Solidity: function IPledge() view returns(address)
func (_Schedule *ScheduleSession) IPledge() (common.Address, error) {
	return _Schedule.Contract.IPledge(&_Schedule.CallOpts)
}

// IPledge is a free data retrieval call binding the contract method 0x605ebc07.
//
// Solidity: function IPledge() view returns(address)
func (_Schedule *ScheduleCallerSession) IPledge() (common.Address, error) {
	return _Schedule.Contract.IPledge(&_Schedule.CallOpts)
}

// OPS is a free data retrieval call binding the contract method 0x472d0447.
//
// Solidity: function OPS() view returns(address)
func (_Schedule *ScheduleCaller) OPS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "OPS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OPS is a free data retrieval call binding the contract method 0x472d0447.
//
// Solidity: function OPS() view returns(address)
func (_Schedule *ScheduleSession) OPS() (common.Address, error) {
	return _Schedule.Contract.OPS(&_Schedule.CallOpts)
}

// OPS is a free data retrieval call binding the contract method 0x472d0447.
//
// Solidity: function OPS() view returns(address)
func (_Schedule *ScheduleCallerSession) OPS() (common.Address, error) {
	return _Schedule.Contract.OPS(&_Schedule.CallOpts)
}

// POSITIONROLE is a free data retrieval call binding the contract method 0x767483ea.
//
// Solidity: function POSITION_ROLE() view returns(bytes32)
func (_Schedule *ScheduleCaller) POSITIONROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "POSITION_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POSITIONROLE is a free data retrieval call binding the contract method 0x767483ea.
//
// Solidity: function POSITION_ROLE() view returns(bytes32)
func (_Schedule *ScheduleSession) POSITIONROLE() ([32]byte, error) {
	return _Schedule.Contract.POSITIONROLE(&_Schedule.CallOpts)
}

// POSITIONROLE is a free data retrieval call binding the contract method 0x767483ea.
//
// Solidity: function POSITION_ROLE() view returns(bytes32)
func (_Schedule *ScheduleCallerSession) POSITIONROLE() ([32]byte, error) {
	return _Schedule.Contract.POSITIONROLE(&_Schedule.CallOpts)
}

// PositionMemberMinValue is a free data retrieval call binding the contract method 0xb80653b8.
//
// Solidity: function PositionMemberMinValue() view returns(uint256)
func (_Schedule *ScheduleCaller) PositionMemberMinValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "PositionMemberMinValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PositionMemberMinValue is a free data retrieval call binding the contract method 0xb80653b8.
//
// Solidity: function PositionMemberMinValue() view returns(uint256)
func (_Schedule *ScheduleSession) PositionMemberMinValue() (*big.Int, error) {
	return _Schedule.Contract.PositionMemberMinValue(&_Schedule.CallOpts)
}

// PositionMemberMinValue is a free data retrieval call binding the contract method 0xb80653b8.
//
// Solidity: function PositionMemberMinValue() view returns(uint256)
func (_Schedule *ScheduleCallerSession) PositionMemberMinValue() (*big.Int, error) {
	return _Schedule.Contract.PositionMemberMinValue(&_Schedule.CallOpts)
}

// Rose is a free data retrieval call binding the contract method 0x3f5a3474.
//
// Solidity: function Rose() view returns(address)
func (_Schedule *ScheduleCaller) Rose(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "Rose")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rose is a free data retrieval call binding the contract method 0x3f5a3474.
//
// Solidity: function Rose() view returns(address)
func (_Schedule *ScheduleSession) Rose() (common.Address, error) {
	return _Schedule.Contract.Rose(&_Schedule.CallOpts)
}

// Rose is a free data retrieval call binding the contract method 0x3f5a3474.
//
// Solidity: function Rose() view returns(address)
func (_Schedule *ScheduleCallerSession) Rose() (common.Address, error) {
	return _Schedule.Contract.Rose(&_Schedule.CallOpts)
}

// StorageAddress is a free data retrieval call binding the contract method 0xa65ab2fd.
//
// Solidity: function StorageAddress() view returns(address)
func (_Schedule *ScheduleCaller) StorageAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "StorageAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageAddress is a free data retrieval call binding the contract method 0xa65ab2fd.
//
// Solidity: function StorageAddress() view returns(address)
func (_Schedule *ScheduleSession) StorageAddress() (common.Address, error) {
	return _Schedule.Contract.StorageAddress(&_Schedule.CallOpts)
}

// StorageAddress is a free data retrieval call binding the contract method 0xa65ab2fd.
//
// Solidity: function StorageAddress() view returns(address)
func (_Schedule *ScheduleCallerSession) StorageAddress() (common.Address, error) {
	return _Schedule.Contract.StorageAddress(&_Schedule.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Schedule *ScheduleCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Schedule *ScheduleSession) USDC() (common.Address, error) {
	return _Schedule.Contract.USDC(&_Schedule.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Schedule *ScheduleCallerSession) USDC() (common.Address, error) {
	return _Schedule.Contract.USDC(&_Schedule.CallOpts)
}

// VerifierAddress is a free data retrieval call binding the contract method 0x874ed5b5.
//
// Solidity: function VerifierAddress() view returns(address)
func (_Schedule *ScheduleCaller) VerifierAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "VerifierAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifierAddress is a free data retrieval call binding the contract method 0x874ed5b5.
//
// Solidity: function VerifierAddress() view returns(address)
func (_Schedule *ScheduleSession) VerifierAddress() (common.Address, error) {
	return _Schedule.Contract.VerifierAddress(&_Schedule.CallOpts)
}

// VerifierAddress is a free data retrieval call binding the contract method 0x874ed5b5.
//
// Solidity: function VerifierAddress() view returns(address)
func (_Schedule *ScheduleCallerSession) VerifierAddress() (common.Address, error) {
	return _Schedule.Contract.VerifierAddress(&_Schedule.CallOpts)
}

// WBTC is a free data retrieval call binding the contract method 0x4dede3de.
//
// Solidity: function WBTC() view returns(address)
func (_Schedule *ScheduleCaller) WBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "WBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WBTC is a free data retrieval call binding the contract method 0x4dede3de.
//
// Solidity: function WBTC() view returns(address)
func (_Schedule *ScheduleSession) WBTC() (common.Address, error) {
	return _Schedule.Contract.WBTC(&_Schedule.CallOpts)
}

// WBTC is a free data retrieval call binding the contract method 0x4dede3de.
//
// Solidity: function WBTC() view returns(address)
func (_Schedule *ScheduleCallerSession) WBTC() (common.Address, error) {
	return _Schedule.Contract.WBTC(&_Schedule.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Schedule *ScheduleCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Schedule *ScheduleSession) WETH() (common.Address, error) {
	return _Schedule.Contract.WETH(&_Schedule.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Schedule *ScheduleCallerSession) WETH() (common.Address, error) {
	return _Schedule.Contract.WETH(&_Schedule.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Schedule *ScheduleCaller) WITHDRAWROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "WITHDRAW_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Schedule *ScheduleSession) WITHDRAWROLE() ([32]byte, error) {
	return _Schedule.Contract.WITHDRAWROLE(&_Schedule.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Schedule *ScheduleCallerSession) WITHDRAWROLE() ([32]byte, error) {
	return _Schedule.Contract.WITHDRAWROLE(&_Schedule.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Schedule *ScheduleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Schedule *ScheduleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Schedule.Contract.GetRoleAdmin(&_Schedule.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Schedule *ScheduleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Schedule.Contract.GetRoleAdmin(&_Schedule.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Schedule *ScheduleCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Schedule *ScheduleSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Schedule.Contract.GetRoleMember(&_Schedule.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Schedule *ScheduleCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Schedule.Contract.GetRoleMember(&_Schedule.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Schedule *ScheduleCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Schedule *ScheduleSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Schedule.Contract.GetRoleMemberCount(&_Schedule.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Schedule *ScheduleCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Schedule.Contract.GetRoleMemberCount(&_Schedule.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Schedule *ScheduleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Schedule *ScheduleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Schedule.Contract.HasRole(&_Schedule.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Schedule *ScheduleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Schedule.Contract.HasRole(&_Schedule.CallOpts, role, account)
}

// IsUsedTag is a free data retrieval call binding the contract method 0x3c390cb2.
//
// Solidity: function isUsedTag(bytes32 ) view returns(bool)
func (_Schedule *ScheduleCaller) IsUsedTag(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "isUsedTag", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUsedTag is a free data retrieval call binding the contract method 0x3c390cb2.
//
// Solidity: function isUsedTag(bytes32 ) view returns(bool)
func (_Schedule *ScheduleSession) IsUsedTag(arg0 [32]byte) (bool, error) {
	return _Schedule.Contract.IsUsedTag(&_Schedule.CallOpts, arg0)
}

// IsUsedTag is a free data retrieval call binding the contract method 0x3c390cb2.
//
// Solidity: function isUsedTag(bytes32 ) view returns(bool)
func (_Schedule *ScheduleCallerSession) IsUsedTag(arg0 [32]byte) (bool, error) {
	return _Schedule.Contract.IsUsedTag(&_Schedule.CallOpts, arg0)
}

// SignAddress is a free data retrieval call binding the contract method 0x0682bdbc.
//
// Solidity: function signAddress() view returns(address)
func (_Schedule *ScheduleCaller) SignAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "signAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignAddress is a free data retrieval call binding the contract method 0x0682bdbc.
//
// Solidity: function signAddress() view returns(address)
func (_Schedule *ScheduleSession) SignAddress() (common.Address, error) {
	return _Schedule.Contract.SignAddress(&_Schedule.CallOpts)
}

// SignAddress is a free data retrieval call binding the contract method 0x0682bdbc.
//
// Solidity: function signAddress() view returns(address)
func (_Schedule *ScheduleCallerSession) SignAddress() (common.Address, error) {
	return _Schedule.Contract.SignAddress(&_Schedule.CallOpts)
}

// UniswapAddress is a free data retrieval call binding the contract method 0x0e2feb05.
//
// Solidity: function uniswapAddress() view returns(address)
func (_Schedule *ScheduleCaller) UniswapAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Schedule.contract.Call(opts, &out, "uniswapAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapAddress is a free data retrieval call binding the contract method 0x0e2feb05.
//
// Solidity: function uniswapAddress() view returns(address)
func (_Schedule *ScheduleSession) UniswapAddress() (common.Address, error) {
	return _Schedule.Contract.UniswapAddress(&_Schedule.CallOpts)
}

// UniswapAddress is a free data retrieval call binding the contract method 0x0e2feb05.
//
// Solidity: function uniswapAddress() view returns(address)
func (_Schedule *ScheduleCallerSession) UniswapAddress() (common.Address, error) {
	return _Schedule.Contract.UniswapAddress(&_Schedule.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Schedule *ScheduleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Schedule *ScheduleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.GrantRole(&_Schedule.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Schedule *ScheduleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.GrantRole(&_Schedule.TransactOpts, role, account)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x0ac545f1.
//
// Solidity: function openPosition(address inviter, address builder, address periodManager, bool rosPay) payable returns(uint256)
func (_Schedule *ScheduleTransactor) OpenPosition(opts *bind.TransactOpts, inviter common.Address, builder common.Address, periodManager common.Address, rosPay bool) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "openPosition", inviter, builder, periodManager, rosPay)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x0ac545f1.
//
// Solidity: function openPosition(address inviter, address builder, address periodManager, bool rosPay) payable returns(uint256)
func (_Schedule *ScheduleSession) OpenPosition(inviter common.Address, builder common.Address, periodManager common.Address, rosPay bool) (*types.Transaction, error) {
	return _Schedule.Contract.OpenPosition(&_Schedule.TransactOpts, inviter, builder, periodManager, rosPay)
}

// OpenPosition is a paid mutator transaction binding the contract method 0x0ac545f1.
//
// Solidity: function openPosition(address inviter, address builder, address periodManager, bool rosPay) payable returns(uint256)
func (_Schedule *ScheduleTransactorSession) OpenPosition(inviter common.Address, builder common.Address, periodManager common.Address, rosPay bool) (*types.Transaction, error) {
	return _Schedule.Contract.OpenPosition(&_Schedule.TransactOpts, inviter, builder, periodManager, rosPay)
}

// OpenPosition0 is a paid mutator transaction binding the contract method 0x1222c146.
//
// Solidity: function openPosition(address builder, address periodManager, bool rosPay) payable returns(uint256)
func (_Schedule *ScheduleTransactor) OpenPosition0(opts *bind.TransactOpts, builder common.Address, periodManager common.Address, rosPay bool) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "openPosition0", builder, periodManager, rosPay)
}

// OpenPosition0 is a paid mutator transaction binding the contract method 0x1222c146.
//
// Solidity: function openPosition(address builder, address periodManager, bool rosPay) payable returns(uint256)
func (_Schedule *ScheduleSession) OpenPosition0(builder common.Address, periodManager common.Address, rosPay bool) (*types.Transaction, error) {
	return _Schedule.Contract.OpenPosition0(&_Schedule.TransactOpts, builder, periodManager, rosPay)
}

// OpenPosition0 is a paid mutator transaction binding the contract method 0x1222c146.
//
// Solidity: function openPosition(address builder, address periodManager, bool rosPay) payable returns(uint256)
func (_Schedule *ScheduleTransactorSession) OpenPosition0(builder common.Address, periodManager common.Address, rosPay bool) (*types.Transaction, error) {
	return _Schedule.Contract.OpenPosition0(&_Schedule.TransactOpts, builder, periodManager, rosPay)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Schedule *ScheduleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Schedule *ScheduleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.RenounceRole(&_Schedule.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Schedule *ScheduleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.RenounceRole(&_Schedule.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Schedule *ScheduleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Schedule *ScheduleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.RevokeRole(&_Schedule.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Schedule *ScheduleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.RevokeRole(&_Schedule.TransactOpts, role, account)
}

// SetBillAddress is a paid mutator transaction binding the contract method 0xf595576e.
//
// Solidity: function setBillAddress(address _bill) returns()
func (_Schedule *ScheduleTransactor) SetBillAddress(opts *bind.TransactOpts, _bill common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "setBillAddress", _bill)
}

// SetBillAddress is a paid mutator transaction binding the contract method 0xf595576e.
//
// Solidity: function setBillAddress(address _bill) returns()
func (_Schedule *ScheduleSession) SetBillAddress(_bill common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetBillAddress(&_Schedule.TransactOpts, _bill)
}

// SetBillAddress is a paid mutator transaction binding the contract method 0xf595576e.
//
// Solidity: function setBillAddress(address _bill) returns()
func (_Schedule *ScheduleTransactorSession) SetBillAddress(_bill common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetBillAddress(&_Schedule.TransactOpts, _bill)
}

// SetFundManageAddress is a paid mutator transaction binding the contract method 0xaf9d070c.
//
// Solidity: function setFundManageAddress(address _fundManageContract, bool _enable) returns()
func (_Schedule *ScheduleTransactor) SetFundManageAddress(opts *bind.TransactOpts, _fundManageContract common.Address, _enable bool) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "setFundManageAddress", _fundManageContract, _enable)
}

// SetFundManageAddress is a paid mutator transaction binding the contract method 0xaf9d070c.
//
// Solidity: function setFundManageAddress(address _fundManageContract, bool _enable) returns()
func (_Schedule *ScheduleSession) SetFundManageAddress(_fundManageContract common.Address, _enable bool) (*types.Transaction, error) {
	return _Schedule.Contract.SetFundManageAddress(&_Schedule.TransactOpts, _fundManageContract, _enable)
}

// SetFundManageAddress is a paid mutator transaction binding the contract method 0xaf9d070c.
//
// Solidity: function setFundManageAddress(address _fundManageContract, bool _enable) returns()
func (_Schedule *ScheduleTransactorSession) SetFundManageAddress(_fundManageContract common.Address, _enable bool) (*types.Transaction, error) {
	return _Schedule.Contract.SetFundManageAddress(&_Schedule.TransactOpts, _fundManageContract, _enable)
}

// SetOPSAddress is a paid mutator transaction binding the contract method 0x2358a64d.
//
// Solidity: function setOPSAddress(address _ops) returns()
func (_Schedule *ScheduleTransactor) SetOPSAddress(opts *bind.TransactOpts, _ops common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "setOPSAddress", _ops)
}

// SetOPSAddress is a paid mutator transaction binding the contract method 0x2358a64d.
//
// Solidity: function setOPSAddress(address _ops) returns()
func (_Schedule *ScheduleSession) SetOPSAddress(_ops common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetOPSAddress(&_Schedule.TransactOpts, _ops)
}

// SetOPSAddress is a paid mutator transaction binding the contract method 0x2358a64d.
//
// Solidity: function setOPSAddress(address _ops) returns()
func (_Schedule *ScheduleTransactorSession) SetOPSAddress(_ops common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetOPSAddress(&_Schedule.TransactOpts, _ops)
}

// SetPositionMemberMinValue is a paid mutator transaction binding the contract method 0xb4ed3c1e.
//
// Solidity: function setPositionMemberMinValue(uint256 _value) returns()
func (_Schedule *ScheduleTransactor) SetPositionMemberMinValue(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "setPositionMemberMinValue", _value)
}

// SetPositionMemberMinValue is a paid mutator transaction binding the contract method 0xb4ed3c1e.
//
// Solidity: function setPositionMemberMinValue(uint256 _value) returns()
func (_Schedule *ScheduleSession) SetPositionMemberMinValue(_value *big.Int) (*types.Transaction, error) {
	return _Schedule.Contract.SetPositionMemberMinValue(&_Schedule.TransactOpts, _value)
}

// SetPositionMemberMinValue is a paid mutator transaction binding the contract method 0xb4ed3c1e.
//
// Solidity: function setPositionMemberMinValue(uint256 _value) returns()
func (_Schedule *ScheduleTransactorSession) SetPositionMemberMinValue(_value *big.Int) (*types.Transaction, error) {
	return _Schedule.Contract.SetPositionMemberMinValue(&_Schedule.TransactOpts, _value)
}

// SetRosPledgeAddr is a paid mutator transaction binding the contract method 0x9fda3eb2.
//
// Solidity: function setRosPledgeAddr(address _rosPledge) returns()
func (_Schedule *ScheduleTransactor) SetRosPledgeAddr(opts *bind.TransactOpts, _rosPledge common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "setRosPledgeAddr", _rosPledge)
}

// SetRosPledgeAddr is a paid mutator transaction binding the contract method 0x9fda3eb2.
//
// Solidity: function setRosPledgeAddr(address _rosPledge) returns()
func (_Schedule *ScheduleSession) SetRosPledgeAddr(_rosPledge common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetRosPledgeAddr(&_Schedule.TransactOpts, _rosPledge)
}

// SetRosPledgeAddr is a paid mutator transaction binding the contract method 0x9fda3eb2.
//
// Solidity: function setRosPledgeAddr(address _rosPledge) returns()
func (_Schedule *ScheduleTransactorSession) SetRosPledgeAddr(_rosPledge common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetRosPledgeAddr(&_Schedule.TransactOpts, _rosPledge)
}

// SetRose is a paid mutator transaction binding the contract method 0x6908362b.
//
// Solidity: function setRose(address _ros) returns()
func (_Schedule *ScheduleTransactor) SetRose(opts *bind.TransactOpts, _ros common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "setRose", _ros)
}

// SetRose is a paid mutator transaction binding the contract method 0x6908362b.
//
// Solidity: function setRose(address _ros) returns()
func (_Schedule *ScheduleSession) SetRose(_ros common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetRose(&_Schedule.TransactOpts, _ros)
}

// SetRose is a paid mutator transaction binding the contract method 0x6908362b.
//
// Solidity: function setRose(address _ros) returns()
func (_Schedule *ScheduleTransactorSession) SetRose(_ros common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetRose(&_Schedule.TransactOpts, _ros)
}

// SetSignAddress is a paid mutator transaction binding the contract method 0x15137045.
//
// Solidity: function setSignAddress(address _addr) returns()
func (_Schedule *ScheduleTransactor) SetSignAddress(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "setSignAddress", _addr)
}

// SetSignAddress is a paid mutator transaction binding the contract method 0x15137045.
//
// Solidity: function setSignAddress(address _addr) returns()
func (_Schedule *ScheduleSession) SetSignAddress(_addr common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetSignAddress(&_Schedule.TransactOpts, _addr)
}

// SetSignAddress is a paid mutator transaction binding the contract method 0x15137045.
//
// Solidity: function setSignAddress(address _addr) returns()
func (_Schedule *ScheduleTransactorSession) SetSignAddress(_addr common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetSignAddress(&_Schedule.TransactOpts, _addr)
}

// SetStorageAddr is a paid mutator transaction binding the contract method 0xa6fa0b7b.
//
// Solidity: function setStorageAddr(address _Fund2MemAddr) returns()
func (_Schedule *ScheduleTransactor) SetStorageAddr(opts *bind.TransactOpts, _Fund2MemAddr common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "setStorageAddr", _Fund2MemAddr)
}

// SetStorageAddr is a paid mutator transaction binding the contract method 0xa6fa0b7b.
//
// Solidity: function setStorageAddr(address _Fund2MemAddr) returns()
func (_Schedule *ScheduleSession) SetStorageAddr(_Fund2MemAddr common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetStorageAddr(&_Schedule.TransactOpts, _Fund2MemAddr)
}

// SetStorageAddr is a paid mutator transaction binding the contract method 0xa6fa0b7b.
//
// Solidity: function setStorageAddr(address _Fund2MemAddr) returns()
func (_Schedule *ScheduleTransactorSession) SetStorageAddr(_Fund2MemAddr common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetStorageAddr(&_Schedule.TransactOpts, _Fund2MemAddr)
}

// SetUSDC is a paid mutator transaction binding the contract method 0xb3e089a2.
//
// Solidity: function setUSDC(address _usdc) returns()
func (_Schedule *ScheduleTransactor) SetUSDC(opts *bind.TransactOpts, _usdc common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "setUSDC", _usdc)
}

// SetUSDC is a paid mutator transaction binding the contract method 0xb3e089a2.
//
// Solidity: function setUSDC(address _usdc) returns()
func (_Schedule *ScheduleSession) SetUSDC(_usdc common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetUSDC(&_Schedule.TransactOpts, _usdc)
}

// SetUSDC is a paid mutator transaction binding the contract method 0xb3e089a2.
//
// Solidity: function setUSDC(address _usdc) returns()
func (_Schedule *ScheduleTransactorSession) SetUSDC(_usdc common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetUSDC(&_Schedule.TransactOpts, _usdc)
}

// SetUniswapAddress is a paid mutator transaction binding the contract method 0x7884e7c6.
//
// Solidity: function setUniswapAddress(address _addr) returns()
func (_Schedule *ScheduleTransactor) SetUniswapAddress(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "setUniswapAddress", _addr)
}

// SetUniswapAddress is a paid mutator transaction binding the contract method 0x7884e7c6.
//
// Solidity: function setUniswapAddress(address _addr) returns()
func (_Schedule *ScheduleSession) SetUniswapAddress(_addr common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetUniswapAddress(&_Schedule.TransactOpts, _addr)
}

// SetUniswapAddress is a paid mutator transaction binding the contract method 0x7884e7c6.
//
// Solidity: function setUniswapAddress(address _addr) returns()
func (_Schedule *ScheduleTransactorSession) SetUniswapAddress(_addr common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetUniswapAddress(&_Schedule.TransactOpts, _addr)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0x17e95526.
//
// Solidity: function setVerifierAddress(address _addr) returns()
func (_Schedule *ScheduleTransactor) SetVerifierAddress(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "setVerifierAddress", _addr)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0x17e95526.
//
// Solidity: function setVerifierAddress(address _addr) returns()
func (_Schedule *ScheduleSession) SetVerifierAddress(_addr common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetVerifierAddress(&_Schedule.TransactOpts, _addr)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0x17e95526.
//
// Solidity: function setVerifierAddress(address _addr) returns()
func (_Schedule *ScheduleTransactorSession) SetVerifierAddress(_addr common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetVerifierAddress(&_Schedule.TransactOpts, _addr)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address _weth) returns()
func (_Schedule *ScheduleTransactor) SetWETH(opts *bind.TransactOpts, _weth common.Address) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "setWETH", _weth)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address _weth) returns()
func (_Schedule *ScheduleSession) SetWETH(_weth common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetWETH(&_Schedule.TransactOpts, _weth)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address _weth) returns()
func (_Schedule *ScheduleTransactorSession) SetWETH(_weth common.Address) (*types.Transaction, error) {
	return _Schedule.Contract.SetWETH(&_Schedule.TransactOpts, _weth)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ff54a04.
//
// Solidity: function withdraw(string _tag, address[2] _contractAddr, address[] _addr, uint256[] tokenNum, uint256[] rewardNum, uint256[] billTokenNum, uint256[] tokenChargeNum, uint256[] tokenBrokerageNum, uint256[] surplusNum, uint256[2][2][2] a, uint256[21] input, bytes32[4] _vrsHash) returns()
func (_Schedule *ScheduleTransactor) Withdraw(opts *bind.TransactOpts, _tag string, _contractAddr [2]common.Address, _addr []common.Address, tokenNum []*big.Int, rewardNum []*big.Int, billTokenNum []*big.Int, tokenChargeNum []*big.Int, tokenBrokerageNum []*big.Int, surplusNum []*big.Int, a [2][2][2]*big.Int, input [21]*big.Int, _vrsHash [4][32]byte) (*types.Transaction, error) {
	return _Schedule.contract.Transact(opts, "withdraw", _tag, _contractAddr, _addr, tokenNum, rewardNum, billTokenNum, tokenChargeNum, tokenBrokerageNum, surplusNum, a, input, _vrsHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ff54a04.
//
// Solidity: function withdraw(string _tag, address[2] _contractAddr, address[] _addr, uint256[] tokenNum, uint256[] rewardNum, uint256[] billTokenNum, uint256[] tokenChargeNum, uint256[] tokenBrokerageNum, uint256[] surplusNum, uint256[2][2][2] a, uint256[21] input, bytes32[4] _vrsHash) returns()
func (_Schedule *ScheduleSession) Withdraw(_tag string, _contractAddr [2]common.Address, _addr []common.Address, tokenNum []*big.Int, rewardNum []*big.Int, billTokenNum []*big.Int, tokenChargeNum []*big.Int, tokenBrokerageNum []*big.Int, surplusNum []*big.Int, a [2][2][2]*big.Int, input [21]*big.Int, _vrsHash [4][32]byte) (*types.Transaction, error) {
	return _Schedule.Contract.Withdraw(&_Schedule.TransactOpts, _tag, _contractAddr, _addr, tokenNum, rewardNum, billTokenNum, tokenChargeNum, tokenBrokerageNum, surplusNum, a, input, _vrsHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5ff54a04.
//
// Solidity: function withdraw(string _tag, address[2] _contractAddr, address[] _addr, uint256[] tokenNum, uint256[] rewardNum, uint256[] billTokenNum, uint256[] tokenChargeNum, uint256[] tokenBrokerageNum, uint256[] surplusNum, uint256[2][2][2] a, uint256[21] input, bytes32[4] _vrsHash) returns()
func (_Schedule *ScheduleTransactorSession) Withdraw(_tag string, _contractAddr [2]common.Address, _addr []common.Address, tokenNum []*big.Int, rewardNum []*big.Int, billTokenNum []*big.Int, tokenChargeNum []*big.Int, tokenBrokerageNum []*big.Int, surplusNum []*big.Int, a [2][2][2]*big.Int, input [21]*big.Int, _vrsHash [4][32]byte) (*types.Transaction, error) {
	return _Schedule.Contract.Withdraw(&_Schedule.TransactOpts, _tag, _contractAddr, _addr, tokenNum, rewardNum, billTokenNum, tokenChargeNum, tokenBrokerageNum, surplusNum, a, input, _vrsHash)
}

// ScheduleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Schedule contract.
type ScheduleRoleAdminChangedIterator struct {
	Event *ScheduleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ScheduleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScheduleRoleAdminChanged)
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
		it.Event = new(ScheduleRoleAdminChanged)
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
func (it *ScheduleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScheduleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScheduleRoleAdminChanged represents a RoleAdminChanged event raised by the Schedule contract.
type ScheduleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Schedule *ScheduleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ScheduleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Schedule.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ScheduleRoleAdminChangedIterator{contract: _Schedule.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Schedule *ScheduleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ScheduleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Schedule.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScheduleRoleAdminChanged)
				if err := _Schedule.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Schedule *ScheduleFilterer) ParseRoleAdminChanged(log types.Log) (*ScheduleRoleAdminChanged, error) {
	event := new(ScheduleRoleAdminChanged)
	if err := _Schedule.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScheduleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Schedule contract.
type ScheduleRoleGrantedIterator struct {
	Event *ScheduleRoleGranted // Event containing the contract specifics and raw log

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
func (it *ScheduleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScheduleRoleGranted)
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
		it.Event = new(ScheduleRoleGranted)
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
func (it *ScheduleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScheduleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScheduleRoleGranted represents a RoleGranted event raised by the Schedule contract.
type ScheduleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Schedule *ScheduleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ScheduleRoleGrantedIterator, error) {

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

	logs, sub, err := _Schedule.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ScheduleRoleGrantedIterator{contract: _Schedule.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Schedule *ScheduleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ScheduleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Schedule.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScheduleRoleGranted)
				if err := _Schedule.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Schedule *ScheduleFilterer) ParseRoleGranted(log types.Log) (*ScheduleRoleGranted, error) {
	event := new(ScheduleRoleGranted)
	if err := _Schedule.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScheduleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Schedule contract.
type ScheduleRoleRevokedIterator struct {
	Event *ScheduleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ScheduleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScheduleRoleRevoked)
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
		it.Event = new(ScheduleRoleRevoked)
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
func (it *ScheduleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScheduleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScheduleRoleRevoked represents a RoleRevoked event raised by the Schedule contract.
type ScheduleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Schedule *ScheduleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ScheduleRoleRevokedIterator, error) {

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

	logs, sub, err := _Schedule.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ScheduleRoleRevokedIterator{contract: _Schedule.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Schedule *ScheduleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ScheduleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Schedule.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScheduleRoleRevoked)
				if err := _Schedule.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Schedule *ScheduleFilterer) ParseRoleRevoked(log types.Log) (*ScheduleRoleRevoked, error) {
	event := new(ScheduleRoleRevoked)
	if err := _Schedule.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScheduleWithdrawLogIterator is returned from FilterWithdrawLog and is used to iterate over the raw logs and unpacked data for WithdrawLog events raised by the Schedule contract.
type ScheduleWithdrawLogIterator struct {
	Event *ScheduleWithdrawLog // Event containing the contract specifics and raw log

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
func (it *ScheduleWithdrawLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScheduleWithdrawLog)
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
		it.Event = new(ScheduleWithdrawLog)
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
func (it *ScheduleWithdrawLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScheduleWithdrawLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScheduleWithdrawLog represents a WithdrawLog event raised by the Schedule contract.
type ScheduleWithdrawLog struct {
	Tag string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawLog is a free log retrieval operation binding the contract event 0xc44cc4edd3e805aa73a169258442d684ceef076329aa12a58dcec2076c5dbf7a.
//
// Solidity: event WithdrawLog(string _tag)
func (_Schedule *ScheduleFilterer) FilterWithdrawLog(opts *bind.FilterOpts) (*ScheduleWithdrawLogIterator, error) {

	logs, sub, err := _Schedule.contract.FilterLogs(opts, "WithdrawLog")
	if err != nil {
		return nil, err
	}
	return &ScheduleWithdrawLogIterator{contract: _Schedule.contract, event: "WithdrawLog", logs: logs, sub: sub}, nil
}

// WatchWithdrawLog is a free log subscription operation binding the contract event 0xc44cc4edd3e805aa73a169258442d684ceef076329aa12a58dcec2076c5dbf7a.
//
// Solidity: event WithdrawLog(string _tag)
func (_Schedule *ScheduleFilterer) WatchWithdrawLog(opts *bind.WatchOpts, sink chan<- *ScheduleWithdrawLog) (event.Subscription, error) {

	logs, sub, err := _Schedule.contract.WatchLogs(opts, "WithdrawLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScheduleWithdrawLog)
				if err := _Schedule.contract.UnpackLog(event, "WithdrawLog", log); err != nil {
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

// ParseWithdrawLog is a log parse operation binding the contract event 0xc44cc4edd3e805aa73a169258442d684ceef076329aa12a58dcec2076c5dbf7a.
//
// Solidity: event WithdrawLog(string _tag)
func (_Schedule *ScheduleFilterer) ParseWithdrawLog(log types.Log) (*ScheduleWithdrawLog, error) {
	event := new(ScheduleWithdrawLog)
	if err := _Schedule.contract.UnpackLog(event, "WithdrawLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
