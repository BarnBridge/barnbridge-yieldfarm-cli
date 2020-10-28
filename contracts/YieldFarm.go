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

// YieldFarmABI is the input ABI used to generate the binding from.
const YieldFarmABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bondTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"susd\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dai\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochsHarvested\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"}],\"name\":\"MassHarvest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NR_OF_EPOCHS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_DISTRIBUTED_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getPoolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"harvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInitializedEpoch\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massHarvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLastEpochIdHarvested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// YieldFarm is an auto generated Go binding around an Ethereum contract.
type YieldFarm struct {
	YieldFarmCaller     // Read-only binding to the contract
	YieldFarmTransactor // Write-only binding to the contract
	YieldFarmFilterer   // Log filterer for contract events
}

// YieldFarmCaller is an auto generated read-only Go binding around an Ethereum contract.
type YieldFarmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldFarmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YieldFarmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldFarmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YieldFarmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldFarmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YieldFarmSession struct {
	Contract     *YieldFarm        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldFarmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YieldFarmCallerSession struct {
	Contract *YieldFarmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// YieldFarmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YieldFarmTransactorSession struct {
	Contract     *YieldFarmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// YieldFarmRaw is an auto generated low-level Go binding around an Ethereum contract.
type YieldFarmRaw struct {
	Contract *YieldFarm // Generic contract binding to access the raw methods on
}

// YieldFarmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YieldFarmCallerRaw struct {
	Contract *YieldFarmCaller // Generic read-only contract binding to access the raw methods on
}

// YieldFarmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YieldFarmTransactorRaw struct {
	Contract *YieldFarmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYieldFarm creates a new instance of YieldFarm, bound to a specific deployed contract.
func NewYieldFarm(address common.Address, backend bind.ContractBackend) (*YieldFarm, error) {
	contract, err := bindYieldFarm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YieldFarm{YieldFarmCaller: YieldFarmCaller{contract: contract}, YieldFarmTransactor: YieldFarmTransactor{contract: contract}, YieldFarmFilterer: YieldFarmFilterer{contract: contract}}, nil
}

// NewYieldFarmCaller creates a new read-only instance of YieldFarm, bound to a specific deployed contract.
func NewYieldFarmCaller(address common.Address, caller bind.ContractCaller) (*YieldFarmCaller, error) {
	contract, err := bindYieldFarm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YieldFarmCaller{contract: contract}, nil
}

// NewYieldFarmTransactor creates a new write-only instance of YieldFarm, bound to a specific deployed contract.
func NewYieldFarmTransactor(address common.Address, transactor bind.ContractTransactor) (*YieldFarmTransactor, error) {
	contract, err := bindYieldFarm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YieldFarmTransactor{contract: contract}, nil
}

// NewYieldFarmFilterer creates a new log filterer instance of YieldFarm, bound to a specific deployed contract.
func NewYieldFarmFilterer(address common.Address, filterer bind.ContractFilterer) (*YieldFarmFilterer, error) {
	contract, err := bindYieldFarm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YieldFarmFilterer{contract: contract}, nil
}

// bindYieldFarm binds a generic wrapper to an already deployed contract.
func bindYieldFarm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YieldFarmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YieldFarm *YieldFarmRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _YieldFarm.Contract.YieldFarmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YieldFarm *YieldFarmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldFarm.Contract.YieldFarmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YieldFarm *YieldFarmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YieldFarm.Contract.YieldFarmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YieldFarm *YieldFarmCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _YieldFarm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YieldFarm *YieldFarmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldFarm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YieldFarm *YieldFarmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YieldFarm.Contract.contract.Transact(opts, method, params...)
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_YieldFarm *YieldFarmCaller) NROFEPOCHS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarm.contract.Call(opts, out, "NR_OF_EPOCHS")
	return *ret0, err
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_YieldFarm *YieldFarmSession) NROFEPOCHS() (*big.Int, error) {
	return _YieldFarm.Contract.NROFEPOCHS(&_YieldFarm.CallOpts)
}

// NROFEPOCHS is a free data retrieval call binding the contract method 0x05917ac0.
//
// Solidity: function NR_OF_EPOCHS() view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) NROFEPOCHS() (*big.Int, error) {
	return _YieldFarm.Contract.NROFEPOCHS(&_YieldFarm.CallOpts)
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_YieldFarm *YieldFarmCaller) TOTALDISTRIBUTEDAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarm.contract.Call(opts, out, "TOTAL_DISTRIBUTED_AMOUNT")
	return *ret0, err
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_YieldFarm *YieldFarmSession) TOTALDISTRIBUTEDAMOUNT() (*big.Int, error) {
	return _YieldFarm.Contract.TOTALDISTRIBUTEDAMOUNT(&_YieldFarm.CallOpts)
}

// TOTALDISTRIBUTEDAMOUNT is a free data retrieval call binding the contract method 0x674247d6.
//
// Solidity: function TOTAL_DISTRIBUTED_AMOUNT() view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) TOTALDISTRIBUTEDAMOUNT() (*big.Int, error) {
	return _YieldFarm.Contract.TOTALDISTRIBUTEDAMOUNT(&_YieldFarm.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_YieldFarm *YieldFarmCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarm.contract.Call(opts, out, "epochDuration")
	return *ret0, err
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_YieldFarm *YieldFarmSession) EpochDuration() (*big.Int, error) {
	return _YieldFarm.Contract.EpochDuration(&_YieldFarm.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) EpochDuration() (*big.Int, error) {
	return _YieldFarm.Contract.EpochDuration(&_YieldFarm.CallOpts)
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_YieldFarm *YieldFarmCaller) EpochStart(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarm.contract.Call(opts, out, "epochStart")
	return *ret0, err
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_YieldFarm *YieldFarmSession) EpochStart() (*big.Int, error) {
	return _YieldFarm.Contract.EpochStart(&_YieldFarm.CallOpts)
}

// EpochStart is a free data retrieval call binding the contract method 0x15e5a1e5.
//
// Solidity: function epochStart() view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) EpochStart() (*big.Int, error) {
	return _YieldFarm.Contract.EpochStart(&_YieldFarm.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_YieldFarm *YieldFarmCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarm.contract.Call(opts, out, "getCurrentEpoch")
	return *ret0, err
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_YieldFarm *YieldFarmSession) GetCurrentEpoch() (*big.Int, error) {
	return _YieldFarm.Contract.GetCurrentEpoch(&_YieldFarm.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _YieldFarm.Contract.GetCurrentEpoch(&_YieldFarm.CallOpts)
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_YieldFarm *YieldFarmCaller) GetEpochStake(opts *bind.CallOpts, userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarm.contract.Call(opts, out, "getEpochStake", userAddress, epochId)
	return *ret0, err
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_YieldFarm *YieldFarmSession) GetEpochStake(userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _YieldFarm.Contract.GetEpochStake(&_YieldFarm.CallOpts, userAddress, epochId)
}

// GetEpochStake is a free data retrieval call binding the contract method 0x43312451.
//
// Solidity: function getEpochStake(address userAddress, uint128 epochId) view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) GetEpochStake(userAddress common.Address, epochId *big.Int) (*big.Int, error) {
	return _YieldFarm.Contract.GetEpochStake(&_YieldFarm.CallOpts, userAddress, epochId)
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_YieldFarm *YieldFarmCaller) GetPoolSize(opts *bind.CallOpts, epochId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarm.contract.Call(opts, out, "getPoolSize", epochId)
	return *ret0, err
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_YieldFarm *YieldFarmSession) GetPoolSize(epochId *big.Int) (*big.Int, error) {
	return _YieldFarm.Contract.GetPoolSize(&_YieldFarm.CallOpts, epochId)
}

// GetPoolSize is a free data retrieval call binding the contract method 0xf7e251f8.
//
// Solidity: function getPoolSize(uint128 epochId) view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) GetPoolSize(epochId *big.Int) (*big.Int, error) {
	return _YieldFarm.Contract.GetPoolSize(&_YieldFarm.CallOpts, epochId)
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_YieldFarm *YieldFarmCaller) LastInitializedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarm.contract.Call(opts, out, "lastInitializedEpoch")
	return *ret0, err
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_YieldFarm *YieldFarmSession) LastInitializedEpoch() (*big.Int, error) {
	return _YieldFarm.Contract.LastInitializedEpoch(&_YieldFarm.CallOpts)
}

// LastInitializedEpoch is a free data retrieval call binding the contract method 0x9c7ec881.
//
// Solidity: function lastInitializedEpoch() view returns(uint128)
func (_YieldFarm *YieldFarmCallerSession) LastInitializedEpoch() (*big.Int, error) {
	return _YieldFarm.Contract.LastInitializedEpoch(&_YieldFarm.CallOpts)
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_YieldFarm *YieldFarmCaller) UserLastEpochIdHarvested(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _YieldFarm.contract.Call(opts, out, "userLastEpochIdHarvested")
	return *ret0, err
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_YieldFarm *YieldFarmSession) UserLastEpochIdHarvested() (*big.Int, error) {
	return _YieldFarm.Contract.UserLastEpochIdHarvested(&_YieldFarm.CallOpts)
}

// UserLastEpochIdHarvested is a free data retrieval call binding the contract method 0xa1c130d4.
//
// Solidity: function userLastEpochIdHarvested() view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) UserLastEpochIdHarvested() (*big.Int, error) {
	return _YieldFarm.Contract.UserLastEpochIdHarvested(&_YieldFarm.CallOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_YieldFarm *YieldFarmTransactor) Harvest(opts *bind.TransactOpts, epochId *big.Int) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "harvest", epochId)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_YieldFarm *YieldFarmSession) Harvest(epochId *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.Harvest(&_YieldFarm.TransactOpts, epochId)
}

// Harvest is a paid mutator transaction binding the contract method 0xa43564eb.
//
// Solidity: function harvest(uint128 epochId) returns(uint256)
func (_YieldFarm *YieldFarmTransactorSession) Harvest(epochId *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.Harvest(&_YieldFarm.TransactOpts, epochId)
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_YieldFarm *YieldFarmTransactor) MassHarvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "massHarvest")
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_YieldFarm *YieldFarmSession) MassHarvest() (*types.Transaction, error) {
	return _YieldFarm.Contract.MassHarvest(&_YieldFarm.TransactOpts)
}

// MassHarvest is a paid mutator transaction binding the contract method 0x290e4544.
//
// Solidity: function massHarvest() returns(uint256)
func (_YieldFarm *YieldFarmTransactorSession) MassHarvest() (*types.Transaction, error) {
	return _YieldFarm.Contract.MassHarvest(&_YieldFarm.TransactOpts)
}

// YieldFarmHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the YieldFarm contract.
type YieldFarmHarvestIterator struct {
	Event *YieldFarmHarvest // Event containing the contract specifics and raw log

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
func (it *YieldFarmHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldFarmHarvest)
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
		it.Event = new(YieldFarmHarvest)
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
func (it *YieldFarmHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldFarmHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldFarmHarvest represents a Harvest event raised by the YieldFarm contract.
type YieldFarmHarvest struct {
	User    common.Address
	EpochId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_YieldFarm *YieldFarmFilterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address, epochId []*big.Int) (*YieldFarmHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _YieldFarm.contract.FilterLogs(opts, "Harvest", userRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &YieldFarmHarvestIterator{contract: _YieldFarm.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x04ad45a69eeed9c390c3a678fed2d4b90bde98e742de9936d5e0915bf3d0ea4e.
//
// Solidity: event Harvest(address indexed user, uint128 indexed epochId, uint256 amount)
func (_YieldFarm *YieldFarmFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *YieldFarmHarvest, user []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _YieldFarm.contract.WatchLogs(opts, "Harvest", userRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldFarmHarvest)
				if err := _YieldFarm.contract.UnpackLog(event, "Harvest", log); err != nil {
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
func (_YieldFarm *YieldFarmFilterer) ParseHarvest(log types.Log) (*YieldFarmHarvest, error) {
	event := new(YieldFarmHarvest)
	if err := _YieldFarm.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// YieldFarmMassHarvestIterator is returned from FilterMassHarvest and is used to iterate over the raw logs and unpacked data for MassHarvest events raised by the YieldFarm contract.
type YieldFarmMassHarvestIterator struct {
	Event *YieldFarmMassHarvest // Event containing the contract specifics and raw log

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
func (it *YieldFarmMassHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldFarmMassHarvest)
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
		it.Event = new(YieldFarmMassHarvest)
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
func (it *YieldFarmMassHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldFarmMassHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldFarmMassHarvest represents a MassHarvest event raised by the YieldFarm contract.
type YieldFarmMassHarvest struct {
	User            common.Address
	EpochsHarvested *big.Int
	TotalValue      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMassHarvest is a free log retrieval operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_YieldFarm *YieldFarmFilterer) FilterMassHarvest(opts *bind.FilterOpts, user []common.Address) (*YieldFarmMassHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YieldFarm.contract.FilterLogs(opts, "MassHarvest", userRule)
	if err != nil {
		return nil, err
	}
	return &YieldFarmMassHarvestIterator{contract: _YieldFarm.contract, event: "MassHarvest", logs: logs, sub: sub}, nil
}

// WatchMassHarvest is a free log subscription operation binding the contract event 0xb68dafc1da13dc868096d0b87347c831d0bda92d178317eb1dec7f788444485c.
//
// Solidity: event MassHarvest(address indexed user, uint256 epochsHarvested, uint256 totalValue)
func (_YieldFarm *YieldFarmFilterer) WatchMassHarvest(opts *bind.WatchOpts, sink chan<- *YieldFarmMassHarvest, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YieldFarm.contract.WatchLogs(opts, "MassHarvest", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldFarmMassHarvest)
				if err := _YieldFarm.contract.UnpackLog(event, "MassHarvest", log); err != nil {
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
func (_YieldFarm *YieldFarmFilterer) ParseMassHarvest(log types.Log) (*YieldFarmMassHarvest, error) {
	event := new(YieldFarmMassHarvest)
	if err := _YieldFarm.contract.UnpackLog(event, "MassHarvest", log); err != nil {
		return nil, err
	}
	return event, nil
}
