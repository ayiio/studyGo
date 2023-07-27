### 定义服务注册接口 Registry，定义方法
<img width="334" alt="image" src="https://github.com/ayiio/studyGo/assets/61615400/72148b56-104a-45f8-b33c-60c823d68353">

  Registry -> etcdRegister/consulRegister -> 调用方  
  Name(): 插件名，例如etcd  
  Init(opts ...Option):初始化，里面用选项设计模式做初始化  
  Register():服务注册  
  Unregister():服务反注册，例如服务器停掉，注册列表销毁  
  GetService():服务发现(ip port[]string)  
### 抽象出一些结构体
  Node:单个节点的结构体，包含id, ip, port, weight(权重)   
  Service：服务名，节点列表，一个服务多台服务器支撑  
### 选项设计模式，实现参数初始化
### 插件管理类
  大map管理，key字符串，value为Register接口对象  
  支持用户自定义调用，自定义插件  
  实现注册中心的初始化，供系统使用  
### etcd插件注册

  
