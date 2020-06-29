package hamming

import "errors"

// Distance calculates hamming distance between 2 DNA strands
func Distance(a, b string) (int, error) {
	la := len(a)
	lb := len(b)
	if lb != la {
		return 0, errors.New("need same length")
	}
	delta := 0
	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			delta++
		}
	}
	return delta, nil
}
