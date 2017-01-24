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

//Prints each key file name to the terminal
func PrintKeyFile(n string) {
	fmt.Println(n)
}

func NoKeysExist(){
	fmt.Println("No keys were found on your device")
}
