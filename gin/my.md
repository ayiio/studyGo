### gin
    获取：`go get -u github.com/gin-gonic/gin`
### gin路由
#### 基本路由
    gin框架中采用的路由库是基于httprouter做的
    地址：https://github.com/julienschmidt/httprouter
#### Restful风格API
    支持Representational State Transfer，即表现层状态转化，一种互联网应用程序的API设计理念，URL定位资源，HTTP描述操作。
#### API参数
    通过Context的Param方法获取
#### URL参数
    通过DefaultQuery() 或 Query() 方法获取
#### 表单参数
    传输为post请求，常见格式有4中：
        - application/json
        - application/X-WWW-form-urlencoded
        - application/xml
        - multipart/form-data
    表单参数可以通过 PostForm() 方法获取，默认解析是 X-WWW-form-urlencoded 或 form-data格式的参数
#### 上传单个文件
    multipart/form-data格式用于文件上传
    gin的文件上传与net/http原生方法类似，不同在于gin将request封装到了c.Request中
#### 上传多个文件
    <input type="file" name="files" multiple>
#### 路由组
    r.Group("/v1")
    {
        v1.Get("/get1", func1)
    }
#### 路由原理
    httprouter会将所有路由规则构造一棵前缀树
#### 参数绑定
    JSON/form/uri参数绑定
        + json: gin.Context.ShouldBindJson(200, gin.H{})
        + form: gin.Context.Bind(200, gin.H{})
        + uri: gin.Context.ShouldBindURI(200, gin.H{})
        + yaml: gin.Context.YAML(200, gin.H{})
        + protobuf:  高效存储读取工具
        ```
        r.GET("/protobuf", func(c *gin.Context) {
            reps := []int64{int64(1), int64(2)}
            //定义数据
            label := "label"
            //传protobuf格式数据
            data := protoexample.Test {
                Label : &label,
                Reps: reps,
            }
            c.ProtoBuf(200, data)
        })
        ```
#### HTML模板渲染
    gin加载HTML模板后，根据模板参数进行配置并返回相应的数据，本质是字符串替换
    LoadHTMLGlob() 方法加载模板文件
    <h1>{{ .xxx }}</h1>
#### 重定向
    gin.Context.Redirect()
#### 同步异步
    goroutine实现异步处理，启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本
#### gin中间件
    构建的中间件只对注册过的路由函数起作用
    对于分组路由，嵌套使用中间件，可以限定中间件的作用范围
    中间件分为全局中间件，单个路由中间件和群组中间件
    中间件必须是gin.HandlerFunc类型
##### 全局中间件
    所有请求都经过此中间件
##### Next方法
    根据handlerFunc列表长度遍历
##### 局部中间件
    将中间件加入到请求参数中
#### 会话控制
##### Cookie是什么
    http是无状态协议，服务器不能记录浏览器的访问状态，也就是说服务器不能区分两次请求是否由同一个客户端发起
    Cookie用于解决http协议无状态的方案之一，实际上是服务器保存在浏览器上的一段信息，浏览器有了Cookie后，每次向服务器发送请求时都会同时将该信息发送给服务器，服务器收到请求后，就可以根据该信息处理请求
    Cookie由服务器创建，并发送给浏览器，最终由浏览器保存
##### Cookie的用途
    保持用户登录状态
    购物车，浏览限制
##### Cookie的使用
    测试服务器发送Cookie给客户端，客户端请求时携带Cookie，Cookie信息存储在客户端
##### Cookie缺点
    不安全，明文
    增加带宽消耗
    可以被禁用
    有上限
##### Session是什么
    session弥补了Cookie的不足，session必须依赖于Cookie才能使用，生成一个sessionId放在Cookie里传给客户端即可(实际Cookie信息存储在服务端，只返回sessionId到客户端)
##### session中间件
    通用session，支持内存存储和Redis缓存
    设计思路：
        + session模块设计(session规范)
            + k-v系统，通过key进行增删查改
            + session可以存储在内存或Redis中
            + 用户与session一对一，session内一对多kv，session与sessionMgr多对一
            + session接口设计
                + Set()
                + Get()
                + Del()
                + Save()，session延迟加载(redis)
        + sessionMgr接口设计(Mgr规范)
            + Init() 初始化，加载redis地址
            + CreateSession() 创建一个新的session
            + GetSession() 通过SessionID获取对应的session对象
        + MemorySession设计(session实现)
            + 定义MemorySession对象(字段：sessionID，存kv map，读写锁)
            + 构造函数，获取对象
            + memorySession接口设计
                + Set()
                + Get()
                + Del()
                + Save()，session延迟加载(redis)
        + MemorySessionMgr设计(Mgr实现)
            + 定义MemorySessionMgr对象(字段：session map，读写锁)
            + 构造函数，获取对象
            + memorySessionMgr接口设计
                + Init()
                + CreateSession()
                + GetSession()
        + RedisSession设计(session实现)
            + 定义RedisSession对象(字段：sessionID，存kv map，读写锁， redis连接池，记录内存中的map是否被修改的标记)
            + 构造函数
            + redisSession接口设计
                + Set() 将session存储到内存中的map
                + Get() 取数据，实现延迟加载
                + Del()
                + Save() 将session存储到redis
        + RedisSessionMgr设计(Mgr实现)
            + 定义RedisSessionMgr对象(字段：redis地址，redis密码，连接池，读写锁，大map)
            + 构造函数
            + RedisSessionMgr接口设计
                + Init()
                + CreateSession()
                + GetSession()
    
