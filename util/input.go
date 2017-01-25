package util

import(
	"strings"
	"path/filepath"
	"bufio"
	"os"
	"fmt"
)

//Container function to wait for user to make one of two selections
func RecurCompareInput(affirmative string , negative string)string{
		var myreturn string
		reader1 := bufio.NewReader(os.Stdin)
		userInput, _ := reader1.ReadString('\n')
		theInput := strings.TrimSpace(userInput)
		switch theInput{
			case affirmative: Print("The original file will be deleted.")
				myreturn = affirmative
			case negative: Print("The original file will NOT be deleted.")
				myreturn = negative
			default:
				fmt.Println("That is not a valid input.  Please type ",affirmative," or ", negative, " to continue.")
				RecurCompareInput(affirmative, negative)
		}
		return myreturn
}

//Prints each key file name to the terminal
func PrintFileName(n string) {
	fname := strings.TrimSuffix(n, filepath.Ext(n))
	Print(fname)
}
