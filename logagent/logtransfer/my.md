### 实现目的
#### logtransfer
从kafka中将日志取出来，写入ES，并使用Kibana做可视化展示。
##### Elastic search
一个基于Lucene构建的开源、分布式、RESTful接口的全文搜索引擎，同时也是分布式文档数据库，其中每个字段都可以被索引，且每个字段的数据均可被搜索。ES可以横向扩展至数以百计的服务器存储以及处理PB级别的数据。可以在极短时间内存储、搜索和分析大量数据，通常做为具有复杂搜索场景的核心。
###### ES可做什么
1.商城产品目录和库存信息，提供精准搜索和推荐。/ 2.日志收集或交易数据分析和挖掘，寻找趋势、进行统计或发现异常。该场景可以使用Logstash等工具进行数据收集，并将数据存储到ES中。/ 3.类似Github的search。
###### 基本概念
1.`Near Realtime(NRT)` 几乎实时：是一个实时的搜索平台，索引一个文档到这个文档可被搜索，这之间只需要毫秒级的延迟。

2.`Cluster`集群：群集是一个或多个节点(服务器)的集合，这些节点共同保存整个数据，并在所有节点上提供联合索引和搜索功能。一个集群有一个唯一集群ID确定，并指定一个集群名(默认为"elasticsearch")，集群名很重要，节点可以通过这个集群名加入群集，一个节点只能是群集的一部分。需要确保在不同的环境中不要使用相同的群集名称，以免错误连接，如logging-dev/logging-stage/logging-prod分别为开发/阶段产品/生产集群做记录。

3.`Node`节点：节点是单个服务器实例，是群集的一部分，可以存储数据，并参与群集的索引和搜索功能。如一个集群，节点的名称默认为一个随机的通用唯一标识符(UUID)，确定在启动时分配给该节点，可以自定义节点名。这个节点名可以确定网络服务器对应的ElasticSearch群集节点。可以通过群集名配置节点以连接特定的群集。默认情况下每个节点设置加入名为"elasticsearch"的集群。

4.`Index`索引：具有相似特性的文档集合。例如客户数据/产品目录/订单数据分别为一个索引。索引由名称(必须全部小写)标识，该名称用于在对其中的文档进行索引、搜索、更新和删除操作时引用索引。

5.`Type`类型：索引中可以定义一个或多个类型。类型是索引的逻辑类别/分区，自定义语义。一般类型定义为具有公共字段集的文档。例如博客平台，所有数据存储在一个索引中，其中用户数据为一种类型，博客数据类型是一种类型，注释数据是另一种类型。

6.`Document`文档：是可以被索引的信息的基本单位。例如为单个客户提供一个文档，单个产品提供另一个文档，以及单个订单提供其他文档。文档表现形式为JSON格式。尽管文档物理驻留在索引中，文档实际上必须索引或分配到索引中的类型。

7.`Shards & Replicas`分片和副本：索引可以存储大量数据，而这些数据可能会超过单个节点的硬件限制。可以使用ES提供的分片功能，例如当创建一个索引时，可以简单定义需要的分片数量。每个分片本身是一个全功能的、独立的"指数"，可以依托在集群中的任意节点。集群网络或云环境中，不可避免会有故障出现，ES允许创建一个或多个拷贝，形成故障转移机制用以防止分片和节点离线或消失，索引分片进入副本或复制品的分片，被称为Replicas。

`Shards分片的重要性`主要体现在如下两个特征：
1.分片允许水平拆分或缩放内容大小。/ 2.分片允许分配和并行操作碎片(可能在多个节点上)，从而提高性能/吞吐量。这个机制中的碎片的分布方式以及文件汇总的搜索请求均由ES管理，用户方透明。

`Replicas副本重要性`主要体现在如下两个特征：
1.副本为分片或节点失败提供了高可用性。一个副本的分片不会分配在同一个节点作为原始的或主分片，副本是从主分片复制而来。/ 2.副本允许用户扩展搜索量或吞吐量，搜索可以在副本上并行执行。

###### 和关系型数据库的比较
ES中Index(索引)支持全文检索，类似关系型数据库中的Database(数据库)。</br>
ES中Type(类型)类似关系型数据库中的Table(表)。</br>
ES中Document(文档)，不同文档可以有不同的字段集合，类似关系型数据库中的Row(数据行)。</br>
ES中Field(字段)类似关系型数据库中的Column(数据列)。</br>
ES中Mapping(映射)类似关系型数据库中的Schema(模式)。

###### ES搭建
1.[ES下载](https://www.elastic.co/cn/downloads/elasticsearch) / [ES旧版本](https://www.elastic.co/cn/downloads/past-releases#elasticsearch)

2.启动：以Windows为例，解压后执行bin\elasticsearch.bat，默认使用本机9200端口。使用浏览器访问elasticsearch服务(localhost:9200)，可以查看对应的节点信息。

###### ES使用
`curl -X GET 127.0.0.1:9200/_cat/health?v`用于查看健康状态。</br>
`curl -X GET 127.0.0.1:9200/_cat/indices?v`用于查询当前es集群中所有的indices索引。</br>
`curl -X PUT 127.0.0.1:9200/www`用于创建索引。</br>
`curl -X DELETE 127.0.0.1:9200/www`用于删除索引。</br>
`curl -H "ContentType:application/json" -X POST 127.0.0.1:9200/user/person -d '{"name":"test", "age":22, "married":true}'`用于插入数据。</br>
`curl -X GET 127.0.0.1:9200/user/person/_search`用于检索。</br>
`curl -H "ContentType:application/json" -X PUT 127.0.0.1:9200/user/person/4 -d '{"query":{"match":{"name":"test"}}}'`用于按条件检索，ES默认一次最多返回10条结果，可以通过size字段设置返回结果的数目：`curl -H "ContentType:application/json" -X PUT 127.0.0.1:9200/user/person/4 -d '{"query":{"match":{"name":"test"}}, "size":2}'`。

###### Go操作ES
[第三方库](https://github.com/olivere/elastic)，使用go.mod管理依赖，注意保持和ES版本一致。`require ( github.com/olivere/elastic/v8 v8.0.0 )`

##### Kibana
开源的分析和图形化展示平台，与ElasticSearch一起工作，需要和ES保持同一版本。可以使用Kibana以搜索、查看等方式和存储在ElasticSearch索引的数据进行交互。</br>
###### Kibana搭建
1.[Kibana下载](https://www.elastic.co/cn/downloads/kibana)

2.配置：以Windows为例，解压后修改config/kibana.yml，将其中的`elasticsearch.hosts`设置为指定地址，例如`elasticsearch.hosts: ["http://127.0.0.1:9200"]`。同时修改语言为简体中文，`i18n.locale: "zh-CN"`。

3.启动：以Windows为例，解压后执行bin\kibana.bat，等待启动过程完成后即可看到对应的图形化界面。

#### 系统监控
gopsutil做系统监控信息的采集，写入influxDB，使用grafana展示。

Prometheus监控，采集性能指标数据，保存并使用grafana展示。
