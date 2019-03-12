# epicker
实现了一个简单的 error 处理器

## 特性
- 当 error 为 nil，不做任何处理；为 error，相应函数处理错误
- picker 是基于 log 的封，底层调用 log.Logger.Output 会在最后追加换行

## 测试 & 文档
```shell
// 单元测试
go test [-v]
// 基准测试
go test -test.bench=".*"
// 本地文档
godoc -http=:6060
```

## 索引
- [Print](#Print)
- [Printf](#Printf)
- [Fatal](#Fatal)
- [Fatalf](#Fatalf)
- [Panic](#Panic)
- [Panicf](#Panicf)
- [SetLogger](#SetLogger)

### 使用
```go
import (
	"github.com/goindow/epicker"
)
```

### Print
- 如果 e 不是 nil，则打印错误
```go
// e := nil
e := errors.New("error info")

epicker.Print(e)
// todo:
```

### Printf
- 如果 e 不是 nil，则使用自定义格式打印错误(会追加 e 的内容)
```go
// e := nil
e := errors.New("error info")

epicker.Printf(e, "format %s", "custom error info")
// todo:
```

### Fatal
- 如果 e 不是 nil，则打印错误，并退出程序
- 本质是 Print() 后，接着执行 os.exit(1)
```go
// e := nil
e := errors.New("error info")

epicker.Fatal(e)
// todo:
```

### Fatalf
- 如果 e 不是 nil，则使用自定义格式打印错误(会追加 e 的内容)，并退出程序
- 本质是 Print() 后，接着执行 os.exit(1)
```go
// e := nil
e := errors.New("error info")

epicker.Fatalf(e, "format %s", "custom error info")
// todo:
```

### Panic
- 如果 e 不是 nil，则抛出一个包含 e 的 panic
```go
// e := nil
e := errors.New("error info")

epicker.Panic(e)
// todo:
```

### Panicf
- 如果 e 不是 nil，则抛出一个自定义格式错误(会追加 e 的内容)的 panic
```go
// e := nil
e := errors.New("error info")

epicker.Panicf(e, "format %s", "custom error info")
// todo:
```

### SetLogger
- 设置 picker.logger
- 本质是调用 log.New(out io.Writer, prefix string, flag int)，方便切换 logger
- 本包有一套默认的 logger 设置，一般无需调用，除非需要自定义日志输出目的地（默认 os.Stderr）、日志前缀（默认 ""）、标志（默认 log.Ltime | log.Lshortfile，时间及文件名和行号）
```go
var buf bytes.buffer

epicker.SetLogger(&buf, "prefix", 0)	// 将错误信息输出到 buf 中（而不是默认的 os.Stderr，一般是显示设备，如显示器）
```

## 一个完整的例子
```go
package epicker_test

import (
	"github.com/goindow/epicker"
	"os"
)

var (
	pwd, _        = os.Getwd()
	nonExistsFile = pwd + "/non_exists.file"
)

func ExamplePrint() {
	epicker.SetLogger(os.Stdout, "", 0)
	_, err := os.Open(nonExistsFile)
	epicker.Print(err)
	// Output:
	// open /usr/local/var/go/src/github.com/goindow/epicker/non_exists.file: no such file or directory
}

func ExamplePrintf() {
	epicker.SetLogger(os.Stdout, "", 0)
	_, err := os.Open(nonExistsFile)
	epicker.Printf(err, "format %s", "custom erorr info")
	// Output:
	// format custom erorr info (open /usr/local/var/go/src/github.com/goindow/epicker/non_exists.file: no such file or directory)
}
```

## 更多信息
- [epicker_test.go](https://github.com/goindow/epicker/blob/master/epicker_test.go)
- [example_test.go](https://github.com/goindow/epicker/blob/master/example_test.go)
