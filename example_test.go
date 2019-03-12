package epicker_test

import (
	"os"
	"github.com/goindow/epicker"
)

var (
	pwd, _        = os.Getwd()
	nonExistsFile = pwd + "/non_exists.file"
)

func setLogger() {
	epicker.SetLogger(os.Stderr, "", 0)
}

func ExamplePrint() {
	_, err := os.Open(nonExistsFile)
	epicker.Print(err)
	// Todo: something
	// Output:
	// open /usr/local/var/go/src/github.com/goindow/epicker/non_exists.file: no such file or directory
}

func ExamplePrintf() {
	_, err := os.Open(nonExistsFile)
	epicker.Printf(err, "format %s", "custom erorr info")
	// Todo: something
	// Output:
	// format custom error info (open /usr/local/var/go/src/github.com/goindow/epicker/non_exists.file: no such file or directory)
}