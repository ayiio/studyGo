## Redis
开源内存数据库，一种KV 数据库。

支持诸如字符串(string)、哈希(hashes)、列表(lists)、集合(sets)、带范式查询的排序集合(sorted set)、位图(bitmaps)、hyperloglogs、带半径查询和流的地理空间索引等数据结构。Redis支持更丰富的5种数据类型，支持RDB持久化，以及master/slave模式。

### 常规用处：

1.cache缓存，减轻主数据库(MySQL)的压力。

2.简单的队列，使用LIST实现。

3.热门排行榜，需要排序的场景特别适合使用ZSET。

4.记数场景，比如微博、抖音中的关注和粉丝数。

### 获取：
`go get -u github.com/go-redis/redis`

或

`go get github.com/redis/go-redis/v9`
