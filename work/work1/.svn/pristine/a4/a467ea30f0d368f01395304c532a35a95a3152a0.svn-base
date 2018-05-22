// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accmanager

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllAccountsAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"accountAddr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"string\"},{\"name\":\"describe\",\"type\":\"string\"},{\"name\":\"telephone\",\"type\":\"string\"}],\"name\":\"addAccount\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"telephone\",\"type\":\"string\"}],\"name\":\"getOneAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getEmployStaffs\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"staffAddr\",\"type\":\"address\"}],\"name\":\"addEmployStaff\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"staffAddr\",\"type\":\"address\"}],\"name\":\"delEmployStaff\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"Accounts\",\"outputs\":[{\"name\":\"accountAddr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"string\"},{\"name\":\"accountDescribe\",\"type\":\"string\"},{\"name\":\"telephone\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountAddr\",\"type\":\"address\"}],\"name\":\"getAccount\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"string\"},{\"name\":\"accountDescribe\",\"type\":\"string\"},{\"name\":\"telephone\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `606060405234156200001057600080fd5b60405162001b7038038062001b708339810160405280805182019190602001805182019190505081600090805190602001906200004f929190620000b2565b50806001908051906020019062000068929190620000b2565b5033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000161565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000f557805160ff191683800117855562000126565b8280016001018555821562000126579182015b828111156200012557825182559160200191906001019062000108565b5b50905062000135919062000139565b5090565b6200015e91905b808211156200015a57600081600090555060010162000140565b5090565b90565b6119ff80620001716000396000f3006060604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100a95780630edee5d31461013757806314e59350146101a15780636bd449aa146102fe5780638538984b146103b057806392cbbcd81461043e57806395d89b411461048f578063d9ce60981461051d578063e203b5061461056e578063fbcbc0f114610797575b600080fd5b34156100b457600080fd5b6100bc61098d565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100fc5780820151818401526020810190506100e1565b50505050905090810190601f1680156101295780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561014257600080fd5b61014a610a2b565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561018d578082015181840152602081019050610172565b505050509050019250505060405180910390f35b34156101ac57600080fd5b6102e4600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610abf565b604051808215151515815260200191505060405180910390f35b341561030957600080fd5b610359600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610d34565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561039c578082015181840152602081019050610381565b505050509050019250505060405180910390f35b34156103bb57600080fd5b6103e7600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610e33565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561042a57808201518184015260208101905061040f565b505050509050019250505060405180910390f35b341561044957600080fd5b610475600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610f06565b604051808215151515815260200191505060405180910390f35b341561049a57600080fd5b6104a2611091565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104e25780820151818401526020810190506104c7565b50505050905090810190601f16801561050f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561052857600080fd5b610554600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061112f565b604051808215151515815260200191505060405180910390f35b341561057957600080fd5b6105a5600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061129c565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001806020018060200180602001858103855289818151815260200191508051906020019080838360005b83811015610623578082015181840152602081019050610608565b50505050905090810190601f1680156106505780820380516001836020036101000a031916815260200191505b50858103845288818151815260200191508051906020019080838360005b8381101561068957808201518184015260208101905061066e565b50505050905090810190601f1680156106b65780820380516001836020036101000a031916815260200191505b50858103835287818151815260200191508051906020019080838360005b838110156106ef5780820151818401526020810190506106d4565b50505050905090810190601f16801561071c5780820380516001836020036101000a031916815260200191505b50858103825286818151815260200191508051906020019080838360005b8381101561075557808201518184015260208101905061073a565b50505050905090810190601f1680156107825780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b34156107a257600080fd5b6107ce600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611552565b6040518080602001806020018060200180602001858103855289818151815260200191508051906020019080838360005b8381101561081a5780820151818401526020810190506107ff565b50505050905090810190601f1680156108475780820380516001836020036101000a031916815260200191505b50858103845288818151815260200191508051906020019080838360005b83811015610880578082015181840152602081019050610865565b50505050905090810190601f1680156108ad5780820380516001836020036101000a031916815260200191505b50858103835287818151815260200191508051906020019080838360005b838110156108e65780820151818401526020810190506108cb565b50505050905090810190601f1680156109135780820380516001836020036101000a031916815260200191505b50858103825286818151815260200191508051906020019080838360005b8381101561094c578082015181840152602081019050610931565b50505050905090810190601f1680156109795780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b60008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a235780601f106109f857610100808354040283529160200191610a23565b820191906000526020600020905b815481529060010190602001808311610a0657829003601f168201915b505050505081565b610a33611845565b6006805480602002602001604051908101604052809291908181526020018280548015610ab557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610a6b575b5050505050905090565b600060a0604051908101604052808773ffffffffffffffffffffffffffffffffffffffff16815260200186815260200185815260200184815260200183815250600460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019080519060200190610ba1929190611859565b506040820151816002019080519060200190610bbe929190611859565b506060820151816003019080519060200190610bdb929190611859565b506080820151816004019080519060200190610bf8929190611859565b509050506005826040518082805190602001908083835b602083101515610c345780518252602082019150602081019050602083039250610c0f565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054806001018281610c7991906118d9565b9160005260206000209001600088909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060068054806001018281610cdc91906118d9565b9160005260206000209001600088909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505095945050505050565b610d3c611845565b6005826040518082805190602001908083835b602083101515610d745780518252602082019150602081019050602083039250610d4f565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020805480602002602001604051908101604052809291908181526020018280548015610e2757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610ddd575b50505050509050919050565b610e3b611845565b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610efa57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610eb0575b50505050509050919050565b600080600080600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091505b8280549050821015610fe4578282815481101515610f6c57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610fd75760009350611089565b8180600101925050610f52565b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480600101828161103591906118d9565b9160005260206000209001600087909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600193505b505050919050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111275780601f106110fc57610100808354040283529160200191611127565b820191906000526020600020905b81548152906001019060200180831161110a57829003601f168201915b505050505081565b600080600080600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091505b828054905082101561129457828281548110151561119557fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561128757828281548110151561120557fe5b906000526020600020900160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905582600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020908054611281929190611905565b50611294565b818060010192505061117b565b505050919050565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561136e5780601f106113435761010080835404028352916020019161136e565b820191906000526020600020905b81548152906001019060200180831161135157829003601f168201915b505050505090806002018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561140c5780601f106113e15761010080835404028352916020019161140c565b820191906000526020600020905b8154815290600101906020018083116113ef57829003601f168201915b505050505090806003018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156114aa5780601f1061147f576101008083540402835291602001916114aa565b820191906000526020600020905b81548152906001019060200180831161148d57829003601f168201915b505050505090806004018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156115485780601f1061151d57610100808354040283529160200191611548565b820191906000526020600020905b81548152906001019060200180831161152b57829003601f168201915b5050505050905085565b61155a611957565b611562611957565b61156a611957565b611572611957565b6000600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905080600101816002018260030183600401838054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561165a5780601f1061162f5761010080835404028352916020019161165a565b820191906000526020600020905b81548152906001019060200180831161163d57829003601f168201915b50505050509350828054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156116f65780601f106116cb576101008083540402835291602001916116f6565b820191906000526020600020905b8154815290600101906020018083116116d957829003601f168201915b50505050509250818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156117925780601f1061176757610100808354040283529160200191611792565b820191906000526020600020905b81548152906001019060200180831161177557829003601f168201915b50505050509150808054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561182e5780601f106118035761010080835404028352916020019161182e565b820191906000526020600020905b81548152906001019060200180831161181157829003601f168201915b505050505090509450945094509450509193509193565b602060405190810160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061189a57805160ff19168380011785556118c8565b828001600101855582156118c8579182015b828111156118c75782518255916020019190600101906118ac565b5b5090506118d5919061196b565b5090565b815481835581811511611900578183600052602060002091820191016118ff919061196b565b5b505050565b8280548282559060005260206000209081019282156119465760005260206000209182015b8281111561194557825482559160010191906001019061192a565b5b5090506119539190611990565b5090565b602060405190810160405280600081525090565b61198d91905b80821115611989576000816000905550600101611971565b5090565b90565b6119d091905b808211156119cc57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101611996565b5090565b905600a165627a7a723058201b01f9b7ccdb414054e8b00340dceb1e8e60f446f1f72b2ade9f5aac27d7a30b0029`

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
// Solidity: function Accounts( address) constant returns(accountAddr address, name string, password string, accountDescribe string, telephone string)
func (_Token *TokenCaller) Accounts(opts *bind.CallOpts, arg0 common.Address) (struct {
	AccountAddr     common.Address
	Name            string
	Password        string
	AccountDescribe string
	Telephone       string
}, error) {
	ret := new(struct {
		AccountAddr     common.Address
		Name            string
		Password        string
		AccountDescribe string
		Telephone       string
	})
	out := ret
	err := _Token.contract.Call(opts, out, "Accounts", arg0)
	return *ret, err
}

// Accounts is a free data retrieval call binding the contract method 0xe203b506.
//
// Solidity: function Accounts( address) constant returns(accountAddr address, name string, password string, accountDescribe string, telephone string)
func (_Token *TokenSession) Accounts(arg0 common.Address) (struct {
	AccountAddr     common.Address
	Name            string
	Password        string
	AccountDescribe string
	Telephone       string
}, error) {
	return _Token.Contract.Accounts(&_Token.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xe203b506.
//
// Solidity: function Accounts( address) constant returns(accountAddr address, name string, password string, accountDescribe string, telephone string)
func (_Token *TokenCallerSession) Accounts(arg0 common.Address) (struct {
	AccountAddr     common.Address
	Name            string
	Password        string
	AccountDescribe string
	Telephone       string
}, error) {
	return _Token.Contract.Accounts(&_Token.CallOpts, arg0)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(accountAddr address) constant returns(name string, password string, accountDescribe string, telephone string)
func (_Token *TokenCaller) GetAccount(opts *bind.CallOpts, accountAddr common.Address) (struct {
	Name            string
	Password        string
	AccountDescribe string
	Telephone       string
}, error) {
	ret := new(struct {
		Name            string
		Password        string
		AccountDescribe string
		Telephone       string
	})
	out := ret
	err := _Token.contract.Call(opts, out, "getAccount", accountAddr)
	return *ret, err
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(accountAddr address) constant returns(name string, password string, accountDescribe string, telephone string)
func (_Token *TokenSession) GetAccount(accountAddr common.Address) (struct {
	Name            string
	Password        string
	AccountDescribe string
	Telephone       string
}, error) {
	return _Token.Contract.GetAccount(&_Token.CallOpts, accountAddr)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(accountAddr address) constant returns(name string, password string, accountDescribe string, telephone string)
func (_Token *TokenCallerSession) GetAccount(accountAddr common.Address) (struct {
	Name            string
	Password        string
	AccountDescribe string
	Telephone       string
}, error) {
	return _Token.Contract.GetAccount(&_Token.CallOpts, accountAddr)
}

// GetAllAccountsAddr is a free data retrieval call binding the contract method 0x0edee5d3.
//
// Solidity: function getAllAccountsAddr() constant returns(address[])
func (_Token *TokenCaller) GetAllAccountsAddr(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getAllAccountsAddr")
	return *ret0, err
}

// GetAllAccountsAddr is a free data retrieval call binding the contract method 0x0edee5d3.
//
// Solidity: function getAllAccountsAddr() constant returns(address[])
func (_Token *TokenSession) GetAllAccountsAddr() ([]common.Address, error) {
	return _Token.Contract.GetAllAccountsAddr(&_Token.CallOpts)
}

// GetAllAccountsAddr is a free data retrieval call binding the contract method 0x0edee5d3.
//
// Solidity: function getAllAccountsAddr() constant returns(address[])
func (_Token *TokenCallerSession) GetAllAccountsAddr() ([]common.Address, error) {
	return _Token.Contract.GetAllAccountsAddr(&_Token.CallOpts)
}

// GetEmployStaffs is a free data retrieval call binding the contract method 0x8538984b.
//
// Solidity: function getEmployStaffs(owner address) constant returns(address[])
func (_Token *TokenCaller) GetEmployStaffs(opts *bind.CallOpts, owner common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getEmployStaffs", owner)
	return *ret0, err
}

// GetEmployStaffs is a free data retrieval call binding the contract method 0x8538984b.
//
// Solidity: function getEmployStaffs(owner address) constant returns(address[])
func (_Token *TokenSession) GetEmployStaffs(owner common.Address) ([]common.Address, error) {
	return _Token.Contract.GetEmployStaffs(&_Token.CallOpts, owner)
}

// GetEmployStaffs is a free data retrieval call binding the contract method 0x8538984b.
//
// Solidity: function getEmployStaffs(owner address) constant returns(address[])
func (_Token *TokenCallerSession) GetEmployStaffs(owner common.Address) ([]common.Address, error) {
	return _Token.Contract.GetEmployStaffs(&_Token.CallOpts, owner)
}

// GetOneAddress is a free data retrieval call binding the contract method 0x6bd449aa.
//
// Solidity: function getOneAddress(telephone string) constant returns(address[])
func (_Token *TokenCaller) GetOneAddress(opts *bind.CallOpts, telephone string) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getOneAddress", telephone)
	return *ret0, err
}

// GetOneAddress is a free data retrieval call binding the contract method 0x6bd449aa.
//
// Solidity: function getOneAddress(telephone string) constant returns(address[])
func (_Token *TokenSession) GetOneAddress(telephone string) ([]common.Address, error) {
	return _Token.Contract.GetOneAddress(&_Token.CallOpts, telephone)
}

// GetOneAddress is a free data retrieval call binding the contract method 0x6bd449aa.
//
// Solidity: function getOneAddress(telephone string) constant returns(address[])
func (_Token *TokenCallerSession) GetOneAddress(telephone string) ([]common.Address, error) {
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

// AddAccount is a paid mutator transaction binding the contract method 0x14e59350.
//
// Solidity: function addAccount(accountAddr address, name string, password string, describe string, telephone string) returns(success bool)
func (_Token *TokenTransactor) AddAccount(opts *bind.TransactOpts, accountAddr common.Address, name string, password string, describe string, telephone string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addAccount", accountAddr, name, password, describe, telephone)
}

// AddAccount is a paid mutator transaction binding the contract method 0x14e59350.
//
// Solidity: function addAccount(accountAddr address, name string, password string, describe string, telephone string) returns(success bool)
func (_Token *TokenSession) AddAccount(accountAddr common.Address, name string, password string, describe string, telephone string) (*types.Transaction, error) {
	return _Token.Contract.AddAccount(&_Token.TransactOpts, accountAddr, name, password, describe, telephone)
}

// AddAccount is a paid mutator transaction binding the contract method 0x14e59350.
//
// Solidity: function addAccount(accountAddr address, name string, password string, describe string, telephone string) returns(success bool)
func (_Token *TokenTransactorSession) AddAccount(accountAddr common.Address, name string, password string, describe string, telephone string) (*types.Transaction, error) {
	return _Token.Contract.AddAccount(&_Token.TransactOpts, accountAddr, name, password, describe, telephone)
}

// AddEmployStaff is a paid mutator transaction binding the contract method 0x92cbbcd8.
//
// Solidity: function addEmployStaff(staffAddr address) returns(bool)
func (_Token *TokenTransactor) AddEmployStaff(opts *bind.TransactOpts, staffAddr common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addEmployStaff", staffAddr)
}

// AddEmployStaff is a paid mutator transaction binding the contract method 0x92cbbcd8.
//
// Solidity: function addEmployStaff(staffAddr address) returns(bool)
func (_Token *TokenSession) AddEmployStaff(staffAddr common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddEmployStaff(&_Token.TransactOpts, staffAddr)
}

// AddEmployStaff is a paid mutator transaction binding the contract method 0x92cbbcd8.
//
// Solidity: function addEmployStaff(staffAddr address) returns(bool)
func (_Token *TokenTransactorSession) AddEmployStaff(staffAddr common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddEmployStaff(&_Token.TransactOpts, staffAddr)
}

// DelEmployStaff is a paid mutator transaction binding the contract method 0xd9ce6098.
//
// Solidity: function delEmployStaff(staffAddr address) returns(bool)
func (_Token *TokenTransactor) DelEmployStaff(opts *bind.TransactOpts, staffAddr common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "delEmployStaff", staffAddr)
}

// DelEmployStaff is a paid mutator transaction binding the contract method 0xd9ce6098.
//
// Solidity: function delEmployStaff(staffAddr address) returns(bool)
func (_Token *TokenSession) DelEmployStaff(staffAddr common.Address) (*types.Transaction, error) {
	return _Token.Contract.DelEmployStaff(&_Token.TransactOpts, staffAddr)
}

// DelEmployStaff is a paid mutator transaction binding the contract method 0xd9ce6098.
//
// Solidity: function delEmployStaff(staffAddr address) returns(bool)
func (_Token *TokenTransactorSession) DelEmployStaff(staffAddr common.Address) (*types.Transaction, error) {
	return _Token.Contract.DelEmployStaff(&_Token.TransactOpts, staffAddr)
}
