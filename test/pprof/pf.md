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



