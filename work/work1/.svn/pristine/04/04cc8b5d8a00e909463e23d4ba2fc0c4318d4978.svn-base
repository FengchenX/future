#API说明

----------
##0. 概述
	本产品具有用户注册，发布·管理·申请工作，以及付款分账三大功能。
	-  其中用户注册为单一交互界面。
	-  注册完成后
	   - manager(管理员)：拥有查看所有已发布列表（别人），自己所有合约（未开始，进行中，已完成）的界面，并可以修改未开始的合约。
	   - emplyor(雇员)  ：拥有查看所有已发布的合约的列表权限，并可以申请相关的工作。以及 自己所有合约（未开始，进行中，已完成）的界面，等两个交互界面。
	- 当线下活动完成后，manager点击结算，使得线上合约完成，并触发分账。
	- 由于本文主要介绍api,其他场景先不赘述。
	   
	ps - 1.客户端和服务器交互协议为protobuf3(3.5.1为佳)
	ps - 2.协议发送以 **前缀** “req” 和 “resp” 作为区分
 
##1. 注册接口
####  18.4.26新加入‘用户密码’和‘用户私钥’客户端自行保存。
	用户通过提供真实姓名，身份，密码以及手机号向api server发起注册,返回注册状态和自身地址。

	
> // rpc调用方法</br>
> rpc Register (ReqRegister) returns (RespRegister) {}; 

    注册交互协议

>// 客户端 注册</br>
message ReqRegister{</br>
  string PassWord = 1;     // 用户密码</br>
}</br>

>// 服务端 注册</br>
message RespRegister{</br>
  uint32 StatusCode          = 1; // 状态码</br>
  string UserAddress         = 2; // 用户地址</br>
  string PassWord            = 3; // 用户密码</br>
  string AccountDescribe     = 4; // 用户私钥</br>
}</br>

##1.1. 绑定信息接口
####  18.4.26 拆分注册和绑定，注册后才自行绑定信息（需要钱）
> rpc Band (ReqBand) returns (RespBand) {};  

> // 客户端 绑定
message ReqBand{
  string Name               = 1 ;    // 用户名
  string Role               = 2;     // 用户身份
  string PassWord           = 3;     // 用户密码
  string Phone              = 4;     // 注册手机号
  string AccountDescribe    = 5;     // 用户私钥
}

> // 服务端 绑定
message RespBand{
  uint32 StatusCode         = 1; // 状态码
}

##2. 查询注册接口
	查询注册信息
	** 注册后不可及时查询，因为上链速度非常慢，注册后5~30秒才可以查询。
> // rpc调用方法</br>
> rpc Register (ReqRegister) returns (RespRegister) {}; </br>
> 
> // 客户端 查询账户</br>
message ReqCheckAccount{</br>
  string Address   = 1; // 查询用户地址</br>
}</br>

>// 服务端 查询账户
message RespCheckAccount{  </br>
  uint32 StatusCode  = 1;     // 状态码      </br>
  string Name        = 2 ;    // 用户名      </br>
  string Role        = 3;     // 用户身份    </br>
  string PassWord    = 4;     // 用户密码    </br>
  string Phone       = 5;     // 注册手机号  </br>
}</br>
 

##3. 查询合约

	用户通过一定的关键字查询合约（未定义）返回合约切片 []contract{}。

> rpc CheckContract (ReqCheckContract) returns (RespCheckContract) {};       // 查询合约服务
> 
> // 客户端 查询合约</br>
message ReqCheckContract{</br>
    uint64   ContractHash   = 1; // 合约哈希</br>
    string   Key            = 2; // 查询key</br>
}</br>

>// 服务端 查询合约</br>
message RespCheckContract{</br>
  uint32 StatusCode = 1; // 状态码</br>
}</br>

##4. 发布排班
####  18.4.26新加入‘用户密码’和‘用户私钥’客户端自行保存，然后在本接口填入。

	只有拥有manage权限的user拥有发布排班的权力。

>rpc SetSchedule (ReqScheduling) returns (RespScheduling) {};                // 发布排班服务

>// 客户端 发布排班</br>
message ReqScheduling{</br>
    string   UserAddress        = 1;   // 用户地址</br>
    string   PassWord           = 2; // 用户密码</br>
    string   AccountDescribe    = 3; // 用户私钥</br>
    string   Company            = 4;   // 工作单位</br>
    repeated Job Jobs           = 5;   // 发布的工作</br>
}</br>

>// 服务端 发布排班
message RespScheduling{</br>
    uint32 StatusCode = 1; // 状态码</br>
    string Address    = 2; // 合约地址</br>
}</br>


##6. 查询排班

>rpc GetSchedule (ReqGetSchedue) returns (RespGetSchedue){};   // 查询排班服务

>// 客户端 查询排班</br>
message ReqGetSchedue{</br>
    string UserAddress = 1; // 用户地址</br>
    string CompanyName = 2; // 公司名或者是发布者名</br>
    string TimeStamp   = 3; // 时间戳</br>
}</br>

>// 服务端 查询排班
message RespGetSchedue{</br>
	uint32 StatusCode = 1; // 状态码</br>
    repeated Schedule Schedules  = 2; // 返回排班数组</br>
}</br>

##5. 申请工作
####  18.4.26新加入‘用户密码’和‘用户私钥’客户端自行保存,然后在本接口填入。
	所有权限的user都拥有申请工作的权力。

>rpc FindJob (ReqFindJob) returns (RespFindJob) {};                         // 申请工作服务


>// 客户端 申请工作</br>
message ReqFindJob{</br>
    string UserAddress      = 1; // 用户地址</br>
    string  PassWord        = 2; // 用户秘钥</br>
    string AccountDescribe  = 3; // 用户私钥</br>
    Job    MyJob            = 4; // 申请的工作</br>
}</br>

>// 服务端 申请工作</br>
	message RespFindJob{</br>
	uint32 StatusCode = 1; // 状态码</br>
}</br>

##6.检查是否申请工作成功
>rpc CheckIsOkApplication (ReqCheckIsOkApplication) returns (RespCheckIsOkApplication) {}; // 15.检查是否申请工作成功


##7. 提交订单
####  18.4.26新加入‘用户密码’和‘用户私钥’客户端自行保存,然后在本接口填入。

>  rpc SetContent(ReqSetContent) returns (RespSetContent) {};                 // 提交订单服务
>  
>  // 客户端 客户下单
message ReqSetContent{</br>
    string  UserAddress     = 1; // 用户地址</br>
    string  PassWord        = 2; // 用户秘钥</br>
    string AccountDescribe  = 3; // 用户私钥</br>
    string  OrderId         = 4; // 订单ID</br>
    Order   Content         = 5; // 订单内容（json字符串也是可以的）</br>
    string  JobAddress      = 6; // 订单地址</br>
}</br>

> // 服务端 客户下单</br>
message RespSetContent{</br>
 	uint32 StatusCode = 1; // 状态码</br>
 	string OrderId    = 2; // 返回的訂單號</br>
}</br>

##8. 查单
>  rpc GetContent(ReqGetContent) returns (RespGetContent) {};                 // 查单

>  // 客户端 客户查单
message ReqGetContent{</br>
    string  UserAddress = 1; // 用户地址</br>
    string  OrderId     = 2; // 订单ID</br>
    string  JobAddress  = 3; // 订单地址</br>
}</br>

>// 服务端 客户查单</br>
message RespGetContent{</br>
  uint32 StatusCode = 1; // 状态码</br>
  uint64 Money      = 2; // 下单金额</br>
}</br>


##9. 付款
####  18.4.26新加入‘用户密码’和‘用户私钥’客户端自行保存,然后在本接口填入。

	只有发布合约的manager拥有操作点击付款的权限。

>  rpc Pay (ReqPay) returns (RespPay) {};                                     // 申请付款服务

>// 客户端 确认付款</br>
 message ReqPay{</br>
    string  UserAddress     = 1; // 用户地址</br>
    string  PassWord        = 2; // 用户秘钥</br>
    string AccountDescribe  = 3; // 用户私钥</br>
    uint64  Money           = 4; // 申请的工作</br>
    string  Address         = 5; // 订单地址</br>
}</br>

>// 服务端 确认付款</br>
 message RespPay{</br>
  uint32 StatusCode = 1; // 状态码</br>
}</br>


##10. 查询某一班的收入
	用以查询某一班的收入
>  rpc GetBalance(ReqGetBalance) returns (RespGetBalance) {};            

>// 请求查询某一排班的收入
message ReqGetBalance{</br>
    string SchedueAddress        = 1; // 排班地址</br>
    string AccountAddress        = 2; // 账户地址</br>
}</br>

>// 请求查询某一排班的收入</br>
message RespGetBalance{</br>
    uint32 StatusCode             = 1; // 状态码</br>
    uint64 Money                  = 2; // 收取的钱</br>
}</br>
 

##11.查询用户在某一公司所有收入
>rpc GetAllMoney (ReqGetAllMoney) returns (RespGetAllMoney) {};  

>// 客户端 查询用户在某一公司所有收入</br>
message ReqGetAllMoney{</br>
    string   CompanyName           = 1; // 公司名</br>
    string   UserAddress           = 2; // 用戶地址</br>
}</br>

>// 服务端 查詢排班下所有訂單</br>
message RespGetAllMoney{</br>
    uint32   StatusCode             = 1; // 状态码</br>
    double   Sum                    = 2; // 订单总价</br>
}</br>                          

##12.查詢排班下所有訂單
>rpc GetAllOrder (ReqGetAllOrder)returns (RespGetAllOrder) {};
 
>// 客户端 查詢排班下所有訂單</br> 
  message ReqGetAllOrder{</br> 
      string   UserAddress       = 1; // 用戶地址</br> 
      string   JobAddress        = 2; // 排班合约地址</br> 
  }</br> 
  
>// 服务端 查詢排班下所有訂單</br> 
  message RespGetAllOrder{</br> 
      uint32   StatusCode             = 1; // 状态码</br> 
      repeated Order Orders           = 2; // 訂單詳情</br> 
      double   Sum                    = 3; // 订单总价</br> 
  }</br> 
  
##13.用户登录服务
  >rpc Login (ReqLogin) returns (RespLogin) {};
 
 >// 客户端 登录服务</br> 
 message ReqLogin{</br> 
     string   Phone            = 1; // 用户电话号码</br> 
 }
 
  >// 服务端 登录服务</br> 
 message RespLogin{</br> 
     uint32 StatusCode             = 1; // 状态码</br> 
     string Name                   = 2; // 用户名</br> 
     string Role                   = 3; // 用户身份</br> 
     string PassWord               = 4; // 用户密码</br> 
     string Phone                  = 5; // 注册手机号</br> 
     string Address                = 6; // 用户地址</br> 
     string AccountDescribe        = 7; // 用户私钥</br> 
 } </br> 
 