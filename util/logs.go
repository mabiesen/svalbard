package util

import (
	"fmt"
)

func Error(err error) {
	fmt.Println("ERROR: ",err)
}

func Notice(n string) {
	fmt.Println("NOTICE: ",n)
}

func Print(s string) {
	fmt.Println(s)
}
