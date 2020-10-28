// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// StakingABI is the input ABI used to generate the binding from.
const StakingABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch1Start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epochDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"ManualEpochInit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prevBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"prevMultiplier\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"currentMultiplier\",\"type\":\"uint128\"}],\"name\":\"computeNewMultiplier\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpochMultiplier\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch1Start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"epochIsInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochPoolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochUserBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"manualEpochInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address user, address token) view returns(uint256)
func (_Staking *StakingCaller) BalanceOf(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "balanceOf", user, token)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address user, address token) view returns(uint256)
func (_Staking *StakingSession) BalanceOf(user common.Address, token common.Address) (*big.Int, error) {
	return _Staking.Contract.BalanceOf(&_Staking.CallOpts, user, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address user, address token) view returns(uint256)
func (_Staking *StakingCallerSession) BalanceOf(user common.Address, token common.Address) (*big.Int, error) {
	return _Staking.Contract.BalanceOf(&_Staking.CallOpts, user, token)
}

// ComputeNewMultiplier is a free data retrieval call binding the contract method 0x4be41dba.
//
// Solidity: function computeNewMultiplier(uint256 prevBalance, uint128 prevMultiplier, uint256 amount, uint128 currentMultiplier) pure returns(uint128)
func (_Staking *StakingCaller) ComputeNewMultiplier(opts *bind.CallOpts, prevBalance *big.Int, prevMultiplier *big.Int, amount *big.Int, currentMultiplier *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "computeNewMultiplier", prevBalance, prevMultiplier, amount, currentMultiplier)
	return *ret0, err
}

// ComputeNewMultiplier is a free data retrieval call binding the contract method 0x4be41dba.
//
// Solidity: function computeNewMultiplier(uint256 prevBalance, uint128 prevMultiplier, uint256 amount, uint128 currentMultiplier) pure returns(uint128)
func (_Staking *StakingSession) ComputeNewMultiplier(prevBalance *big.Int, prevMultiplier *big.Int, amount *big.Int, currentMultiplier *big.Int) (*big.Int, error) {
	return _Staking.Contract.ComputeNewMultiplier(&_Staking.CallOpts, prevBalance, prevMultiplier, amount, currentMultiplier)
}

// ComputeNewMultiplier is a free data retrieval call binding the contract method 0x4be41dba.
//
// Solidity: function computeNewMultiplier(uint256 prevBalance, uint128 prevMultiplier, uint256 amount, uint128 currentMultiplier) pure returns(uint128)
func (_Staking *StakingCallerSession) ComputeNewMultiplier(prevBalance *big.Int, prevMultiplier *big.Int, amount *big.Int, currentMultiplier *big.Int) (*big.Int, error) {
	return _Staking.Contract.ComputeNewMultiplier(&_Staking.CallOpts, prevBalance, prevMultiplier, amount, currentMultiplier)
}

// CurrentEpochMultiplier is a free data retrieval call binding the contract method 0xce58a2a8.
//
// Solidity: function currentEpochMultiplier() view returns(uint128)
func (_Staking *StakingCaller) CurrentEpochMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "currentEpochMultiplier")
	return *ret0, err
}

// CurrentEpochMultiplier is a free data retrieval call binding the contract method 0xce58a2a8.
//
// Solidity: function currentEpochMultiplier() view returns(uint128)
func (_Staking *StakingSession) CurrentEpochMultiplier() (*big.Int, error) {
	return _Staking.Contract.CurrentEpochMultiplier(&_Staking.CallOpts)
}

// CurrentEpochMultiplier is a free data retrieval call binding the contract method 0xce58a2a8.
//
// Solidity: function currentEpochMultiplier() view returns(uint128)
func (_Staking *StakingCallerSession) CurrentEpochMultiplier() (*big.Int, error) {
	return _Staking.Contract.CurrentEpochMultiplier(&_Staking.CallOpts)
}

// Epoch1Start is a free data retrieval call binding the contract method 0xf4a4341d.
//
// Solidity: function epoch1Start() view returns(uint256)
func (_Staking *StakingCaller) Epoch1Start(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "epoch1Start")
	return *ret0, err
}

// Epoch1Start is a free data retrieval call binding the contract method 0xf4a4341d.
//
// Solidity: function epoch1Start() view returns(uint256)
func (_Staking *StakingSession) Epoch1Start() (*big.Int, error) {
	return _Staking.Contract.Epoch1Start(&_Staking.CallOpts)
}

// Epoch1Start is a free data retrieval call binding the contract method 0xf4a4341d.
//
// Solidity: function epoch1Start() view returns(uint256)
func (_Staking *StakingCallerSession) Epoch1Start() (*big.Int, error) {
	return _Staking.Contract.Epoch1Start(&_Staking.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Staking *StakingCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "epochDuration")
	return *ret0, err
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Staking *StakingSession) EpochDuration() (*big.Int, error) {
	return _Staking.Contract.EpochDuration(&_Staking.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Staking *StakingCallerSession) EpochDuration() (*big.Int, error) {
	return _Staking.Contract.EpochDuration(&_Staking.CallOpts)
}

// EpochIsInitialized is a free data retrieval call binding the contract method 0xaa579154.
//
// Solidity: function epochIsInitialized(address token, uint128 epochId) view returns(bool)
func (_Staking *StakingCaller) EpochIsInitialized(opts *bind.CallOpts, token common.Address, epochId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "epochIsInitialized", token, epochId)
	return *ret0, err
}

// EpochIsInitialized is a free data retrieval call binding the contract method 0xaa579154.
//
// Solidity: function epochIsInitialized(address token, uint128 epochId) view returns(bool)
func (_Staking *StakingSession) EpochIsInitialized(token common.Address, epochId *big.Int) (bool, error) {
	return _Staking.Contract.EpochIsInitialized(&_Staking.CallOpts, token, epochId)
}

// EpochIsInitialized is a free data retrieval call binding the contract method 0xaa579154.
//
// Solidity: function epochIsInitialized(address token, uint128 epochId) view returns(bool)
func (_Staking *StakingCallerSession) EpochIsInitialized(token common.Address, epochId *big.Int) (bool, error) {
	return _Staking.Contract.EpochIsInitialized(&_Staking.CallOpts, token, epochId)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint128)
func (_Staking *StakingCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "getCurrentEpoch")
	return *ret0, err
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint128)
func (_Staking *StakingSession) GetCurrentEpoch() (*big.Int, error) {
	return _Staking.Contract.GetCurrentEpoch(&_Staking.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint128)
func (_Staking *StakingCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _Staking.Contract.GetCurrentEpoch(&_Staking.CallOpts)
}

// GetEpochPoolSize is a free data retrieval call binding the contract method 0x2ca32d7e.
//
// Solidity: function getEpochPoolSize(address tokenAddress, uint128 epochId) view returns(uint256)
func (_Staking *StakingCaller) GetEpochPoolSize(opts *bind.CallOpts, tokenAddress common.Address, epochId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "getEpochPoolSize", tokenAddress, epochId)
	return *ret0, err
}

// GetEpochPoolSize is a free data retrieval call binding the contract method 0x2ca32d7e.
//
// Solidity: function getEpochPoolSize(address tokenAddress, uint128 epochId) view returns(uint256)
func (_Staking *StakingSession) GetEpochPoolSize(tokenAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _Staking.Contract.GetEpochPoolSize(&_Staking.CallOpts, tokenAddress, epochId)
}

// GetEpochPoolSize is a free data retrieval call binding the contract method 0x2ca32d7e.
//
// Solidity: function getEpochPoolSize(address tokenAddress, uint128 epochId) view returns(uint256)
func (_Staking *StakingCallerSession) GetEpochPoolSize(tokenAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _Staking.Contract.GetEpochPoolSize(&_Staking.CallOpts, tokenAddress, epochId)
}

// GetEpochUserBalance is a free data retrieval call binding the contract method 0x8c028dd0.
//
// Solidity: function getEpochUserBalance(address user, address token, uint128 epochId) view returns(uint256)
func (_Staking *StakingCaller) GetEpochUserBalance(opts *bind.CallOpts, user common.Address, token common.Address, epochId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "getEpochUserBalance", user, token, epochId)
	return *ret0, err
}

// GetEpochUserBalance is a free data retrieval call binding the contract method 0x8c028dd0.
//
// Solidity: function getEpochUserBalance(address user, address token, uint128 epochId) view returns(uint256)
func (_Staking *StakingSession) GetEpochUserBalance(user common.Address, token common.Address, epochId *big.Int) (*big.Int, error) {
	return _Staking.Contract.GetEpochUserBalance(&_Staking.CallOpts, user, token, epochId)
}

// GetEpochUserBalance is a free data retrieval call binding the contract method 0x8c028dd0.
//
// Solidity: function getEpochUserBalance(address user, address token, uint128 epochId) view returns(uint256)
func (_Staking *StakingCallerSession) GetEpochUserBalance(user common.Address, token common.Address, epochId *big.Int) (*big.Int, error) {
	return _Staking.Contract.GetEpochUserBalance(&_Staking.CallOpts, user, token, epochId)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address tokenAddress, uint256 amount) returns()
func (_Staking *StakingTransactor) Deposit(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "deposit", tokenAddress, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address tokenAddress, uint256 amount) returns()
func (_Staking *StakingSession) Deposit(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Deposit(&_Staking.TransactOpts, tokenAddress, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address tokenAddress, uint256 amount) returns()
func (_Staking *StakingTransactorSession) Deposit(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Deposit(&_Staking.TransactOpts, tokenAddress, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address tokenAddress) returns()
func (_Staking *StakingTransactor) EmergencyWithdraw(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "emergencyWithdraw", tokenAddress)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address tokenAddress) returns()
func (_Staking *StakingSession) EmergencyWithdraw(tokenAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.EmergencyWithdraw(&_Staking.TransactOpts, tokenAddress)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address tokenAddress) returns()
func (_Staking *StakingTransactorSession) EmergencyWithdraw(tokenAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.EmergencyWithdraw(&_Staking.TransactOpts, tokenAddress)
}

// ManualEpochInit is a paid mutator transaction binding the contract method 0xea2c38ae.
//
// Solidity: function manualEpochInit(address[] tokens, uint128 epochId) returns()
func (_Staking *StakingTransactor) ManualEpochInit(opts *bind.TransactOpts, tokens []common.Address, epochId *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "manualEpochInit", tokens, epochId)
}

// ManualEpochInit is a paid mutator transaction binding the contract method 0xea2c38ae.
//
// Solidity: function manualEpochInit(address[] tokens, uint128 epochId) returns()
func (_Staking *StakingSession) ManualEpochInit(tokens []common.Address, epochId *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ManualEpochInit(&_Staking.TransactOpts, tokens, epochId)
}

// ManualEpochInit is a paid mutator transaction binding the contract method 0xea2c38ae.
//
// Solidity: function manualEpochInit(address[] tokens, uint128 epochId) returns()
func (_Staking *StakingTransactorSession) ManualEpochInit(tokens []common.Address, epochId *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ManualEpochInit(&_Staking.TransactOpts, tokens, epochId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address tokenAddress, uint256 amount) returns()
func (_Staking *StakingTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdraw", tokenAddress, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address tokenAddress, uint256 amount) returns()
func (_Staking *StakingSession) Withdraw(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts, tokenAddress, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address tokenAddress, uint256 amount) returns()
func (_Staking *StakingTransactorSession) Withdraw(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts, tokenAddress, amount)
}

// StakingDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Staking contract.
type StakingDepositIterator struct {
	Event *StakingDeposit // Event containing the contract specifics and raw log

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
func (it *StakingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDeposit)
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
		it.Event = new(StakingDeposit)
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
func (it *StakingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDeposit represents a Deposit event raised by the Staking contract.
type StakingDeposit struct {
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Staking *StakingFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, tokenAddress []common.Address) (*StakingDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Deposit", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingDepositIterator{contract: _Staking.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Staking *StakingFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StakingDeposit, user []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Deposit", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDeposit)
				if err := _Staking.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Staking *StakingFilterer) ParseDeposit(log types.Log) (*StakingDeposit, error) {
	event := new(StakingDeposit)
	if err := _Staking.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Staking contract.
type StakingEmergencyWithdrawIterator struct {
	Event *StakingEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *StakingEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingEmergencyWithdraw)
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
		it.Event = new(StakingEmergencyWithdraw)
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
func (it *StakingEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingEmergencyWithdraw represents a EmergencyWithdraw event raised by the Staking contract.
type StakingEmergencyWithdraw struct {
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xf24ef89f38eadc1bde50701ad6e4d6d11a2dc24f7cf834a486991f3883328504.
//
// Solidity: event EmergencyWithdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Staking *StakingFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, tokenAddress []common.Address) (*StakingEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingEmergencyWithdrawIterator{contract: _Staking.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xf24ef89f38eadc1bde50701ad6e4d6d11a2dc24f7cf834a486991f3883328504.
//
// Solidity: event EmergencyWithdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Staking *StakingFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *StakingEmergencyWithdraw, user []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingEmergencyWithdraw)
				if err := _Staking.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xf24ef89f38eadc1bde50701ad6e4d6d11a2dc24f7cf834a486991f3883328504.
//
// Solidity: event EmergencyWithdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Staking *StakingFilterer) ParseEmergencyWithdraw(log types.Log) (*StakingEmergencyWithdraw, error) {
	event := new(StakingEmergencyWithdraw)
	if err := _Staking.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManualEpochInitIterator is returned from FilterManualEpochInit and is used to iterate over the raw logs and unpacked data for ManualEpochInit events raised by the Staking contract.
type StakingManualEpochInitIterator struct {
	Event *StakingManualEpochInit // Event containing the contract specifics and raw log

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
func (it *StakingManualEpochInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManualEpochInit)
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
		it.Event = new(StakingManualEpochInit)
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
func (it *StakingManualEpochInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManualEpochInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManualEpochInit represents a ManualEpochInit event raised by the Staking contract.
type StakingManualEpochInit struct {
	Caller  common.Address
	EpochId *big.Int
	Tokens  []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManualEpochInit is a free log retrieval operation binding the contract event 0xb85c32b8d9cecc81feba78646289584a693e9a8afea40ab2fd31efae4408429f.
//
// Solidity: event ManualEpochInit(address indexed caller, uint128 indexed epochId, address[] tokens)
func (_Staking *StakingFilterer) FilterManualEpochInit(opts *bind.FilterOpts, caller []common.Address, epochId []*big.Int) (*StakingManualEpochInitIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ManualEpochInit", callerRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingManualEpochInitIterator{contract: _Staking.contract, event: "ManualEpochInit", logs: logs, sub: sub}, nil
}

// WatchManualEpochInit is a free log subscription operation binding the contract event 0xb85c32b8d9cecc81feba78646289584a693e9a8afea40ab2fd31efae4408429f.
//
// Solidity: event ManualEpochInit(address indexed caller, uint128 indexed epochId, address[] tokens)
func (_Staking *StakingFilterer) WatchManualEpochInit(opts *bind.WatchOpts, sink chan<- *StakingManualEpochInit, caller []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ManualEpochInit", callerRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManualEpochInit)
				if err := _Staking.contract.UnpackLog(event, "ManualEpochInit", log); err != nil {
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

// ParseManualEpochInit is a log parse operation binding the contract event 0xb85c32b8d9cecc81feba78646289584a693e9a8afea40ab2fd31efae4408429f.
//
// Solidity: event ManualEpochInit(address indexed caller, uint128 indexed epochId, address[] tokens)
func (_Staking *StakingFilterer) ParseManualEpochInit(log types.Log) (*StakingManualEpochInit, error) {
	event := new(StakingManualEpochInit)
	if err := _Staking.contract.UnpackLog(event, "ManualEpochInit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Staking contract.
type StakingWithdrawIterator struct {
	Event *StakingWithdraw // Event containing the contract specifics and raw log

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
func (it *StakingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWithdraw)
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
		it.Event = new(StakingWithdraw)
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
func (it *StakingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWithdraw represents a Withdraw event raised by the Staking contract.
type StakingWithdraw struct {
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Staking *StakingFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, tokenAddress []common.Address) (*StakingWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Withdraw", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWithdrawIterator{contract: _Staking.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Staking *StakingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StakingWithdraw, user []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Withdraw", userRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWithdraw)
				if err := _Staking.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed tokenAddress, uint256 amount)
func (_Staking *StakingFilterer) ParseWithdraw(log types.Log) (*StakingWithdraw, error) {
	event := new(StakingWithdraw)
	if err := _Staking.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
