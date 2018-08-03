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
        uint subWay; // 分配方式，0表示按比例分配，1表示按定额分配。
        bool isowner;
        uint quotaWay; // 数据置空方式，0不置空，1按日置空，2按月置空。
        uint resetTime; // 分配数据重置数据，按日的话，每天的这个时间，按月的话，每月1号的这个时间；1-6之间，表1点到6点之间更新
    }

    mapping(address => uint256) public balances;
    mapping(string => joinerDesired[]) numberSubAccount; // map[gxc001]=[{address:10%}]
    string[] public subAccountKeys; // 所有的分配编号
    mapping(address => bytes32[]) bindIssueToAddr;
    mapping(string => mapping(address => uint)) subCodeQuotaData; // 每个分配表每个账户按定额分配的已分配数据
    // new add
    mapping(string => address) subCodeIssuers; // 每个分配编号的发布者

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
    }

    mapping(address => accountInfo) public Accounts;
    mapping(string => address[]) AccAddressList; // 一个电话号码可以绑定多个账户
    address[] allAccountsAddr;

    /*
        分账流水信息记录，存证
    */
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
        2. 传入是数组顺序按设置的权重排序
    */
    function issueSubCxt(string issueCode, address[] subaccounts, uint[] rtaios, uint[] subWays, uint[] quotaWays, uint[] resetTimes) {
        require(subaccounts.length == rtaios.length && subaccounts.length > 0);

        joinerDesired[] storage sc_arry = numberSubAccount[issueCode];

        if (numberSubAccount[issueCode].length == 0) {
            subAccountKeys.push(issueCode);
            bindIssueToAddr[msg.sender].push(stringToBytesVer2(issueCode));

            subCodeIssuers[issueCode] = msg.sender;
        } else {
            // 必须是发布者才能更新
            require(subCodeIssuers[issueCode] == msg.sender);
            sc_arry.length = 0;
        }

        for (uint i = 0; i < subaccounts.length; i++) {
            bool flag = false;
            if (msg.sender == subaccounts[i]) {
                flag = true;
            }
            sc_arry.push(joinerDesired(issueCode, subaccounts[i], rtaios[i], subWays[i], flag, quotaWays[i], resetTimes[i]));
        }
        numberSubAccount[issueCode] = sc_arry;
    }

    function getIssueSubCxt(string issueCode) constant returns (address[], uint[], uint[], uint[], uint[]){
        address[] memory subaccounts = new address[](numberSubAccount[issueCode].length);
        uint[] memory rtaios = new uint[](numberSubAccount[issueCode].length);
        uint[] memory subWays = new uint[](numberSubAccount[issueCode].length);
        uint[] memory quotaWays = new uint[](numberSubAccount[issueCode].length);
        uint[] memory resetTimes = new uint[](numberSubAccount[issueCode].length);

        for (uint i = 0; i < numberSubAccount[issueCode].length; i++) {
            subaccounts[i] = numberSubAccount[issueCode][i].joinAddr;
            rtaios[i] = numberSubAccount[issueCode][i].ratio;
            subWays[i] = numberSubAccount[issueCode][i].subWay;
            quotaWays[i] = numberSubAccount[issueCode][i].quotaWay;
            resetTimes[i] = numberSubAccount[issueCode][i].resetTime;
        }
        return (subaccounts, rtaios, subWays, quotaWays, resetTimes);
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
            Accounts[accountAddr] = accountInfo(accountAddr, name, telephone, bankCard, weChat, alipay);
        } else {
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

    function settleAccounts(string issueCode, uint totalConsume, string transferId) {
        require(msg.sender == payer);
        require((checkLedgerById(transferId)) == false);

        uint sum = 0;
        uint calculate = 0;
        uint c_totalConsume = totalConsume;

        joinerDesired[] storage sc_arry = numberSubAccount[issueCode];
        uint tmp_num = sc_arry.length + 1;

        for (uint j = 0; j < sc_arry.length; j++) {
            if (totalConsume == 0) {
                break;
            }
            if (sc_arry[j].ratio == 0) {
                continue;
            }
            calculate = 0;

            if (sc_arry[j].subWay == 1) {
                // 按定额分配
                if (sc_arry[j].ratio > subCodeQuotaData[issueCode][sc_arry[j].joinAddr]) {
                    calculate = (sc_arry[j].ratio - subCodeQuotaData[issueCode][sc_arry[j].joinAddr]);

                    if (calculate > totalConsume) {
                        calculate = totalConsume;
                    }

                    totalConsume = totalConsume - calculate;
                    subCodeQuotaData[issueCode][sc_arry[j].joinAddr] = subCodeQuotaData[issueCode][sc_arry[j].joinAddr] + calculate;
                }
            } else {
                // 按比例分配
                if (totalConsume > 0) {
                    if (sc_arry[j].isowner == true) {
                        tmp_num = j;
                    } else {
                        uint tmp = totalConsume * sc_arry[j].ratio;
                        calculate = (tmp / 10000 - (((tmp / 10000) % 100)));
                        sum += calculate;
                    }
                }
            }

            //update
            if (calculate > 0) {
                updateOneLedger(msg.sender, sc_arry[j].joinAddr, calculate, transferId);

                // 更新账本，给财务平台查询
                ledgerDesList[transferId].push(calulateDecired(sc_arry[j].joinAddr, sc_arry[j].ratio, sc_arry[j].subWay, calculate, issueCode, transferId, c_totalConsume, false, ""));
            }
        }
        if (tmp_num >= 0 && tmp_num < sc_arry.length) {
            calculate = totalConsume - sum;
            updateOneLedger(msg.sender, sc_arry[tmp_num].joinAddr, calculate, transferId);

            // 更新账本，给财务平台查询
            ledgerDesList[transferId].push(calulateDecired(sc_arry[tmp_num].joinAddr, sc_arry[tmp_num].ratio, sc_arry[tmp_num].subWay, calculate, issueCode, transferId, c_totalConsume, false, ""));
        }
    }

    function updateOneLedger(address sender, address joiner, uint calculate, string transferId) returns (bool){
        bytes memory idg = bytes(transferId);
        if (idg.length == 0) {
            return false;
        }

        balances[joiner] = balances[joiner] + calculate;
        Transfer(sender, joiner, calculate);
        ledgerSubAddress[transferId].push(joiner);
        return true;
    }

    function checkLedgerById(string transferId) returns (bool) {
        if (ledgerDesList[transferId].length > 0) {
            return true;
        }
        return false;
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
    function updateCalulateLedger(address uaddr, string transferId, string transferDetails) {
        bytes memory idb = bytes(transferId);
        require(idb.length > 0);
        calulateDecired[] storage lg = ledgerDesList[transferId];
        for (uint j = 0; j < lg.length; j++) {
            if (lg[j].joinAddr == uaddr) {
                lg[j].transferDetails = transferDetails;
                lg[j].rflag = true;
                ledgerDesList[transferId] = lg;
                return;
            }
        }
    }

    function getAllSubKeysLen() constant returns (uint){
        return subAccountKeys.length;
    }
    /*
       5. 定额分配已分配数据重置。
    */
    function resetSubCodeQuotaData(string subCode, address setAddr) {
        subCodeQuotaData[subCode][setAddr] = 0;
    }

    function getSubCodeQuotaData(string subCode, address setAddr) constant returns (uint){
        return subCodeQuotaData[subCode][setAddr];
    }
    // 获取每个地址获取的分账总数
    function getBalance(address account) constant returns (uint256) {
        return balances[account];
    }
}