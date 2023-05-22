### ELK
应用服务端(AppServer) -> 日志收集器(Logstash Agent) -> ES检索集群(ElasticSearch Cluster) -> 展示器(Kibana) -> 客户端(Browser)

存在的问题：
运维成本高，每增加一个日志收集器，都需要手动修改配置 / 监控缺失，无法准确获取logstash的状态 / 无法做到定制化开发和维护

### 改进方案
[架构](https://github.com/ayiio/StudyGo/blob/master/kafka/logAgent.jpg)

logAgent从etcd中获取配置信息，如果有新配置再通过kafka分发，写入到ES/Hadoop等，再通过Kibana进行检索。监控系统直接或通过kafka写入数据库，再使用Grafana做客户端展示。

### kafka补充命令
客户端消费: `bin\windows\kafka-console-consumer.bat --bootstrap-server 127.0.0.1:9091 --topic topic1 --from-beginning`

创建消费者时如果出现out of memory, map failed的错误，可以修改`bin\windows\kafka-server-start.bat`中的以下片段进行解决，即将默认的1G都改为512M即可。[参考出处](https://debugah.com/kafka-error-caused-by-java-lang-outofmemoryerror-map-failed-how-to-solve-21161/)
```sh
IF ["%KAFKA_HEAP_OPTS%"] EQU [""] (
    rem detect OS architecture
    wmic os get osarchitecture | find /i "32-bit" >nul 2>&1
    IF NOT ERRORLEVEL 1 (
        rem 32-bit OS
        set KAFKA_HEAP_OPTS=-Xmx512M -Xms512M
    ) ELSE (
        rem 64-bit OS
        set KAFKA_HEAP_OPTS=-Xmx512M -Xms512M    
    )
)
```

