package goLanguageDetection

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type WordsCount struct {
	language string
	count    int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// StringInSlice returns if the string is past of the given list.
//
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// CleanInput takes the request string, removes unwanted values and returns an
// Array of words.
//
func CleanInput(input string) []string {
	var words []string

	for _, element := range strings.Split(input, " ") {
		word := strings.ToLower(element)

		// TODO : Words such as "allez-vous" will be transformed to allezvous which
		// isn't correct. A better Regexp must be found.
		r := regexp.MustCompile(`\W`)
		word = r.ReplaceAllString(word, "")

		if word != "" {
			words = append(words, word)
		}
	}

	return words
}

// CountOccurences counts the number words belonging to the given language and
// returns the value.
//
func CountOccurences(language string, words []string, messages chan WordsCount) {
	file, err := os.Open(os.Getenv("GOPATH") + "/src/github.com/AntoineFinkelstein/go-language-detection/wordlists/" + language)
	defer file.Close()
	check(err)

	var wordLists []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordLists = append(wordLists, scanner.Text())
	}

	result := 0
	for _, word := range words {
		if StringInSlice(word, wordLists) {
			result++
		}
	}
	messages <- WordsCount{language: language, count: result}
}

// Find returns the language of the text as well as the percentage of words that
// were found.
//
func Find(input string) (string, float64) {
	languages := []string{"danish", "dutch", "english", "farsi", "french", "german", "italian", "pinyin", "portuguese", "russian", "spanish", "swedish"}
	results := make(map[string]int)
	messages := make(chan WordsCount)

	words := CleanInput(input)
	for _, language := range languages {
		go CountOccurences(language, words, messages)
	}

	for _ = range languages {
		wordsCount := <-messages
		results[wordsCount.language] = wordsCount.count
	}

	bestMatch := WordsCount{language: "none", count: 0}
	for language, count := range results {
		if count > bestMatch.count {
			bestMatch = WordsCount{language: language, count: count}
		}
	}

	return bestMatch.language, float64(bestMatch.count) / float64(len(words))
}
