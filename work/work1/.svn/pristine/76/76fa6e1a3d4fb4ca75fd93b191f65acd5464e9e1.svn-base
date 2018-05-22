pragma solidity ^0.4.21;

/*
    1.该合约主要实现账户管理和索引功能。
*/

contract AccountManager {
    string public name; // 合约的名字
    string public symbol; // 合约的标记或者符号，如 早中晚班
    address issuer; // 部署合约的账户地址

    struct AccountInfo {
        address accountAddr;
        string name;
        string password;
        string accountDescribe;
        int role; // 角色，0表示无身份（默认），1表示经理，2表示厨师，3表示服务员，后面持续新增
        string telephone; // 电话号码
    }

    mapping(address => AccountInfo) public Accounts;
    mapping(string => address) AccAddressList;

    function AccountManager(string tokenName, string tokenSymbol) {
        name = tokenName;
        symbol = tokenSymbol;
        issuer = msg.sender;
    }

    function addAccount(address accountAddr, string name, string password, string describe, int role, string telephone) returns (bool success) {
        Accounts[accountAddr] = AccountInfo(accountAddr, name, password, describe, role, telephone);

        AccAddressList[telephone] = accountAddr;
    }

    function getAccount(address accountAddr) constant returns (string name, string password, string accountDescribe, int role, string telephone){
        AccountInfo account = Accounts[accountAddr];
        return (account.name, account.password, account.accountDescribe, account.role, account.telephone);
    }

    function getOneAddress(string telephone) constant returns (address) {
        return AccAddressList[telephone];
    }
}