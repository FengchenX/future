package api

const (
	ORDER_FAIL    = uint32(00001) // 开始时间或结束时间为0
	ORDER_NOTFIND = uint32(00002) // 未查到结果

	ACCOUNT_ACCOUNT_NIL = uint32(10001) //账户为空
	ACCOUNT_NAME_NIL    = uint32(10002) //姓名为空
	ACCOUNT_ALI_NIL     = uint32(10003) //支付宝为空
	ACCOUNT_APPLY_NIL   = uint32(10004) //添加账户失败
	ACCOUNT_APPLY_ERR   = uint32(10005) //添加账户失败
	ACCOUNT_GET_ERR     = uint32(10006) //获取账户失败

	SCHEDULE_CANT_FIND_ID         = uint32(20001) //未找到相应的分账编号
	SCHEDULE_NO_PEOPLE            = uint32(20002) //排班人数不足
	SCHEDULE_NOT_CONTAIN_PUBLSHER = uint32(20003) //排班必须包括发布者
	SCHEDULE_SUM_MUSTBE_100       = uint32(20004) //比例总数必须为100
	SCHEDULE_SHOULD_HAS_RATIO     = uint32(20005) //不可以全部都是定额
	SCHEDULE_ACCOUNT_UNIQUE       = uint32(20006) //参与分账的账号必须唯一
	SCHEDULE_FAIL                 = uint32(20007) //比例表发布失败
)

const (
	//GetAllSchedule
	AS_DB_ERR = 40001 //数据库错误
	AS_DOPOST_ERR = 40002 //执行doPost函数时发生错误
	AS_SUCCESS = 0

	//GetMoney
	M_TIME_ERR = 50001 //起始时间填写错误
	M_NO_HAVE_MORE = 50002 //没有更多
	M_SUCCESS = 0

	//SetSchedule
	S_NO_FIND = 60001 //从数据库中没有找到相应编号
	S_DOPOST_ERR = 60002 //执行doPost函数错误
	S_SUM_NOT_100 = 60003 //比例之和不是100
	S_ALL_QUO = 60004 //全部是定额
	S_NO_JOB = 60005 //没有设置工作
	S_TWO_JOB_SAME = 60006 //分配表有两相同工作
	S_API_RETURN_ERR = 70007 //API返回错误
	S_SUCCESS = 0

	//SetPaiBan
	PB_NO_YOURSELF = 70001 //没有包含自己
	PB_HAD_NIL = 70002 //有空排班
	PB_DB_ERR = 70003 //db错误	
	PB_SCH_PB_LEN_NOTSAME = 70004 //分配表和排班表人数不一致
	PB_OWER_BEFORE_QUOTA = 70005 //自己在定额前面
	PB_DOPOST_ERR = 70006 //执行doPost错误
	PB_SUCCESS = 0
	PB_API_RETURN_ERR = 70007 //API返回错误
)
