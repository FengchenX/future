package entity

import "time"

// AAA050 商品表
type AAA050 struct {

}
//SELECT [o01330] --流水号
//,[o01329] --前台销售单号
//,[o00935] --仓库编号
//,[o01292] --商品编号
//,[b00193] --原价
//,[o00557] --实际价格
//,[b00741] --数量
//,[b00517] --小计金额
//,[b00495] --成本价
//,[b01058] --销售方式
//,[b01264] --收银员
//,[o01068] --营业员
//,[o00493] --专柜号
//,[o00781] --时间
//,[b00264] --日结标识
//,[O80074] --备注
//,[o00680] --条码
//,[b00152] --特价类型
//,[O80142] --机号
//,[b00569] --班次
//,[b00583] --
//,[b01382] --金卡号
//,[b00558] --
//,[b00501] --
//,[b01339] --
//,[b00627]
//,[up_index]
//FROM [jxc2010_V9].[dbo].[BBB070]
// BBB070 商品明细
type BBB070 struct {
	ProductCode string `xorm:"'o01292'"` //[o01292] --商品编号
	OriginPrice float64 `xorm:"'b00193'"`//,[b00193] --原价
	Price float64 `xorm:"'o00557'"`//,[o00557] --实际价格
	Count float64 `xorm:"'b00741'"`//,[b00741] --数量
	SubTotal float64  `xorm:"'b00517'"`//,[b00517] --小计金额
	CostPrice float64 	`xorm:"'b00495'"`//,[b00495] --成本价
	SellType string	`xorm:"'b01058'"`//,[b01058] --销售方式
	SellMan string `xorm:"o01068"`  //售货员
	//--商品表
	ProductName string `xorm:"a00827"` //产品名称
	Unit string `xorm:"b01214"` //单位
}

//SELECT [o01330] -- 流水号
//,[o01329] -- 收银单号
//,[b00390] -- 单据金额
//,[o00935] -- 仓库编号
//,[b01253] -- 付款方式
//,[b01058] -- 销售方式
//,[b01357] -- 信用卡号
//,[b01382] -- 金卡号码
//,[b01348] -- 币种号码
//,[b00932] -- 汇率
//,[b00659] -- 原币付款金额
//,[o00781] -- 时间
//,[b01264] -- b01264
//,[O80112] -- 备注
//,[o00575] --
//,[O80142] -- 机号
//,[b00264] -- 日结标识
//,[b00569] --班次
//,[b00583] --
//,[b00609] --客户付款金额
//,[b01339] --
//,[b00627]
//,[b01205]
//,[up_index]
//,[tradeno]
//FROM [jxc2010_V9].[dbo].[BBB077]

// BBB077 付款明细
type BBB077 struct {
	ThirdTradeNo string `xorm:"'o01330'"` //三方交易号
	OrderNo 	 string `xorm:"'o01329'"` //自己平台的交易单号
	OrderTime    time.Time `xorm:"'o00781'"` // 创建时间
	PayWay       string `xorm:"'b01348'"` // 支付方式 币种号码
	Price        float64 `xorm:"'b00390'"`// 订单价格单据金额
	SellType	 string `xorm:"'b01058'"`//订单状态
	ScheduleNo	string `xorm:"b00569"` //班次
	Salesman	    string `xorm:"b01264"` //售货员
	Remarks string `xorm:"'O80112'"` //备注
}