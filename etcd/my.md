### etcd介绍
etcd是使用Go开发的开源、高可用的分布式key-value存储系统，用于配置共享和服务的注册和发现。类似项目有zookeeper和consul。</br>
特点：</br>
1.完全复制：集群中的每个节点都可以使用完整的存档

2.高可用性：etcd可用于避免硬件的单点故障和网格问题

3.一致性：每次读取都会返回跨多主机的最新写入

4.简单：包含一个定义良好、面向用户的API(gRPC)/zookeeper的zad协议

5.安全：实现了带有可选的客户端证书身份验证的自动化TLS

6.快速：每秒10000次写入的基本速度

7.可靠：使用Raft算法实现了强一致性、高可用的服务存储目录

Raft协议：选举、日志复制机制、异常处理/脑裂

#### 应用场景
##### 服务发现
分布式系统常见问题之一，即在同一个分布式集群中的进程或服务，要如何才能找到对方并建立连接。服务发现本质上是想要了解集群中是否有进程在监听udp或tcp端口，并且通过名字就可以查找和连接。
##### 配置中心
将一些配置信息放到etcd上进行集中管理。应用在启动时主动从etcd上获取一次配置信息，同时在etcd节点上注册一个Watcher并等待，以后每次配置有更新时，etcd都会实时通知订阅者，以此达到获取最新配置信息的目的。
##### 分布式锁
etcd使用的Raft算法保证了数据的强一致性，某次操作存储到集群中的值必然是全局一致的，所以很容易实现分布式锁。锁服务有两种使用方式：保持独占和控制时序。
  + `保持独占即所有获取锁的用户最终只有一个可以得到`。etcd为此提供了一套实现分布式锁原子操作CAS(`compareAndSwap`)的API。
通过设置`prevExist`值，可以保证在多个节点同时去创建某个目录时，只有一个成功，而创建成功的用户就可以认为是获得了锁。
  + `控制时序`，即所有想要获取锁的用户都会被安排执行，但是`获得锁的顺序是全局唯一的，同时决定了执行顺序`。
etcd为此提供了一套API(自动创建有序键)，对一个目录建值时指定为`POST`动作，这样etcd会自动在目录下生成一个当前最大的值为键，存储这个新的值(客户端编号)。
同时还可以使用API按顺序列出所有当前目录下的键值，此时这些键的值就是客户端的时序，而这些键中存储的值可以是代表客户端的编号。

#### 为什么使用etcd而不是zookeeper
etcd可实现的功能zookeeper也可实现，但相较起来zookeeper存在如下缺点：

1.复杂。zookeeper的`部署维护复杂`，管理员需要掌握一系列的知识和技能，而Paxos强一致性算法也是复杂难懂。另外zookeeper使用也比较复杂，需要安装客户端，官方只提供了Java和C两种语言的接口。

2.Java编写。属于重型应用，会引入大量依赖，而运维更偏向希望保持强一致性、高可用的机器集群尽可能简单，维护起来不易出错。

3.发展缓慢。apache基金会项目特有的apache way在开源界饱受争议，一大原因是由于基金会庞大的结构以及松散的管理导致项目发展缓慢。

etcd具备的优点(CoreOS/Kubernetes/CloudFoundry等已在生产环境用到etcd，仍缺少大项目长时间的检验)：

1.简单。使用Go编写-部署简单，使用HTTP接口-使用简单，使用Raft算法保证强一致性-易于理解。

2.数据持久化。etcd默认数据一更新就进行持久化。

3.安全。etcd支持SSL客户端安全认证。

#### etcd架构
etcd架构主要分为四个部分：
  + HTTP server：用于处理用户发送的API请求以及其他etcd节点的同步与心跳信息请求。
  + Store：用于处理etcd支持的各类功能的事务，包括数据索引、节点状态变更、监控与反馈、事件处理与执行等，是etcd对用户提供大多数API功能的具体实现。
  + Raft：etcd核心，是强一致性算法的具体实现。
  + WAL：Write Ahead Log(预写式日志)，是etcd的数据存储方式。除了在内存中存有所有数据的状态以及节点的索引之外，etcd通过WAL进行持久化存储。WAL中所有数据提交前都会事先记录日志。
Snapshot是为了防止数据过多而进行的状态快照，Entry表示存储的具体日志内容。

#### etcd集群
Raft算法在做决策时需要多数节点的投票，所以etcd一般部署集群推荐奇数个节点，推荐数为3、5或7个节点构成一个集群。

搭建一个3个节点集群，在每个etcd节点指定集群成员，为了区分不同的集群最好同时配置一个独一无二的token，以下示例n1、n2和n3表示3个不同的etcd节点：

方式1：
```bash
TOKEN=token-01
CLUSTER_STATE=new
CLUSTER=n1=http://xxx.xxx.xxx.11:2380,n2=http://xxx.xxx.xxx.22:2380,n3=http://xxx.xxx.xxx.33:2380
```
在n1机器上启动etcd，
--initial-advertise-peer-urls自节点，
--listen-peer-urls集群多节点间通信，
--advertise-client-urls客户端节点，
--listen-client-urls客户端访问节点，
--initial-cluster归属集群，
--initial-cluster-state状态，
--initial-cluster-token集群中用于区分的token:
```bash
etcd --data-dir=data.etcd --name n1 \
   --initial-advertise-peer-urls http://xxx.xxx.xxx.11:2380 --listen-peer-urls http://xxx.xxx.xxx.11:2380 \
   --advertise-client-urls http://xxx.xxx.xxx.11:2379 --listen-client-urls http://xxx.xxx.xxx.11:2379 \
   --initial-cluster ${CLUSTER} \
   --initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN}
```
在n2机器上启动etcd:
```bash
etcd --data-dir=data.etcd --name n2 \
   --initial-advertise-peer-urls http://xxx.xxx.xxx.22:2380 --listen-peer-urls http://xxx.xxx.xxx.22:2380 \
   --advertise-client-urls http://xxx.xxx.xxx.22:2379 --listen-client-urls http://xxx.xxx.xxx.22:2379 \
   --initial-cluster ${CLUSTER} \
   --initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN}
```
在n3机器上启动etcd:
```bash
etcd --data-dir=data.etcd --name n3 \
   --initial-advertise-peer-urls http://xxx.xxx.xxx.33:2380 --listen-peer-urls http://xxx.xxx.xxx.33:2380 \
   --advertise-client-urls http://xxx.xxx.xxx.33:2379 --listen-client-urls http://xxx.xxx.xxx.33:2379 \
   --initial-cluster ${CLUSTER} \
   --initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN}
```

方式2：

etcd提供了一个可以公网访问的etcd存储地址，可以通过以下命令得到etcd服务的目录，并把它作为-discovery参数使用：
```bash
curl https://discovery.etcd.io/new?size=3
https://discovery.etcd.io/a81xxxxxxxecbc92

# grab this token
TOKEN=token-01
CLUSTER_STATE=new
DISCOVERY=https://discovery.etcd.io/a81xxxxxxxecbc92

etcd --data-dir=data.etcd --name n1 \
   --initial-advertise-peer-urls http://xxx.xxx.xxx.11:2380 --listen-peer-urls http://xxx.xxx.xxx.11:2380 \
   --advertise-client-urls http://xxx.xxx.xxx.11:2379 --listen-client-urls http://xxx.xxx.xxx.11:2379 \
   --initial-cluster ${CLUSTER} \
   --initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN}
  
etcd --data-dir=data.etcd --name n2 \
   --initial-advertise-peer-urls http://xxx.xxx.xxx.22:2380 --listen-peer-urls http://xxx.xxx.xxx.22:2380 \
   --advertise-client-urls http://xxx.xxx.xxx.22:2379 --listen-client-urls http://xxx.xxx.xxx.22:2379 \
   --initial-cluster ${CLUSTER} \
   --initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN}
   
etcd --data-dir=data.etcd --name n3 \
   --initial-advertise-peer-urls http://xxx.xxx.xxx.33:2380 --listen-peer-urls http://xxx.xxx.xxx.33:2380 \
   --advertise-client-urls http://xxx.xxx.xxx.33:2379 --listen-client-urls http://xxx.xxx.xxx.33:2379 \
   --initial-cluster ${CLUSTER} \
   --initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN}
```

集群搭建好后，可以使用`etcdctl`来连接etcd。
```bash
exprot ETCDCTL_API=3
HOST_1=xxx.xxx.xxx.11
HOST_2=xxx.xxx.xxx.22
HOST_3=xxx.xxx.xxx.33
ENDPOINTS=$HOST_1:2379,$HOST_2:2379,$HOST_3:2379

etcdctl --endpoints=$ENDPOINTS member list
```

#### etcd搭建
[下载地址](https://github.com/etcd-io/etcd/releases)

以Windows平台为例，下载zip文件后解压到本地即可，双击`etcd.exe`即启动了etcd，其他平台需要解压后在bin目录下寻找etcd可执行文件。默认在2379端口监听客户端通信，在2380端口监听节点间通信。
etcdctl.exe可以理解为客户端或本机的etcd控制端。

连接etcd：旧版本默认etcdctl使用了v2版本的命令，即环境变量设置为ETCDCTL_API=2，如果put/set命令遇到提示`No help topic for 'put'`时，需要设置环境变量SET ETCDCTL_API=3来使用v3版本的API。

put: `etcdctl --endpoints=http://127.0.0.1:2379 put kfc "v50"`

get：`etcdctl --endpoints=http://127.0.0.1:2379 get kfc`

delte：`etcdctl --endpoints=http://127.0.0.1:2379 del kfc`

#### Go 操作etcd
安装：`go get go.etcd.io/etcd/clientv3`

#### Go mod 注意事项
```go
replace (
    google.golang.org/grpc v1.55.0 => google.golang.org/grpc v1.26.0
)
```

