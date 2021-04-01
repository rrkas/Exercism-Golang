// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	words := strings.Fields(s)
	a := ""
	for _, w := range words {
		if w[0] == '-' || w[0] == '_' {
			if len(w) != 1 {
				a = a + string(w[1])
			}
		} else {
			if strings.Contains(w, "-") {
				for _, sw := range strings.Split(w, "-") {
					a = a + string(sw[0])
				}
			} else {
				a = a + string(w[0])
			}
		}
	}
	return strings.ToUpper(a)
}
