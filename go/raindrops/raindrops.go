package raindrops

import (
	"strconv"
	"strings"
)

// Convert returns raindrop version of input
func Convert(input int) string {
	factor := false

	var sb strings.Builder

	if input%3 == 0 {
		sb.WriteString("Pling")
		factor = true
	}
	if input%5 == 0 {
		sb.WriteString("Plang")
		factor = true
	}
	if input%7 == 0 {
		sb.WriteString("Plong")
		factor = true
	}

	if !factor {
		sb.WriteString(strconv.Itoa(input))
	}

	return sb.String()
}
