USE finance_orders;
DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
    id INT(11) NOT NULL PRIMARY KEY,
    third_trade_no VARCHAR(64) NOT NULL,
    sub_account_no VARCHAR(64) NOT NULL FOREIGN KEY REFERENCES persons(ip)
)