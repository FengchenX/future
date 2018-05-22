pragma solidity ^0.4.21;

/*
    1.该合约主要实线上部署好的合约地址管理和索引功能。
*/

contract AddressManager {
    string public name; // 合约的名字
    string public symbol; // 合约的标记或者符号，如 早中晚班
    address issuer; // 部署合约的账户地址

    event globallog(address sender, string logstr);

    struct AddressInfo {
        address smartAddr;
        string name;
        string symbol;
    }

    //    AddressInfo[] addresslist;
    mapping(string => address[]) listIndex;
    mapping(address => AddressInfo) public addresslist;

    function AddressManager(string tokenName, string tokenSymbol) {
        name = tokenName;
        symbol = tokenSymbol;
        issuer = msg.sender;
    }

    function strConcat(string _a, string _b, string _c, string _d, string _e) internal returns (string){
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        bytes memory _bc = bytes(_c);
        bytes memory _bd = bytes(_d);
        bytes memory _be = bytes(_e);
        string memory abcde = new string(_ba.length + _bb.length + _bc.length + _bd.length + _be.length);
        bytes memory babcde = bytes(abcde);
        uint k = 0;
        for (uint i = 0; i < _ba.length; i++) babcde[k++] = _ba[i];
        for (i = 0; i < _bb.length; i++) babcde[k++] = _bb[i];
        for (i = 0; i < _bc.length; i++) babcde[k++] = _bc[i];
        for (i = 0; i < _bd.length; i++) babcde[k++] = _bd[i];
        for (i = 0; i < _be.length; i++) babcde[k++] = _be[i];
        return string(babcde);
    }

    function strConcat(string _a, string _b) internal returns (string) {
        return strConcat(_a, _b, "", "", "");
    }

    function addAddress(address accountAddr, string name, string symbol) returns (bool success) {
        addresslist[accountAddr] = AddressInfo(accountAddr, name, symbol);

        // 添加索引
        //        addAddressIndes(accountAddr, name, symbol);
    }

    function addAddressIndex(address accountAddr, string name, string symbol) {
        listIndex[name].push(accountAddr);
        listIndex[symbol].push(accountAddr);
        listIndex[strConcat(name, symbol)].push(accountAddr);

        globallog(accountAddr, name);
        globallog(accountAddr, symbol);
        globallog(accountAddr, strConcat(name, symbol));
    }

    function getAddressIndex(string findkey) constant returns (address[]){
        return listIndex[findkey];
        //        address[] findSmartAddrs;
        //        address[] indexs = listIndex[findkey];

        //        for (uint i = 1; i < indexs.length; i++) {
        //            address smartAddr = addresslist[indexs[i]].smartAddr;
        //            findSmartAddrs.push(smartAddr);
        //        }
        //        return findSmartAddrs;
    }

    function getAddressArray(address accountAddr) constant returns (address, string, string){
        return (addresslist[accountAddr].smartAddr, addresslist[accountAddr].name, addresslist[accountAddr].symbol);
    }
}