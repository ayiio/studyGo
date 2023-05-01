基准测试以`Benchmark`为前缀，需要一个`*testing.B`类型的参数b，基准测试必须要执行`b.N`次来保证有对照性。`b.N`的值是根据实际情况调整，以保证稳定性。

对比方式的基准测试，如同一个函数处理元素个数不同所对应的耗时差别可能不同，再或者相同的输入不同的算法实现，对应的性能也不同。可以通过性能比较函数，即`Benchamark`函数传入不同的参数来实现。

默认情况下，每个基准测试至少运行1秒。`Benchmark`函数返回时如果还没到1秒，则b.N的值会按1, 2, 5, 10, 20, 50...增加，并且函数再次执行。可以使用`-benchtime`标志增加最少基准时间，以产生更准确的结果。

重置时间，`b.ResetTimer`之前的处理不会累计到执行时间里，也不会报告到最终结果里，因此有些需要不计入测试报告的过程(连接数据库)，可以在执行对应过程后使用该方法重置累计时间。

并行测试，`func (b *B) RunParallel(body func(*PB))`会以并行的方式执行给定的基准测试。 `RunParallel`会创建出多个`goroutine(默认GOMAXPROCS)`，并将`b.N`分配给这些`goroutine`执行。如果需要增加非CPU受限(non-CPU-bound)基准测试的并行性，可以在`RunParallel`之前调用`SetParallelism`，`RunParallel`通常会和`-cpu`标志一起使用。

Setup和TearDown，测试准备和测试后恢复(如数据库插入+删除操作)。通常在`*_test.go`中定义`TestMain`函数，进而在测试之前进行额外的设置`setup`或在测试之后进行拆卸`teardown`操作。如果在测试文件中写入函数`func TestMain(m *testing.M)`，那么测试将先调用`TestMain(m)`，然后再运行具体的测试。`TestMain`运行在主`goroutine`中，可以在调用`m.Run`前后做任何`setup`或`teardown`，退出测试时应使用`m.Run`的返回值作为参数调用`os.Exit`。

示例函数，以`Example`为前缀，无参数无返回值。示例函数能够作为文档直接使用，例如基于web的godoc。示例函数只要包含了`// Output:`也是可以通过`go test`执行测试。
