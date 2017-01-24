package util

import (
	"fmt"
	"strings"
	"path/filepath"
)

func Error(err error) {
	fmt.Println("ERROR: ",err)
}

func Notice(n string) {
	fmt.Println("NOTICE: ",n)
}

//Prints each key file name to the terminal
func PrintKeyFile(n string) {
	fname := strings.TrimSuffix(n, filepath.Ext(n))
	fmt.Println(fname)
}
