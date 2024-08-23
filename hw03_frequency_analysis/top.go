package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	words := strings.Fields(str)
	mapWords := map[string]int{}
	for _, el := range words {
		if el != "" {
			mapWords[el]++
		}
	}
	result := make([]string, 0, len(mapWords))
	for key := range mapWords {
		result = append(result, key)
	}
	sort.Slice(result, func(i, j int) bool {
		if mapWords[result[i]] == mapWords[result[j]] {
			return result[i] < result[j]
		}
		return mapWords[result[i]] > mapWords[result[j]]
	})

	if len(result) > 9 {
		return result[:10]
	}
	return result
}
