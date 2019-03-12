package epicker_test

import (
	"github.com/goindow/epicker"
	"os"
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
