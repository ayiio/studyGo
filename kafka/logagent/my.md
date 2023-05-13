### ELK
应用服务端(AppServer) -> 日志收集器(Logstash Agent) -> ES检索集群(ElasticSearch Cluster) -> 展示器(Kibana) -> 客户端(Browser)

存在的问题：
运维成本高，每增加一个日志收集器，都需要手动修改配置 / 监控缺失，无法准确获取logstash的状态 / 无法做到定制化开发和维护

### 改进方案
[架构](https://github.com/ayiio/StudyGo/blob/master/kafka/logAgent.jpg)

logAgent从etcd中获取配置信息，如果有新配置再通过kafka分发，写入到ES/Hadoop等，再通过Kibana进行检索。监控系统直接或通过kafka写入数据库，再使用Grafana做客户端展示。

