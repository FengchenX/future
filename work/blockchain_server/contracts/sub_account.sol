pragma solidity ^0.4.21;

/*
    实时分账系统三期需求：
    1.一个合约搞定分账的所有事情；
*/

contract SubAccount {
    address public owner;
    address public payer;      // 财务系统的绑定账户，可以写死在合约里，固定不变
    string  public name;       // 自己要发行的货币的名称
    string  public symbol;     // 要发行的货币的符号
    uint8   public decimals;   // 货币的小数位
    string  public postscript; // 备注信息

    event Transfer(address from, address to, uint amount);
    //    event globallog(address sender, string logstr);

    /*
        分配比例
    */
    struct joinerDesired {
        string subCode;
        address joinAddr;
        uint ratio;
        uint subWay; // 分配方式，0表示按比例分配，1表示按定额分配。定额分配涉及优先级问题，待讨论。
        bool isowner;
    }

    mapping(address => uint256) public balances;
    mapping(string => joinerDesired[]) numberSubAccount; // map[gxc001]=[{address:10%}]
    string[] public subAccountKeys; // 所有的分配编号
    mapping(address => bytes32[]) bindIssueToAddr;

    /*
        账户信息绑定
    */
    struct accountInfo {
        address accountAddr;
        string name;
        string telephone; // 电话号码
        string bankCard;  // 银行卡，保持只有一个，可更换。
        string weChat;    // 微信
        string alipay;    // 支付宝
        //        string password;
        //        string accountDescribe;
    }

    mapping(address => accountInfo) public Accounts;
    mapping(string => address[]) AccAddressList; // 一个电话号码可以绑定多个账户
    address[] allAccountsAddr;

    /*
        分账流水信息记录，存证
    */
    struct orderInfo {
        string orderId;
        string content;
        uint256 creatAt;
        bool rflag;
    }

    mapping(string => orderInfo) ordersInfoList; // 订单存证数据
    string[] public ordersHashKeys; // 订单存证的索引hash

    /*
        分账账本信息
    */
    //    struct ledgerDesired {
    //            string orderId; //要分账的三方流水号
    //            uint totalConsume;
    //            bool rflag;
    //            string transferDetails; // 转账成功后的详情，分账后为空，三方支付成功后，调用接口更新。
    //            //        calulateDecired[] ledgerDesList;
    //        }

    struct calulateDecired {
        address joinAddr;
        uint ratio;
        uint subWay;       // 分配方式，0表示按比例分配，1表示按定额分配。定额分配涉及优先级问题，待讨论。
        uint calculate;    // 分账所得
        string subCode;      // 分配编号
        string orderId;      //要分账的三方流水号
        uint totalConsume;
        bool rflag;        // 是否由平台确认过信息
        string transferDetails;
    }

    mapping(string => calulateDecired[]) ledgerDesList;
    mapping(string => address[]) ledgerSubAddress;

    // 一个构造函数，会在创建合约的时候运行，后面不可调用。会永久的存储合约创建者的地址 --- owner。
    function SubAccount(string tokenName, string tokenSymbol, address initpayer, string init_postscript, uint8 decimalUnits) {
        name = tokenName;
        symbol = tokenSymbol;
        decimals = decimalUnits;
        owner = msg.sender;
        payer = initpayer;
        postscript = init_postscript;
    }

    function StringToBytesVer1(string memory source) returns (bytes result) {
        return bytes(source);
    }

    function stringToBytesVer2(string memory source) returns (bytes32 result) {
        assembly {
            result := mload(add(source, 32))
        }
    }

    /*
        1. 发布编号对应的分账人和比例
        ps: 分配人里没自己和比例和超过100，失败，调用合约之前判断。
    */
    function issueSubCxt(string issueCode, address[] subaccounts, uint[] rtaios) {
        if (subaccounts.length != rtaios.length && subaccounts.length > 0) throw;

        joinerDesired[] storage sc_arry = numberSubAccount[issueCode];
        if (sc_arry.length == 0) {
            subAccountKeys.push(issueCode);
            bindIssueToAddr[msg.sender].push(stringToBytesVer2(issueCode));
        }

        for (uint i = 0; i < subaccounts.length; i++) {
            bool flag = false;
            if (msg.sender == subaccounts[i]){
                flag = true;
            }
            sc_arry.push(joinerDesired(issueCode, subaccounts[i], rtaios[i], 0, flag));
        }
        numberSubAccount[issueCode] = sc_arry;
    }

    function getIssueSubCxtLen(string issueCode) constant returns (uint, address, uint){
        return (numberSubAccount[issueCode].length, numberSubAccount[issueCode][0].joinAddr, numberSubAccount[issueCode][0].ratio);
    }

    function getIssueSubCxt(string issueCode) constant returns (address[], uint[], uint){
        address[] memory subaccounts = new address[](numberSubAccount[issueCode].length);
        uint[] memory rtaios = new uint[](numberSubAccount[issueCode].length);

        for (uint i = 0; i < numberSubAccount[issueCode].length; i++) {
            subaccounts[i] = numberSubAccount[issueCode][i].joinAddr;
            rtaios[i] = numberSubAccount[issueCode][i].ratio;
        }
        return (subaccounts, rtaios, numberSubAccount[issueCode].length);
    }

    function getSubAccountKeysLen() constant returns (uint){
        return subAccountKeys.length;
    }

    /*
        查询自己发布了多少个分配比例编号
    */
    function getbindIssueToAddr(address faddr) constant returns (bytes32[]){
        return bindIssueToAddr[faddr];
    }

    /*
        2. 信息绑定
    */
    function addAccount(address accountAddr, string name, string bankCard, string weChat, string alipay, string telephone) returns (bool success) {
        accountInfo storage account = Accounts[accountAddr];
        bytes memory nameb = bytes(account.name);
        if (nameb.length > 0) {
            // update
            Accounts[accountAddr] = accountInfo(accountAddr, name, telephone, bankCard, weChat, alipay);
        } else {
            // new add
            Accounts[accountAddr] = accountInfo(accountAddr, name, telephone, bankCard, weChat, alipay);

            AccAddressList[telephone].push(accountAddr);
            allAccountsAddr.push(accountAddr);
        }
    }

    function getAccount(address accountAddr) constant returns (string, string, string, string, string){
        accountInfo account = Accounts[accountAddr];
        return (account.name, account.bankCard, account.weChat, account.alipay, account.telephone);
    }

    function getOneAddress(string telephone) constant returns (address[]) {
        return AccAddressList[telephone];
    }

    function getAllAccountsAddr() constant returns (address[]){
        return allAccountsAddr;
    }

    /*
        3. 实时分账
        issueCode：发布的分配编号；
        transferId：三方转入的流水号
        totalConsume：交易的总数；

        1.  settleAccounts()
        2.  生成账本信息 ledgerDesList;
    */
    function settleAccounts(string issueCode, uint256 totalConsume, string transferId) returns (bool success){
        if (msg.sender != payer) throw;

        uint sum = 0;
        uint tmp_num = 0;

        joinerDesired[] storage sc_arry = numberSubAccount[issueCode];
        for (uint j = 0; j < sc_arry.length; j++) {
            if (!sc_arry[j].isowner){
                uint tmp = totalConsume * sc_arry[j].ratio;
                uint calculate = (tmp-((tmp%100))) / 100;
                sum += calculate;

                balances[sc_arry[j].joinAddr] = balances[sc_arry[j].joinAddr] + calculate;
                Transfer(msg.sender, sc_arry[j].joinAddr, calculate);

                // 更新账本，给财务平台查询
                ledgerDesList[transferId].push(calulateDecired(sc_arry[j].joinAddr, sc_arry[j].ratio, sc_arry[j].subWay, calculate, issueCode, transferId, totalConsume, false, ""));

                ledgerSubAddress[transferId].push(sc_arry[j].joinAddr);
            }else{
                tmp_num = j;
            }
        }

        uint money = totalConsume - sum;
        balances[sc_arry[tmp_num].joinAddr] = balances[sc_arry[tmp_num].joinAddr]+money;
        Transfer(msg.sender, sc_arry[tmp_num].joinAddr, money);

        // 更新账本，给财务平台查询
        ledgerDesList[transferId].push(calulateDecired(sc_arry[tmp_num].joinAddr, sc_arry[tmp_num].ratio, sc_arry[tmp_num].subWay, money, issueCode, transferId, totalConsume, false, ""));

        ledgerSubAddress[transferId].push(sc_arry[tmp_num].joinAddr);
    }

    function getOneLedgerCxt(address uaddr, string transferId) constant returns (uint, string, bool, string){
        calulateDecired[] storage lg = ledgerDesList[transferId];
        for (uint j = 0; j < lg.length; j++) {
            if (lg[j].joinAddr == uaddr) {
                return (lg[j].calculate, lg[j].orderId, lg[j].rflag, lg[j].transferDetails);
            }
        }
    }

    function getLedgerSubAddrs(string transferId) constant returns (address[]){
        return ledgerSubAddress[transferId];
    }

    /*
        4. 支付成功后，三方支付流水信息更新
    */
    function updateCalulateLedger(address uaddr, string transferId, string transferDetails) returns (bool success){
        calulateDecired[] storage lg = ledgerDesList[transferId];
        if (lg.length > 0) {
            for (uint j = 0; j < lg.length; j++) {
                if (lg[j].joinAddr == uaddr) {
                    lg[j].transferDetails = transferDetails;
                    lg[j].rflag = true;
                    ledgerDesList[transferId] = lg;
                    break;
                }
            }
        }
    }

    function getAllSubKeysLen() constant returns (uint){
        return subAccountKeys.length;
    }

    /*
        4. 以下order相关函数，做数据存证，接口备用。
    */
    function saveContHash(string hash, string content) returns (bool success){
        ordersInfoList[hash] = orderInfo(hash, content, now, true);
    }

    function getContHash(string hash) constant returns (string, uint256, bool){
        return (ordersInfoList[hash].content, ordersInfoList[hash].creatAt, ordersInfoList[hash].rflag);
    }

    function getContentHashKeysLen() constant returns (uint) {
        return ordersHashKeys.length;
    }

    function getOneContentHashKeys(uint idx) constant returns (string) {
        return ordersHashKeys[idx];
    }

    // 获取每个地址获取的分账总数
    function getBalance(address account) constant returns (uint) {
        return balances[account];
    }
}