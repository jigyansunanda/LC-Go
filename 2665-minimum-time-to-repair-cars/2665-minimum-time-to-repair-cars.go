func isPossible(ranks []int, x int64, cars int64) bool {
	totalCars := int64(0)
	for _, val := range ranks {
		squared := x / int64(val)
		root := int64(math.Sqrt(float64(squared)))
		totalCars += root
	}
	return (totalCars >= cars)
}

func repairCars(ranks []int, cars int) int64 {
	var c int64 = int64(cars)
	var lo int64 = 0
	var hi int64 = c * c * 101
	for hi > lo+1 {
		mid := lo + (hi-lo)/2
		if isPossible(ranks, mid, c) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return hi
}