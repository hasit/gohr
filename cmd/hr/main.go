package main

import (
	"os"

	"github.com/hasit/gohr"
)

func getArgs() []string {
	return os.Args[1:]
}

func main() {
	args := getArgs()
	gohr.DrawHr(args)
}
