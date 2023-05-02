### 参数类型
flag包支持的命令行参数类型有：
`bool`：1/0/t/f/T/F/true/false/TRUE/FALSE/True/False

`int`/`int64`/`unit`/`uint64`：1234/0644/0x123/-1

`float`/`float64`：合法浮点数

`string`：合法字符串

`duration`：合法时间字符串，单位有ns/us/ms/s/m/h，如"300ms"/"-1.5h"/"2h30m"

#### 定义flag参数
`flag.Type()`: `flag.Type(flag名, 默认值, 帮助信息) *Type`，例如要定义姓名，年龄，是否已婚的参数，可以使用如下方式，注意`生成的字段参数都是对应类型的指针`：
```go
name := flag.String("name", "zs", "姓名")
age := flag.Int("age", 19, "年龄")
married := flag.Bool("married", false, "是否已婚")
delay := flag.Duration("d", time.Second, "间隔")
flag.Parse()
fmt.Printf("姓名:%v, 年龄:%v, 是否已婚:%v, 时间间隔:%v\n", name, age, married, delay)
```
可以使用编译后的执行文件+`--help`来查看用法提示，或+`-name=值`来指定某一参数值，或+`--age=值`来指定某一参数值。

`flag.TypeVal()`: `flag.TypeVal(Type指针, flag名, 默认值, 帮助信息)`，例如要定义姓名，年龄，是否已婚的参数，可以使用如下方式：
```go
var (
    name    string
    age     int
    married bool
    delay   time.Duration
)
flag.StringVal(&name, "name", "zs", "姓名")
flag.IntVal(&age, "age", 19, "年龄")
flag.BoolVal(&married, "married", false, "是否已婚")
flag.Duration(&delay, "delay", 0, "间隔")
flag.Parse()
```

`flag.Parse()`: 解析命令行参数。flag解析在第一个非falg参数(单个 - 不是flag参数)之前停止，或在终止符 `-` 之后停止。

支持的命令行参数格式：`-flag xxx(使用空格，一个 - 符号)`/`--flag xxx(使用空格，两个 -- 符号)`/`-flag=xxx(使用=，一个 - 符号)`/`--flag=xxx(使用等号，两个 - 符号)`。其中布尔型参数必须使用等号方式指定。

#### flag其他函数
`flag.Args()`: 返回命令行参数后的其他参数，[]string类型

`flag.NArg()`: 返回命令行参数后的其他参数个数

`flag.NFlag()`: 返回使用的命令行参数个数
