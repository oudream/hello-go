
### gin 日誌
- https://cloud.tencent.com/developer/article/1468169
- https://blog.gtwang.org/web-development/network-lantency-and-bandwidth/
- https://www.bookstack.cn/read/gin-doc/middleware.md
```shell
// 日志格式
logger.WithFields(logrus.Fields{
 "status_code"  : statusCode,
 "latency_time" : latencyTime,
 "client_ip"    : clientIP,
 "req_method"   : reqMethod,
 "req_uri"      : reqUri,
}).Info()
```