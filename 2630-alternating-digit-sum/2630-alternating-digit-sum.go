func alternateDigitSum(n int) int {
	alternateDigitSum := 0
	digitCount := 0
	isAlternate := true
	for n > 0 {
		remainder := n % 10
		if isAlternate {
			alternateDigitSum += remainder
		} else {
			alternateDigitSum -= remainder
		}
		isAlternate = !isAlternate
		digitCount++
		n /= 10
	}
	if digitCount%2 == 0 {
		alternateDigitSum = -alternateDigitSum
	}
	return alternateDigitSum
}