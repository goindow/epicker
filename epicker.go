package epicker

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	picker *Picker
)

type Picker struct {
	logger *log.Logger
}

func init() {
	picker = &Picker{
		logger: log.New(os.Stderr, "", log.Ltime|log.Lshortfile),
	}
}

func pick(e error) (error, bool) {
	if e != nil {
		return e, true
	}
	return e, false
}

func dump(s string) {
	// 调用栈追溯要设置成 3，显示调用 picker 的行号，而不是显示 picker 中函数的行号(默认 2，因为我们又包装了一层，故需要再向上一层)，log.Logger.Output 会自动追加换行
	picker.logger.Output(3, s)
}

func Print(e error) {
	if err, ok := pick(e); ok {
		dump(err.Error())
	}
}

func Printf(e error, format string, v ...interface{}) {
	if _, ok := pick(e); ok {
		dump(fmt.Sprintf(format, v...) + " (" + e.Error() + ")")
	}
}

func Fatal(e error) {
	if err, ok := pick(e); ok {
		dump(err.Error())
		os.Exit(1)
	}
}

func Fatalf(e error, format string, v ...interface{}) {
	if _, ok := pick(e); ok {
		dump(fmt.Sprintf(format, v...) + " (" + e.Error() + ")")
		os.Exit(1)
	}
}

func Panic(e error) {
	if _, ok := pick(e); ok {
		panic(e)
	}
}

func Panicf(e error, format string, v ...interface{}) {
	if _, ok := pick(e); ok {
		panic(errors.New(fmt.Sprintf(format, v...) + " (" + e.Error() + ")"))
	}
}
