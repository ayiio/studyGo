## 工作流程
1.读日志 - tailf第三方库`github.com/hpcloud/tail`
2.写日志到kafka - sarama第三方库`github.com/Shopify/sarama`，v1.20后加入了zstd算法，需要用到cgo，windows平台编译会提示错误：`exec: "gcc": executable file not found in %PATH%`，可使用v1.19版本的sarama避开此问题。
```sh
    require (
        github.com/shopify/sarama v1.19.0
    )
```

### go.mod
`require github.com/Shopify/sarama v1.19.0`

go mod init + go mod download/go mod tidy

### kafka log
INFO [Partition web_log-0 broker=0] Log loaded for partition web_log-0 with initial high watermark 0 (kafka.cluster.Partition)

### kafka-logs
kafka-logs\web_log-0: </br>
00000000000000000000.index</br>
00000000000000000000.log</br>
00000000000000000000.timeindex</br>
leader-epoch-checkpoint</br>
partition.metadata
