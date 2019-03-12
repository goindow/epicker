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
