package output

import (
	"fmt"
	"path/filepath"
	"strings"
	"stx/internal/analyzer"
)

const (
	filenameWidth = 20
	columnWidth   = 10
)

type Formatter struct {
	showLines bool
	showWords bool
	showChars bool
}

func New(lines, words, chars bool) *Formatter {
	return &Formatter{
		showLines: lines,
		showWords: words,
		showChars: chars,
	}
}

func (f *Formatter) Print(results []analyzer.FileStats, total analyzer.FileStats, showTotal bool) {
	fmt.Printf("%-*s", filenameWidth, "File")

	if f.showLines {
		fmt.Printf("%*s", columnWidth, "Lines")
	}
	if f.showWords {
		fmt.Printf("%*s", columnWidth, "Words")
	}
	if f.showChars {
		fmt.Printf("%*s", columnWidth, "Chars")
	}
	fmt.Println()

	f.printSeparator()

	for _, result := range results {
		filename := filepath.Base(result.Filename)
		if len(filename) > 15 {
			filename = filename[:12] + "..."
		}
		fmt.Printf("%-*s", filenameWidth, filename)

		if f.showLines {
			fmt.Printf("%*d", columnWidth, result.Lines)
		}
		if f.showWords {
			fmt.Printf("%*d", columnWidth, result.Words)
		}
		if f.showChars {
			fmt.Printf("%*d", columnWidth, result.Chars)
		}
		fmt.Println()
	}

	if showTotal {
		f.printSeparator()
		fmt.Printf("%-*s", filenameWidth, "TOTAL")

		if f.showLines {
			fmt.Printf("%*d", columnWidth, total.Lines)
		}
		if f.showWords {
			fmt.Printf("%*d", columnWidth, total.Words)
		}
		if f.showChars {
			fmt.Printf("%*d", columnWidth, total.Chars)
		}
		fmt.Println()
	}
}

func (f *Formatter) printSeparator() {
	fmt.Print(strings.Repeat("-", filenameWidth))

	if f.showLines {
		fmt.Print(strings.Repeat("-", columnWidth))
	}
	if f.showWords {
		fmt.Print(strings.Repeat("-", columnWidth))
	}
	if f.showChars {
		fmt.Print(strings.Repeat("-", columnWidth))
	}
	fmt.Println()
}
