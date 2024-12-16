func isPossible(candies []int, x int, k int64) bool {
	var total int64 = 0
	for _, val := range candies {
		total += int64(val / x)
	}
	return (total >= k)
}

func maximumCandies(candies []int, k int64) int {
	lo, hi := 0, 0
	for _, val := range candies {
		hi = max(val, hi)
	}
    hi += 1
	for hi > lo+1 {
		mid := lo + (hi-lo)/2
		if isPossible(candies, mid, k) {
			lo = mid
		} else {
			hi = mid
		}
	}
	return lo
}