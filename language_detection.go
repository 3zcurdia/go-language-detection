package goLanguageDetection

import (
	"github.com/willf/bloom"
	"regexp"
	"strings"
)

type Detect struct {
	languages map[string]*bloom.BloomFilter
}

func New() *Detect {
	detect := Detect{}

	detect.languages = map[string]*bloom.BloomFilter{
		"danish":     DefineFilter("danish"),
		"dutch":      DefineFilter("dutch"),
		"english":    DefineFilter("english"),
		"farsi":      DefineFilter("farsi"),
		"french":     DefineFilter("french"),
		"german":     DefineFilter("german"),
		"italian":    DefineFilter("italian"),
		"pinyin":     DefineFilter("pinyin"),
		"portuguese": DefineFilter("portuguese"),
		"russian":    DefineFilter("russian"),
		"spanish":    DefineFilter("spanish"),
		"swedish":    DefineFilter("swedish"),
	}

	return &detect
}

func (d *Detect) Text(input string) (string, float64) {

	results := make(map[string]int)
	messages := make(chan WordsCount)

	words := CleanInput(input)
	for language, filter := range d.languages {
		go CountOccurences(language, filter, words, messages)
	}

	for _, _ = range d.languages {
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

type WordsCount struct {
	language string
	count    int
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
