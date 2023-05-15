### 物料
#### JDK 
##### 安装: [JAVA_JDK](https://www.oracle.com/java/technologies/downloads/)
##### 配置 : jdk解压目录配置到PATH：%JAVA_HOME% 和 %JAVA_HOME%\bin
#### zookeeper 
##### 安装 : [zookeeper](https://github.com/apache/zookeeper/tags)
##### 配置 : 将conf\zoo_sample.cfg复制一份，并重命名为zoo.cfg，随后修改zoo.cfg中的dataDir值为本地解压目录\data。启动使用bin\zkServer.cmd
#### kafka 
##### 安装: [kafka](https://mirrors.tuna.tsinghua.edu.cn/apache/kafka/)
##### 配置 : 修改config\server.properties，更改log.dirs值为本地解压目录\kafkalogs。启动使用bin\windows\kafka-server-start.bat config\server.properties

### 消息队列的通信模式
#### 点对点模式(queue)
消息生产者生产消息发送到queue中，消费者从queue中取出并消费消息，一条消息被消费后queue中不再存有重复消息。

#### 发布/订阅(topic)
消息生产者(发布)将消息发布到topic中，同时有多个消息订阅者(订阅)消费该消息，发布到topic中的消息会被所有订阅者消费(类似微信公众号文章等)。

发布订阅模式下，当发布者消息量很大时，单个订阅者处理能力将不足。实际场景是多个订阅者节点组成一个订阅组负载均衡消费topic消息，即分组订阅，从而实现消费能力线性扩展。可以看作一个topic下有多个queue，每个queue是点对点方式，queue之间是发布订阅方式。

#### kafka
分布式数据流平台，可单台或集群部署，提供了发布和订阅功能，使用者可以发送数据到kafka，也可以从kafka中读取数据做后续处理。kafka具有高吞吐、低延迟、高容错等特点。
##### 构成
`Producer`：生产者。消息的产生者，消息入口。
`kafka cluster`：kafka集群，一台或多台服务器组成。
    + `Broker`：部署了kafka实例的服务器节点。每个服务器上有一个或多个kafka实例，每个kafka集群内的broker都有一个`不重复`的编号，例如broker-0, broker-1等
    + `Topic`：消息的主题，或称消息的分类。kafka数据就保存在topic。每个broker上可以创建多个topic，实际应用是一个业务线创建一个topic
    + `Partition`：topic的分区，每个topic可以有多个分区，分区作用是做负载，提高kafka的吞吐量。同一个topic在不同的分区的数据是不重复的，partition表现形式是一个个的文件夹
    + `Replication`：每个分区都有多个副本，用于备份。主分区(Leader)故障时会选择一个备份(Follower)成为Leader。默认副本最大数量时10个，且副本数量不能大于broker数量，follower和leader绝对是在不同机器，同一个机器对同一个分区也只可能存放一个副本(包括自己)
`Consumer`：消费者。消息的消费方，消息出口。
    + `Consumer Group`：可以将多个消费者组成一个消费者组，同一个分区的数据只能被消费者组中的某一个消费者消费。同一个消费者组的消费者可以消费同一个topic的不同分区的数据，提高了kafka的吞吐量。
##### 工作流程
producer是生产者，即数据入口。producer写入数据时会把数据写入leader中而不会写入follower，数据写入流程：
    + 1.生产者从kafka集群获取分区leader信息
    + 2.生产者将信息发送给leader
    + 3.leader将信息写入本地磁盘
    + 4.follower从leader拉取消息数据
    + 5.follower将消息写入本地磁盘后向leader发送ACK
    + 6.leader收到所有follower的ACK后向生产者发送ACK
##### 选择partition原则
某个topic中有多个partition，producer发数据到对应partition的原则：
    + 1.partition写入时可以被指定，并按指定写入
    + 2.没有指定partition，但设置了数据的key，则会根据key的值hash出一个partition
    + 3.没有指定partition，且没有设置key，则采用轮询方式，即每次取一小段时间的数据写入某个partition，下一小段时间写入下一个partition。
##### ACK应答机制
producer向kafka写入消息，可以使用参数`0`,`1`,`all`进行设置来确定是否确认kafka接收到数据。
    + 0代表producer往集群发送数据不需要等到集群的返回，不确保消息发送成功。安全性最低但效率最高。
    + 1代表producer往集群发送数据只要等待leader应答就可以发送下一条，只确保leader发送成功。
    + all代表producer往集群发送数据需要所有follower都完成从leader的同步才会发送下一条，确保leader发送成功和所有副本都完成备份。安全性最高，但效率最低。
如果往不存在的topic写数据，kafka会自动创建topic，partition和replication的数量默认配置都是1.
##### topic和数据日志
topic是同一类别的消息记录(record)的集合。kafka中一个主题通常有多个订阅者。对于每个主题，kafka集群维护了一个分区数据日志文件结构。

每个partition都是一个有序且不可变的消息记录集合。当新的数据写入时，就被追加到partition的末尾。在每个partition中，每条消息都会被分配一个顺序的唯一标识，这个标识被成为offset，即偏移量。kafka只保证在同一个partition内部消息是有序的，在不同的partition之间，并不能保证消息有序。

Kafka可以配置一个保留期限，用来标识日志会在kafka集群内保留多长时间。kafka集群会保留仍在保留期限内所有已被发布的消息，不管消息是否被消费过。例如保留期限是1天，则数据被发布到kafka集群的一天内都可以被消费，超过1天后将被清空，以便为后续数据腾出空间。kafka可以将数据写入磁盘进行持久化，因此保留的数据大小可以设置为一个比较大的值。
##### partition结构
表现形式是一个个的文件夹，每个文件夹下会有多组segment文件，每组segment文件又包含`.index`文件、`.log`文件和`timeindex`文件。`.log`文件是实际存储message的地方，而`.index`和`.timeindex`文件为索引文件，用于检索消息。
##### 消费数据
多个消费者实例可以组成一个消费者组，并用一个标签来标识这个消费者组，一个消费者组中的不同消费者实例可以运行在不同的进程甚至不同的服务器上。

如果所有的消费者实例都在同一个消费者组中，那么消息记录会被很好均衡的发送到每个消费者实例。如果所有消费者实例都不在同一个消费者组，那么每一条消息记录会被广播到每一个消费者实例。


例如一个由2个broker(server1/2)组成的kafka集群上，1个topic拥有4个partition(p0-p3)，有两个消费者组都在消费这个topic中的数据，消费者组A有两个消费者实例，消费者组B有四个消费者实例。同一个消费者组中，每个消费者实例可以消费多个分区，但每个分区最多只能被消费者组中的一个实例消费，即如果有一个4个分区的主题，那么消费者组中最多只能有4个消费者实例去消费，多出来的都不会被分配到分区。如果在消费者组中动态的上线或下线消费者，那么kafka集群会自动调整分区与消费者实例间的对应关系。
##### 使用场景
###### 消息队列(MQ)
跨进程的通信机制，用于上下游的消息传递，可以使用MQ将上下游解耦。MQ使用场景如流量削峰、数据驱动的任务依赖等，类似有kafka、ActiveMQ、RabbitMQ等。
###### 追踪网站活动
对网站活动(PV，UV，搜索记录等)进行追踪。可以将不同的活动放到不同的主题，供后续的实时计算/实时监控等程序使用，也可以将数据导入到数据仓库中进行后续的离线处理和报表生成。
###### Metrics
传输监控数据。用来聚合分布式应用程序的统计数据，将数据集中后进行统一的分析和展示等。其中日志聚合是指将不同服务器上的日志收集起来并放入到一个日志中心，例如一台文件服务器或者HDFS中的一个目录，供后续进行分析处理。

