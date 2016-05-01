package gohr

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

// DrawHr fills a row with '#' by default (if no arguments are provided) or take command line arguments and print each pattern on a new line
func DrawHr(args ...string) {
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
			// Fill ups the remaining columns in the row with part of the pattern
			fmt.Printf("%v\n", arg[:cols%l])
		}
	}
}
