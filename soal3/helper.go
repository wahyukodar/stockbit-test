package soal3

import "strings"

func findFirstStringInBracket(str string) string {
	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return ""
	}

	runes := []rune(str)
	wordsAfterFirstBracket := string(runes[indexFirstBracketFound+1 : len(str)])
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound < 0 {
		return ""
	}
	runes = []rune(wordsAfterFirstBracket)
	return string(runes[0:indexClosingBracketFound])
}