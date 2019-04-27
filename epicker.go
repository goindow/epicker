// Package epicker 封装了一个简单的 error 处理器
package epicker

import (
    "errors"
    "fmt"
    "io"
    "log"
    "os"
)

const (
    FLAG = log.Ltime | log.Lshortfile // 默认的日志输出标记，时间及文件名和行号
)

// Picker 是对 log.Logger 的一层包装
type Picker struct {
    logger *log.Logger
}

var (
    picker *Picker
)

func init() {
    picker = &Picker{
        logger: log.New(os.Stderr, "", FLAG),
    }
}

func pick(e error) (error, bool) {
    if e != nil {
        return e, true
    }
    return e, false
}

func dump(s string) {
    picker.logger.Output(3, s) // 调用栈追溯要设置成 3，显示调用 picker 的行号，而不是显示 picker 中函数的行号(默认 2，因为我们又包装了一层，故需要再向上一层)，log.Logger.Output 会自动追加换行
}

// SetLogger 设置 picker.logger，
// 本质是调用 log.New(out io.Writer, prefix string, flag int)，方便切换 logger，包有一套默认设置，一般无需调用
func SetLogger(out io.Writer, prefix string, flag int) {
    picker.logger = log.New(out, prefix, flag)
}

// Print 如果 e 不是 nil，则打印错误
func Print(e error) {
    if err, ok := pick(e); ok {
        dump(err.Error())
    }
}

// Printf 如果 e 不是 nil，则使用自定义格式打印错误(会追加 e 的内容)
func Printf(e error, format string, v ...interface{}) {
    if _, ok := pick(e); ok {
        dump(fmt.Sprintf(format, v...) + " (" + e.Error() + ")")
    }
}

// Fatal 如果 e 不是 nil，则打印错误，并退出程序，
// 本质是 Print() 后，接着执行 os.exit(1)
func Fatal(e error) {
    if err, ok := pick(e); ok {
        dump(err.Error())
        os.Exit(1)
    }
}

// Fatalf 如果 e 不是 nil，则使用自定义格式打印错误(会追加 e 的内容)，并退出程序，
// 本质是 Printf() 后，接着执行 os.exit(1)
func Fatalf(e error, format string, v ...interface{}) {
    if _, ok := pick(e); ok {
        dump(fmt.Sprintf(format, v...) + " (" + e.Error() + ")")
        os.Exit(1)
    }
}

// Panic 如果 e 不是 nil，则抛出一个包含 e 的 panic
func Panic(e error) {
    if _, ok := pick(e); ok {
        panic(e)
    }
}

// Panicf 如果 e 不是 nil，则抛出一个自定义格式错误(会追加 e 的内容)的 panic
func Panicf(e error, format string, v ...interface{}) {
    if _, ok := pick(e); ok {
        panic(errors.New(fmt.Sprintf(format, v...) + " (" + e.Error() + ")"))
    }
}
