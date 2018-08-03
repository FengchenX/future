package contracts

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"sub_account_service/fabric_chaincode/arguments"
	"encoding/json"
	"sub_account_service/fabric_chaincode/store"
)

type AccountBookContract struct {
}

func (s *AccountBookContract) GetOneLedgerCxt(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success([]byte(""))
}

func (s *AccountBookContract) GetIssueSubCxt(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success([]byte(""))
}

func (s *AccountBookContract) GetLedgerSubAddrs(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success([]byte(""))
}

func (s *AccountBookContract) UpdateCalulateLedger(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success([]byte(""))
}

func (s *AccountBookContract) SettleAccounts(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args)!=1{
		return shim.Error("argument error")
	}
	req:=arguments.ReqSettleAccounts{}
	err:=json.Unmarshal([]byte(args[0]),&req)
	if err!=nil{
		return shim.Error(err.Error())
	}
	issue,err:=store.GetIssue(stub,req.IssueCode)
	if err!=nil{
		return shim.Error(err.Error())
	}
	schedule,err:=store.GetSchedule(stub,req.IssueCode)
	if err!=nil{
		return shim.Error(err.Error())
	}
	quota,err:=store.GetQuota(stub,req.IssueCode)
	if err!=nil{
		return shim.Error(err.Error())
	}
	totalConsume:=req.TotalConsume
	sc_arry:=issue.Content

	var  sum uint;
	var calculate uint;
	c_totalConsume := totalConsume;
	var c_ratio uint= 10000; //剩余比率
	var r_flag uint;
	tmp_num := len(sc_arry)+ 1;
	for j := 0; j< len(sc_arry); j++ {
		if totalConsume == 0{
			break;
		}
		if sc_arry[j].Ratio == 0 {
			continue;
		}
		calculate = 0;
		if (sc_arry[j].SubWay == 1) {// 按定额分配
			// 还未分满
			rquota:=quota[schedule[sc_arry[j].Role]]
			if (sc_arry[j].Ratio >rquota) {
				calculate = sc_arry[j].Ratio - rquota
				//本次分配的金额为未分满的部分，如果未分满的部分大于总额，那么所有钱全部分给他
				if (calculate > totalConsume) {
					calculate = totalConsume
				}
				rquota= rquota + calculate
				//数据回写
				quota[schedule[sc_arry[j].Role]]=rquota
				err=store.PutQuota(stub,req.IssueCode,quota)
				if (r_flag == 1) { //标志已分配过定额
					r_flag = 2
				}
			}
		} else {
			// 按比例分配
			if (issue.Issuer == schedule[sc_arry[j].Role]) {
				tmp_num = j;
			} else {
				if (r_flag >= 2) {
					if (r_flag == 2) {
						c_totalConsume = totalConsume
					}
					r_flag = 3
					calculate = afterCalculateByRatio(c_totalConsume, sc_arry[j].Ratio, c_ratio)
				} else {
					r_flag = 1
					calculate = calculateByRatio(c_totalConsume, sc_arry[j].Ratio)
					c_ratio = c_ratio - sc_arry[j].Ratio
				}
			}
		}

		//update
		if calculate > 0 {
			totalConsume = totalConsume - calculate//剩余
			updateOneLedger("", schedule[sc_arry[j].Role], calculate, req.TransferId)

			// 更新账本，给财务平台查询
			insertLedgerCxt(schedule[sc_arry[j].Role], sc_arry[j].Ratio, sc_arry[j].SubWay, calculate, req.IssueCode, req.TransferId, c_totalConsume)
		}
	}
	if tmp_num >= 0 && tmp_num < len(sc_arry) {
		calculate = totalConsume - sum;
		updateOneLedger("", schedule[sc_arry[tmp_num].Role], calculate, req.TransferId)

		// 更新账本，给财务平台查询
		insertLedgerCxt(schedule[sc_arry[tmp_num].Role], sc_arry[tmp_num].Ratio, sc_arry[tmp_num].SubWay, calculate, req.IssueCode, req.TransferId, c_totalConsume)
	}
	return shim.Success([]byte("success"))
}

//按比例计算
func calculateByRatio(totalConsume uint ,ratio uint) uint {
	var tmp uint = (totalConsume * ratio);
	var calculate uint = (tmp / 10000 - (((tmp / 10000) % 100)));
	return calculate;
}

//定额之后的余额按比例计算
func afterCalculateByRatio(totalConsume uint, ratio uint,c_ratio uint )uint  {
	 var tmp uint= totalConsume * ratio;
	 var calculate uint = ((tmp / 10000 - ((tmp / 10000) % c_ratio)) / c_ratio) * 10000;
	return calculate;
}

func updateOneLedger(sender string,joiner string,calculate uint, transferId string ) error {
	return nil
}

func insertLedgerCxt(joiner string,ratio uint, subWay uint,calculate uint,issueCode string ,transferId string,c_totalConsume uint)  {

}

func checkLedgerById(transferId string) error {
	return nil
}




func (s *AccountBookContract) GetbindIssueToAddr(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success([]byte(""))
}
