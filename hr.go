//make it available as a command
package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

//get number of columns of terminal from crypto subdirectory ssh/terminal
func getCols() int {
	c, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}
	return c
}

//fill a row with '#' by default (if no arguments are provided) or take command line arguments and print each pattern on a new line
func drawHr(args []string) {
	cols := getCols()

	if len(args) == 0 {
		for i := 0; i < cols; i++ {
			fmt.Printf("#")
		}
		fmt.Println()
	} else {
		for _, arg := range args {
			l := len(arg)
			for i := 0; i < cols/l; i++ {
				fmt.Printf(arg)
			}
			fmt.Printf("%v\n", arg[:cols%l]) //this fills ups the remaining columns in the row with part of the pattern
		}
	}
}

//get command line arguments from 1 onwards (exclude the file name)
func getArgs() {
	args := os.Args[1:]
	drawHr(args)
}

func main() {
	getArgs()
}
