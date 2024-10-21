func climbStairs(n int) int {
	cache := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i <= 1 {
			cache[i] = 1
		} else {
			cache[i] = cache[i-1] + cache[i-2]
		}
	}
	return cache[n]
}