package scrabble

import "strings"

func Score(text string) int {
	s := 0
	for i := 0; i < len(text); i++ {
		switch strings.ToUpper(string(text[i])) {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			s = s + 1
		case "D", "G":
			s = s + 2
		case "B", "C", "M", "P":
			s = s + 3
		case "F", "H", "V", "W", "Y":
			s = s + 4
		case "K":
			s = s + 5
		case "J", "X":
			s = s + 8
		case "Q", "Z":
			s = s + 10
		}
	}
	return s
}
