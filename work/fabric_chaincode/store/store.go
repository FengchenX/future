package store

import (
	."sub_account_service/fabric_chaincode/arguments"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"sub_account_service/fabric_chaincode/models"
	"encoding/json"
)



//发布分配比例
func PutIssue(stub shim.ChaincodeStubInterface ,req ReqIssueSubCxt) error {
	//1.存证排班发起人
	owner:=req.AccountAddr
	err:=stub.PutState("Issuer:"+req.IssueCode,[]byte(owner))
	if err!=nil{
		return err
	}
	ledger:=&IssueBook{
		Issuer:owner,
		IssueCode:req.IssueCode,
	}
	//2.存证分配比率表
	for i, _ := range req.SubRoles {
		item:= models.Issue{
			SubCode:   req.IssueCode,
			Role:      req.SubRoles[i],
			Ratio:     req.Rtaios[i],
			SubWay:    req.SubWays[i],
			QuotaWay:  req.QuotaWays[i],
			ResetTime: req.ResetTimes[i],
		}
		ledger.Content=append(ledger.Content,item)
	}
	jsonBytes,err:=json.Marshal(ledger)
	if err!=nil{
		return err
	}
	err=stub.PutState("IssueBook:"+req.IssueCode,jsonBytes)
	if err!=nil{
		return err
	}
	return nil
}

func GetIssue(stub shim.ChaincodeStubInterface,issueCode string) ( issue IssueBook , err error ) {
	jsonBytes,err:=stub.GetState("IssueBook:"+issueCode)
	if err!=nil{
		return issue,err
	}
	err=json.Unmarshal(jsonBytes,&issue)
	if err!=nil{
		return issue,err
	}
	return
}


//定额分配要特殊存储一下余额
func PutQuota(stub shim.ChaincodeStubInterface,subCode string,quota map[string]uint) error {
	jsonBytes,err:=json.Marshal(quota)
	if err!=nil{
		return err
	}
	err=stub.PutState("Quota:"+subCode,jsonBytes)
	if err!=nil{
		return err
	}
	return nil
}
func GetQuota(stub shim.ChaincodeStubInterface,subCode string)(map[string]uint,error)  {
	jsonBytes,err:=stub.GetState("Quota:"+subCode)
	if err!=nil{
		return nil,err
	}
	var quota map[string]uint
	err=json.Unmarshal(jsonBytes,&quota)
	if err!=nil{
		return nil,err
	}
	return quota,err
}
// 每个职位对应的人
func PutSchedule(stub shim.ChaincodeStubInterface,req ReqIssueScheduling) error {
	var schedule=make(map[[32]byte]string)
	for i,_ := range req.Roles {
		schedule[req.Roles[i]]=req.Joiners[i]
	}
	jsonBytes,err:=json.Marshal(&schedule)
	if err!=nil{
		return err
	}
	err=stub.PutState("Schedule:"+req.IssueCode,jsonBytes)
	if err!=nil{
		return err
	}
	return nil
}

func GetSchedule(stub shim.ChaincodeStubInterface,subCode string) (map[[32]byte]string, error) {
	jsonBytes,err:=stub.GetState("Schedule:"+subCode)
	if err!=nil{
		return nil,err
	}
	var schedule map[[32]byte]string
	err=json.Unmarshal(jsonBytes,&schedule)
	if err!=nil{
		return nil,err
	}
	return schedule, nil
}

func PutAccount(stub shim.ChaincodeStubInterface,req ReqAccount) error {
	acc:=models.Account{
		AccountAddr: req.AccountAddr,
		Name:        req.Name,
		Telephone:   req.Telephone,
		BankCard:    req.BankCard,
		WeChat:      req.WeChat,
		Alipay:      req.Alipay,
		Balance:     0,
	}
	jsonBytes,err:=json.Marshal(acc)
	if err!=nil{
		return err
	}
	err=stub.PutState("Account:"+req.AccountAddr,jsonBytes)
	if err!=nil{
		return err
	}
	return nil
}

func GetAccount(stub shim.ChaincodeStubInterface,account string ) ([]byte, error) {
	jsonBytes,err:=stub.GetState("Account:"+account)
	if err!=nil{
		return nil,err
	}
	return jsonBytes, nil
}


type IssueBook struct {
	Issuer string//发起人
	IssueCode string //编号
	Content   []models.Issue //分配表详情
}