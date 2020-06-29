package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains in cell n
func Square(n int) (uint64, error) {
	if n <= 64 && n > 0 {
		return uint64(math.Pow(2, float64(n-1))), nil
	} else {
		return 0, errors.New("n should be less than 65")
	}
}

// Total returns the total number of grains on the board
func Total() uint64 {
	return 18446744073709551615
}
