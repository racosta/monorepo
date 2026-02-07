// Package iteration provides functions for repeating strings.
package iteration

import "strings"

const repeatCount = 5

// Repeat returns the given string repeated a fixed number of times.
func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
