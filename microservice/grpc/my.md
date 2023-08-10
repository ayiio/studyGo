## gRPC简介
* 由Google开发的一款语言中立、开源的远程过程调用系统   
* gRPC客户端和服务端可以在多种环境中运行和交互，例如Java实现服务端，go实现客户端调用  
## gRPC 和 Protobuf简介
* 微服务框架中，由于每个服务对应的代码库是独立运行的，无法直接调用，彼此间的通信是一个大问题  
* gRPC可以实现微服务，将大的项目拆分为多个小且独立的业务模块，也就是服务，各服务间使用高效的protobuf协议进行RPC调用，gRPC默认使用protocol buffers，即由google开源的一套成熟的结构数据序列化机制。也可以使用其他数据格式如JSON
* 可以用proto files创建gRPC服务，用message类型来定义方法参数和返回类型
## 安装gRPC 和 Protobuf
* go get github.com/golang/protobuf/proto
* go get google.golang.org/grpc (无法使用时替换为如下命令)   
  ```
  git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
  git clone https://github.com/goloang/net.git $GOPATH/src/golang.org/x/net
  git clone https://github.com/golang/text.git $GOPATH/src/golang.ort/x/text
  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
  go clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
  cd $GOPATH/src
  go install google.golang.org/grpc
  ```
* go get github.com/golang/protobuf/protoc-gen-go ==> go install google.golang.org/protobuf/cmd/protoc-gen-go@latest [参考](https://www.cnblogs.com/shiding/p/16608117.html)
* 上面安装好后，会在GOPATH/bin下生成protoc-gen-go.exe
* 但还需要一个protoc.exe，windows较难实现手动编译，可以下载 `https://github.com/protocolbuffers/protobuf/releases/tag/v3.9.0` 放到GOPATH/bin下  
## Protobuf语法
### 1.基本规范
* 文件以`.proto`作为后缀名，除结构定义外的语句以分号结尾
* 结构定义可以包含：`message` / `service` / `enum`
* rpc方法定义结尾的分号可有可无
* Message命名采用驼峰命名方式，字段命名采用小写字母加下划线分割方式
```
message TestServerRequest {
    required string test_name = 1;
}
```
* Enums类型名采用驼峰命名方式，字段命名采用大写字母加下划线分割方式
```
enum Foo {
    FIRST_VALUE = 1;
    SENCOND_VALUE = 2;
}
```
* Service 与 rpc 方法名统一采用驼峰式命名
### 2.字段规则
* 字段格式：限定修饰符 | 数据类型 | 字段名称 | = | 字段编码值 | [字段默认值]
* 限定修饰符包含 `required` / `optional` / `repeated`
  * Required：表示是一个必须字段，必须相对应于发送方，在发送消息之前必须设置该字段的值，对于接收方，必须能够识别该字段的意思。发送之前没有设置required字段或者无法识别required字段都会引发编码异常，导致消息被丢弃
  * Optional：表示是一个可选字段，可选对于发送方，在发送消息时，可以有选择性的设置或者不设置该字段的值，对于接收方，如果能够识别可选字段就进行相应的处理，如果无法识别，则忽略该字段，消息中的其他字段正常处理。
    因为Optional字段的特性，很多接口在升级中都把后来添加的字段都统一设置为optional字段，这样老的版本无需升级程序也可以正常与新的程序进行通信，只不过新的字段无法识别而已，以为并不是每一个节点都需要新的功能，因此可以做到按需升级和平滑过渡。
  * Repeated：表示该字段可以包含0~N个元素。其特性和optional一样，但是每一次可以包含多个值。可以看作是在传递一个数组的值。
* 数据类型

