package epicker

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

const (
	ERR               = "error info"
	CUSTOM_ERR        = "custom error info"
	CUSTOM_ERR_FORMAT = "format %s"

	ENV_FATAL_ERR  = "GO_TEST_EPICKER_FATAL_ERR"
	ENV_FATAL_NIL  = "GO_TEST_EPICKER_FATAL_NIL"
	ENV_FATALF_ERR = "GO_TEST_EPICKER_FATALF_ERR"
	ENV_FATALF_NIL = "GO_TEST_EPICKER_FATALF_NIL"
)

var (
	buf       bytes.Buffer
	err       = errors.New(ERR)
	customErr = fmt.Sprintf(CUSTOM_ERR_FORMAT, CUSTOM_ERR) + " (" + ERR + ")"
)

func init() {
	SetLogger(&buf, "", 0) // 设置 logger，默认是输出到 os.stderr，为了测试，将输出放到 buf 变量中，并且不比较 flag 的输出(0，关闭标志位输出)
}

func Test_Print_Err(t *testing.T) {
	buf.Reset()
	if Print(err); (string)(bytes.TrimSpace(buf.Bytes())) != ERR {
		fail(t, "epicker.Print() should print error("+ERR+")")
	}
}

func Test_Print_Nil(t *testing.T) {
	buf.Reset()
	if Print(nil); buf.Len() != 0 {
		fail(t, "epicker.Print() should do nothing")
	}
}

func Test_Printf_Err(t *testing.T) {
	buf.Reset()
	if Printf(err, CUSTOM_ERR_FORMAT, CUSTOM_ERR); (string)(bytes.TrimSpace(buf.Bytes())) != customErr {
		fail(t, "epicker.Printf() should print error("+customErr+")")
	}
}

func Test_Printf_Nil(t *testing.T) {
	buf.Reset()
	if Printf(nil, CUSTOM_ERR_FORMAT, CUSTOM_ERR); buf.Len() != 0 {
		fail(t, "epicker.Printf() should do nothing")
	}
}

func Test_Fatal_Err(t *testing.T) {
	// cmd 内执行的代码，读取标志，防止死循环
	if os.Getenv(ENV_FATAL_ERR) == "1" {
		SetLogger(os.Stdout, "", 0) // 由于 cmd 另起一个线程，cmd.Output 只能捕获标准输出，所以切换回来 os.Stdout（buf 只存在于各自的堆栈上，并不能通信）
		Fatal(err)
		return // return 不能少，测试成功会跳出循环，如果测试失败的话，可能不会执行 Fatal 内的 os.exit(1)，将会导致死循环
	}
	// 相当于起了另外一个线程执行命令，go test -test.run=Test_Fatal，创造一个封闭的环境，不会因为被测试代码中含有 os.exit() 而退出测试代码
	cmd := exec.Command(os.Args[0], "-test.run=Test_Fatal_Err")
	cmd.Env = append(os.Environ(), ENV_FATAL_ERR+"=1") // 设置标志
	out, er := cmd.Output()                            // 执行命令并返回标准输出的切片
	// 没有识别 err
	if er == nil {
		fail(t, "epicker.Fatal() should print error")
	}
	// 没有正确 exit(1)
	if e, ok := er.(*exec.ExitError); !ok || e.Success() {
		fail(t, "epicker.Fatal() should exit with code 1")
	}
	// 没有正确输出错误信息
	if string(bytes.TrimSpace(out)) != ERR {
		fail(t, "epicker.Fatal() should print error("+ERR+")")
	}
}

func Test_Fatal_Nil(t *testing.T) {
	if os.Getenv(ENV_FATAL_NIL) == "1" {
		SetLogger(os.Stdout, "", 0)
		Fatal(nil)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=Test_Fatal_Nil")
	cmd.Env = append(os.Environ(), ENV_FATAL_NIL+"=1")
	_, er := cmd.Output()
	if er != nil {
		fail(t, "epicker.Fatal() should do nothing")
	}
}

func Test_Fatalf_Err(t *testing.T) {
	if os.Getenv(ENV_FATALF_ERR) == "1" {
		SetLogger(os.Stdout, "", 0)
		Fatalf(err, CUSTOM_ERR_FORMAT, CUSTOM_ERR)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=Test_Fatalf_Err")
	cmd.Env = append(os.Environ(), ENV_FATALF_ERR+"=1")
	out, er := cmd.Output()
	if er == nil {
		fail(t, "epicker.Fatalf() should print error")
	}
	if e, ok := er.(*exec.ExitError); !ok || e.Success() {
		fail(t, "epicker.Fatalf() should exit with code 1")
	}
	if string(bytes.TrimSpace(out)) != customErr {
		fail(t, "epicker.Fatalf() should print error("+customErr+")")
	}
}

func Test_Fatalf_Nil(t *testing.T) {
	if os.Getenv(ENV_FATALF_NIL) == "1" {
		SetLogger(os.Stdout, "", 0)
		Fatalf(nil, CUSTOM_ERR_FORMAT, CUSTOM_ERR)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=Test_Fatalf_Nil")
	cmd.Env = append(os.Environ(), ENV_FATALF_NIL+"=1")
	_, er := cmd.Output()
	if er != nil {
		fail(t, "epicker.Fatalf() should do nothing")
	}
}

func Test_Panic_Err(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			fail(t, "epicker.Panic() should print error with a painc")
		} else if fmt.Sprint(p) != ERR {
			fail(t, "epicker.Panic() should print error("+ERR+")")
		}
	}()
	Panic(err)
}

func Test_Panic_Nil(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fail(t, "epicker.Panic() should do nothing")
		}
	}()
	Panic(nil)
}

func Test_Panicf_Err(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			fail(t, "epicker.Panicf() should print error with a painc")
		} else if fmt.Sprint(p) != customErr {
			fail(t, "epicker.Panicf() should print error("+customErr+")")
		}
	}()
	Panicf(err, CUSTOM_ERR_FORMAT, CUSTOM_ERR)
}

func Test_Panicf_Nil(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fail(t, "epicker.Panicf() should do nothing")
		}
	}()
	Panicf(nil, CUSTOM_ERR_FORMAT, CUSTOM_ERR)
}

func Benchmark_Print(b *testing.B) {
	buf.Reset()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Print(err)
	}
}

func Benchmark_Printf(b *testing.B) {
	buf.Reset()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Printf(err, CUSTOM_ERR_FORMAT, CUSTOM_ERR)
	}
}

func fail(t *testing.T, s string) {
	echo(t, s, 1)
}

func ok(t *testing.T, s string) {
	echo(t, s, 2)
}

func echo(t *testing.T, s string, level uint) {
	switch level {
	case 1:
		t.Error("[fail] " + s)
	case 2:
		t.Log("[ok] " + s)
	}
}
