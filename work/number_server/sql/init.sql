
-- 建表语句会将之前的表直接删除掉
USE finance_orders;
DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
	id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
	third_trade_no VARCHAR(64) NOT NULL COMMENT '第三方交易号',
	order_no VARCHAR(64) NOT NULL COMMENT '商户订单编号',
	sub_account_no VARCHAR(64) NOT NULL COMMENT '用户分账编号',
	company VARCHAR(64) NOT NULL COMMENT '所属公司',
	order_type TINYINT(1) NOT NULL COMMENT '订单类型，0:支付宝，1:微信',
	payment_type TINYINT(1) NOT NULL DEFAULT 0 COMMENT '付款类型，0，支付，1，转账',
	transfer_amount FLOAT NOT NULL DEFAULT 0 COMMENT '转账金额',
	transfer_info TEXT NOT NULL COMMENT '转账信息',
    auto_transfer TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否自动打款，0，否，1是',
    order_time  TIMESTAMP NOT NULL DEFAULT NOW() COMMENT '下单时间',
    order_state TINYINT(1) NOT NULL DEFAULT 0 COMMENT '订单状态',
	create_time DATETIME NOT NULL DEFAULT NOW() COMMENT '创建时间',
	update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	INDEX(third_trade_no,order_no)
)  DEFAULT charset=utf8  COMMENT '订单表';




ALTER TABLE orders ADD COLUMN mch_id VARCHAR(64) NOT NULL   COMMENT '商户ID';
ALTER TABLE orders ADD COLUMN branch_shop VARCHAR(64) NOT NULL  COMMENT '分店';


DROP TABLE IF EXISTS service_develop_keys;
CREATE TABLE service_develop_keys(
	key_id INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
	app_id CHAR(5) NOT NULL COMMENT '开发的appid',
	develop_key CHAR(32) NOT NULL COMMENT '开发的key',
	service VARCHAR(255) NOT NULL COMMENT '服务名称',
	create_time DATETIME NOT NULL COMMENT '创建时间',
	UNIQUE(app_id)
) DEFAULT charset=utf8 COMMENT '开发者key';


INSERT INTO service_develop_keys VALUES (1, "10001", 'golo10000', 'golo', NOW());

INSERT INTO service_develop_keys VALUES (2, "20001", 'golo10001', 'golo', NOW());
INSERT INTO service_develop_keys VALUES (2, "30001", 'golo10001', 'golo', NOW());
INSERT INTO service_develop_keys VALUES (4, "00000", 'golo10000', 'golo', NOW());

DROP TABLE IF EXISTS query_versions;
CREATE TABLE query_versions(
	id INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
	app_id CHAR(5) NOT NULL COMMENT 'app_id',
	version VARCHAR(50) NOT NULL COMMENT '版本号',
	begin_order_id INT(11) NOT NULL COMMENT '该版本查询开始的ID',
	end_order_id INT(11) NOT NULL COMMENT '该版本查询结束的订单ID',
	create_time DATETIME NOT NULL COMMENT '创建时间',
	INDEX(app_id,version)
) DEFAULT charset=utf8 COMMENT '版本号';

ALTER TABLE query_versions ADD COLUMN company_id VARCHAR(64) NOT NULL  COMMENT '商户ID'

ALTER TABLE query_versions ADD COLUMN company_id VARCHAR(64) NOT NULL  COMMENT '商户ID';

DROP TABLE IF EXISTS shop_indexes;
CREATE TABLE shop_indexes(
	key_id  INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
	shop_id INT(20) NOT NULL COMMENT '商户ID 用来设计分表算法',
	status TINYINT(1) NOT NULL DEFAULT 0 COMMENT '账户状态，用来控制商家能否拉去数据',
	UNIQUE(shop_id)
)DEFAULT charset=utf8 COMMENT '商户表';