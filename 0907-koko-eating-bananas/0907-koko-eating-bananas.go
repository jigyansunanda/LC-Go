func isPossible(piles []int, h int, x int) bool {
	totalHours := 0
	for _, val := range piles {
		if val%x == 0 {
			totalHours += val / x
		} else {
			totalHours += (val / x) + 1
		}
	}
	return (totalHours <= h)
}

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 0, piles[0]
	for _, val := range piles {
		hi = max(hi, val)
	}
	for hi > lo+1 {
		mid := lo + (hi-lo)/2
		if isPossible(piles, h, mid) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return hi
}