package luhn

import (
	"strings"
	"unicode"
)

// Valid returns true if the input is valid a/c to Luhn's algorithm
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")

	if len(s) <= 1 {
		return false
	}

	for _, chr := range s {
		if !unicode.IsDigit(chr) {
			return false
		}
	}

	sum := 0
	for i, j := len(s)-1, 0; j < len(s); i, j = i-1, j+1 {
		if j%2 == 0 {
			sum += int(s[i] - '0')
		} else {
			dbl := int(s[i]-'0') * 2
			if dbl > 9 {
				dbl -= 9
			}
			sum += dbl
		}
	}

	return sum%10 == 0
}
