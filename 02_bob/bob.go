// The bob package handles forming responses for Alice
package bob

import (
	"regexp"
	"strings"
)

// Hey takes in a remark and returns a response to the remark as a string
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	// If there is nothing left after trimming whitespace, the string is blank
	if remark == "" {
		return "Fine. Be that way!"
	}

	letterRegex, _ := regexp.Compile(`[a-zA-Z]`)

	if remark == strings.ToUpper(remark) {
		// An all uppercase question, note '42?' will land here
		if remark[len(remark)-1:] == "?" {
			// Check if there are letters in the question
			if len(letterRegex.FindString(remark)) > 0 {
				return "Calm down, I know what I'm doing!"
			} else {
				return "Sure."
			}
		}

		// Check if there are letters or if it's e.g. '1,2,3'
		if len(letterRegex.FindString(remark)) < 1 {
			return "Whatever."
		} else {
			return "Whoa, chill out!"
		}
	}

	if remark[len(remark)-1:] == "?" {
		return "Sure."
	}

	return "Whatever."
}
