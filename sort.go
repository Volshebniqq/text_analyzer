package textAnalyzer

import (
	"sort"
	"sync"
)

type KV struct {
	Key   string
	Value int
}

func sortWordsResult(m sync.Map)[]KV {

	var ss []KV

	m.Range(func(k interface{}, v interface{}) bool {
		ss = append(ss, KV{k.(string), v.(int)})
		return true
	})

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}

func sortCharsResult(m sync.Map)[]KV {

	var ss []KV

	m.Range(func(k interface{}, v interface{}) bool {
		ss = append(ss, KV{string(k.(uint8)), v.(int)})
		return true
	})

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}
