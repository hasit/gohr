package gohr

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// getWidth gets number of width of terminal from crypto subdirectory ssh/terminal
func getWidth() (int, error) {
	w, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return -1, err
	}
	return w, nil
}

func drawPatterns(width int, patterns []string, wr io.Writer) {
	if len(patterns) == 0 {
		for i := 0; i < width; i++ {
			fmt.Fprintf(wr, "#")
		}
		fmt.Fprintf(wr, "\n")
	} else {
		for _, pattern := range patterns {
			patternLength := len(pattern)
			for i := 0; i < width/patternLength; i++ {
				fmt.Fprintf(wr, pattern)
			}
			// Fills up the remaining columns in the row with part of the pattern
			fmt.Fprintf(wr, "%s\n", pattern[:width%patternLength])
		}
	}
}

// Draw fills a row with '#' by default (if no arguments are provided) or takes arguments and prints each pattern on a new line.
func Draw(patterns ...string) {
	w, err := getWidth()
	if err != nil {
		log.Fatalf("Error getting terminal width: %s\n", err)
	}

	var p []string

	for _, pattern := range patterns {
		p = append(p, pattern)
	}

	drawPatterns(w, p, os.Stdout)
}
