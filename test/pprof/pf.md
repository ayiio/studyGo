## 性能调优
profiling是指对应程序的画像，反映了程序使用CPU和内存的情况。

## 性能优化分类：

CPU Profile: 报告程序的CPU使用情况，按照一定频率去采集应用程序在CPU和寄存器上的数据

Memory Profile(Heap Profile)：报告程序的内存使用情况

Block Profiling：报告goroutines不在运行状态的情况，可以用来分析和查找死锁等性能瓶颈

Goroutine Profiling：报告goroutines的使用情况，有哪些goroutine以及它们的调用关系

## 采集性能数据
`runtime/pprof`：采集`工具型`应用数据进行分析

`net/http/pporf`：采集`服务型`应用运行时数据进行分析

pprof开启后，每隔一段时间(10ms)就会收集当前的堆栈信息，获取各函数占用CPU和内存状况，最后通过分析采集到的数据，产出性能分析报告。应只在做性能测试时才引入pprof。

### 工具型应用
非持续运行的程序，建议在应用程序退出时将profiling报告保存在文件中。可在代码中引入`runtime/pprof`库。

#### CPU性能分析
开启CPU性能分析：`pprof.StartCPUProfile(w io.Writer)`

停止CPU性能分析：`pprof.StopCPUProfile()`

应用程序执行结束后会生成保存有CPU使用情况的文件，使用`go tool pprof`工具进行CPU性能分析。

#### 内存性能分析
记录堆栈信息：`pprof.WriteHeapProfile(w io.Writer)`

得到采样数据后，使用`go tool pprof`工具进行内存性能分析。默认使用`-inuse_space`进行统计，还可以使用`-inuse-objects`查看分配对象的数量。

### 服务型应用
持续运行的程序，例如web应用，如果使用默认的`http.DefaultServeMux`(直接使用`http.ListenAndServe("0.0.0.0:8080", nil)`)，只需要在web server端代码中引入`net/http/pprof`：`import _ "net/http/pprof"`。

如果使用自定义Mux，需要手动注册一些路由规则：
```go
r.HandleFunc("/debug/pprof/", pprof.Index)
r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
r.HandleFunc("/debug/pprof/profile", pprof.Profile)
r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
r.HandleFunc("/debug/pprof/trace", pprof.Trace)
```

如果使用gin框架，推荐使用`"github.com/DeanThompson/ginpprof"`。

无论哪种方式，HTTP服务都会多出`/debug/pprof` endpoint。


