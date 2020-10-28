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

// YieldFarmLPABI is the input ABI used to generate the binding from.
const YieldFarmLPABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bondTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniLP\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochsHarvested\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"}],\"name\":\"MassHarvest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NR_OF_EPOCHS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_DISTRIBUTED_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getPoolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"harvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInitializedEpoch\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massHarvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLastEpochIdHarvested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// YieldFarmLP is an auto generated Go binding around an Ethereum contract.
type YieldFarmLP struct {
	YieldFarmLPCaller     // Read-only binding to the contract
	YieldFarmLPTransactor // Write-only binding to the contract
	YieldFarmLPFilterer   // Log filterer for contract events
}

// YieldFarmLPCaller is an auto generated read-only Go binding around an Ethereum contract.
type YieldFarmLPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldFarmLPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YieldFarmLPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldFarmLPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YieldFarmLPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldFarmLPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YieldFarmLPSession struct {
	Contract     *YieldFarmLP      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldFarmLPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YieldFarmLPCallerSession struct {
	Contract *YieldFarmLPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// YieldFarmLPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YieldFarmLPTransactorSession struct {
	Contract     *YieldFarmLPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YieldFarmLPRaw is an auto generated low-level Go binding around an Ethereum contract.
type YieldFarmLPRaw struct {
	Contract *YieldFarmLP // Generic contract binding to access the raw methods on
}

// YieldFarmLPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YieldFarmLPCallerRaw struct {
	Contract *YieldFarmLPCaller // Generic read-only contract binding to access the raw methods on
}

// YieldFarmLPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YieldFarmLPTransactorRaw struct {
	Contract *YieldFarmLPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYieldFarmLP creates a new instance of YieldFarmLP, bound to a specific deployed contract.
func NewYieldFarmLP(address common.Address, backend bind.ContractBackend) (*YieldFarmLP, error) {
	contract, err := bindYieldFarmLP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YieldFarmLP{YieldFarmLPCaller: YieldFarmLPCaller{contract: contract}, YieldFarmLPTransactor: YieldFarmLPTransactor{contract: contract}, YieldFarmLPFilterer: YieldFarmLPFilterer{contract: contract}}, nil
}

// NewYieldFarmLPCaller creates a new read-only instance of YieldFarmLP, bound to a specific deployed contract.
func NewYieldFarmLPCaller(address common.Address, caller bind.ContractCaller) (*YieldFarmLPCaller, error) {
	contract, err := bindYieldFarmLP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YieldFarmLPCaller{contract: contract}, nil
}

// NewYieldFarmLPTransactor creates a new write-only instance of YieldFarmLP, bound to a specific deployed contract.
func NewYieldFarmLPTransactor(address common.Address, transactor bind.ContractTransactor) (*YieldFarmLPTransactor, error) {
	contract, err := bindYieldFarmLP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YieldFarmLPTransactor{contract: contract}, nil
}

// NewYieldFarmLPFilterer creates a new log filterer instance of YieldFarmLP, bound to a specific deployed contract.
func NewYieldFarmLPFilterer(address common.Address, filterer bind.ContractFilterer) (*YieldFarmLPFilterer, error) {
	contract, err := bindYieldFarmLP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YieldFarmLPFilterer{contract: contract}, nil
}

// bindYieldFarmLP binds a generic wrapper to an already deployed contract.
func bindYieldFarmLP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YieldFarmLPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YieldFarmLP *YieldFarmLPRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _YieldFarmLP.Contract.YieldFarmLPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YieldFarmLP *YieldFarmLPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldFarmLP.Contract.YieldFarmLPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YieldFarmLP *YieldFarmLPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YieldFarmLP.Contract.YieldFarmLPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YieldFarmLP *YieldFarmLPCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _YieldFarmLP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YieldFarmLP *YieldFarmLPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldFarmLP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YieldFarmLP *YieldFarmLPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YieldFarmLP.Contract.contract.Transact(opts, method, params...)
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCaller) NROFEPOCHS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarmLP.contract.Call(opts, out, "NR_OF_EPOCHS")
	return *ret0, err
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPSession) NROFEPOCHS() (*big.Int, error) {
	return _YieldFarmLP.Contract.NROFEPOCHS(&_YieldFarmLP.CallOpts)
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCallerSession) NROFEPOCHS() (*big.Int, error) {
	return _YieldFarmLP.Contract.NROFEPOCHS(&_YieldFarmLP.CallOpts)
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCaller) TOTALDISTRIBUTEDAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarmLP.contract.Call(opts, out, "TOTAL_DISTRIBUTED_AMOUNT")
	return *ret0, err
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPSession) TOTALDISTRIBUTEDAMOUNT() (*big.Int, error) {
	return _YieldFarmLP.Contract.TOTALDISTRIBUTEDAMOUNT(&_YieldFarmLP.CallOpts)
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCallerSession) TOTALDISTRIBUTEDAMOUNT() (*big.Int, error) {
	return _YieldFarmLP.Contract.TOTALDISTRIBUTEDAMOUNT(&_YieldFarmLP.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarmLP.contract.Call(opts, out, "epochDuration")
	return *ret0, err
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPSession) EpochDuration() (*big.Int, error) {
	return _YieldFarmLP.Contract.EpochDuration(&_YieldFarmLP.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCallerSession) EpochDuration() (*big.Int, error) {
	return _YieldFarmLP.Contract.EpochDuration(&_YieldFarmLP.CallOpts)
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCaller) EpochStart(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarmLP.contract.Call(opts, out, "epochStart")
	return *ret0, err
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPSession) EpochStart() (*big.Int, error) {
	return _YieldFarmLP.Contract.EpochStart(&_YieldFarmLP.CallOpts)
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCallerSession) EpochStart() (*big.Int, error) {
	return _YieldFarmLP.Contract.EpochStart(&_YieldFarmLP.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarmLP.contract.Call(opts, out, "getCurrentEpoch")
	return *ret0, err
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPSession) GetCurrentEpoch() (*big.Int, error) {
	return _YieldFarmLP.Contract.GetCurrentEpoch(&_YieldFarmLP.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _YieldFarmLP.Contract.GetCurrentEpoch(&_YieldFarmLP.CallOpts)
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCaller) GetEpochStake(opts *bind.CallOpts, userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarmLP.contract.Call(opts, out, "getEpochStake", userAddress, epochId)
	return *ret0, err
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_YieldFarmLP *YieldFarmLPSession) GetEpochStake(userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _YieldFarmLP.Contract.GetEpochStake(&_YieldFarmLP.CallOpts, userAddress, epochId)
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCallerSession) GetEpochStake(userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _YieldFarmLP.Contract.GetEpochStake(&_YieldFarmLP.CallOpts, userAddress, epochId)
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCaller) GetPoolSize(opts *bind.CallOpts, epochId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarmLP.contract.Call(opts, out, "getPoolSize", epochId)
	return *ret0, err
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_YieldFarmLP *YieldFarmLPSession) GetPoolSize(epochId *big.Int) (*big.Int, error) {
	return _YieldFarmLP.Contract.GetPoolSize(&_YieldFarmLP.CallOpts, epochId)
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCallerSession) GetPoolSize(epochId *big.Int) (*big.Int, error) {
	return _YieldFarmLP.Contract.GetPoolSize(&_YieldFarmLP.CallOpts, epochId)
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_YieldFarmLP *YieldFarmLPCaller) LastInitializedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarmLP.contract.Call(opts, out, "lastInitializedEpoch")
	return *ret0, err
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_YieldFarmLP *YieldFarmLPSession) LastInitializedEpoch() (*big.Int, error) {
	return _YieldFarmLP.Contract.LastInitializedEpoch(&_YieldFarmLP.CallOpts)
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_YieldFarmLP *YieldFarmLPCallerSession) LastInitializedEpoch() (*big.Int, error) {
	return _YieldFarmLP.Contract.LastInitializedEpoch(&_YieldFarmLP.CallOpts)
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCaller) UserLastEpochIdHarvested(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarmLP.contract.Call(opts, out, "userLastEpochIdHarvested")
	return *ret0, err
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPSession) UserLastEpochIdHarvested() (*big.Int, error) {
	return _YieldFarmLP.Contract.UserLastEpochIdHarvested(&_YieldFarmLP.CallOpts)
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_YieldFarmLP *YieldFarmLPCallerSession) UserLastEpochIdHarvested() (*big.Int, error) {
	return _YieldFarmLP.Contract.UserLastEpochIdHarvested(&_YieldFarmLP.CallOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_YieldFarmLP *YieldFarmLPTransactor) Harvest(opts *bind.TransactOpts, epochId *big.Int) (*types.Transaction, error) {
	return _YieldFarmLP.contract.Transact(opts, "harvest", epochId)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_YieldFarmLP *YieldFarmLPSession) Harvest(epochId *big.Int) (*types.Transaction, error) {
	return _YieldFarmLP.Contract.Harvest(&_YieldFarmLP.TransactOpts, epochId)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_YieldFarmLP *YieldFarmLPTransactorSession) Harvest(epochId *big.Int) (*types.Transaction, error) {
	return _YieldFarmLP.Contract.Harvest(&_YieldFarmLP.TransactOpts, epochId)
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_YieldFarmLP *YieldFarmLPTransactor) MassHarvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldFarmLP.contract.Transact(opts, "massHarvest")
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_YieldFarmLP *YieldFarmLPSession) MassHarvest() (*types.Transaction, error) {
	return _YieldFarmLP.Contract.MassHarvest(&_YieldFarmLP.TransactOpts)
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_YieldFarmLP *YieldFarmLPTransactorSession) MassHarvest() (*types.Transaction, error) {
	return _YieldFarmLP.Contract.MassHarvest(&_YieldFarmLP.TransactOpts)
}

// YieldFarmLPHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the YieldFarmLP contract.
type YieldFarmLPHarvestIterator struct {
	Event *YieldFarmLPHarvest // Event containing the contract specifics and raw log

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
func (it *YieldFarmLPHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldFarmLPHarvest)
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
		it.Event = new(YieldFarmLPHarvest)
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
func (it *YieldFarmLPHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldFarmLPHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldFarmLPHarvest represents a Harvest event raised by the YieldFarmLP contract.
type YieldFarmLPHarvest struct {
	User    common.Address
	EpochId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_YieldFarmLP *YieldFarmLPFilterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address, epochId []*big.Int) (*YieldFarmLPHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _YieldFarmLP.contract.FilterLogs(opts, "Harvest", userRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &YieldFarmLPHarvestIterator{contract: _YieldFarmLP.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_YieldFarmLP *YieldFarmLPFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *YieldFarmLPHarvest, user []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _YieldFarmLP.contract.WatchLogs(opts, "Harvest", userRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldFarmLPHarvest)
				if err := _YieldFarmLP.contract.UnpackLog(event, "Harvest", log); err != nil {
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

// ParseHarvest is a log parse operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_YieldFarmLP *YieldFarmLPFilterer) ParseHarvest(log types.Log) (*YieldFarmLPHarvest, error) {
	event := new(YieldFarmLPHarvest)
	if err := _YieldFarmLP.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// YieldFarmLPMassHarvestIterator is returned from FilterMassHarvest and is used to iterate over the raw logs and unpacked data for MassHarvest events raised by the YieldFarmLP contract.
type YieldFarmLPMassHarvestIterator struct {
	Event *YieldFarmLPMassHarvest // Event containing the contract specifics and raw log

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
func (it *YieldFarmLPMassHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldFarmLPMassHarvest)
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
		it.Event = new(YieldFarmLPMassHarvest)
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
func (it *YieldFarmLPMassHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldFarmLPMassHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldFarmLPMassHarvest represents a MassHarvest event raised by the YieldFarmLP contract.
type YieldFarmLPMassHarvest struct {
	User            common.Address
	EpochsHarvested *big.Int
	TotalValue      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMassHarvest is a free log retrieval operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_YieldFarmLP *YieldFarmLPFilterer) FilterMassHarvest(opts *bind.FilterOpts, user []common.Address) (*YieldFarmLPMassHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YieldFarmLP.contract.FilterLogs(opts, "MassHarvest", userRule)
	if err != nil {
		return nil, err
	}
	return &YieldFarmLPMassHarvestIterator{contract: _YieldFarmLP.contract, event: "MassHarvest", logs: logs, sub: sub}, nil
}

// WatchMassHarvest is a free log subscription operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_YieldFarmLP *YieldFarmLPFilterer) WatchMassHarvest(opts *bind.WatchOpts, sink chan<- *YieldFarmLPMassHarvest, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YieldFarmLP.contract.WatchLogs(opts, "MassHarvest", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldFarmLPMassHarvest)
				if err := _YieldFarmLP.contract.UnpackLog(event, "MassHarvest", log); err != nil {
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

// ParseMassHarvest is a log parse operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_YieldFarmLP *YieldFarmLPFilterer) ParseMassHarvest(log types.Log) (*YieldFarmLPMassHarvest, error) {
	event := new(YieldFarmLPMassHarvest)
	if err := _YieldFarmLP.contract.UnpackLog(event, "MassHarvest", log); err != nil {
		return nil, err
	}
	return event, nil
}
