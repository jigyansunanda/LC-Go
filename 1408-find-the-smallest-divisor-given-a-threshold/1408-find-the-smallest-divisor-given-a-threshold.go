func checkDivisorSum(nums []int, x int, threshold int) bool {
	divisorSum := 0
	for _, val := range nums {
		if val%x == 0 {
			divisorSum += (val / x)
		} else {
			divisorSum += (val / x) + 1
		}
	}
	return divisorSum <= threshold
}

func smallestDivisor(nums []int, threshold int) int {
	lo, hi := 0, 0
	for _, val := range nums {
		hi = max(hi, val)
	}
	for hi > lo+1 {
		mid := lo + (hi-lo)/2
		if checkDivisorSum(nums, mid, threshold) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return hi
}