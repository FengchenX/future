package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"sub_account_service/fabric_chaincode/contracts"
)

var logger = shim.NewLogger("main")

var accountContract = new(contracts.AccountContract)
var accountBookContract = new(contracts.AccountBookContract)
var scheduleContract = new(contracts.ScheduleContract)
var eventContract = new(contracts.EventContract)

type EntryPoint struct {
}

// Init : implementation for shim.Chaincode interface.
func (s *EntryPoint) Init(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Info("instantiated chaincode")
	return shim.Success([]byte("instantiated chaincode"))
}

// Invoke : implementation for shim.Chaincode interface.
func (s *EntryPoint) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "addAccount":
		return accountContract.AddAccount(stub, args)
	case "getAccount":
		return accountContract.GetAccount(stub, args)
	case "getBalance":
		return accountContract.GetBalance(stub, args)
	case "changePayer":
		return accountContract.ChangePayer(stub, args)
	case "subAccountKeys":
		return accountContract.SubAccountKeys(stub, args)
	case "issueSubCxt":
		return scheduleContract.IssueSubCxt(stub, args)
	case "getIssueSubCxt":
		return scheduleContract.GetIssueSubCxt(stub, args)
	case "resetSubCodeQuotaData":
		return scheduleContract.ResetSubCodeQuotaData(stub, args)
	case "getSubCodeQuotaData":
		return scheduleContract.GetSubCodeQuotaData(stub, args)
	case "issueScheduling":
		return scheduleContract.IssueScheduling(stub, args)
	case "getSchedulingCxt":
		return scheduleContract.GetSchedulingCxt(stub, args)
	case "setSubCodeQuotaData":
		return scheduleContract.SetSubCodeQuotaData(stub, args)
	case "getOneLedgerCxt":
		return accountBookContract.GetOneLedgerCxt(stub, args)
	case "getLedgerSubAddrs":
		return accountBookContract.GetLedgerSubAddrs(stub, args)
	case "updateCalulateLedger":
		return accountBookContract.UpdateCalulateLedger(stub, args)
	}
	msg := fmt.Sprintf("No such function. function = %s, args = %s", function, args)
	logger.Error(msg)
	return shim.Error(msg)
}

func main() {
	if err := shim.Start(new(EntryPoint)); err != nil {
		logger.Errorf("Error creating new Chaincode. Error = %s\n", err)
	}
}
