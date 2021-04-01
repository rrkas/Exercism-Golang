package luhn

import (
	"regexp"
	"strconv"
)

func Valid(text string) bool {
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, "")
	if regexp.MustCompile(`[\D]`).MatchString(text) {
		return false
	}
	if len(text) <= 1 {
		return false
	}
	s := 0
	for i, d := range text {
		di, err := strconv.Atoi(string(d))
		if err != nil {
			continue
		}
		if len(text)%2 == 0 {
			if i%2 == 0 {
				di = di * 2
				if di > 9 {
					di = di - 9
				}
			}
		} else {
			if i%2 == 1 {
				di = di * 2
				if di > 9 {
					di = di - 9
				}
			}
		}
		s = s + di
	}
	return s%10 == 0
}
