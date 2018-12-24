package textAnalyzer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func write(output string, wordsCount int, averageWordLength float64, charsResult []KV, wordsResult []KV) {
	f, err := os.Create(output)
	check(err)

	w := bufio.NewWriter(f)
	avg := fmt.Sprintf("%f", averageWordLength)
	words := strconv.Itoa(wordsCount)
	w.WriteString("Words count: " + words + "\n")
	w.WriteString("Average word length: " + avg + "\n\n")

	w.WriteString("Chars statistic:\n")
	for _, kv := range charsResult {
		value := strconv.Itoa(kv.Value)
		w.WriteString(kv.Key + " - " + value + "\n")
	}
	w.WriteString("\n\n Words statistic: \n")
	for _, kv := range wordsResult {
		value := strconv.Itoa(kv.Value)
		w.WriteString(kv.Key + " - " + value + "\n")
	}

	w.Flush()

}
