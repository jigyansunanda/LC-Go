func trailingZeroes(n int) int {
	totalZeroes := 0
	denominator := 5
	for n/denominator != 0 {
		totalZeroes += n / denominator
		denominator *= 5
	}
	return totalZeroes
}