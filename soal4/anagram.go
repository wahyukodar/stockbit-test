package soal4

import (
	"sort"
)

func checkAnagram(words []string) [][]string {
	temp := map[string][]string{}

	for _, word := range words {
		r := []rune(word)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		sortedWord := string(r)

		temp[sortedWord] = append(temp[sortedWord], word)
	}

	var result [][]string
	for _, item := range temp {
		result = append(result, item)
	}
	return result
}
