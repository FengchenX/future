pragma solidity ^0.4.21;

/*
    by GrApple 2018.04.28
    1.该合约主要实现账户管理和索引功能。
    2.实时分账系统二期需求
    3.添加用人白名单
*/

contract AccountManager {
    string public name; // 合约的名字
    string public symbol; // 合约的标记或者符号，如 早中晚班
    address issuer; // 部署合约的账户地址

    /* 用人白名单 */
    mapping(address => address[]) employList;

    struct AccountInfo {
        address accountAddr;
        string name;
        string password;
        string accountDescribe;
        string telephone; // 电话号码
    }

    mapping(address => AccountInfo) public Accounts;
    mapping(string => address[]) AccAddressList; // 一个电话号码可以绑定多个账户
    address[] allAccountsAddr;

    function AccountManager(string tokenName, string tokenSymbol) {
        name = tokenName;
        symbol = tokenSymbol;
        issuer = msg.sender;
    }

    function addAccount(address accountAddr, string name, string password, string describe, string telephone) returns (bool success) {
        Accounts[accountAddr] = AccountInfo(accountAddr, name, password, describe, telephone);

        AccAddressList[telephone].push(accountAddr);
        allAccountsAddr.push(accountAddr);
    }

    function getAccount(address accountAddr) constant returns (string name, string password, string accountDescribe, string telephone){
        AccountInfo account = Accounts[accountAddr];
        return (account.name, account.password, account.accountDescribe, account.telephone);
    }

    function getOneAddress(string telephone) constant returns (address[]) {
        return AccAddressList[telephone];
    }

    /* 用人白名单 */
    function addEmployStaff(address staffAddr) returns (bool) {
        address[] storage sdCxt = employList[msg.sender];
        for (uint i = 0; i < sdCxt.length; i++) {
            address accountAddr = sdCxt[i];
            if (accountAddr == staffAddr) {
                return false;
            }
        }
        employList[msg.sender].push(staffAddr);
        return true;
    }

    function delEmployStaff(address staffAddr) returns (bool) {
        address[] storage sdCxt = employList[msg.sender];

        for (uint i = 0; i < sdCxt.length; i++) {
            address accountAddr = sdCxt[i];
            if (accountAddr == staffAddr) {
                delete sdCxt[i];

                employList[msg.sender] = sdCxt;
                break;
            }
        }
    }

    function getEmployStaffs(address owner) constant returns (address[]){
//        if (owner != msg.sender) throw;
        return employList[owner];
    }

    function getAllAccountsAddr() constant returns (address[]){
        return allAccountsAddr;
    }
}