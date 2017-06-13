package main

import (
	"os"

	"github.com/hasit/gohr"
)

func main() {
	args := os.Args[1:]
	gohr.Draw(args...)
}
