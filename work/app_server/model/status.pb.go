// Code generated by protoc-gen-go. DO NOT EDIT.
// source: status.proto

package model

type Status int32

const (
	Status_Success            Status = 0
	Status_ManageContractFail Status = 20001
	Status_FindContractFail   Status = 20002
	Status_UndateContractFail Status = 20003
	Status_AddContractFail    Status = 20004
	Status_DelContractFail    Status = 20005
	Status_SetSchedueFail     Status = 30001
	Status_GetSchedueFail     Status = 30002
	Status_FindJobFail        Status = 30003
	Status_JobLimitFail       Status = 30004
	Status_GetContentFail     Status = 30005
	Status_SetContentFail     Status = 30006
	Status_GetBalanceFail     Status = 30007
	Status_RoleNotFitFail     Status = 30008
	Status_HasAppliedJobFail  Status = 30009
	Status_LoginFail          Status = 30010
	Status_GetAllIncomeFail   Status = 30011
	Status_AlreadyApply       Status = 30012
	Status_ApplyFull          Status = 30013
	Status_ReqNotFull         Status = 30014
	Status_SetAccountFail     Status = 30015
	Status_GetAccountFail     Status = 30016
	Status_GetAccountBookFail Status = 30017
	Status_ThreeConfirmFail   Status = 30018
	Status_GetByHashFail      Status = 30019
	Status_ThreeSetBillFail   Status = 30020
	Status_ABByIdFail         Status = 30021
)

var Status_name = map[int32]string{
	0:     "Success",
	20001: "ManageContractFail",
	20002: "FindContractFail",
	20003: "UndateContractFail",
	20004: "AddContractFail",
	20005: "DelContractFail",
	30001: "SetSchedueFail",
	30002: "GetSchedueFail",
	30003: "FindJobFail",
	30004: "JobLimitFail",
	30005: "GetContentFail",
	30006: "SetContentFail",
	30007: "GetBalanceFail",
	30008: "RoleNotFitFail",
	30009: "HasAppliedJobFail",
	30010: "LoginFail",
	30011: "GetAllIncomeFail",
	30012: "AlreadyApply",
	30013: "ApplyFull",
	30014: "ReqNotFull",
	30015: "SetAccountFail",
	30016: "GetAccountFail",
	30017: "GetAccountBookFail",
	30018: "ThreeConfirmFail",
	30019: "GetByHashFail",
	30020: "ThreeSetBillFail",
	30021: "ABByIdFail",
}
var Status_value = map[string]int32{
	"Success":            0,
	"ManageContractFail": 20001,
	"FindContractFail":   20002,
	"UndateContractFail": 20003,
	"AddContractFail":    20004,
	"DelContractFail":    20005,
	"SetSchedueFail":     30001,
	"GetSchedueFail":     30002,
	"FindJobFail":        30003,
	"JobLimitFail":       30004,
	"GetContentFail":     30005,
	"SetContentFail":     30006,
	"GetBalanceFail":     30007,
	"RoleNotFitFail":     30008,
	"HasAppliedJobFail":  30009,
	"LoginFail":          30010,
	"GetAllIncomeFail":   30011,
	"AlreadyApply":       30012,
	"ApplyFull":          30013,
	"ReqNotFull":         30014,
	"SetAccountFail":     30015,
	"GetAccountFail":     30016,
	"GetAccountBookFail": 30017,
	"ThreeConfirmFail":   30018,
	"GetByHashFail":      30019,
	"ThreeSetBillFail":   30020,
	"ABByIdFail":         30021,
}

type Hash int32

const (
	Hash_Successful Hash = 0
	Hash_Pending    Hash = 1
	Hash_Fail       Hash = 2
)

var Hash_name = map[int32]string{
	0: "Successful",
	1: "Pending",
	2: "Fail",
}
var Hash_value = map[string]int32{
	"Successful": 0,
	"Pending":    1,
	"Fail":       2,
}
