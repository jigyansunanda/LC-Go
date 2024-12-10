func isPossible(weights []int, days int, x int) bool {
	current_total, day_count := 0, 0
	for _, weight := range weights {
		if weight > x {
			return false
		}
		if current_total+weight < x {
			current_total += weight
			continue
		} else if current_total+weight == x {
			day_count += 1
			current_total = 0
		} else {
			day_count += 1
			current_total = weight
		}
	}
	if current_total > 0 {
		day_count += 1
	}
	return days >= day_count
}

func shipWithinDays(weights []int, days int) int {
	lo, hi := 0, 100000000
	for hi > lo+1 {
		mid := lo + (hi-lo)/2
		if isPossible(weights, days, mid) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return hi
}