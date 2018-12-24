package textAnalyzer

func averageWordLength(result []KV) float64 {
	wordsCount := 0
	lettersCount := 0
	for _, kv := range result {
		wordsCount += kv.Value
		lettersCount += kv.Value * len(kv.Key)
	}
	var averageLength float64 = float64(lettersCount) / float64(wordsCount)
	return averageLength
}
