

var sc_arry=[
    {
        "subCode":"a1",
        "role":"10001",
        "ratio":30,
        "subWay":0,
        "quotaWay":0,
        "resetTime":0
    },
    {
        "subCode":"a1",
        "role":"10002",
        "ratio":300*10000,
        "subWay":1,
        "quotaWay":0,
        "resetTime":0
    },
    {
        "subCode":"a1",
        "role":"10003",
        "ratio":70,
        "subWay":0,
        "quotaWay":0,
        "resetTime":0
    }
]
var subCodeIssuers=[]
subCodeIssuers["a1"]="0x10001"


var subCodeQuotaData=[]

subCodeQuotaData["a1"]=[]
subCodeQuotaData["a1"]["0x10001"]=0;
subCodeQuotaData["a1"]["0x10002"]=0;
subCodeQuotaData["a1"]["0x10003"]=0;

var schedulingJoiners=[]
schedulingJoiners["a1"]=[]

schedulingJoiners["a1"]["10001"]="0x10001"
schedulingJoiners["a1"]["10002"]="0x10002"
schedulingJoiners["a1"]["10003"]="0x10003"


var balances=[]
balances["0x10001"]=0;
balances["0x10002"]=0;
balances["0x10003"]=0;




function afterCalculateByRatio(totalConsume,ratio ,c_ratio){
    var tmp = totalConsume * ratio;
    var calculate = ((tmp / 10000 - ((tmp / 10000) % c_ratio)) / c_ratio) * 10000;
    return calculate;
}

function calculateByRatio(totalConsume,ratio){
    var calculate=afterCalculateByRatio(totalConsume,ratio,100)
    return calculate;
}


function updateOneLedger(sender,joiner,calculate,transferId){
    console.log("[ 账户变化： 钱包地址="+joiner+"  变化前="+balances[joiner]+"  变化后="+calculate+"  ]")
    balances[joiner] = balances[joiner] + calculate;
    print()
}

function print(){
  console.log("[ 0x10001:"+balances["0x10001"]/10000,"0x10002:"+balances["0x10002"]/10000,"0x10003:"+balances["0x10003"]/10000+" ]")
}

function settleAccounts(issueCode,totalConsume,transferId){
    xsum=totalConsume/10000
    console.log("分账开始:  transferId="+transferId+"  sum="+xsum)
    var sum = 0;
    var calculate = 0;
    var c_totalConsume = totalConsume;
    var c_ratio = 100; //剩余比率
    var r_flag = 0;

    var tmp_num = sc_arry.length + 1;

    for( var j=0;j<sc_arry.length;j++){
        if (totalConsume == 0) {
            break;
        }
        if (sc_arry[j].ratio == 0) {
            continue;
        }
        calculate = 0;
        console.log("分配方式====>"+(sc_arry[j].subWay==0?" 比率分配":" 定额分配")+"  ,"+j)
        if (sc_arry[j].subWay == 1) {// 按定额分配
            // 还未分满
            if (sc_arry[j].ratio > subCodeQuotaData[issueCode][schedulingJoiners[issueCode][sc_arry[j].role]]) {
                calculate = (sc_arry[j].ratio - subCodeQuotaData[issueCode][schedulingJoiners[issueCode][sc_arry[j].role]]);
                //本次分配的金额为未分满的部分，如果未分满的部分大于总额，那么所有钱全部分给他
                if (calculate > totalConsume) {
                    calculate = totalConsume;
                }
                subCodeQuotaData[issueCode][schedulingJoiners[issueCode][sc_arry[j].role]] = subCodeQuotaData[issueCode][schedulingJoiners[issueCode][sc_arry[j].role]] + calculate;
                if (r_flag == 1) { //标志已分配过定额
                    r_flag = 2;
                }
            }
        } else {
            // 按比例分配
            if (subCodeIssuers[issueCode] == schedulingJoiners[issueCode][sc_arry[j].role]) {
                tmp_num = j;
            } else {
                if (r_flag >= 2) {
                    if (r_flag == 2) {
                        c_totalConsume = totalConsume;
                    }
                    r_flag = 3;
                    calculate = afterCalculateByRatio(c_totalConsume, sc_arry[j].ratio, c_ratio);
                } else {
                    r_flag = 1;
                    calculate = calculateByRatio(c_totalConsume, sc_arry[j].ratio);
                    c_ratio = c_ratio - sc_arry[j].ratio;
                }
            }
        }
         //update
         if (calculate > 0) {
            totalConsume = totalConsume - calculate;//剩余
            updateOneLedger("正常"+schedulingJoiners[issueCode][sc_arry[j].role], schedulingJoiners[issueCode][sc_arry[j].role], calculate, transferId);

        }
    }
    if (tmp_num >= 0 && tmp_num < sc_arry.length) {
        calculate = totalConsume - sum;
        updateOneLedger("发布者:"+tmp_num, schedulingJoiners[issueCode][sc_arry[tmp_num].role], calculate, transferId);
    }
    console.log("分账结束:  transferId="+transferId+"  sum="+xsum)
}


settleAccounts("a1",10*10000,1)
settleAccounts("a1",200*10000,2)
settleAccounts("a1",300*10000,3)