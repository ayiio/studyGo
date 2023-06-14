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
#### 中间件

