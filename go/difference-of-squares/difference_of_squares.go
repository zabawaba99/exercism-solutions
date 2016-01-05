package diffsquares

func SquareOfSums(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares sums the squares of the first n natural numbers
// Shorthand sum taken from:
// http://www.trans4mind.com/personal_development/mathematics/series/sumNaturalSquares.htm
func SumOfSquares(n int) int {
	sum := n * (n + 1) * (2*n + 1) / 6
	return sum
}

// Difference takes the difference between the SquareOfSums and SumofSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
