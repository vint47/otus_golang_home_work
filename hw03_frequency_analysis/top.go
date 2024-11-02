package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	resMap := make(map[string]int, 0)

	sl := strings.Fields(s)

	for _, word := range sl {
		resMap[word]++
	}

	type frequency struct {
		word  string
		count int
	}

	fr := make([]frequency, 0, len(resMap))
	for k, v := range resMap {
		fr = append(fr, frequency{k, v})
	}

	sort.Slice(fr, func(i, j int) bool {
		if fr[i].count == fr[j].count {
			return strings.Compare(fr[i].word, fr[j].word) < 0
		}
		return fr[i].count > fr[j].count
	})

	res := make([]string, 0)
	for i := 0; i < min(10, len(fr)); i++ {
		res = append(res, fr[i].word)
	}

	return res
}
