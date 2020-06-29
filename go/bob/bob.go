// Package bob implements response of Bob for a remark
package bob

import (
	"strings"
)

// Hey returns a response from bob based on remark
func Hey(remark string) (response string) {

	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	capCount := 0
	for _, chr := range remark {
		if chr >= 0x41 && chr <= 0x5A {
			capCount++
		}
	}

	allCaps := strings.Compare(remark, strings.ToUpper(remark)) == 0 && capCount > 0
	question := remark[len(remark)-1] == 0x3f

	response = "Whatever."
	if question && allCaps {
		response = "Calm down, I know what I'm doing!"
	} else if question && !allCaps {
		response = "Sure."
	} else if !question && allCaps {
		response = "Whoa, chill out!"
	}
	return response
}
