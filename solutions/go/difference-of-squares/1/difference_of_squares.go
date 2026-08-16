package differenceofsquares

func exponent(base int, power int) int {
	result := 1
	for range power {
		result *= base
	}
	return base
}

func SquareOfSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n
		n--
	}
	return exponent(sum, 2)
}

func SumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		sum += exponent(n, 2)
		n--
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
