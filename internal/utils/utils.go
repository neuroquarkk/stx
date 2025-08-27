package utils

import (
	"fmt"
	"os"
)

func PrintUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] <file1> [file2...]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Options:\n")
	fmt.Fprintf(os.Stderr, "  -l      count lines\n")
	fmt.Fprintf(os.Stderr, "  -w      count words\n")
	fmt.Fprintf(os.Stderr, "  -c      count characters\n")
	fmt.Fprintf(os.Stderr, "  -u      count unique words\n")
	fmt.Fprintf(os.Stderr, "  -h      show detailed help\n")
	fmt.Fprintf(os.Stderr, "  -lwc    count lines, words, and characters (flags can be combined)\n")
	fmt.Fprintf(os.Stderr, "  (no flags = count all)\n")
}
