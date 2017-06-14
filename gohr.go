package gohr

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// getWidth gets number of width of terminal from crypto subdirectory ssh/terminal
func getTerminalWidth() (int, error) {
	w, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return -1, err
	}
	return w, nil
}

// Draw fills a row with '#' by default (if no arguments are provided) or takes arguments and prints each pattern on a new line.
func Draw(patterns ...string) {
	w, err := getTerminalWidth()
	if err != nil {
		log.Fatalf("Error getting terminal width: %s\n", err)
	}

	if len(patterns) == 0 {
		for i := 0; i < w; i++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	} else {
		for _, pattern := range patterns {
			l := len(pattern)
			for i := 0; i < w/l; i++ {
				fmt.Printf(pattern)
			}
			// Fills up the remaining columns in the row with part of the pattern
			fmt.Printf("%s\n", pattern[:w%l])
		}
	}
}
