package manager

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/tealeg/xlsx"

	"sub_account_service/finance/config"
	"sub_account_service/finance/db"
	"sub_account_service/finance/lib"
	_ "sub_account_service/finance/memory" //引用memory包初始化函数
	"sub_account_service/finance/models"
	"sub_account_service/finance/protocol"
	"sub_account_service/finance/session"
	"sub_account_service/finance/third_part_pay"
	"sub_account_service/finance/utils"
)

//GlobalSessions 全局session
var GlobalSessions *session.Manager

func init() {
	GlobalSessions, _ = session.NewSessionManager("memory", "token", 3600)

	//go GlobalSessions.GC()
}

//Login 登录处理器
func Login(c *gin.Context) {
	glog.Infoln("Login******************start")
	login := &models.Login{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		glog.Errorln("Login**************读request错误")
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "request错误"))
		return
	}
	err = json.Unmarshal(body, login)
	if err != nil {
		glog.Errorln("Login***********json 解码错误")
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "解码错误"))
		return
	}
	//数据库取用户数据
	db := db.DbClient.Client
	if db == nil {
		glog.Errorln("db****************is nil")
		return
	}
	// 获取第一个匹配记录
	user := &models.User{}
	mydb := db.Where("user_name = ?", login.UserName).First(user)
	if mydb.Error != nil {
		glog.Errorln("db***********err", mydb.Error)

		return
	}
	if login.Password == user.Password {
		//登录成功
		glog.Infoln("Login***********success")
		//告诉客户端登录成功, 将token传给客户端， 并保存session对象
		sess := GlobalSessions.SessionStart(c)
		if sess == nil {
			return
		}
		var resp struct {
			Token     string `json:"token"`
			Authority string `json:"authority"`
		}
		resp.Token = sess.SessionID()
		resp.Authority = user.Authority
		c.JSON(http.StatusOK, lib.Result.Success("登录成功", resp))
		//将用户名保存到session对象中
		sess.Set("username", login.UserName)
	} else {
		//登录失败
		glog.Errorln("Login**********fail")
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "登录失败"))

	}
}

//QueryBill 查询账单
func QueryBill(c *gin.Context) {
	var qb struct {
		Option       uint8 //0 查进账流水， 1 查出帐流水， 2 未上链， 3 已上链, 4 转账中， 5 转账完成
		PageNum      int
		PageSize     int
		BeginPayDate string
		EndPayDate   string
		Platform     int
		SubAccountNo string `json:"sub_account_no"`
		TradeNo      string `json:"trade_no"`
	}
	err := parseReq(c, "QueryBill", &qb)
	if err != nil {
		return
	}

	//查询表中满足条件总数量, 并输出到out
	var f1 = func(count *int, limit, offset int, model interface{}, out interface{}, str string, args ...interface{}) {
		reg, _ := regexp.Compile(`(?i:^\s*and)`)
		str = reg.ReplaceAllString(strings.Trim(str, " "), "")
		if str != "" {
			temp := db.DbClient.Client.Model(model).Where(str, args...).Count(count)
			if temp.Error != nil {
				glog.Errorln("QueryBill*************查询错误", temp.Error)
				return
			}
			temp = db.DbClient.Client.Where(str, args...).Order("send_pay_date desc").Limit(limit).Offset(offset).Find(out)
			if temp.Error != nil {
				glog.Errorln("QueryBill*************查询错误", temp.Error)
				return
			}
		} else {
			db.DbClient.Client.Model(model).Count(count)
			db.DbClient.Client.Order("send_pay_date desc").Limit(limit).Offset(offset).Find(out)
		}
	}
	var f2 = func(resp seter, model interface{}, out interface{}, str string, args ...interface{}) {
		var count int
		f1(&count, qb.PageSize, (qb.PageNum-1)*qb.PageSize, model, out, str, args...)
		resp.Set(out, qb.PageNum, qb.PageSize, count)
		c.JSON(http.StatusOK, lib.Result.Success("查询流水成功", resp))
	}

	var querystr string
	var param []interface{}
	if qb.BeginPayDate != "" {
		qb.BeginPayDate += " 00:00:00"
		querystr += " AND send_pay_date >= ?"
		param = append(param, formToUnix(qb.BeginPayDate))
	}
	if qb.EndPayDate != "" {
		qb.EndPayDate += " 23:59:59"
		querystr += " AND send_pay_date <= ?"
		param = append(param, formToUnix(qb.EndPayDate))
	}
	if qb.Platform != -1 {
		querystr += " AND order_type = ?"
		param = append(param, qb.Platform)
	}
	if qb.SubAccountNo != "" {
		querystr += " AND sub_account_no = ?"
		qb.SubAccountNo = strings.Trim(qb.SubAccountNo, " ")
		param = append(param, qb.SubAccountNo)
	}
	if qb.TradeNo != "" {
		querystr += " AND trade_no = ?"
		qb.TradeNo = strings.Trim(qb.TradeNo, " ")
		param = append(param, qb.TradeNo)
	}

	//按状态查询
	var f3 = func(status uint8) {
		if status == 1 {
			var income []models.IncomeStatement
			resp := &respIncomeBill{}
			querystr += " AND blockchain_status = ?"
			param = append(param, 1)
			f2(resp, &models.IncomeStatement{}, &income, querystr, param...)
		} else if status > 1 {
			//支出
			var expenses []models.ExpensesBill
			resp := &respExpBill{}
			querystr += " AND blockchain_status = ?"
			param = append(param, status)
			f2(resp, &models.ExpensesBill{}, &expenses, querystr, param...)
		}
	}

	switch qb.Option {
	case 0:
		//查询进账流水
		var income []models.IncomeStatement
		resp := &respIncomeBill{}
		f2(resp, &models.IncomeStatement{}, &income, querystr, param...)
	case 1:
		//查询出帐流水
		//支出
		var expenses []models.ExpensesBill
		resp := &respExpBill{}
		f2(resp, &models.ExpensesBill{}, &expenses, querystr, param...)
	case 2:
		//查询未上链流水
		f3(1)
	case 3:
		//查询已上链流水
		f3(2)
	case 4:
		//查询转账中流水
		f3(3)
	case 5:
		//查询转账完成流水
		f3(4)
	default:

	}
}

type respIncomeBill struct {
	CurrentPage int
	PageSize    int
	Count       int
	In          []models.IncomeStatement
}
type income struct {
	In models.IncomeStatement
	//PayDate string
}
type respExpBill struct {
	CurrentPage int
	PageSize    int
	Count       int
	Exp         []models.ExpensesBill
}
type expense struct {
	ExpBill models.ExpensesBill
	//PayDate string
}
type seter interface {
	Set(obj interface{}, curPage, pageSize, count int)
}

func (ib *respIncomeBill) Set(obj interface{}, curPage, pageSize, count int) {
	x, ok := obj.(*[]models.IncomeStatement)
	if !ok {
		return
	}
	// temp := *x
	// for _, in := range temp {
	// 	ib.In = append(ib.In, *in)
	// }
	ib.In = *x
	ib.Count = count
	ib.CurrentPage = curPage
	ib.PageSize = pageSize
}
func (eb *respExpBill) Set(obj interface{}, curPage, pageSize, count int) {
	x, ok := obj.(*[]models.ExpensesBill)
	if !ok {
		return
	}
	// temp := *x
	// for _, ex := range temp {
	// 	eb.Exp = append(eb.Exp, *ex)
	// }
	eb.Exp = *x
	eb.Count = count
	eb.CurrentPage = curPage
	eb.PageSize = pageSize
}

//DownIncomeExcel 下载income
func DownIncomeExcel(c *gin.Context) {
	fmt.Println("DownIncomeExcel******************start")
	sub_account_no := c.DefaultQuery("sub_account_no", "no sub_account_no")
	trade_no := c.DefaultQuery("trade_no", "no trade_no")
	fmt.Println("DownIncomeExcel******************", sub_account_no, trade_no)
	var querystr string
	var params []interface{}
	if sub_account_no != "" {
		querystr += " AND sub_account_no = ?"
		params = append(params, sub_account_no)
	}
	if trade_no != "" {
		querystr += " AND trade_no = ?"
		params = append(params, trade_no)
	}
	reg, _ := regexp.Compile(`(?i:^\s*and)`)
	querystr = reg.ReplaceAllString(strings.Trim(querystr, " "), "")
	db := db.DbClient.Client

	var incomes []models.IncomeStatement
	mydb := db.Where(querystr, params...).Find(&incomes)
	if mydb.Error != nil {
		fmt.Println("DownIncomeExcel******************db", mydb.Error)
	}
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row := sheet.AddRow()
	row.SetHeightCM(1)
	cell := row.AddCell()
	cell.Value = "ID"
	cell = row.AddCell()
	cell.Value = "订单类型"
	cell = row.AddCell()
	cell.Value = "商家账户"
	cell = row.AddCell()
	cell.Value = "买家支付宝账号"
	cell = row.AddCell()
	cell.Value = "时间"
	cell = row.AddCell()
	cell.Value = "订单金额"
	cell = row.AddCell()
	cell.Value = "手续费"
	cell = row.AddCell()
	cell.Value = "交易号"
	cell = row.AddCell()
	cell.Value = "分账信息编号"
	cell = row.AddCell()
	cell.Value = "付款方"
	cell = row.AddCell()
	cell.Value = "交易流水是否上链"

	for _, v := range incomes {
		row := sheet.AddRow()
		row.SetHeightCM(1)
		cell := row.AddCell()
		cell.Value = fmt.Sprintf("%d", v.ID)

		var T string
		if v.OrderType == 0 {
			T = "支付宝"
		} else {
			T = "微信"
		}
		cell = row.AddCell()
		cell.Value = T

		cell = row.AddCell()
		cell.Value = v.AlipaySeller
		cell = row.AddCell()
		cell.Value = v.BuyerLogonID
		cell = row.AddCell()
		cell.Value = unixToForm(v.SendPayDate)
		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%f", v.TotalAmount)
		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%f", v.Fees)
		cell = row.AddCell()
		cell.Value = v.TradeNo
		cell = row.AddCell()
		cell.Value = v.SubAccountNo
		cell = row.AddCell()
		cell.Value = v.PayerName
		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", v.BlockchainStatus)
	}
	err = file.Save("income_write.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	c.File("income_write.xlsx")
}

//DownExpExcel 下载DownExpExcel
func DownExpExcel(c *gin.Context) {
	fmt.Println("DownExpExcel******************start")
	sub_account_no := c.DefaultQuery("sub_account_no", "no sub_account_no")
	trade_no := c.DefaultQuery("trade_no", "no trade_no")
	fmt.Println("DownExpExcel******************", sub_account_no, trade_no)
	var querystr string
	var params []interface{}
	if sub_account_no != "" {
		querystr += " AND sub_account_no = ?"
		params = append(params, sub_account_no)
	}
	if trade_no != "" {
		querystr += " AND trade_no = ?"
		params = append(params, trade_no)
	}
	reg, _ := regexp.Compile(`(?i:^\s*and)`)
	querystr = reg.ReplaceAllString(strings.Trim(querystr, " "), "")
	db := db.DbClient.Client

	var exps []models.ExpensesBill
	mydb := db.Where(querystr, params...).Find(&exps)
	if mydb.Error != nil {
		fmt.Println("DownExpExcel******************db", mydb.Error)
	}
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row := sheet.AddRow()
	row.SetHeightCM(1)
	cell := row.AddCell()
	cell.Value = "ID"
	cell = row.AddCell()
	cell.Value = "订单类型"
	cell = row.AddCell()
	cell.Value = "商家账户"
	cell = row.AddCell()
	cell.Value = "买家支付宝账号"
	cell = row.AddCell()
	cell.Value = "时间"
	cell = row.AddCell()
	cell.Value = "订单金额"
	cell = row.AddCell()
	cell.Value = "手续费"
	cell = row.AddCell()
	cell.Value = "交易号"
	cell = row.AddCell()
	cell.Value = "分账信息编号"
	cell = row.AddCell()
	cell.Value = "付款方"
	cell = row.AddCell()
	cell.Value = "交易状态"

	for _, v := range exps {
		row := sheet.AddRow()
		row.SetHeightCM(1)
		cell := row.AddCell()
		cell.Value = fmt.Sprintf("%d", v.ID)

		var T string
		if v.OrderType == 0 {
			T = "支付宝"
		} else {
			T = "微信"
		}
		cell = row.AddCell()
		cell.Value = T

		cell = row.AddCell()
		cell.Value = v.AlipaySeller
		cell = row.AddCell()
		cell.Value = v.BuyerLogonID
		cell = row.AddCell()
		cell.Value = unixToForm(v.SendPayDate)
		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%f", v.TotalAmount)
		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%f", v.Fees)
		cell = row.AddCell()
		cell.Value = v.TradeNo
		cell = row.AddCell()
		cell.Value = v.SubAccountNo
		cell = row.AddCell()
		cell.Value = v.PayerName
		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", v.TradeStatus)
	}
	err = file.Save("exp_write.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	c.File("exp_write.xlsx")
}

//GetBook 获取账本
func GetBook(c *gin.Context) {
	var gb struct {
		ID int `json:"id"`
	}
	if err := parseReq(c, "GetBook", &gb); err != nil {
		return
	}
	var ubs []models.UserBill
	mydb := db.DbClient.Client.Table("user_bills").Joins("join expenses_bills on user_bills.bill_id = expenses_bills.id").
		Where(" expenses_bills.id = ?", gb.ID).Find(&ubs)
	if mydb.Error != nil {
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "查询失败!"))
		fmt.Println(mydb.Error)
		return
	}
	c.JSON(http.StatusOK, lib.Result.Success("查询账本成功", ubs))
}

//PageTransMoney 页面转账接口
func PageTransMoney(c *gin.Context) {
	var ptm struct {
		ID []uint
		//Status uint
	}
	if err := parseReq(c, "PageTransMoney", &ptm); err != nil {
		return
	}
	glog.Infoln("PageTransMoney*********************调用转账接口前")
	sa := &third_part_pay.SubAccountInfo{}

	temp, _ := sa.InitiativeSubAccount(ptm.ID)
	glog.Infoln("PageTransMoney*********************调用转账接口后")
	result := true
	for _, v := range temp {
		if v == 0 {
			result = false
		}
	}
	if result {
		c.JSON(http.StatusOK, lib.Result.Success("转账成功", &struct{}{}))
	} else {
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "转账失败"))
	}
}

//AccountList 账本列表
func AccountList(c *gin.Context) {
	var al struct {
		CurrentPage int    `json:"currentPage"`
		PageSize    int    `json:"pageSize"`
		StartDate   string `json:"startDate"`
		EndDate     string `json:"endDate"`
		Name        string `json:"name"`
		NeedTotal   bool   `json:"needTotal"`
	}
	if err := parseReq(c, "AccountList", &al); err != nil {
		return
	}

	var resp respAL
	var querystr string
	var params []interface{}

	if al.StartDate != "" {
		al.StartDate += " 00:00:00"
		querystr += " AND pay_date > ?"
		params = append(params, al.StartDate)
	}
	if al.EndDate != "" {
		al.EndDate += " 23:59:59"
		querystr += " AND pay_date < ?"
		params = append(params, al.EndDate)
	}
	if al.Name != "" {
		querystr += " AND name = ?"
		params = append(params, strings.Trim(al.Name, " "))
	}

	if db.DbClient.Client == nil {
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "db.DbClient.Client==nil"))
		glog.Errorln("AccountList*******************************db == nil")
		return
	}
	db := db.DbClient.Client

	resp.Pagination.PageSize = al.PageSize
	resp.Pagination.CurrentPage = al.CurrentPage
	re := regexp.MustCompile(`(?i:^and)`)
	querystr = re.ReplaceAllString(strings.Trim(querystr, " "), "")
	var Models []respUB //todo 换成对应模型

	// mydb := db.Table("user_bills").Joins("join expenses_bills on user_bills.bill_id = expenses_bills.id").
	// 	Where(querystr, params...).
	// 	Order("send_pay_date desc").Limit(al.PageSize).Offset((al.CurrentPage - 1) * al.PageSize).Find(&Models)

	mydb := db.Select("expenses_bills.send_pay_date, user_bills.id, user_bills.name, user_bills.bank_card,"+
		"user_bills.wechat, user_bills.alipay, user_bills.telephone, user_bills.address, user_bills.order_id,"+
		"user_bills.trade_no, user_bills.money, user_bills.rflag, user_bills.transfer_details, user_bills.radio,"+
		"user_bills.sub_way, user_bills.trade_status, user_bills.on_chain_time, user_bills.on_chain_num").
		Table("user_bills").
		Joins("join expenses_bills on user_bills.bill_id = expenses_bills.id").
		Where(querystr, params...).
		Order("send_pay_date desc").
		Limit(al.PageSize).Offset((al.CurrentPage - 1) * al.PageSize).
		Find(&Models)
	if mydb.Error != nil {
		glog.Errorln("AccountList**********************", mydb.Error)
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "查询失败"))
		return
	}

	if al.NeedTotal {
		var totalPrice float64
		row := db.Select("sum(money)").
			Table("user_bills").
			Joins("join expenses_bills on user_bills.bill_id = expenses_bills.id").
			Where(querystr, params...).
			Row()
		if row == nil {
			fmt.Println("AccountList**************************row is nil")
			c.JSON(http.StatusOK, lib.Result.Fail(-1, "查询失败"))
			return
		}
		row.Scan(&totalPrice)
		resp.Amount = totalPrice
		fmt.Println("Amount****************price", totalPrice)
	}

	var total int
	mydb = db.Table("user_bills").Joins("join expenses_bills on user_bills.bill_id = expenses_bills.id").
		Where(querystr, params...).Count(&total)

	if mydb.Error != nil {
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "查询失败"))
		return
	}

	resp.Pagination.Total = total

	resp.List = Models
	glog.Infoln("AccountList**********************", resp)

	c.JSON(http.StatusOK, resp)

}

type respAL struct {
	List       interface{} `json:"list"`
	Pagination pagination  `json:"pagination"`
	Amount     float64     `json:"amount"`
}

type respUB struct {
	models.UserBill
	SendPayDate int64 `json:"sendPayDate"`
}

type pagination struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
	Total       int `json:"total"`
}

//DownBookExcel 下载DownBookExcel
func DownBookExcel(c *gin.Context) {
	fmt.Println("DownBookExcel******************start")
	name := c.DefaultQuery("name", "no name")
	startDate := c.DefaultQuery("startDate", "no startDate")
	endDate := c.DefaultQuery("endDate", "no endDate")
	fmt.Println("DownBookExcel******************", name, startDate, endDate)
	var querystr string
	var params []interface{}
	if name != "" {
		querystr += " AND name = ?"
		params = append(params, name)
	}
	if startDate != "" {
		querystr += " AND send_pay_date > ?"
		params = append(params, formToUnix(startDate+" 00:00:00"))
	}
	if endDate != "" {
		querystr += " AND send_pay_date < ?"
		params = append(params, formToUnix(endDate+" 23:59:59"))
	}
	reg, _ := regexp.Compile(`(?i:^\s*and)`)
	querystr = reg.ReplaceAllString(strings.Trim(querystr, " "), "")
	db := db.DbClient.Client

	var resp []respUB
	mydb := db.Select("expenses_bills.send_pay_date, user_bills.id, user_bills.name, user_bills.bank_card,"+
		"user_bills.wechat, user_bills.alipay, user_bills.telephone, user_bills.address, user_bills.order_id,"+
		"user_bills.trade_no, user_bills.money, user_bills.rflag, user_bills.transfer_details, user_bills.radio,"+
		"user_bills.sub_way, user_bills.trade_status, user_bills.on_chain_time, user_bills.on_chain_num").
		Table("user_bills").
		Joins("join expenses_bills on user_bills.bill_id = expenses_bills.id").
		Where(querystr, params...).
		Order("send_pay_date desc").
		Find(&resp)

	if mydb.Error != nil {
		fmt.Println("DownBookExcel******************db", mydb.Error)
	}
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row := sheet.AddRow()
	row.SetHeightCM(1)
	cell := row.AddCell()
	cell.Value = "用户名"
	cell = row.AddCell()
	cell.Value = "收益"
	cell = row.AddCell()
	cell.Value = "分账比例(%)"
	cell = row.AddCell()
	cell.Value = "时间"
	cell = row.AddCell()
	cell.Value = "用户支付宝"
	cell = row.AddCell()
	cell.Value = "用户电话"
	cell = row.AddCell()
	cell.Value = "平台交易编号"
	cell = row.AddCell()
	cell.Value = "流水号"
	cell = row.AddCell()
	cell.Value = "是否支付"

	for _, v := range resp {
		row := sheet.AddRow()
		row.SetHeightCM(1)
		cell := row.AddCell()
		cell.Value = v.Name
		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%f", v.Money)
		cell = row.AddCell()
		var radio string
		if v.SubWay == 1 {
			//定额
			radio = fmt.Sprintf("%f", v.Radio) + "(定额)"
		} else {
			radio = fmt.Sprintf("%f", v.Radio)
		}
		cell.Value = radio
		cell = row.AddCell()
		cell.Value = unixToForm(v.SendPayDate)
		cell = row.AddCell()
		cell.Value = v.Alipay
		cell = row.AddCell()
		cell.Value = v.Telephone
		cell = row.AddCell()
		cell.Value = v.OrderId
		cell = row.AddCell()
		cell.Value = v.TradeNo
		var pay string
		if v.TradeStatus > 6 {
			pay = "已支付"
		} else {
			pay = "未支付"
		}
		cell = row.AddCell()
		cell.Value = pay
	}
	err = file.Save("book_write.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	c.File("book_write.xlsx")
}

//DisProfile 分配详情
func DisProfile(c *gin.Context) {
	var rq struct {
		SubAccountNo string `json:"sub_account_no"`
	}
	if err := parseReq(c, "DisProfile", &rq); err != nil {
		return
	}

	in := protocol.ReqGetSchedue{
		ScheduleName: strings.Trim(rq.SubAccountNo, " "),
	}

	out, err := json.Marshal(in)
	if err != nil {
		glog.Errorln(lib.Log("json marshal error", "", "CheckAccountFromBlockchain"), "", nil)
		return
	}

	reqBody := bytes.NewReader(out)
	reqURL := fmt.Sprintf("http://%s:%s/getschedule", config.Optional.ApiAddress, config.Optional.ApiPort)
	rspBody, err := utils.SendHttpRequest("POST", reqURL, reqBody, nil, nil)
	if err != nil {
		glog.Errorln(lib.Log("get schedule err", "", "getschedule"), "", err, in.ScheduleName)
		return
	}

	//respSch := protocol.RespGetSchedue{}
	respSch := respGetSchedue{}
	if err := json.Unmarshal(rspBody, &respSch); err != nil {
		glog.Errorln(lib.Log("json unmarshal error", "", "getschedule"), "", nil)
		return
	}
	glog.Infoln("DisProfile******************resp:", respSch)
	accs := respSch.Accounts
	rs := respSch.Schedules

	var resp respDis

	if len(accs) != len(rs) || len(accs) <= 0 {
		fmt.Println("DisProfile******************accs != rs || len <= 0", accs, rs)
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "DisProfile******************accs != rs || len <= 0"))
		return
	}
	for i, v := range accs {
		var ds disSch
		ds.userAccount = v
		ds.rs = rs[i]
		resp.List = append(resp.List, ds)
	}
	glog.Infoln("DisProfile resp", resp)
	c.JSON(http.StatusOK, resp)
}

type userAccount struct {
	Address   string
	Name      string
	BankCard  string
	WeChat    string
	Alipay    string
	Telephone string
}
type rs struct {
	Accounts  string
	Level     int64
	Radios    float64
	SubWay    int64
	ResetWay  int64
	ResetTime int64
	GetMoney  float64
	Job       string
}
type respGetSchedue struct {
	StatusCode uint32
	Accounts   []userAccount
	Schedules  []rs
	Msg        string
}

type disSch struct {
	//protocol.UserAccount
	//protocol.Rs
	userAccount
	rs
}

type respDis struct {
	List []disSch `json:"list"`
}

//Refund 退款
func Refund(c *gin.Context) {
	var refund struct {
		CurrentPage  int    `json:"currentPage"`
		PageSize     int    `json:"pageSize"`
		StartDate    string `json:"startDate"`
		EndDate      string `json:"endDate"`
		SubAccountNo string `json:"sub_account_no"`
		TradeNo      string `json:"trade_no"`
	}
	if err := parseReq(c, "Refund", &refund); err != nil {
		return
	}
	var resp respRefund
	var query string
	var params []interface{}

	if refund.StartDate != "" {
		refund.StartDate += " 00:00:00"
		query += " AND send_pay_date > ?"
		params = append(params, formToUnix(refund.StartDate))
	}
	if refund.EndDate != "" {
		refund.EndDate += " 23:59:59"
		query += " AND send_pay_date < ?"
		params = append(params, formToUnix(refund.EndDate))
	}
	if refund.TradeNo != "" {
		query += " AND trade_no = ?"
		params = append(params, refund.TradeNo)
	}
	if refund.SubAccountNo != "" {
		query += " AND sub_account_no = ?"
		params = append(params, refund.SubAccountNo)
	}

	if db.DbClient.Client == nil {
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "db.DbClient.Client == nil"))
		glog.Errorln("AccountList****************************db == nil")
		return
	}
	DB := db.DbClient.Client

	resp.Pagination.PageSize = refund.PageSize
	resp.Pagination.CurrentPage = refund.CurrentPage
	re := regexp.MustCompile(`(?i:^and)`)
	query = re.ReplaceAllString(strings.Trim(query, " "), "")
	var Models []refundTrade
	mydb := DB.Select("*").
		Table("refund_trades").
		Where(query, params...).
		Limit(refund.PageSize).Offset((refund.CurrentPage - 1) * refund.PageSize).
		Find(&Models)

	if mydb.Error != nil {
		glog.Errorln("Refund************************db err:", mydb.Error)
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "查询失败"))
		return
	}

	var total int
	mydb = DB.Table("refund_trades").
		Where(query, params...).
		Count(&total)
	if mydb.Error != nil {
		glog.Errorln("Refund*************************db err:", mydb.Error)
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "查询失败"))
		return
	}
	resp.Pagination.Total = total

	resp.List = Models
	glog.Infoln("Refund**************************resp: ", resp)
	c.JSON(http.StatusOK, resp)
}

type respRefund struct {
	List       interface{} `json:"list"`
	Pagination pagination  `json:"pagination"`
}
type refundTrade struct {
	models.RefundTrade
}

//parseReq 解析请求消息, funcName: 调用这个解析函数的名字， obj: 接受json数据的结构体指针对象
func parseReq(c *gin.Context, funcName string, obj interface{}) error {
	glog.Infoln(funcName + "***********start")
	sess := GlobalSessions.Session(c)
	if sess == nil {
		glog.Errorln(funcName + "*******未找到session")
		c.JSON(332, lib.Result.Fail(-1, "请先登录"))
		err := errors.New("未登录")
		return err
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		glog.Errorln(funcName + "**************读request错误")
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "解析错误"))
		err := errors.New("解析错误")
		return err
	}
	err = json.Unmarshal(body, obj)
	if err != nil {
		glog.Errorln(funcName + "***********json 解码错误")
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "解码错误"))
		err := errors.New("解码错误")
		return err
	}
	val := reflect.ValueOf(obj)
	glog.Infoln(funcName+"*********************req: ", val.Elem().Interface())
	return nil
}

//parseReqNotIntercept 解析请求消息不验证用户名
func parseReqNotIntercept(c *gin.Context, funcName string, obj interface{}) error {
	glog.Infoln(funcName + "***********start")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		glog.Errorln(funcName + "**************读request错误")
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "解析错误"))
		err := errors.New("解析错误")
		return err
	}
	err = json.Unmarshal(body, obj)
	if err != nil {
		glog.Errorln(funcName + "***********json 解码错误")
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "解码错误"))
		err := errors.New("解码错误")
		return err
	}
	return nil
}

func formToUnix(t string) int64 {
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	//获取本地location
	loc, _ := time.LoadLocation("Local")                   //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, t, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                   //转化为时间戳 类型是int64
	return sr
}

func unixToForm(t int64) string {
	//获取本地location
	timeLayout := "2006-01-02 15:04:05" //转化所需模板

	//时间戳转日期
	dataTimeStr := time.Unix(t, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	return dataTimeStr
}
