## 服务注册组件开发
### 服务注册和发现原理  
  client -> pull -> **registry**  
  client <- push <- **registry**  
  client -> call -> server  
  server -> register -> **registry**  
### 注册中心选型
<img width="243" alt="image" src="https://github.com/ayiio/studyGo/assets/61615400/e1b80c19-fbe0-4f56-a644-9b7e31a0326f"> 

euerka(aws)主从式存储，存在一定的延迟。  
consul/zookeeper/etcd 分布式存储，zk重量级。  
### 选项设计模式
主结构体字段变更，但不改变随后的初始化逻辑  
1.声明一个函数类型的变量，用于传参  
2.初始化函数中遍历参数，得到每一个函数  
3.调用函数，在函数里给传入的对象赋值  
4.对象字段赋值完成  
### 注册接口开发
1.支持多注册中心，既支持consul也支持etcd  
2.支持可扩展  
3.提供基于名字的插件管理函数，用来注册插件  
