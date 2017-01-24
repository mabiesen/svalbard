package util

import (
	"fmt"
	"strings"
	"path/filepath"
	"bufio"
	"os"
)

func Error(err error) {
	fmt.Println("ERROR: ",err)
}

func Notice(n string) {
	fmt.Println("NOTICE: ",n)
}

//Prints each key file name to the terminal
func PrintFileName(n string) {
	fname := strings.TrimSuffix(n, filepath.Ext(n))
	fmt.Println(fname)
}

//@func - recurCompareInput
//@desc - Container function to wait for user to make one of two selections
//Note:  combined scan and switch functions to create this comprehensive func
//@params 1. affirmative String
//@affirmative is the string user must match to create a 'yes' response
//@params 2. negative String
//@negative is the string user must match to create a 'no' response
//@return - none
//@type - none
//@pkg - "strings", "fmt", "bufio", "os"
func RecurCompareInput(affirmative string , negative string)string{
		var myreturn string
		fmt.Println("Waiting for input...")
		reader1 := bufio.NewReader(os.Stdin)
		userInput, _ := reader1.ReadString('\n')
		theInput := strings.TrimSpace(userInput)
		switch theInput{
			case affirmative: fmt.Println("The original file will be deleted.")
				myreturn = affirmative
			case negative: fmt.Println("The original file will NOT be deleted.")
				myreturn = negative
			default:
				fmt.Println("That is not a valid input.  Please type ",affirmative," or ", negative, " to continue.")
				RecurCompareInput(affirmative, negative)
		}
		return myreturn
}
