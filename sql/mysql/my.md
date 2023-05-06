## 数据库
常见关系型数据库有SQLlite，MySQL，SQLServer，postgreSQL，Oracle。

关系型数据库用于存放一类数据。

KV数据库有Redis。

### MySQL
#### SQL语句
DDL：操作数据库</br>
DML：表的增删改查</br>
DCL：用户和权限</br>
#### 存储引擎
支持插件式存储引擎。

常见存储引擎：MyISAM和InnoDB。其中MyISAM查询速度快，支持表锁，不支持事务。InnoDB整体速度快，支持表锁和行锁，支持事务。

事务：把多个SQL操作当作一个整体。
ACID特点：1.原子性(事务要么成功要么失败，没有中间状态)，2.一致性(事务开始到结束，数据库中的数据保持完整性)，3.隔离性(事务之间是互相不影响，隔离级别)，4.持久性(事务操作结果落盘，不会丢失)。

表结构设计三大范式：
1.遵从原子性(表中的数据字段不能再拆分，遵从唯一性，消除部分依赖)，
2.一张表只能描述一件事(任意主键或联合主键可以确认除主键外的所有非主键值，消除部分传递依赖)，
3.不能存在用非主键a去确认非主键b的值(消除传递依赖)。

索引：B树和B+树。有唯一索引和联合索引等，索引命中规则等。

分库分表

SQL注入

SQL慢查询优化

MySQL主从复制：binlog。

MySQL读写分离：从库读主库写，数据一致性。

#### 创建库表
```mysql
CREATE DATABASE sql_test;

use sql_test;

CREATE TABLE `user` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(20) DEFAULT '',
  `age` INT(11) DEFAULT '0',
  PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
```
