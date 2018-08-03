package contracts

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"sub_account_service/fabric_chaincode/arguments"
	"sub_account_service/fabric_chaincode/store"
)


type AccountContract struct {
}

func (s *AccountContract) AddAccount(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("parameters are wrong")
	}
	req:=arguments.ReqAccount{}
	err:=json.Unmarshal([]byte(args[0]),&req)
	if err!=nil{
		return shim.Error(err.Error())
	}
	err=store.PutAccount(stub,req)
	if err!=nil{
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("AddAccount success"))
}

func (s *AccountContract) GetAccount(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("")
	}
	accountNo := args[0]
	jsonBytes, err :=store.GetAccount(stub,accountNo)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(jsonBytes)
}

func (s *AccountContract) GetBalance(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	return shim.Success([]byte(""))
}

func (s *AccountContract) ChangePayer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	return shim.Success([]byte(""))
}

func (s *AccountContract) SubAccountKeys(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	return shim.Success([]byte(""))
}
