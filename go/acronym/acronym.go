// Package acronym implements converting a string to its abbreviation
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns the abbreviated form of s
func Abbreviate(s string) string {
	s = strings.ToUpper(s)

	f := func(c rune) bool {
		return unicode.IsSpace(c) || c == 0x2d
	}

	splits := strings.FieldsFunc(s, f)
	var sb strings.Builder

	for _, word := range splits {
		for _, chr := range word {
			if unicode.IsUpper(chr) {
				sb.WriteRune(chr)
				break
			}
		}
	}
	return sb.String()
}
