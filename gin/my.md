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
