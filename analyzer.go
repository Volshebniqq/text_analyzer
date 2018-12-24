package textAnalyzer

import (
	"io/ioutil"
	"strings"
	"sync"
)

type Result struct {
	wordsCount int
	averageWordLength float64
	wordsStatistic []KV
	charsStatistic []KV
}


func Analyze(input string, output string, wordPerRoutine int, charsPerRoutine int) Result {
	dat, err := ioutil.ReadFile(input)
	check(err)
	data := removeSpecialChars(string(dat))

	var wg sync.WaitGroup

	charsResult := sync.Map{}
	routinesCount := len(data) / charsPerRoutine

	for i := 0; i <= routinesCount; i++ {
		wg.Add(1)
		if i == routinesCount {
			go parseCharsCount(&wg, data, &charsResult, charsPerRoutine * i, routinesCount)
		} else {
			go parseCharsCount(&wg, data, &charsResult, charsPerRoutine * i, charsPerRoutine * (i + 1) - 1)
		}
	}


	words := strings.Fields(data)
	wordsResult := sync.Map{}
	wordsCount := len(words)
	routinesCount = wordsCount / wordPerRoutine


	for i := 0; i <= routinesCount; i++ {
		wg.Add(1)
		if i == routinesCount {
			go parseWordsCount(&wg, words, &wordsResult, wordPerRoutine * i, wordsCount)
		} else {
			go parseWordsCount(&wg, words, &wordsResult, wordPerRoutine * i,wordPerRoutine * (i + 1) - 1)
		}

	}

	wg.Wait()

	sortedChars := sortCharsResult(charsResult)
	sortedWords := sortWordsResult(wordsResult)
	averageLength := averageWordLength(sortedWords)

	if output != "" {
		write(output, wordsCount, averageLength, sortedChars, sortedWords)
	}

	result := Result{
		wordsCount: wordsCount,
		averageWordLength: averageLength,
		wordsStatistic: sortedWords,
		charsStatistic: sortedChars}
	return result
}