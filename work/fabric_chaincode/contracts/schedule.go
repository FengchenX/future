package contracts

import (
	"encoding/json"
	"sub_account_service/fabric_chaincode/arguments"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"sub_account_service/fabric_chaincode/store"
)

type ScheduleContract struct {
}

/*
	1. 发布编号对应的分账人和比例
	ps: 分配人里没自己和比例和超过100，失败，调用合约之前判断。
	2. 传入是数组顺序按设置的权重排序
*/
func (s *ScheduleContract) IssueSubCxt(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("argument error")
	}
	req := arguments.ReqIssueSubCxt{}
	err := json.Unmarshal([]byte(args[0]), &req)
	if err != nil {
		return shim.Error("Failure of parameter parsing")
	}
	if len(req.SubRoles) != len(req.Rtaios) {
		return shim.Error("argument error")
	}
	err=store.PutIssue(stub,req)
	if err!=nil{
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("IssueSubCxt success"))
}
func (s *ScheduleContract) GetIssueSubCxt(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args)!=1{
		return shim.Error("argument error")
	}
	issueCode:=args[0]
	issue,err:=store.GetIssue(stub,issueCode)
	if err!=nil{
		return shim.Error(err.Error())
	}
	jsonBytes,err:=json.Marshal(issue)
	if err!=nil{
		return shim.Error(err.Error())
	}
	return shim.Success(jsonBytes)
}

func (s *ScheduleContract) ResetSubCodeQuotaData(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success([]byte(""))
}

func (s *ScheduleContract) GetSubCodeQuotaData(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success([]byte(""))
}

func (s *ScheduleContract) IssueScheduling(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args)!=1{
		return shim.Error("argument error")
	}
	req:=arguments.ReqIssueScheduling{}
	err:=json.Unmarshal([]byte(args[0]),&req)
	if err!=nil{
		return shim.Error("json Unmarshal error")
	}
	err=store.PutSchedule(stub,req)
	return shim.Success([]byte("IssueScheduling success"))
}


func (s *ScheduleContract) GetSchedulingCxt(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args)!=1{
		return  shim.Error("argument error")
	}
	issueCode:=args[0]
	jsonBytes,err:=stub.GetState("schedule:"+issueCode)
	if err!=nil{
		return shim.Error(err.Error())
	}
	return shim.Success(jsonBytes)
}


func (s *ScheduleContract) SetSubCodeQuotaData(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success([]byte(""))
}
