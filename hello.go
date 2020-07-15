// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hello

import (
	"math/big"
	"strings"

	fiscobcos "github.com/chislab/go-fiscobcos"
	"github.com/chislab/go-fiscobcos/accounts/abi"
	"github.com/chislab/go-fiscobcos/accounts/abi/bind"
	"github.com/chislab/go-fiscobcos/common"
	"github.com/chislab/go-fiscobcos/common/hexutil"
	"github.com/chislab/go-fiscobcos/core/types"
	"github.com/chislab/go-fiscobcos/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = fiscobcos.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HelloABI is the input ABI used to generate the binding from.
const HelloABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"EvtSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HelloBin is the compiled bytecode used for deploying new contracts.
var HelloBin = "0x608060405234801561001057600080fd5b506040518060400160405280601481526020017f68656c6c6f2066726f6d2064656e67206d696d690000000000000000000000008152506000908051906020019061005c929190610062565b50610107565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a357805160ff19168380011785556100d1565b828001600101855582156100d1579182015b828111156100d05782518255916020019190600101906100b5565b5b5090506100de91906100e2565b5090565b61010491905b808211156101005760008160009055506001016100e8565b5090565b90565b6103ff806101166000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634ed3885e1461003b5780636d4ce63c14610057575b600080fd5b6100556004803603610050919081019061021a565b610076565b005b61005f6100ca565b60405161006d929190610358565b60405180910390f35b806000908051906020019061008c929190610175565b507f54b0facb31be7c81ea1c4dd380c803b0e379d4e9e1c9799087926c8d841d15c83360006040516100bf9291906102ba565b60405180910390a150565b60606000806001818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101665780601f1061013b57610100808354040283529160200191610166565b820191906000526020600020905b81548152906001019060200180831161014957829003601f168201915b50505050509150915091509091565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106101b657805160ff19168380011785556101e4565b828001600101855582156101e4579182015b828111156101e35782518255916020019190600101906101c8565b5b5090506101f191906101f5565b5090565b61021791905b808211156102135760008160009055506001016101fb565b5090565b90565b60006020828403121561022b578081fd5b813567ffffffffffffffff80821115610242578283fd5b81840185601f820112610253578384fd5b8035925081831115610263578384fd5b6040516020601f19601f8601168201018181108482111715610283578586fd5b806040525083815286602085840101111561029c578485fd5b6102ad8460208301602085016103ba565b8094505050505092915050565b60006040820160018060a01b03851683526040602084015281845460018116600081146102ee576001811461030c5761034a565b607f6002830416845260ff198216606087015260808601925061034a565b6002820480855287865260208620865b8281101561033e5781546060828b01015260018201915060208101905061031c565b6060818a010195505050505b505080925050509392505050565b6000604082528351806040840152815b81811015610389576020818701015160608286010152602081019050610368565b8181111561039a5782606083860101525b506060601f19601f83011684010191505082151560208301529392505050565b8281833760008383015250505056fea2646970667358221220306ebc6ee119efd64d2596869b67f4b45d7c3cb1869123f8c81058079390fc3864736f6c63430006000033"

// DeployHello deploys a new FiscoBcos contract, binding an instance of Hello to it.
func DeployHello(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Hello, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HelloBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hello{HelloCaller: HelloCaller{contract: contract}, HelloTransactor: HelloTransactor{contract: contract}, HelloFilterer: HelloFilterer{contract: contract}, ABI: parsed}, nil
}

// Hello is an auto generated Go binding around an FiscoBcos contract.
type Hello struct {
	HelloCaller             // Read-only binding to the contract
	HelloTransactor         // Write-only binding to the contract
	HelloFilterer           // Log filterer for contract events
	ABI             abi.ABI // contract abi
}

// HelloCaller is an auto generated read-only Go binding around an FiscoBcos contract.
type HelloCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloTransactor is an auto generated write-only Go binding around an FiscoBcos contract.
type HelloTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloFilterer is an auto generated log filtering Go binding around an FiscoBcos contract events.
type HelloFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloSession is an auto generated Go binding around an FiscoBcos contract,
// with pre-set call and transact options.
type HelloSession struct {
	Contract     *Hello            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloCallerSession is an auto generated read-only Go binding around an FiscoBcos contract,
// with pre-set call options.
type HelloCallerSession struct {
	Contract *HelloCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HelloTransactorSession is an auto generated write-only Go binding around an FiscoBcos contract,
// with pre-set transact options.
type HelloTransactorSession struct {
	Contract     *HelloTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloRaw is an auto generated low-level Go binding around an FiscoBcos contract.
type HelloRaw struct {
	Contract *Hello // Generic contract binding to access the raw methods on
}

// HelloCallerRaw is an auto generated low-level read-only Go binding around an FiscoBcos contract.
type HelloCallerRaw struct {
	Contract *HelloCaller // Generic read-only contract binding to access the raw methods on
}

// HelloTransactorRaw is an auto generated low-level write-only Go binding around an FiscoBcos contract.
type HelloTransactorRaw struct {
	Contract *HelloTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHello creates a new instance of Hello, bound to a specific deployed contract.
func NewHello(address common.Address, backend bind.ContractBackend) (*Hello, error) {
	contract, err := bindHello(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(HelloABI))
	if err != nil {
		return nil, err
	}
	return &Hello{HelloCaller: HelloCaller{contract: contract}, HelloTransactor: HelloTransactor{contract: contract}, HelloFilterer: HelloFilterer{contract: contract}, ABI: parsed}, nil
}

// NewHelloCaller creates a new read-only instance of Hello, bound to a specific deployed contract.
func NewHelloCaller(address common.Address, caller bind.ContractCaller) (*HelloCaller, error) {
	contract, err := bindHello(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloCaller{contract: contract}, nil
}

// NewHelloTransactor creates a new write-only instance of Hello, bound to a specific deployed contract.
func NewHelloTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloTransactor, error) {
	contract, err := bindHello(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloTransactor{contract: contract}, nil
}

// NewHelloFilterer creates a new log filterer instance of Hello, bound to a specific deployed contract.
func NewHelloFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloFilterer, error) {
	contract, err := bindHello(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloFilterer{contract: contract}, nil
}

// bindHello binds a generic wrapper to an already deployed contract.
func bindHello(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hello *HelloRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hello.Contract.HelloCaller.contract.Call(opts, result, method, params...)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hello *HelloRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hello.Contract.HelloTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hello *HelloCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hello.Contract.contract.Call(opts, result, method, params...)
}

func (_Hello *HelloCallerRaw) ReadCall(result interface{}, method string, output []byte) error {
	return _Hello.Contract.contract.ReadCall(result, method, output)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hello *HelloTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hello.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(string, bool)
func (_Hello *HelloCaller) Get(opts *bind.CallOpts) (string, bool, error) {
	var (
		ret0 = new(string)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Hello.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(string, bool)
func (_Hello *HelloSession) Get() (string, bool, error) {
	return _Hello.Contract.Get(&_Hello.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(string, bool)
func (_Hello *HelloCallerSession) Get() (string, bool, error) {
	return _Hello.Contract.Get(&_Hello.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string name) returns()

/******************************************************************************************************************************/
func (_Hello *HelloCaller) ReadSet(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Hello.contract.ReadCall(out, "set", hexutil.MustDecode(output))
	return err
}

// Set is a free data retrieval call binding the contract method 0x4ed3885e.
//
// Solidity: function set(string name) returns()
func (_Hello *HelloSession) ReadSet(output string) error {
	return _Hello.Contract.ReadSet(output)
}

// Set is a free data retrieval call binding the contract method 0x4ed3885e.
//
// Solidity: function set(string name) returns()
func (_Hello *HelloCallerSession) ReadSet(output string) error {
	return _Hello.Contract.ReadSet(output)
}

/******************************************************************************************************************************/

func (_Hello *HelloTransactor) Set(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Hello.contract.Transact(opts, "set", name)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string name) returns()
func (_Hello *HelloSession) Set(name string) (*types.Transaction, error) {
	return _Hello.Contract.Set(&_Hello.TransactOpts, name)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string name) returns()
func (_Hello *HelloTransactorSession) Set(name string) (*types.Transaction, error) {
	return _Hello.Contract.Set(&_Hello.TransactOpts, name)
}

// HelloEvtSetIterator is returned from FilterEvtSet and is used to iterate over the raw logs and unpacked data for EvtSet events raised by the Hello contract.
type HelloEvtSetIterator struct {
	Event *HelloEvtSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log         // Log channel receiving the found contract events
	sub  fiscobcos.Subscription // Subscription for errors, completion and termination
	done bool                   // Whether the subscription completed delivering logs
	fail error                  // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HelloEvtSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloEvtSet)
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
		it.Event = new(HelloEvtSet)
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
func (it *HelloEvtSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloEvtSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloEvtSet represents a EvtSet event raised by the Hello contract.
type HelloEvtSet struct {
	From common.Address
	Msg  string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEvtSet is a free log retrieval operation binding the contract event 0x54b0facb31be7c81ea1c4dd380c803b0e379d4e9e1c9799087926c8d841d15c8.
//
// Solidity: event EvtSet(address from, string msg)
func (_Hello *HelloFilterer) FilterEvtSet(opts *bind.FilterOpts) (*HelloEvtSetIterator, error) {

	logs, sub, err := _Hello.contract.FilterLogs(opts, "EvtSet")
	if err != nil {
		return nil, err
	}
	return &HelloEvtSetIterator{contract: _Hello.contract, event: "EvtSet", logs: logs, sub: sub}, nil
}

// WatchEvtSet is a free log subscription operation binding the contract event 0x54b0facb31be7c81ea1c4dd380c803b0e379d4e9e1c9799087926c8d841d15c8.
//
// Solidity: event EvtSet(address from, string msg)
func (_Hello *HelloFilterer) WatchEvtSet(opts *bind.WatchOpts, sink chan<- *HelloEvtSet) (event.Subscription, error) {

	logs, sub, err := _Hello.contract.WatchLogs(opts, "EvtSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloEvtSet)
				if err := _Hello.contract.UnpackLog(event, "EvtSet", log); err != nil {
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

// ParseEvtSet is a log parse operation binding the contract event 0x54b0facb31be7c81ea1c4dd380c803b0e379d4e9e1c9799087926c8d841d15c8.
//
// Solidity: event EvtSet(address from, string msg)
func (_Hello *HelloFilterer) ParseEvtSet(log types.Log) (*HelloEvtSet, error) {
	event := new(HelloEvtSet)
	if err := _Hello.contract.UnpackLog(event, "EvtSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
