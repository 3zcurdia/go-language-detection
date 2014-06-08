package goLanguageDetection

import (
	"bufio"
	"github.com/willf/bloom"
	"os"
)

// CountOccurences counts the number words belonging to the given language and
// returns the value.
//
func CountOccurences(language string, words []string, messages chan WordsCount) {
	file, _ := os.Open(os.Getenv("GOPATH") + "/src/github.com/AntoineFinkelstein/go-language-detection/wordlists/" + language)

	filter := bloom.NewWithEstimates(700000, 0.05)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		filter.Add([]byte(scanner.Text()))
	}

	file.Close()

	result := 0
	for _, word := range words {

		if filter.Test([]byte(word)) {
			result++
		}
	}
	messages <- WordsCount{language: language, count: result}
}
