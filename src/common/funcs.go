package common

import "strings"

func IsWordWholeMatched(line string, word string, index int) bool {
    leftIndex := index - 1
	rightIndex := index + len(word)
	if leftIndex >= 0 && !strings.ContainsRune(SPLITE_CHARACTORS, rune(line[leftIndex])) { return false }
	if rightIndex < len(line) && !strings.ContainsRune(SPLITE_CHARACTORS, rune(line[rightIndex])) { return false }
	return true
}