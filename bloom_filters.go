package goLanguageDetection

import (
	"github.com/willf/bloom"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"
)

// FileExist search for the existence of a bloom filter for the given language.
//
func FileExist(language string) bool {
	usr, _ := user.Current()
	if _, err := os.Stat(usr.HomeDir + "/tmp/" + language); os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateBloomFilter search for the existence of a bloom filter for the given
// language.
//
func CreateBloomFilter(language string) {
	filter := bloom.NewWithEstimates(700000, 0.05)
	resp, _ := http.Get("https://raw.githubusercontent.com/AntoineFinkelstein/go-language-detection/master/wordlists/" + language)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	for _, element := range strings.Split(string(body), "\n") {
		word := strings.ToLower(element)
		filter.Add([]byte(word))
	}

	content, _ := filter.MarshalJSON()
	usr, _ := user.Current()
	err = ioutil.WriteFile(usr.HomeDir+"/tmp/"+language, content, 0600)
	if err != nil {
		panic(err)
	}
}

// CountOccurences counts the number words belonging to the given language and
// returns the value.
//
func CountOccurences(language string, filter *bloom.BloomFilter, words []string, messages chan WordsCount) {
	result := 0
	for _, word := range words {

		if filter.Test([]byte(word)) {
			result++
		}
	}
	messages <- WordsCount{language: language, count: result}
}

func DefineFilter(language string) *bloom.BloomFilter {
	// If the bloomfilters don't exist. They're created.
	if FileExist(language) == false {
		CreateBloomFilter(language)
	}

	filter := bloom.NewWithEstimates(700000, 0.05)
	usr, _ := user.Current()
	file, _ := ioutil.ReadFile(usr.HomeDir + "/tmp/" + language)
	_ = filter.UnmarshalJSON(file)
	return filter
}
