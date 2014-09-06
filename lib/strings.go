// Created to test using different packages.
package lib

import (
	"unicode/utf8"
)


////////////////////////////////////////////////////////////////////////


// Function to reverse a string.
func Reverse(s string) string {
	cs := make([]rune, utf8.RuneCountInString(s))
	i := len(cs)
	for _, c := range s {
		i--
		cs[i] = c
	}
	return string(cs)
}


////////////////////////////////////////////////////////////////////////