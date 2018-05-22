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
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"addresslist\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"smartAddr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"storesNumber\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"findkey\",\"type\":\"string\"}],\"name\":\"getAddressIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"storesNumber\",\"type\":\"string\"}],\"name\":\"getAllAddressList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"smartAddr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"addAddressIndex\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"smartAddr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"storesNumber\",\"type\":\"string\"}],\"name\":\"addAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"smartAddr\",\"type\":\"address\"}],\"name\":\"getAddressArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"logstr\",\"type\":\"string\"}],\"name\":\"globallog\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `606060405234156200001057600080fd5b60405162001f1838038062001f188339810160405280805182019190602001805182019190505081600090805190602001906200004f929190620000b2565b50806001908051906020019062000068929190620000b2565b5033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000161565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000f557805160ff191683800117855562000126565b8280016001018555821562000126579182015b828111156200012557825182559160200191906001019062000108565b5b50905062000135919062000139565b5090565b6200015e91905b808211156200015a57600081600090555060010162000140565b5090565b90565b611da780620001716000396000f30060606040526004361061008e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde031461009357806370a7287b1461012157806395d89b411461031157806398285d3b1461039f578063b1088ba314610451578063cf89b5cb146104ee578063cfa474c3146105ad578063da97754e146106c7575b600080fd5b341561009e57600080fd5b6100a6610884565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100e65780820151818401526020810190506100cb565b50505050905090810190601f1680156101135780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561012c57600080fd5b610158600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610922565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156102045780820151818401526020810190506101e9565b50505050905090810190601f1680156102315780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b8381101561026a57808201518184015260208101905061024f565b50505050905090810190601f1680156102975780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b838110156102d05780820151818401526020810190506102b5565b50505050905090810190601f1680156102fd5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b341561031c57600080fd5b610324610b60565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610364578082015181840152602081019050610349565b50505050905090810190601f1680156103915780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103aa57600080fd5b6103fa600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610bfe565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561043d578082015181840152602081019050610422565b505050509050019250505060405180910390f35b341561045c57600080fd5b6104ac600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610cfd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156104f957600080fd5b6105ab600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610d92565b005b34156105b857600080fd5b6106ad600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190505061127d565b604051808215151515815260200191505060405180910390f35b34156106d257600080fd5b6106fe600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506114c3565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561077857808201518184015260208101905061075d565b50505050905090810190601f1680156107a55780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b838110156107de5780820151818401526020810190506107c3565b50505050905090810190601f16801561080b5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015610844578082015181840152602081019050610829565b50505050905090810190601f1680156108715780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b60008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561091a5780601f106108ef5761010080835404028352916020019161091a565b820191906000526020600020905b8154815290600101906020018083116108fd57829003601f168201915b505050505081565b60056020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a1a5780601f106109ef57610100808354040283529160200191610a1a565b820191906000526020600020905b8154815290600101906020018083116109fd57829003601f168201915b505050505090806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ab85780601f10610a8d57610100808354040283529160200191610ab8565b820191906000526020600020905b815481529060010190602001808311610a9b57829003601f168201915b505050505090806004018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b565780601f10610b2b57610100808354040283529160200191610b56565b820191906000526020600020905b815481529060010190602001808311610b3957829003601f168201915b5050505050905085565b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bf65780601f10610bcb57610100808354040283529160200191610bf6565b820191906000526020600020905b815481529060010190602001808311610bd957829003601f168201915b505050505081565b610c06611c6e565b6004826040518082805190602001908083835b602083101515610c3e5780518252602082019150602081019050602083039250610c19565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020805480602002602001604051908101604052809291908181526020018280548015610cf157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610ca7575b50505050509050919050565b60006003826040518082805190602001908083835b602083101515610d375780518252602082019150602081019050602083039250610d12565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6004826040518082805190602001908083835b602083101515610dca5780518252602082019150602081019050602083039250610da5565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054806001018281610e0f9190611c82565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506004816040518082805190602001908083835b602083101515610e965780518252602082019150602081019050602083039250610e71565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054806001018281610edb9190611c82565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506004610f3683836117e9565b6040518082805190602001908083835b602083101515610f6b5780518252602082019150602081019050602083039250610f46565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054806001018281610fb09190611c82565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550507f4013cd45e5626ff307481e189e2a154e6035bf11795c12ac2010e5275b2f242e8383604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611094578082015181840152602081019050611079565b50505050905090810190601f1680156110c15780820380516001836020036101000a031916815260200191505b50935050505060405180910390a17f4013cd45e5626ff307481e189e2a154e6035bf11795c12ac2010e5275b2f242e8382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611164578082015181840152602081019050611149565b50505050905090810190601f1680156111915780820380516001836020036101000a031916815260200191505b50935050505060405180910390a17f4013cd45e5626ff307481e189e2a154e6035bf11795c12ac2010e5275b2f242e836111cb84846117e9565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561123d578082015181840152602081019050611222565b50505050905090810190601f16801561126a5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a1505050565b6000611287611cae565b60a0604051908101604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff16815260200186815260200185815260200184815250600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190805190602001906113c4929190611cc2565b5060608201518160030190805190602001906113e1929190611cc2565b5060808201518160040190805190602001906113fe929190611cc2565b509050508290506000815111156114ba57856003846040518082805190602001908083835b6020831015156114485780518252602082019150602081019050602083039250611423565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50949350505050565b60006114cd611d42565b6114d5611d42565b6114dd611d42565b600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600401828054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561169b5780601f106116705761010080835404028352916020019161169b565b820191906000526020600020905b81548152906001019060200180831161167e57829003601f168201915b50505050509250818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156117375780601f1061170c57610100808354040283529160200191611737565b820191906000526020600020905b81548152906001019060200180831161171a57829003601f168201915b50505050509150808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156117d35780601f106117a8576101008083540402835291602001916117d3565b820191906000526020600020905b8154815290600101906020018083116117b657829003601f168201915b5050505050905093509350935093509193509193565b6117f1611d42565b61182e8383602060405190810160405280600081525060206040519081016040528060008152506020604051908101604052806000815250611836565b905092915050565b61183e611d42565b611846611cae565b61184e611cae565b611856611cae565b61185e611cae565b611866611cae565b61186e611d42565b611876611cae565b6000808e98508d97508c96508b95508a94508451865188518a518c51010101016040518059106118a35750595b9080825280601f01601f1916602001820160405250935083925060009150600090505b88518110156119795788818151811015156118dd57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002838380600101945081518110151561193c57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506118c6565b600090505b8751811015611a3157878181518110151561199557fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f01000000000000000000000000000000000000000000000000000000000000000283838060010194508151811015156119f457fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808060010191505061197e565b600090505b8651811015611ae9578681815181101515611a4d57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000028383806001019450815181101515611aac57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050611a36565b600090505b8551811015611ba1578581815181101515611b0557fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000028383806001019450815181101515611b6457fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050611aee565b600090505b8451811015611c59578481815181101515611bbd57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000028383806001019450815181101515611c1c57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050611ba6565b82995050505050505050505095945050505050565b602060405190810160405280600081525090565b815481835581811511611ca957818360005260206000209182019101611ca89190611d56565b5b505050565b602060405190810160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611d0357805160ff1916838001178555611d31565b82800160010185558215611d31579182015b82811115611d30578251825591602001919060010190611d15565b5b509050611d3e9190611d56565b5090565b602060405190810160405280600081525090565b611d7891905b80821115611d74576000816000905550600101611d5c565b5090565b905600a165627a7a72305820fffeaad211dd4905aa14457beee1399d58e257df830edfca88e672aeae6366fd0029`

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
// Solidity: function addresslist( address) constant returns(owner address, smartAddr address, name string, symbol string, storesNumber string)
func (_Token *TokenCaller) Addresslist(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner        common.Address
	SmartAddr    common.Address
	Name         string
	Symbol       string
	StoresNumber string
}, error) {
	ret := new(struct {
		Owner        common.Address
		SmartAddr    common.Address
		Name         string
		Symbol       string
		StoresNumber string
	})
	out := ret
	err := _Token.contract.Call(opts, out, "addresslist", arg0)
	return *ret, err
}

// Addresslist is a free data retrieval call binding the contract method 0x70a7287b.
//
// Solidity: function addresslist( address) constant returns(owner address, smartAddr address, name string, symbol string, storesNumber string)
func (_Token *TokenSession) Addresslist(arg0 common.Address) (struct {
	Owner        common.Address
	SmartAddr    common.Address
	Name         string
	Symbol       string
	StoresNumber string
}, error) {
	return _Token.Contract.Addresslist(&_Token.CallOpts, arg0)
}

// Addresslist is a free data retrieval call binding the contract method 0x70a7287b.
//
// Solidity: function addresslist( address) constant returns(owner address, smartAddr address, name string, symbol string, storesNumber string)
func (_Token *TokenCallerSession) Addresslist(arg0 common.Address) (struct {
	Owner        common.Address
	SmartAddr    common.Address
	Name         string
	Symbol       string
	StoresNumber string
}, error) {
	return _Token.Contract.Addresslist(&_Token.CallOpts, arg0)
}

// GetAddressArray is a free data retrieval call binding the contract method 0xda97754e.
//
// Solidity: function getAddressArray(smartAddr address) constant returns(address, string, string, string)
func (_Token *TokenCaller) GetAddressArray(opts *bind.CallOpts, smartAddr common.Address) (common.Address, string, string, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Token.contract.Call(opts, out, "getAddressArray", smartAddr)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetAddressArray is a free data retrieval call binding the contract method 0xda97754e.
//
// Solidity: function getAddressArray(smartAddr address) constant returns(address, string, string, string)
func (_Token *TokenSession) GetAddressArray(smartAddr common.Address) (common.Address, string, string, string, error) {
	return _Token.Contract.GetAddressArray(&_Token.CallOpts, smartAddr)
}

// GetAddressArray is a free data retrieval call binding the contract method 0xda97754e.
//
// Solidity: function getAddressArray(smartAddr address) constant returns(address, string, string, string)
func (_Token *TokenCallerSession) GetAddressArray(smartAddr common.Address) (common.Address, string, string, string, error) {
	return _Token.Contract.GetAddressArray(&_Token.CallOpts, smartAddr)
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

// GetAllAddressList is a free data retrieval call binding the contract method 0xb1088ba3.
//
// Solidity: function getAllAddressList(storesNumber string) constant returns(address)
func (_Token *TokenCaller) GetAllAddressList(opts *bind.CallOpts, storesNumber string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getAllAddressList", storesNumber)
	return *ret0, err
}

// GetAllAddressList is a free data retrieval call binding the contract method 0xb1088ba3.
//
// Solidity: function getAllAddressList(storesNumber string) constant returns(address)
func (_Token *TokenSession) GetAllAddressList(storesNumber string) (common.Address, error) {
	return _Token.Contract.GetAllAddressList(&_Token.CallOpts, storesNumber)
}

// GetAllAddressList is a free data retrieval call binding the contract method 0xb1088ba3.
//
// Solidity: function getAllAddressList(storesNumber string) constant returns(address)
func (_Token *TokenCallerSession) GetAllAddressList(storesNumber string) (common.Address, error) {
	return _Token.Contract.GetAllAddressList(&_Token.CallOpts, storesNumber)
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

// AddAddress is a paid mutator transaction binding the contract method 0xcfa474c3.
//
// Solidity: function addAddress(smartAddr address, name string, symbol string, storesNumber string) returns(success bool)
func (_Token *TokenTransactor) AddAddress(opts *bind.TransactOpts, smartAddr common.Address, name string, symbol string, storesNumber string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addAddress", smartAddr, name, symbol, storesNumber)
}

// AddAddress is a paid mutator transaction binding the contract method 0xcfa474c3.
//
// Solidity: function addAddress(smartAddr address, name string, symbol string, storesNumber string) returns(success bool)
func (_Token *TokenSession) AddAddress(smartAddr common.Address, name string, symbol string, storesNumber string) (*types.Transaction, error) {
	return _Token.Contract.AddAddress(&_Token.TransactOpts, smartAddr, name, symbol, storesNumber)
}

// AddAddress is a paid mutator transaction binding the contract method 0xcfa474c3.
//
// Solidity: function addAddress(smartAddr address, name string, symbol string, storesNumber string) returns(success bool)
func (_Token *TokenTransactorSession) AddAddress(smartAddr common.Address, name string, symbol string, storesNumber string) (*types.Transaction, error) {
	return _Token.Contract.AddAddress(&_Token.TransactOpts, smartAddr, name, symbol, storesNumber)
}

// AddAddressIndex is a paid mutator transaction binding the contract method 0xcf89b5cb.
//
// Solidity: function addAddressIndex(smartAddr address, name string, symbol string) returns()
func (_Token *TokenTransactor) AddAddressIndex(opts *bind.TransactOpts, smartAddr common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addAddressIndex", smartAddr, name, symbol)
}

// AddAddressIndex is a paid mutator transaction binding the contract method 0xcf89b5cb.
//
// Solidity: function addAddressIndex(smartAddr address, name string, symbol string) returns()
func (_Token *TokenSession) AddAddressIndex(smartAddr common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Token.Contract.AddAddressIndex(&_Token.TransactOpts, smartAddr, name, symbol)
}

// AddAddressIndex is a paid mutator transaction binding the contract method 0xcf89b5cb.
//
// Solidity: function addAddressIndex(smartAddr address, name string, symbol string) returns()
func (_Token *TokenTransactorSession) AddAddressIndex(smartAddr common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Token.Contract.AddAddressIndex(&_Token.TransactOpts, smartAddr, name, symbol)
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
