# es_service
全文索引服务

op:
 - [x] 增: 后期通过kafka消息队列,解耦，防止影响业务主逻辑
 - [x] 删: 通过kafka消息队列
 - [x] 改: 通过kafka消息队列
 - [x] 查: 直接调用服务返回
