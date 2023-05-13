### 物料
#### JDK 
##### 安装: [JAVA_JDK](https://www.oracle.com/java/technologies/downloads/)
##### 配置 : jdk解压目录配置到PATH：%JAVA_HOME% 和 %JAVA_HOME%\bin
#### zookeeper 
##### 安装 : [zookeeper](https://github.com/apache/zookeeper/tags)
##### 配置 : 将conf\zoo_sample.cfg复制一份，并重命名为zoo.cfg，随后修改zoo.cfg中的dataDir值为本地解压目录\data。启动使用bin\zkServer.cmd
#### kafka 
##### 安装: [kafka](https://github.com/apache/kafka/tags)
##### 配置 : 修改config\server.properties，更改log.dirs值为本地解压目录\kafkalogs。启动使用bin\windows\kafka-server-start.bat config\server.properties
