package textAnalyzer

import "strings"

func removeSpecialChars(data string) string {
	var chars = []string{"/", "\"", "\\", "&", "#", "(", "�", ")", "_", "'",  "[", "]", ".", ",", "!", ";", ":", "?", "-", "=", "+", "«", "»", "—"}
	for _, char := range chars {
		data = strings.Replace(data, char, " ", -1)
	}
	return data
}
