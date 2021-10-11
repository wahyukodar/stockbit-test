package soal4

import (
	"sort"
	"strings"
)

func sortString(word string) string {
	stringSlice := strings.Split(word, "")

	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i] < stringSlice[j]
	})

	return strings.Join(stringSlice, "")
}

func checkAndListAnagram(words []string, result [][]string) [][]string {
	var anagram []string
	var currentWords []string

	if len(words) <= 0 {
		return result
	}

	var word1 = sortString(words[0])
	anagram = append(anagram, words[0])
	for i := 1; i < len(words); i++ {
		var word2 = sortString(words[i])

		if word1 == word2 {
			anagram = append(anagram, words[i])
		} else {
			currentWords = append(currentWords, words[i])
		}
	}
	result = append(result, anagram)

	return checkAndListAnagram(currentWords, result)
}
