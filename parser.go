package textAnalyzer

import (
	"regexp"
	"strings"
	"sync"
)

func parseWordsCount(wg *sync.WaitGroup, words []string, result *sync.Map, start int, finish int) {
	for i := start; i < finish; i++ {
		word := strings.ToLower(words[i])
		value, exists := result.Load(word)
		if exists {
			result.Store(word, value.(int) + 1)
		} else {
			result.Store(word, 1)
		}
	}
	defer wg.Done()
}

func parseCharsCount(wg *sync.WaitGroup, data string, result *sync.Map, start int, finish int) {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	for i := start; i < finish; i++ {
		char := data[i]
		if string(char) == " " || !isAlpha(string(char)) { continue }
		value, exists := result.Load(char)
		if exists {
			result.Store(char, value.(int) + 1)
		} else {
			result.Store(char, 1)
		}
	}
	defer wg.Done()
}