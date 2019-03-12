package epicker_test

import (
	"os"
	"fmt"
	"epicker"
)

var (
	pwd, _ = os.Getwd()
	nonExistsFile = pwd + "/non_exists.file"
)

func setLogger() {
	
}

func ExamplePrint() {
	_, err := os.Open(nonExistsFile)
	epicker.Print(err)
	// Todo:
	// Output:
	// open /usr/local/var/go/src/goutils/epicker/non_exists.file: no such file or directory
}
