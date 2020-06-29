package isogram

import "strings"

// IsIsogram returns true if word is isogram
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	var count [128]int

	for _, chr := range word {
		if chr >= 'a' && chr <= 'z' {
			if count[chr] > 0 {
				return false
			}
			count[chr]++
		}
	}
	return true
}
