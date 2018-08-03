

编号服务器

接受三方支付订单，对已经完成的交易，送入nsq中入库处理（nsq可以做成分布式，不同商家，订阅不同的topic，也可以并发去轮训）

produce_server 对三方过来的订单进行验证(不管是哪一方的数据过来，都做一层验证),提供Http接口供三方提交订单数据

consumer_server 对达成的订单进行写入操作

query_server 提供API查询，供内部RPC和三方http使用



数据库分表

根据商家的shopId来分表,在注册时分配,查询时需要提供shopId字段


orders_1  表示shopId为1的商家，提供ID分配算法，查询时直接对应orders_1，查询做分页处理



压力测试