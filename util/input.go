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
		var answer string

		for{
			theInput := ScanForInput()
			if theInput == affirmative || theInput == negative {
				answer = theInput
				break
			}
			fmt.Println("That is not a valid input.  Please type ",affirmative," or ", negative, " to continue.")
		}

		return answer

}

// function scans for user inputs.  Formats and returns the input
func ScanForInput() string{
	fmt.Println("Waiting for input...")
	reader1 := bufio.NewReader(os.Stdin)
	userInput, _ := reader1.ReadString('\n')
	return strings.TrimSpace(userInput)
}

//Prints each key file name to the terminal
func PrintFileName(n string) {
	fname := strings.TrimSuffix(n, filepath.Ext(n))
	Print(fname)
}
