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

##### 命令行方式分析：

应用程序执行结束后会生成保存有CPU使用情况的文件，使用`go tool pprof`工具进行CPU性能分析。例如生成的文件为cpu.pporf，可以使用`go tool pprof cpu.pprof`进入交互式界面进行查看，并使用`top 3`查看CPU使用占比前三的函数，交互界面通过`q`退出。

cup.pprof参数说明：

flat: 当前函数占用CPU的耗时。

flat%: 当前函数占用CPU的耗时百分比。

sum%: 函数占用CPU的耗时累计百分比。

cum: 当前函数加上调用当前函数的函数占用CPU的总耗时。

cum%: 当前函数加上调用当前函数的函数占用CPU的总耗时百分比。

最后一列: 函数名称。

还可以使用`list 函数名`的方式查看具体的函数分析。

##### 图形化方式分析：
调用图：需要安装`graphviz`图形化工具。Windows访问 [graphviz](https://graphviz.org/) 下载后将bin文件夹添加到Path环境变量，通过终端命令`dot -version`查看是否安装成功； mac通过`brew install graphviz`进行安装。安装完成后在`go tool pprof cpu.pprof`的交互式页面中输入`web`就可以通过图形化方式在浏览器中查看调用关系，其中每个框代表一个函数，框越大表示占用CPU资源越多，方框之间的线条表示函数调用关系，线条数字表示函数调用次数，方框中第一行数字表示当前函数占用CPU的百分比，第二行数字表示当前函数累计占用CPU的百分比。

火焰图：或者通过开源工具`go-torch`读取profiling数据生成火焰图来更直观查看结果。安装：`go get -v github.com/uber/go-torch`。可以点击每个方块动态分析对应的内容。火焰图调用顺序从下至上，每个方块表示一个函数，上一层表示这个函数会调用哪些函数，方块大小表示占用CPU时长。该工具若为传入任何参数，则会尝试从`http://localhost:8080/debug/pprof/profile`获取profiling数据，它有三个常用可调参数：

-u -url: 要访问的URL，主机和端口部分

-s -suffix: pprof profile的路径，默认/debug/pprof/profile

-seconds: 要执行profiling的时间长度，默认30秒


`go-torch`需要配合`FlameGraph`工具使用，通过以下方式安装：

1.安装perl: [perl donwload](https://strawberryperl.com/)

2.下载FlameGraph：git clone https://github.com/brendangregg/FlameGraph.git

3.将FlameGraph目录加入到操作系统环境变量中

4.windows平台需要更改go-torch/render/flamegraph.go中的GenerateFlameGraph，然后在go-torch目录下执行go install即可
```go
flameGraph := findInPath(flameGraphScripts)
if flameGraph == "" {
   return nil, errNoPerlScript
}

// add
if runtime.GOOS == "windows" {
   return runScript("perl", append([]string(flameGraph), args...), graphInput)
}
return runScript(flameGraph, args, graphInput)
}
```
5.配合压测工具[wrk](https://github.com/wg/wrk)或[wrk2](https://github.com/adjust/go-wrk)

6.使用wrk进行压测，`go-wrk -n 50000 http://127.0.0.1:8080/test/list`压测同时在另一个终端执行`go-torch -u http://127.0.0.1:8080 -t 30`，30秒后终端会提示`Writing svg to torch.svg`，最后通过浏览器可以查看`torch.svg`就可以看到火焰图。



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



