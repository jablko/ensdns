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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DNSResolverABI is the input ABI used to generate the binding from.
const DNSResolverABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"dnsrr\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setDnsrr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DNSResolverBin is the compiled bytecode used for deploying new contracts.
const DNSResolverBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556103bb806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806301ffc9a714610051578063126a710e1461008c57806376196c881461011e5780638da5cb5b146101cd575b600080fd5b6100786004803603602081101561006757600080fd5b50356001600160e01b0319166101f1565b604080519115158252519081900360200190f35b6100a9600480360360208110156100a257600080fd5b503561020a565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100e35781810151838201526020016100cb565b50505050905090810190601f1680156101105780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101cb6004803603604081101561013457600080fd5b8135919081019060408101602082013564010000000081111561015657600080fd5b82018360208201111561016857600080fd5b8035906020019184600183028401116401000000008311171561018a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102aa945050505050565b005b6101d56102e5565b604080516001600160a01b039092168252519081900360200190f35b6001600160e01b031916600160e11b6309353887021490565b60008181526001602081815260409283902080548451600294821615610100026000190190911693909304601f8101839004830284018301909452838352606093909183018282801561029e5780601f106102735761010080835404028352916020019161029e565b820191906000526020600020905b81548152906001019060200180831161028157829003601f168201915b50505050509050919050565b6000546001600160a01b031633146102c157600080fd5b600082815260016020908152604090912082516102e0928401906102f4565b505050565b6000546001600160a01b031681565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061033557805160ff1916838001178555610362565b82800160010185558215610362579182015b82811115610362578251825591602001919060010190610347565b5061036e929150610372565b5090565b61038c91905b8082111561036e5760008155600101610378565b9056fea165627a7a72305820822fbe25291ff84bdcd6ff2d0dcb6a775bfc781f6d24c741517158f2295173860029`

// DeployDNSResolver deploys a new Ethereum contract, binding an instance of DNSResolver to it.
func DeployDNSResolver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DNSResolver, error) {
	parsed, err := abi.JSON(strings.NewReader(DNSResolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DNSResolverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DNSResolver{DNSResolverCaller: DNSResolverCaller{contract: contract}, DNSResolverTransactor: DNSResolverTransactor{contract: contract}, DNSResolverFilterer: DNSResolverFilterer{contract: contract}}, nil
}

// DNSResolver is an auto generated Go binding around an Ethereum contract.
type DNSResolver struct {
	DNSResolverCaller     // Read-only binding to the contract
	DNSResolverTransactor // Write-only binding to the contract
	DNSResolverFilterer   // Log filterer for contract events
}

// DNSResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type DNSResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DNSResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DNSResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DNSResolverSession struct {
	Contract     *DNSResolver      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DNSResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DNSResolverCallerSession struct {
	Contract *DNSResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DNSResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DNSResolverTransactorSession struct {
	Contract     *DNSResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DNSResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type DNSResolverRaw struct {
	Contract *DNSResolver // Generic contract binding to access the raw methods on
}

// DNSResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DNSResolverCallerRaw struct {
	Contract *DNSResolverCaller // Generic read-only contract binding to access the raw methods on
}

// DNSResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DNSResolverTransactorRaw struct {
	Contract *DNSResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDNSResolver creates a new instance of DNSResolver, bound to a specific deployed contract.
func NewDNSResolver(address common.Address, backend bind.ContractBackend) (*DNSResolver, error) {
	contract, err := bindDNSResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DNSResolver{DNSResolverCaller: DNSResolverCaller{contract: contract}, DNSResolverTransactor: DNSResolverTransactor{contract: contract}, DNSResolverFilterer: DNSResolverFilterer{contract: contract}}, nil
}

// NewDNSResolverCaller creates a new read-only instance of DNSResolver, bound to a specific deployed contract.
func NewDNSResolverCaller(address common.Address, caller bind.ContractCaller) (*DNSResolverCaller, error) {
	contract, err := bindDNSResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DNSResolverCaller{contract: contract}, nil
}

// NewDNSResolverTransactor creates a new write-only instance of DNSResolver, bound to a specific deployed contract.
func NewDNSResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*DNSResolverTransactor, error) {
	contract, err := bindDNSResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DNSResolverTransactor{contract: contract}, nil
}

// NewDNSResolverFilterer creates a new log filterer instance of DNSResolver, bound to a specific deployed contract.
func NewDNSResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*DNSResolverFilterer, error) {
	contract, err := bindDNSResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DNSResolverFilterer{contract: contract}, nil
}

// bindDNSResolver binds a generic wrapper to an already deployed contract.
func bindDNSResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DNSResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DNSResolver *DNSResolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DNSResolver.Contract.DNSResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DNSResolver *DNSResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNSResolver.Contract.DNSResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DNSResolver *DNSResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DNSResolver.Contract.DNSResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DNSResolver *DNSResolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DNSResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DNSResolver *DNSResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNSResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DNSResolver *DNSResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DNSResolver.Contract.contract.Transact(opts, method, params...)
}

// Dnsrr is a free data retrieval call binding the contract method 0x126a710e.
//
// Solidity: function dnsrr(bytes32 node) constant returns(bytes)
func (_DNSResolver *DNSResolverCaller) Dnsrr(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DNSResolver.contract.Call(opts, out, "dnsrr", node)
	return *ret0, err
}

// Dnsrr is a free data retrieval call binding the contract method 0x126a710e.
//
// Solidity: function dnsrr(bytes32 node) constant returns(bytes)
func (_DNSResolver *DNSResolverSession) Dnsrr(node [32]byte) ([]byte, error) {
	return _DNSResolver.Contract.Dnsrr(&_DNSResolver.CallOpts, node)
}

// Dnsrr is a free data retrieval call binding the contract method 0x126a710e.
//
// Solidity: function dnsrr(bytes32 node) constant returns(bytes)
func (_DNSResolver *DNSResolverCallerSession) Dnsrr(node [32]byte) ([]byte, error) {
	return _DNSResolver.Contract.Dnsrr(&_DNSResolver.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DNSResolver *DNSResolverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DNSResolver.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DNSResolver *DNSResolverSession) Owner() (common.Address, error) {
	return _DNSResolver.Contract.Owner(&_DNSResolver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DNSResolver *DNSResolverCallerSession) Owner() (common.Address, error) {
	return _DNSResolver.Contract.Owner(&_DNSResolver.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) constant returns(bool)
func (_DNSResolver *DNSResolverCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DNSResolver.contract.Call(opts, out, "supportsInterface", interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) constant returns(bool)
func (_DNSResolver *DNSResolverSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _DNSResolver.Contract.SupportsInterface(&_DNSResolver.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) constant returns(bool)
func (_DNSResolver *DNSResolverCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _DNSResolver.Contract.SupportsInterface(&_DNSResolver.CallOpts, interfaceID)
}

// SetDnsrr is a paid mutator transaction binding the contract method 0x76196c88.
//
// Solidity: function setDnsrr(bytes32 node, bytes data) returns()
func (_DNSResolver *DNSResolverTransactor) SetDnsrr(opts *bind.TransactOpts, node [32]byte, data []byte) (*types.Transaction, error) {
	return _DNSResolver.contract.Transact(opts, "setDnsrr", node, data)
}

// SetDnsrr is a paid mutator transaction binding the contract method 0x76196c88.
//
// Solidity: function setDnsrr(bytes32 node, bytes data) returns()
func (_DNSResolver *DNSResolverSession) SetDnsrr(node [32]byte, data []byte) (*types.Transaction, error) {
	return _DNSResolver.Contract.SetDnsrr(&_DNSResolver.TransactOpts, node, data)
}

// SetDnsrr is a paid mutator transaction binding the contract method 0x76196c88.
//
// Solidity: function setDnsrr(bytes32 node, bytes data) returns()
func (_DNSResolver *DNSResolverTransactorSession) SetDnsrr(node [32]byte, data []byte) (*types.Transaction, error) {
	return _DNSResolver.Contract.SetDnsrr(&_DNSResolver.TransactOpts, node, data)
}
