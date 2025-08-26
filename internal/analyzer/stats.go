package analyzer

import (
	"os"
	"strings"
	"unicode"
)

type FileStats struct {
	Filename string
	Lines    int
	Words    int
	Chars    int
}

func AnalyzeFile(filename string) (*FileStats, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var stats FileStats
	stats.Filename = filename
	stats.Chars = len(content)

	contentStr := string(content)

	if len(contentStr) == 0 {
		stats.Lines = 0
	} else {
		stats.Lines = strings.Count(contentStr, "\n")
	}

	if !strings.HasSuffix(contentStr, "\n") {
		stats.Lines++
	}

	stats.Words = countWords(contentStr)

	return &stats, nil
}

func countWords(txt string) int {
	if len(strings.TrimSpace(txt)) == 0 {
		return 0
	}

	wordCount := 0
	inWord := false

	for _, char := range txt {
		if unicode.IsSpace(char) {
			if inWord {
				wordCount++
				inWord = false
			}
		} else {
			inWord = true
		}
	}

	if inWord {
		wordCount++
	}

	return wordCount
}
