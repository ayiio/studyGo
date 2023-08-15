## 1.go-micro简介
* Go Micro是一个插件化的基础框架，基于此可以构建微服务，Micro的设计哲学是可插拔的插件化架构
* 默认实现了consul 作为服务发现，2019年源码修改默认使用mdns，通过http进行通信，通过protobuf和json进行编码
## 2.go-micro的主要功能
* 服务发现：自动服务注册和名称解析。服务发现是微服务开发的核心。当服务A需要与服务B通话时，它需要该服务的位置，默认发现机制是多播DNS(mdsn)一种零配系统。可以选择使用SWIM协议为p2p网站设置广播，或者为弹性云原生设置consul。
* 负载均衡：基于服务发现构建的客户端负载均衡。一旦我们获得了服务的任意数量实例的地址，需要一种方法决定路由到哪个节点，可以使用随机散列负载均衡来提供跨服务的均匀分布，并在出现问题时重试不同的节点。
* 消息编码：基于内容类型的动态消息编码，客户端和服务端将使用编码器和内容类型来无缝编码和解码Go类型。可以编码任何种类的消息并从不同的客户端发送。客户端和服务端默认处理此问题。包括默认的protobuf和json。
* 请求/响应：基于RPC的请求/响应，支持双向流。提供了同步通信的抽象，对于服务的请求将自动解决，负载均衡，拨号和流式传输。启用tls时，默认传输为http/1.1或http2。
* Async Messaging：Pubsub是异步通信和事件驱动架构的重要支持，事件通知是微服务开发的核心模式，启用tls时，默认消息传递是点对点http/1.1或http2。
* 可插拔接口：Go Micro为每个分布式系统抽象使用Go接口，实现可插拔。并允许Go Micro为运行时无关，可以插入任何基础技术。
  * 插件地址：https://github.com/micro/go-plugins
## 3.go-micro通信流程
* Server监听客户端的调用，和Broker推送过来的消息进行处理。server端需要向register注册自己的存在或消亡，基于此client才能知道自己的状态。
* Register服务的注册发现，client端向register中得到server的信息，然后每次调用都基于算法选择一个server进行通信，通信经过编码/解码和选择传输协议等一系列过程。
* 如果有需要通知所有的server端可以使用broker进行信息的推送，broker信息队列进行信息的推送和发布。
## 4.go-micro核心接口
* go-micro由8个核心的interface构成，每个接口根据自己的需求重新实现。
service -> (client/server) -> (Broker/Codec/Registry/Selector/Transort)
## 5.go-micro接口核心实现
* Transort通信接口
```
type Scoker interface {
    Recv(*Message) error
    Send(*Message) error
    Close() error
}

type Client interface {
    Socket
}

type Listener interface {
    Addr() string
    Close() error
    Accept(func(Socket)) error
}

type Transport interface {
    Dial(addr string, opts ...DialOption) (Client, error)
    Listen(addr string, opts ...ListenOption) (Listener, error)
    String() string
}
```
* Codec编码接口
```
type Codec interface {
  ReadHeader(*Message, MessageType) error
  ReadBody(interface{}) error
  Write(*Message, interface{}) error
  Close() error
  String() string
}
```
* Registry注册接口
```
type Registry interface {
  Register(*Service, ...RegisterOption) error
  Deregister(*Service) error
  GetService(string) ([]*Service, error)
  ListService() ([]*Service, error)
  Watch(...WatchOption) (Watcher, error)
  String() string
  Options() Options
}
```
* Selector负载均衡
```
type Selector interface {
  Init(opts ...Option) error
  Options() Options
  Select(service string, opts ...SelectOption) (Next, error)
  Mark(service string, node *registry.Node, err error)
  Reset(service string)
  Close() error
  String() string
}
```
* Broker发布订阅接口
```
type Broker interface {
  Options() Options
  Address() string
  Connect() error
  Disconnect() error
  Init(...Option) error
  Publish(string, *Message, ...PublishOption) error
  Subscribe(string, Handler, ...SubscribeOption) (Subscriber, error)
  String() string
}
```
* Client客户端接口
```
type Client interface {
  Init(...Option) error
  Options() Options
  NewMessage(topic string, msg interface{}, opts ...MessageOption) Message
  NewRequest(service, method string, req interface{}, reqOpts ...RequestOption) Request
  Call(ctx context.Context, req Request, rsp interface{}, opts ...CallOption) error
  Stream(ctx context.Context, req Request, opts ...CallOption) (Stream, error)
  Publish(ctx context.Context, msg Message, opts ...PublishOption) error
  String() string
}
```
* Server服务端接口
```
type Server interface {
  Options() Options
  Init(...Option) error
  Handle(Handler) error
  NewHandler(interface{}, ...HandlerOption) Handler
  NewSubscriber(string, interface{}, ...SubscriberOption) Subscriber
  Subscribe(Subscriber) error
  Register() error
  Deregister() error
  Start() error
  Stop() error
  String() string
}
```
* Service接口
```
type Service interface {
  Init(...Option)
  Options() Options
  Client() client.Client
  Server() server.Server
  Run() error
  String() string
}
```
## 6.go-micro安装
* 查看网址：https://github.com/micro/go-micro
* cmd中输入如下3条命名，自动下载关联包
  * go get github.com/micro/micro
  * go get github.com/micro/go-micro
  * go get github.com/micro/protoc-gen-micro  
* 安装go-micro：go install github.com/micro/protoc-gen-micro
* 安装micro.exe:
  * git clone https://github.com/go-micro/cli.git
  * cd cmd/go-micro
  * go build -o micro.exe
 
