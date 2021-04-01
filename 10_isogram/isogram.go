package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(text string) bool {

	for i, c := range text {
		if strings.Contains(strings.ToUpper(string(text[i+1:])), strings.ToUpper(string(c))) && unicode.IsLetter(c) {
			return false
		}
	}
	return true
}
