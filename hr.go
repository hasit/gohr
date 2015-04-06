package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func getCols() int {
	c, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}
	return c
}

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
			fmt.Printf("%v\n", arg[:cols%l])
		}
	}
}

func getArgs() {
	args := os.Args[1:]
	drawHr(args)
}

func main() {
	getArgs()
}
