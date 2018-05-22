// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accmanager

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"telephone\",\"type\":\"string\"}],\"name\":\"getOneAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"Accounts\",\"outputs\":[{\"name\":\"accountAddr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"string\"},{\"name\":\"accountDescribe\",\"type\":\"string\"},{\"name\":\"role\",\"type\":\"int256\"},{\"name\":\"telephone\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"accountAddr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"string\"},{\"name\":\"describe\",\"type\":\"string\"},{\"name\":\"role\",\"type\":\"int256\"},{\"name\":\"telephone\",\"type\":\"string\"}],\"name\":\"addAccount\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountAddr\",\"type\":\"address\"}],\"name\":\"getAccount\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"string\"},{\"name\":\"accountDescribe\",\"type\":\"string\"},{\"name\":\"role\",\"type\":\"int256\"},{\"name\":\"telephone\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `606060405234156200001057600080fd5b604051620013ab380380620013ab8339810160405280805182019190602001805182019190505081600090805190602001906200004f929190620000b2565b50806001908051906020019062000068929190620000b2565b5033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000161565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000f557805160ff191683800117855562000126565b8280016001018555821562000126579182015b828111156200012557825182559160200191906001019062000108565b5b50905062000135919062000139565b5090565b6200015e91905b808211156200015a57600081600090555060010162000140565b5090565b90565b61123a80620001716000396000f300606060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde031461007d5780636bd449aa1461010b57806395d89b41146101a8578063e203b50614610236578063eed3af8114610466578063fbcbc0f1146105cc575b600080fd5b341561008857600080fd5b6100906107c9565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100d05780820151818401526020810190506100b5565b50505050905090810190601f1680156100fd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561011657600080fd5b610166600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610867565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156101b357600080fd5b6101bb6108fc565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101fb5780820151818401526020810190506101e0565b50505050905090810190601f1680156102285780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561024157600080fd5b61026d600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061099a565b604051808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200180602001806020018681526020018060200185810385528a818151815260200191508051906020019080838360005b838110156102f15780820151818401526020810190506102d6565b50505050905090810190601f16801561031e5780820380516001836020036101000a031916815260200191505b50858103845289818151815260200191508051906020019080838360005b8381101561035757808201518184015260208101905061033c565b50505050905090810190601f1680156103845780820380516001836020036101000a031916815260200191505b50858103835288818151815260200191508051906020019080838360005b838110156103bd5780820151818401526020810190506103a2565b50505050905090810190601f1680156103ea5780820380516001836020036101000a031916815260200191505b50858103825286818151815260200191508051906020019080838360005b83811015610423578082015181840152602081019050610408565b50505050905090810190601f1680156104505780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390f35b341561047157600080fd5b6105b2600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610c56565b604051808215151515815260200191505060405180910390f35b34156105d757600080fd5b610603600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610e57565b604051808060200180602001806020018681526020018060200185810385528a818151815260200191508051906020019080838360005b8381101561065557808201518184015260208101905061063a565b50505050905090810190601f1680156106825780820380516001836020036101000a031916815260200191505b50858103845289818151815260200191508051906020019080838360005b838110156106bb5780820151818401526020810190506106a0565b50505050905090810190601f1680156106e85780820380516001836020036101000a031916815260200191505b50858103835288818151815260200191508051906020019080838360005b83811015610721578082015181840152602081019050610706565b50505050905090810190601f16801561074e5780820380516001836020036101000a031916815260200191505b50858103825286818151815260200191508051906020019080838360005b8381101561078757808201518184015260208101905061076c565b50505050905090810190601f1680156107b45780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b60008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561085f5780601f106108345761010080835404028352916020019161085f565b820191906000526020600020905b81548152906001019060200180831161084257829003601f168201915b505050505081565b60006004826040518082805190602001908083835b6020831015156108a1578051825260208201915060208101905060208303925061087c565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109925780601f1061096757610100808354040283529160200191610992565b820191906000526020600020905b81548152906001019060200180831161097557829003601f168201915b505050505081565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a6c5780601f10610a4157610100808354040283529160200191610a6c565b820191906000526020600020905b815481529060010190602001808311610a4f57829003601f168201915b505050505090806002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b0a5780601f10610adf57610100808354040283529160200191610b0a565b820191906000526020600020905b815481529060010190602001808311610aed57829003601f168201915b505050505090806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ba85780601f10610b7d57610100808354040283529160200191610ba8565b820191906000526020600020905b815481529060010190602001808311610b8b57829003601f168201915b505050505090806004015490806005018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c4c5780601f10610c2157610100808354040283529160200191610c4c565b820191906000526020600020905b815481529060010190602001808311610c2f57829003601f168201915b5050505050905086565b600060c0604051908101604052808873ffffffffffffffffffffffffffffffffffffffff16815260200187815260200186815260200185815260200184815260200183815250600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019080519060200190610d3e929190611155565b506040820151816002019080519060200190610d5b929190611155565b506060820151816003019080519060200190610d78929190611155565b506080820151816004015560a0820151816005019080519060200190610d9f929190611155565b50905050866004836040518082805190602001908083835b602083101515610ddc5780518252602082019150602081019050602083039250610db7565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509695505050505050565b610e5f6111d5565b610e676111d5565b610e6f6111d5565b6000610e796111d5565b6000600360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050806001018160020182600301836004015484600501848054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f665780601f10610f3b57610100808354040283529160200191610f66565b820191906000526020600020905b815481529060010190602001808311610f4957829003601f168201915b50505050509450838054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110025780601f10610fd757610100808354040283529160200191611002565b820191906000526020600020905b815481529060010190602001808311610fe557829003601f168201915b50505050509350828054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561109e5780601f106110735761010080835404028352916020019161109e565b820191906000526020600020905b81548152906001019060200180831161108157829003601f168201915b50505050509250808054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561113a5780601f1061110f5761010080835404028352916020019161113a565b820191906000526020600020905b81548152906001019060200180831161111d57829003601f168201915b50505050509050955095509550955095505091939590929450565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061119657805160ff19168380011785556111c4565b828001600101855582156111c4579182015b828111156111c35782518255916020019190600101906111a8565b5b5090506111d191906111e9565b5090565b602060405190810160405280600081525090565b61120b91905b808211156112075760008160009055506001016111ef565b5090565b905600a165627a7a72305820027b10a3839ce8338a11f54d4b8ef6976d3f0bcc0dc1338f15923e20650534340029`

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

// Accounts is a free data retrieval call binding the contract method 0xe203b506.
//
// Solidity: function Accounts( address) constant returns(accountAddr address, name string, password string, accountDescribe string, role int256, telephone string)
func (_Token *TokenCaller) Accounts(opts *bind.CallOpts, arg0 common.Address) (struct {
	AccountAddr     common.Address
	Name            string
	Password        string
	AccountDescribe string
	Role            *big.Int
	Telephone       string
}, error) {
	ret := new(struct {
		AccountAddr     common.Address
		Name            string
		Password        string
		AccountDescribe string
		Role            *big.Int
		Telephone       string
	})
	out := ret
	err := _Token.contract.Call(opts, out, "Accounts", arg0)
	return *ret, err
}

// Accounts is a free data retrieval call binding the contract method 0xe203b506.
//
// Solidity: function Accounts( address) constant returns(accountAddr address, name string, password string, accountDescribe string, role int256, telephone string)
func (_Token *TokenSession) Accounts(arg0 common.Address) (struct {
	AccountAddr     common.Address
	Name            string
	Password        string
	AccountDescribe string
	Role            *big.Int
	Telephone       string
}, error) {
	return _Token.Contract.Accounts(&_Token.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xe203b506.
//
// Solidity: function Accounts( address) constant returns(accountAddr address, name string, password string, accountDescribe string, role int256, telephone string)
func (_Token *TokenCallerSession) Accounts(arg0 common.Address) (struct {
	AccountAddr     common.Address
	Name            string
	Password        string
	AccountDescribe string
	Role            *big.Int
	Telephone       string
}, error) {
	return _Token.Contract.Accounts(&_Token.CallOpts, arg0)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(accountAddr address) constant returns(name string, password string, accountDescribe string, role int256, telephone string)
func (_Token *TokenCaller) GetAccount(opts *bind.CallOpts, accountAddr common.Address) (struct {
	Name            string
	Password        string
	AccountDescribe string
	Role            *big.Int
	Telephone       string
}, error) {
	ret := new(struct {
		Name            string
		Password        string
		AccountDescribe string
		Role            *big.Int
		Telephone       string
	})
	out := ret
	err := _Token.contract.Call(opts, out, "getAccount", accountAddr)
	return *ret, err
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(accountAddr address) constant returns(name string, password string, accountDescribe string, role int256, telephone string)
func (_Token *TokenSession) GetAccount(accountAddr common.Address) (struct {
	Name            string
	Password        string
	AccountDescribe string
	Role            *big.Int
	Telephone       string
}, error) {
	return _Token.Contract.GetAccount(&_Token.CallOpts, accountAddr)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(accountAddr address) constant returns(name string, password string, accountDescribe string, role int256, telephone string)
func (_Token *TokenCallerSession) GetAccount(accountAddr common.Address) (struct {
	Name            string
	Password        string
	AccountDescribe string
	Role            *big.Int
	Telephone       string
}, error) {
	return _Token.Contract.GetAccount(&_Token.CallOpts, accountAddr)
}

// GetOneAddress is a free data retrieval call binding the contract method 0x6bd449aa.
//
// Solidity: function getOneAddress(telephone string) constant returns(address)
func (_Token *TokenCaller) GetOneAddress(opts *bind.CallOpts, telephone string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getOneAddress", telephone)
	return *ret0, err
}

// GetOneAddress is a free data retrieval call binding the contract method 0x6bd449aa.
//
// Solidity: function getOneAddress(telephone string) constant returns(address)
func (_Token *TokenSession) GetOneAddress(telephone string) (common.Address, error) {
	return _Token.Contract.GetOneAddress(&_Token.CallOpts, telephone)
}

// GetOneAddress is a free data retrieval call binding the contract method 0x6bd449aa.
//
// Solidity: function getOneAddress(telephone string) constant returns(address)
func (_Token *TokenCallerSession) GetOneAddress(telephone string) (common.Address, error) {
	return _Token.Contract.GetOneAddress(&_Token.CallOpts, telephone)
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

// AddAccount is a paid mutator transaction binding the contract method 0xeed3af81.
//
// Solidity: function addAccount(accountAddr address, name string, password string, describe string, role int256, telephone string) returns(success bool)
func (_Token *TokenTransactor) AddAccount(opts *bind.TransactOpts, accountAddr common.Address, name string, password string, describe string, role *big.Int, telephone string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addAccount", accountAddr, name, password, describe, role, telephone)
}

// AddAccount is a paid mutator transaction binding the contract method 0xeed3af81.
//
// Solidity: function addAccount(accountAddr address, name string, password string, describe string, role int256, telephone string) returns(success bool)
func (_Token *TokenSession) AddAccount(accountAddr common.Address, name string, password string, describe string, role *big.Int, telephone string) (*types.Transaction, error) {
	return _Token.Contract.AddAccount(&_Token.TransactOpts, accountAddr, name, password, describe, role, telephone)
}

// AddAccount is a paid mutator transaction binding the contract method 0xeed3af81.
//
// Solidity: function addAccount(accountAddr address, name string, password string, describe string, role int256, telephone string) returns(success bool)
func (_Token *TokenTransactorSession) AddAccount(accountAddr common.Address, name string, password string, describe string, role *big.Int, telephone string) (*types.Transaction, error) {
	return _Token.Contract.AddAccount(&_Token.TransactOpts, accountAddr, name, password, describe, role, telephone)
}
