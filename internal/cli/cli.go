package cli

import (
	"fmt"
	"strings"
	"stx/internal/analyzer"
	"stx/internal/output"
	"stx/internal/utils"
)

type App struct {
	countLines bool
	countWords bool
	countChars bool
	showHelp   bool
	files      []string
}

func New() *App {
	return &App{}
}

func (a *App) Run(args []string) error {
	if err := a.parseArgs(args); err != nil {
		return err
	}

	if a.showHelp {
		utils.PrintUsage()
		return nil
	}

	if !a.countLines && !a.countChars && !a.countWords {
		a.countLines = true
		a.countChars = true
		a.countWords = true
	}

	results := make([]analyzer.FileStats, 0, len(a.files))
	var totalStats analyzer.FileStats

	for _, file := range a.files {
		stats, err := analyzer.AnalyzeFile(file)
		if err != nil {
			return fmt.Errorf("failed to analyze %s: %v", file, err)
		}

		results = append(results, *stats)
		totalStats.Lines += stats.Lines
		totalStats.Words += stats.Words
		totalStats.Chars += stats.Chars
	}

	formatter := output.New(a.countLines, a.countWords, a.countChars)
	formatter.Print(results, totalStats, len(a.files) > 1)

	return nil
}

func (a *App) parseArgs(args []string) error {
	for i, arg := range args {
		if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			for _, flag := range arg[1:] {
				switch flag {
				case 'l':
					a.countLines = true
				case 'w':
					a.countWords = true
				case 'c':
					a.countChars = true
				case 'h':
					a.showHelp = true
				default:
					utils.PrintUsage()
					return fmt.Errorf("unknown flag: -%c", flag)
				}
			}
		} else {
			a.files = args[i:]
			break
		}
	}

	if a.showHelp {
		return nil
	}

	if len(a.files) == 0 {
		return fmt.Errorf("no input files specified")
	}

	return nil
}
