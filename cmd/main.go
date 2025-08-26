package main

import (
	"fmt"
	"os"
	"stx/internal/cli"
	"stx/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		utils.PrintUsage()
		os.Exit(1)
	}

	app := cli.New()
	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
