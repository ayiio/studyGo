### context
1.7后加入的标准库，用于简化对于处理单个请求的多个goroutine之间与请求域的数据、取消信号、截至时间等相关操作。

服务器传入的请求应该创建上下文，调用传出应该接收上下文，之间的函数调用链需要传递上下文。或者可以使用`WithCancel`、`WithDeadline`、`WithTimeout`或`WithValue`创建派生上下文，当一个上下文被取消时，它所派生的所有上下文也将被取消。

### context接口
`Deadline() (deadline Time, ok bool)`：返回context被取消的时间/完工时间

`Done() <-chan struct{}`：返回一个Channel，会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel

`Err() error`：返回当前context结束原因，只会在Done返回的channel被关闭时才会返回非空的值，如果当前context被取消就会返回canceled错误，如果当前context超时就会返回DeadlineExceeded错误

`Value(key interface{}) interface{}`：会从context中返回键对应的值，对于同一个上下文，多次调用value并传入相同的key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据

### Background()和TODO()
分别返回一个实现context接口的`background`和`todo`，都是`emptyCtx`结构体类型，是一个不可被取消，没有设置截止时间，没有携带任何值的context，被视为最顶层`parent context`，用于衍生更多的子上下文对象。
其中`Background()`主要用于main函数、初始化以及测试代码中，作为上下文最顶层的根context。TODO()用于未知context的占位。

### With系列函数
`WithCancel`：返回带有新Done通道的父节点副本，当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论发生什么情况。
```go
func gen(ctx context.Context) <-chan int {
  dst := make(chan int)
  n := 1
  go func() {
    for {
      select {
        case <-ctx.Done():
          return //结束goroutine
        case dst<- n:
          n++
      }
    }
  }()
  return dst
}

func main() {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  
  for n := range gen(ctx) {
    fmt.Println(n)
    if n == 5 {
      break
    }
  }
}
```

`WithDeadline`：返回父上下文的副本，并将deadline调整为不迟于d。如果父上下文的deadline已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。
当截至日过期时/调用返回的cancel函数时/父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。
```go
func main() {
  d := time.Now().Add(50 * time.Millisecond)
  ctx, cancel := context.WithDeadline(context.Background(), d)
  
  //调用cacel用于更好维护ctx，而不必必须等待ctx超时
  defer cancel()
  
  select {
    case <-time.After(1 * time.Second):
      fmt.Println("overslept")
    case <-ctx.Done():
      fmt.Println(ctx.Err())
  }
}
```

`WithTimeout`：返回`WithDeadline(parent, time.Now().Add(timeout))`，取消此上下文将释放与其相关的资源，应该在此上下文运行结束后立即调用cancel，通常用于数据库或网络连接超时控制。
```go
var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
   for {
    fmt.Println("db connecting...")
    time.Sleep(time.Millisecond * 10)  // 模拟数据库连接
    select {
      case <-ctx.Done():  // 50毫秒后自动调用
        break LOOP
      default:
    }
   }
}

func main() {
  ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)  // 50毫秒参数
  wg.Add(1)
  go worker(ctx)
  time.Sleep(time.Second * 5)
  cancel() // 通知子goroutine结束
  wg.Wait()
  fmt.Println("Done")
}
```

`WithValue`：返回父节点副本，其中与key关联的值为val，能将请求作用域的数据和Context对象建立关系，仅对API和进程间传递请求域的数据使用上下文值，而不是用于传递可选参数给函数。
提供的键必须是可比较的，且不应是`string`或其他任何内置类型，以免包之间使用上下文发生冲突。为避免参数interface{}的再次分配，键应该用自定义类型，通常是有具体类型的`struct{}`，导出的上下文关键变量的静态类型应该是指针或接口。
```go
type TraceCode string // 自定义类型，避免和内置类型发生冲突

var wg sync.WatiGroup

func worker(ctx context.Context) {
  key := TraceCode("TRACE_CODE")
  traceCode, ok := ctx.Value(key).(string) // 子goroutine中获取trace code
  if !ok {
    fmt.Println("invalid trace code")
  }

LOOP:
  for {
    fmt.Printf("worker, trace code:%s\n", traceCode)
    time.Sleep(time.Millisecond, 10) // 模拟数据库连接
    select {
      case <-ctx.Done():
        break LOOP
      default:
    }
  }
  
  fmt.Println("worker done!")
  wg.Done()
}

func main() {
  ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
  ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "id123")  // 系统入口中设置trace code传递给后续启动的goroutine实现据库聚合
  wg.Add(1)
  go worker(ctx)
  time.Sleep(time.Second * 5)
  cancel() // 通知子goroutine结束
  wg.Wait()
  fmt.Println("done")
}
```

### 注意事项
推荐以参数的方式显示传递Context</br>
以Context作为参数的函数方法，应该把Context作为第一个参数</br>
给一个函数传递Context时，不要传递nil，如果不清楚传递什么，可以使用context.TODO()</br>
Context的Value相关方法应该传递请求域的必要数据，不要用于传递可选参数</br>
Context是线程安全的，可以在多个goroutine间传递


