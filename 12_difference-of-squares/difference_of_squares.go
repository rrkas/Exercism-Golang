package diffsquares

func SumOfSquares(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s = s + i*i
	}
	return s
}

func SquareOfSum(n int) int {
	s := (n * (n + 1)) / 2
	s = s * s
	return s
}

func Difference(n int) int {
	s := SquareOfSum(n) - SumOfSquares(n)
	if s < 0 {
		return s * -1
	}
	return s
}
