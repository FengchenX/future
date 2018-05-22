# go与智能合约交互接口
* 所有的接口新增返回参数 error err，为接口的错误信息。

## 1.	合约地址索引管理合约接口（address_manager）:
```$xslt
Deploy(token_name string, token_symbol string) (address string, token_name string)

```

* 1.传入两个参数：token_name :当前合约的名字，token_symbol :当前合约的标志，后面检索合约的时候区别使用。  
* 2.返回参数：address:当前合约部署成功后的地址, token_name :当前合约部署成功后，通过合约接口获得的名字，如果返回的名字为空并且传入的名字不为空，需要在外部调用接口获取合约名字，直到成功为止。

* Ps：所有合约部署都是相同接口，前面的package区别（不相同）。

## 2.	部署好的合约，做索引关联接口（address_manager）:
```$xslt
NewAddressAdd(smartAddress string, name string, symbol string) hashcode string 

```
* 1.传入参数：smartAddress:要做索引关联的合约地址；name :要做索引关联的合约名字；symbol : 要做索引关联的合约标志。   		  			  
* 2.返回参数：hashcode 关联成功后返回的hashcode，可以在区块链上通过这个hashcode查看状态。  

## 3.通过合约地址查询关联好的合约地址（address_manager）：
```$xslt
GetAddressArray(smartAddress string) (string,string,string)
```
* 1.传入参数：smartAddress:要做查询的合约地址；
* 2.返回参数：smart_address:合约地址, smart_name:合约名字, smart_sy :合约标识。

## 4.通过合约名字或者标识查询合约地址（address_manager）：
```$xslt
GetAddressByKey(fkey string) []string
```
* 1.传入参数：fkey ：检索条件。
* 2.返回参数：字符串数组，addressList 检索出来的合约地址。


## 5.创建新账户接口（account_manager）:
```$xslt
NewAccount(c_password string) (address string, describe string) 
```

* 1.传入参数：c_password:创建账户的密码，后面合约部署和调用，解锁接口时需要用到，app端需要在本地存一份，后面用到时传入。
* 2.返回参数：address:账户生成后的唯一地址，describe:账户生成后的描述，里面包括账户地址和hash信息，json字符串。  

## 6.关联账户管理账户管理合约接口（account_manager）:
```$xslt
type Account struct {
   AccountAddr     string
   Name            string
   Password        string
   AccountDescribe string
   Role            *big.Int
   Telephone       string
}

NewAccountAdd(account Account) hashcode string

```

* 1.传入参数：accuontAddress:要做关联的账户地址；  
			 name:账户名字；  
			 password :创建账户时的密码；  
			 describe: 账户生成后的描述，里面包括账户地址和hash信息，json字符串。  
			 Role:角色，0表示无身份（默认），1表示经理，2表示厨师，3表示服务员，后面持续新增;
			 Telephone:电话号码；    
* 2.返回参数：hashcode 关联成功后返回的hashcode，可以在区块链上通过这个hashcode查看状态。  

## 7.查询账户信息接口（account_manager）：
```$xslt
type FindAccount struct {
   Name            string
   Password        string
   AccountDescribe string
   Role            *big.Int
   Telephone       string
}
```
```$xslt
1. 通过账户地址查询账户信息：
GetAccountByAddr(operationAddress string,address string) (FindAccount, error)

2. 通过电话号码查询账户信息（对应电话号码登录等）：
GetAccountByTel(operationAddress string, telephone string) (FindAccount, error)
```

## 8. 排班发布（sub_account,后面的接口全部为sub_account模板函数）:
```$xslt
PublishScheduleing(smartAddress string,roles []*big.Int, counts []*big.Int, ratio []*big.Int) hashcode string

```
##### 传入参数：三个参数数组的值一一对应。
```$xslt
roles := []*big.Int{big.NewInt(1), big.NewInt(2)} // 发布职位的角色
counts := []*big.Int{big.NewInt(2), big.NewInt(3)} // 发布职位角色对应的数量
ratio := []*big.Int{big.NewInt(15), big.NewInt(25)} // 发布角色职位对应的分成比例

subaccount.PublishScheduleing(roles, counts, ratio)
```
#### smartAddress ： 表示当前排班合约的合约地址，后面排班合约操作接口都需要这个参数，可能同时存在多个排班合约，需要合约地址唯一确认。
##### 返回参数：hashcode 关联成功后返回的hashcode，可以在区块链上通过这个hashcode查看状态。

## 9.	查询排班发布数据：
```$xslt
type scheduleDesired struct {
    Role  *big.Int
    Count *big.Int
    Ratio *big.Int
}

GetScheduleingCxt(roles []*big.Int) []scheduleDesired
```

## 10.岗位申请(申请之前，先调用一下17 判断是否已经申请过岗位 接口判断一下)：
```$xslt
ApplicationJob(smartAddress string,staffAddr string, role *big.Int) hashcode string
```
* 1.传入参数：staffAddr:申请这个职位的账户地址，role:申请这个职位的账户的角色

## 11.查询岗位申请信息：
```$xslt
type scheduleStaff struct {
    StaffAddr common.Address
    Role      *big.Int
    Ratio     *big.Int
}

GetApplicationCxt() []scheduleStaff
```

## 12.订单信息存证：
```$xslt
SetOrdersContent(smartAddress string,ordersId string, content string) hashcode string
```
* 1.传入参数：ordersId:订单ID，content :订单内容，可以是字符串拼接也可以是json字符串。

## 13.订单信息查询：
```$xslt
GetOrdersContent(smartAddress string,ordersId string) (content  string, creatAt *big.Int)
```
* 1.传入参数：ordersId :订单ID
* 2.返回参数：content :订单内容, creatAt: 存储的时间戳

## 14.每单分账接口:
```$xslt
SettleAccounts(smartAddress string,totalConsume *big.Int) hashcode string
```
* 1.传入参数：totalConsume :要分账的总数.

## 15.查询账户积分（余额）：
```$xslt
GetBalance(smartAddress string,faddress string) *big.Int
```
* 1. 传入参数：faddress：要查询的账户地址;  
* 2. 返回参数：balance ：当前账户积分（余额）。   

## 16.查询合约名字接口：
```$xslt
GetTokenName(operationAddress string) name string
```
* 1.传入参数：operationAddress ： 当前要查询的合约地址。
* 2.返回参数：name ： 合约地址对应的名字。

## 17. 判断是否已经申请过岗位
```$xslt
CheckIsOkApplication(smartAddress string, staffAddr string) bool
```
* `1.` 传入参数：smartAddress：要查询的排班合约地址；staffAddr：申请职位的账户地址；
* `2.` 返回参数：isok为bool类型，true表示已经加入过，false表示为加入。

### （2018.04.20）
## 18. 判断当前账户地址是否是合约的发布者(sub_account) 
```$xslt
CheckOwnerIsOk(smartAddress string, accountAddr string) (bool, error)
```
* `1.` 传入参数：smartAddress：要查询的排班合约地址；accountAddr：要验证的账户地址；
* `2.` 返回参数：isok为bool类型，true表示是，false表示不是。err为error类型，非nil时，要做错误判断和返回。

## 19.获取当前合约的所有订单信息(sub_account)  
```$xslt
GetAllContentHashCxt(smartAddress string) ([]string, error)
```

* `1.`传入参数：smartAddress：要查询的排班合约地址；
* `2.`返回参数：为字符串数组，表示查询到的订单信息。err为error类型，非nil时，要做错误判断

#### (2018.07.26) 以下接口在接口调用新增两个参数（最前面）：
```$xslt
operationKey string  // 当前这个操作的账户私钥
operationPhrase string // 当前这个操作的账户密码
```

* `1.` addrmanager模块的 NewAccountAdd ；
* `2.` accmanager模块的 NewAddressAdd ；
* `3.` subaccount 模块的 所有合约操作接口（包括合约部署，排班发布，申请岗位，订单存证等）

## 分账系统二期需求新增接口 2018.05.08
### 1.添加用人白名单（accmanager）：
```$xslt
AddEmployStaff(operationKey string, operationPhrase string, operationAddress string, staffAddr string) (string, error)
```
* `1.` 传入参数：perationKey：当前这个操作的账户私钥;operationPhrase：当前这个操作的账户密码，operationAddress：账户管理合约的合约地址；staffAddr：要添加的白名单账户地址；
* `2.` 返回参数：hashcode关联成功后返回的hashcode，可以在区块链上通过这个hashcode查看状态。

### 2.根据账户地址查询用人白名单（accmanager）：
```$xslt
GetEmployStaffs(operationAddress string, accountAddress string) ([]FindAccount, []string, error)
```
* `1.` 传入参数：operationAddress：账户管理合约的合约地址；accountAddress：要查询的白名单账户地址；
* `2.` 返回参数：hashcode关联成功后返回的[]FindAccount,做个账户的信息，[]string为多个账户的地址。

### 3.删除用人白名单（accmanager）：
```$xslt
DelEmployStaff(operationKey string, operationPhrase string, operationAddress string, staffAddr string) (string, error)
```
* `1.` 传入参数：perationKey：当前这个操作的账户私钥;operationPhrase：当前这个操作的账户密码，operationAddress：账户管理合约的合约地址；staffAddr：要删除的白名单账户地址；
* `2.` 返回参数：hashcode关联成功后返回的hashcode，可以在区块链上通过这个hashcode查看状态。

### 4.通过门店编号查询当前店的最新合约地址（addrmanager）：
```$xslt
模块：
GetPaySamrtAddress(operationAddress string, stores_number string) (string, error)
```
* `1.` 传入参数：operationAddress：合约地址管理合约的合约地址；stores_number：门店编号；
* `2.` 返回参数：返回查找到的排班合约地址和错误信息。

### 5.获取当前合约的百分比，测试用（subaccount）：
```$xslt
模块：
GetAllRatio(smartAddress string) *big.Int
```
* `1.` 传入参数：smartAddress：当前要查询的排班合约的合约地址；
* `2.` 返回参数：返回查找到的排班合约的比例总和，正常情况为100。

### 6.排班合约发布和排班发布相关接口有参数的修改和删减，具体参看sub_deploy.go的测试文件。
### 7.白名单相关功能接口，参照deploy.go的测试文件。

### 8. 设置排班备注信息：
```$xslt
UpdatePostscriptCxt(operationKey string, operationPhrase string, smartAddress string, u_postscript string) (string, error)

* `1.` 传入参数：u_postscript：要设置的备注信息；
* `2.` 返回参数：关联成功后返回的hashcode和错误信息。
```

### 9. 查询排班备注信息：
```$xslt
GetPostscriptCxt(smartAddress string) string
* `1.` 传入参数：smartAddress：要查询的排班地址；
* `2.` 返回参数：返回备注信息。
```

### 10. 查询发布排班时设置的公司相关信息：
```$xslt
GetCompanyCxt(smartAddress string) (string, string, int64)
* `1.` 传入参数：smartAddress：当前要查询的排班合约的合约地址；
* `2.` 返回参数：companyPayee：公司的收款地址, payer：实时分账的付款地址, companyRatio：公司的比例。
```

