// Package diffsquares contains utilities that compute squares and sums
package diffsquares

// SquareOfSum computes the square of sum of first n natural numbers
func SquareOfSum(n int) (sum int) {
	sum = (((n + 1) * n) / 2)
	sum *= sum
	return
}

// SumOfSquares computes the sum of squares of first n natural numbers
func SumOfSquares(n int) (sum int) {
	sum = (n * (n + 1) * (2*n + 1)) / 6
	return
}

// Difference calculates SquareOfSum(n) - SumOfSquares(n)
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
