>1.获取账户例子:http://192.168.234.130:8082/getaccount
```json
{
   "UserAddress": "" 
}
```

>返回：
```json
{
    "StatusCode": 0,
    "UserAccount": {
        "Address": "",
        "Name": "",
        "BankCard": "",
        "WeChat": "",
        "Alipay": "",
        "Telephone": "",
    },
    "Msg": "",
}
```
>2.设置账户:http://192.168.234.130:8082/setaccount
```json
{
   "UserKeyStore": "", 
   "UserParse": "",
   "KeyString": "",
   "UserAccount": {
        "Address": "",
        "Name": "",
        "BankCard": "",
        "WeChat": "",
        "Alipay": "",
        "Telephone": ""
   }
}
```

>返回：
```json
{
    "StatusCode": 0,
    "Msg": "", 
}
```



>6.获取以太币余额:http://192.168.234.130:8082/getethbalance
```json
{
    "UserAddress": ""
}
```

>返回：
```json
{
    "StatusCode": 0,
    "Balance": "", 
    "Msg": "",
}
```
>7.查询排班接口:http://192.168.234.130:8082/getallschedule
```json
{
    "UserAddress": "",
    "Pages": 128
}
```

>返回：
```json
{
    "StatusCode": 0,
    "Schedules": [{
        "Owner": "",
        "CreateTime": 157299223,
        "Status": 0,
        "Rss": [{
            "Accounts": "",
            "Level": 0,
            "Radios": 12.07,
            "SubWay": 0,
            "ResetWay": 0,
            "ResetTime": 0,
            "GetMoney": 12.07,
            "Job": ""
        }],
        "UserAccounts": [{
            "Address": "",
            "Name": "",
            "BankCard": "",
            "WeChat": "",
            "Alipay": "",
            "Telephone": ""
        }],
        "SubCode": "",
        "Message": "",
        "HasPaiBan": true
    }],
    "Pages": 10,
    "PagesCount": 100,
    "Msg": "",
}
```

>9.GetMoney例子:http://192.168.234.130:8082/getmoney
```json
{
    "UserAddress": "",
    "StartTime": 13000,
    "EndTime": 130001,
    "Page": 100
}
```

>返回：
```json
{
    "StatusCode": 0, 
    "AllMoney": 12.07,
    "Month": 12.07,
    "Date": 12.07,
    "Bills":[{
        "Name": "",
        "Money": 12.07,
        "Ratio": 12.07,
        "SubWay": 0,
        "PayAcco": ""
    }],
    "PageCount": 0,
    "Msg": ""
}
```
>10.发布分配表服务例子:http://192.168.234.130:8082/setschedule
```json
{
    "UserAddress": "",
    "UserKeyStore": "",
    "UserParse": "",
    "KeyString": "",
    "ScheduleName": "",
    "Rss": [{
        "Accounts": "",
        "Level": 0,
        "Radios": 12.07,
        "SubWay": 0,
        "ResetWay": 0,
        "ResetTime": 0,
        "GetMoney": 12.07,
        "Job": ""
    }],
    "Message": ""
}
```

>返回：
```json
{
    "StatusCode": 0, 
    "Hash": "", 
    "ScheduleName": "",
    "Msg": ""
}
```
>11查询排班接口例子:http://192.168.234.130:8082/getschedule
```json
{
    "UserAddress": "",
    "ScheduleName": "",
}
```

>返回：
```json
{
    "StatusCode": 0, 
    "Accounts": [{
        "Address": "",
        "Name": "",
        "BankCard": "",
        "WeChat": "",
        "Alipay": "",
        "Telephone": ""
    }], 
    "Schedules":[{
        "Accounts": "",
        "Level": 0,
        "Radios": 12.07,
        "SubWay": 0,
        "ResetWay": 0,
        "ResetTime": 0,
        "GetMoney": 12.07,
        "Job": ""
    }],
    "Msg": ""
}
```

>15发布排班 todo <--（客户端）v2新增:http://192.168.234.130:8082/setpaiban
```json
{
    "OwnerAddress": "",
    "UserKeyStore": "",
    "UserParse": "",
    "UserAddress": "",
    "SubCode": "",
    "PaiBans": [{
        "UserAddress": "",
        "JobName": ""
    }]
}
```

>返回：
```json
{
    "StatusCode": 0, 
    "Msg": ""
}
```
>16查询排班  todo <--（客户端）v2新增:http://192.168.234.130:8082/getpaiban
```json
{
    "OwnerAddress": "",
    "SubCode": "",
}
```

>返回：
```json
{
    "StatusCode": 0, 
    "Msg": "",
    "Roles": ["",""],
    "AddressArray": ["", ""]
}
```
