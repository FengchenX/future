// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package addrmanager

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"addresslist\",\"outputs\":[{\"name\":\"smartAddr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"findkey\",\"type\":\"string\"}],\"name\":\"getAddressIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"accountAddr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"addAddressIndex\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountAddr\",\"type\":\"address\"}],\"name\":\"getAddressArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"accountAddr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"addAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"logstr\",\"type\":\"string\"}],\"name\":\"globallog\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `606060405234156200001057600080fd5b60405162001999380380620019998339810160405280805182019190602001805182019190505081600090805190602001906200004f929190620000b2565b50806001908051906020019062000068929190620000b2565b5033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000161565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000f557805160ff191683800117855562000126565b8280016001018555821562000126579182015b828111156200012557825182559160200191906001019062000108565b5b50905062000135919062000139565b5090565b6200015e91905b808211156200015a57600081600090555060010162000140565b5090565b90565b61182880620001716000396000f300606060405260043610610083576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde031461008857806370a7287b1461011657806395d89b411461026757806398285d3b146102f5578063cf89b5cb146103a7578063da97754e14610466578063e0136ea9146105b7575b600080fd5b341561009357600080fd5b61009b61068e565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100db5780820151818401526020810190506100c0565b50505050905090810190601f1680156101085780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561012157600080fd5b61014d600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061072c565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156101c35780820151818401526020810190506101a8565b50505050905090810190601f1680156101f05780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561022957808201518184015260208101905061020e565b50505050905090810190601f1680156102565780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b341561027257600080fd5b61027a6108a6565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102ba57808201518184015260208101905061029f565b50505050905090810190601f1680156102e75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561030057600080fd5b610350600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610944565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610393578082015181840152602081019050610378565b505050509050019250505060405180910390f35b34156103b257600080fd5b610464600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610a43565b005b341561047157600080fd5b61049d600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610f2e565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156105135780820151818401526020810190506104f8565b50505050905090810190601f1680156105405780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561057957808201518184015260208101905061055e565b50505050905090810190601f1680156105a65780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34156105c257600080fd5b610674600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190505061116c565b604051808215151515815260200191505060405180910390f35b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107245780601f106106f957610100808354040283529160200191610724565b820191906000526020600020905b81548152906001019060200180831161070757829003601f168201915b505050505081565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107fe5780601f106107d3576101008083540402835291602001916107fe565b820191906000526020600020905b8154815290600101906020018083116107e157829003601f168201915b505050505090806002018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561089c5780601f106108715761010080835404028352916020019161089c565b820191906000526020600020905b81548152906001019060200180831161087f57829003601f168201915b5050505050905083565b60018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561093c5780601f106109115761010080835404028352916020019161093c565b820191906000526020600020905b81548152906001019060200180831161091f57829003601f168201915b505050505081565b61094c6116ef565b6003826040518082805190602001908083835b602083101515610984578051825260208201915060208101905060208303925061095f565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020805480602002602001604051908101604052809291908181526020018280548015610a3757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116109ed575b50505050509050919050565b6003826040518082805190602001908083835b602083101515610a7b5780518252602082019150602081019050602083039250610a56565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054806001018281610ac09190611703565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506003816040518082805190602001908083835b602083101515610b475780518252602082019150602081019050602083039250610b22565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054806001018281610b8c9190611703565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506003610be7838361126a565b6040518082805190602001908083835b602083101515610c1c5780518252602082019150602081019050602083039250610bf7565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054806001018281610c619190611703565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550507f4013cd45e5626ff307481e189e2a154e6035bf11795c12ac2010e5275b2f242e8383604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610d45578082015181840152602081019050610d2a565b50505050905090810190601f168015610d725780820380516001836020036101000a031916815260200191505b50935050505060405180910390a17f4013cd45e5626ff307481e189e2a154e6035bf11795c12ac2010e5275b2f242e8382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610e15578082015181840152602081019050610dfa565b50505050905090810190601f168015610e425780820380516001836020036101000a031916815260200191505b50935050505060405180910390a17f4013cd45e5626ff307481e189e2a154e6035bf11795c12ac2010e5275b2f242e83610e7c848461126a565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610eee578082015181840152602081019050610ed3565b50505050905090810190601f168015610f1b5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a1505050565b6000610f3861172f565b610f4061172f565b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110bc5780601f10611091576101008083540402835291602001916110bc565b820191906000526020600020905b81548152906001019060200180831161109f57829003601f168201915b50505050509150808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111585780601f1061112d57610100808354040283529160200191611158565b820191906000526020600020905b81548152906001019060200180831161113b57829003601f168201915b505050505090509250925092509193909250565b60006060604051908101604052808573ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815250600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019080519060200190611242929190611743565b50604082015181600201908051906020019061125f929190611743565b509050509392505050565b61127261172f565b6112af83836020604051908101604052806000815250602060405190810160405280600081525060206040519081016040528060008152506112b7565b905092915050565b6112bf61172f565b6112c76117c3565b6112cf6117c3565b6112d76117c3565b6112df6117c3565b6112e76117c3565b6112ef61172f565b6112f76117c3565b6000808e98508d97508c96508b95508a94508451865188518a518c51010101016040518059106113245750595b9080825280601f01601f1916602001820160405250935083925060009150600090505b88518110156113fa57888181518110151561135e57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f01000000000000000000000000000000000000000000000000000000000000000283838060010194508151811015156113bd57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050611347565b600090505b87518110156114b257878181518110151561141657fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002838380600101945081518110151561147557fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506113ff565b600090505b865181101561156a5786818151811015156114ce57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002838380600101945081518110151561152d57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506114b7565b600090505b855181101561162257858181518110151561158657fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f01000000000000000000000000000000000000000000000000000000000000000283838060010194508151811015156115e557fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808060010191505061156f565b600090505b84518110156116da57848181518110151561163e57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002838380600101945081518110151561169d57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050611627565b82995050505050505050505095945050505050565b602060405190810160405280600081525090565b81548183558181151161172a5781836000526020600020918201910161172991906117d7565b5b505050565b602060405190810160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061178457805160ff19168380011785556117b2565b828001600101855582156117b2579182015b828111156117b1578251825591602001919060010190611796565b5b5090506117bf91906117d7565b5090565b602060405190810160405280600081525090565b6117f991905b808211156117f55760008160009055506001016117dd565b5090565b905600a165627a7a72305820c20905ed9c07459b2fe256d78522e5fece99d53ba23d2f21e747a42d403041150029`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, tokenName string, tokenSymbol string) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend, tokenName, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Addresslist is a free data retrieval call binding the contract method 0x70a7287b.
//
// Solidity: function addresslist( address) constant returns(smartAddr address, name string, symbol string)
func (_Token *TokenCaller) Addresslist(opts *bind.CallOpts, arg0 common.Address) (struct {
	SmartAddr common.Address
	Name      string
	Symbol    string
}, error) {
	ret := new(struct {
		SmartAddr common.Address
		Name      string
		Symbol    string
	})
	out := ret
	err := _Token.contract.Call(opts, out, "addresslist", arg0)
	return *ret, err
}

// Addresslist is a free data retrieval call binding the contract method 0x70a7287b.
//
// Solidity: function addresslist( address) constant returns(smartAddr address, name string, symbol string)
func (_Token *TokenSession) Addresslist(arg0 common.Address) (struct {
	SmartAddr common.Address
	Name      string
	Symbol    string
}, error) {
	return _Token.Contract.Addresslist(&_Token.CallOpts, arg0)
}

// Addresslist is a free data retrieval call binding the contract method 0x70a7287b.
//
// Solidity: function addresslist( address) constant returns(smartAddr address, name string, symbol string)
func (_Token *TokenCallerSession) Addresslist(arg0 common.Address) (struct {
	SmartAddr common.Address
	Name      string
	Symbol    string
}, error) {
	return _Token.Contract.Addresslist(&_Token.CallOpts, arg0)
}

// GetAddressArray is a free data retrieval call binding the contract method 0xda97754e.
//
// Solidity: function getAddressArray(accountAddr address) constant returns(address, string, string)
func (_Token *TokenCaller) GetAddressArray(opts *bind.CallOpts, accountAddr common.Address) (common.Address, string, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Token.contract.Call(opts, out, "getAddressArray", accountAddr)
	return *ret0, *ret1, *ret2, err
}

// GetAddressArray is a free data retrieval call binding the contract method 0xda97754e.
//
// Solidity: function getAddressArray(accountAddr address) constant returns(address, string, string)
func (_Token *TokenSession) GetAddressArray(accountAddr common.Address) (common.Address, string, string, error) {
	return _Token.Contract.GetAddressArray(&_Token.CallOpts, accountAddr)
}

// GetAddressArray is a free data retrieval call binding the contract method 0xda97754e.
//
// Solidity: function getAddressArray(accountAddr address) constant returns(address, string, string)
func (_Token *TokenCallerSession) GetAddressArray(accountAddr common.Address) (common.Address, string, string, error) {
	return _Token.Contract.GetAddressArray(&_Token.CallOpts, accountAddr)
}

// GetAddressIndex is a free data retrieval call binding the contract method 0x98285d3b.
//
// Solidity: function getAddressIndex(findkey string) constant returns(address[])
func (_Token *TokenCaller) GetAddressIndex(opts *bind.CallOpts, findkey string) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getAddressIndex", findkey)
	return *ret0, err
}

// GetAddressIndex is a free data retrieval call binding the contract method 0x98285d3b.
//
// Solidity: function getAddressIndex(findkey string) constant returns(address[])
func (_Token *TokenSession) GetAddressIndex(findkey string) ([]common.Address, error) {
	return _Token.Contract.GetAddressIndex(&_Token.CallOpts, findkey)
}

// GetAddressIndex is a free data retrieval call binding the contract method 0x98285d3b.
//
// Solidity: function getAddressIndex(findkey string) constant returns(address[])
func (_Token *TokenCallerSession) GetAddressIndex(findkey string) ([]common.Address, error) {
	return _Token.Contract.GetAddressIndex(&_Token.CallOpts, findkey)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// AddAddress is a paid mutator transaction binding the contract method 0xe0136ea9.
//
// Solidity: function addAddress(accountAddr address, name string, symbol string) returns(success bool)
func (_Token *TokenTransactor) AddAddress(opts *bind.TransactOpts, accountAddr common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addAddress", accountAddr, name, symbol)
}

// AddAddress is a paid mutator transaction binding the contract method 0xe0136ea9.
//
// Solidity: function addAddress(accountAddr address, name string, symbol string) returns(success bool)
func (_Token *TokenSession) AddAddress(accountAddr common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Token.Contract.AddAddress(&_Token.TransactOpts, accountAddr, name, symbol)
}

// AddAddress is a paid mutator transaction binding the contract method 0xe0136ea9.
//
// Solidity: function addAddress(accountAddr address, name string, symbol string) returns(success bool)
func (_Token *TokenTransactorSession) AddAddress(accountAddr common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Token.Contract.AddAddress(&_Token.TransactOpts, accountAddr, name, symbol)
}

// AddAddressIndex is a paid mutator transaction binding the contract method 0xcf89b5cb.
//
// Solidity: function addAddressIndex(accountAddr address, name string, symbol string) returns()
func (_Token *TokenTransactor) AddAddressIndex(opts *bind.TransactOpts, accountAddr common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addAddressIndex", accountAddr, name, symbol)
}

// AddAddressIndex is a paid mutator transaction binding the contract method 0xcf89b5cb.
//
// Solidity: function addAddressIndex(accountAddr address, name string, symbol string) returns()
func (_Token *TokenSession) AddAddressIndex(accountAddr common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Token.Contract.AddAddressIndex(&_Token.TransactOpts, accountAddr, name, symbol)
}

// AddAddressIndex is a paid mutator transaction binding the contract method 0xcf89b5cb.
//
// Solidity: function addAddressIndex(accountAddr address, name string, symbol string) returns()
func (_Token *TokenTransactorSession) AddAddressIndex(accountAddr common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Token.Contract.AddAddressIndex(&_Token.TransactOpts, accountAddr, name, symbol)
}

// TokenGloballogIterator is returned from FilterGloballog and is used to iterate over the raw logs and unpacked data for Globallog events raised by the Token contract.
type TokenGloballogIterator struct {
	Event *TokenGloballog // Event containing the contract specifics and raw log

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
func (it *TokenGloballogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGloballog)
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
		it.Event = new(TokenGloballog)
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
func (it *TokenGloballogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenGloballogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenGloballog represents a Globallog event raised by the Token contract.
type TokenGloballog struct {
	Sender common.Address
	Logstr string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGloballog is a free log retrieval operation binding the contract event 0x4013cd45e5626ff307481e189e2a154e6035bf11795c12ac2010e5275b2f242e.
//
// Solidity: event globallog(sender address, logstr string)
func (_Token *TokenFilterer) FilterGloballog(opts *bind.FilterOpts) (*TokenGloballogIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "globallog")
	if err != nil {
		return nil, err
	}
	return &TokenGloballogIterator{contract: _Token.contract, event: "globallog", logs: logs, sub: sub}, nil
}

// WatchGloballog is a free log subscription operation binding the contract event 0x4013cd45e5626ff307481e189e2a154e6035bf11795c12ac2010e5275b2f242e.
//
// Solidity: event globallog(sender address, logstr string)
func (_Token *TokenFilterer) WatchGloballog(opts *bind.WatchOpts, sink chan<- *TokenGloballog) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "globallog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenGloballog)
				if err := _Token.contract.UnpackLog(event, "globallog", log); err != nil {
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
