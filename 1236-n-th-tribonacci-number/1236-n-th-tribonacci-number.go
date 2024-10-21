func tribonacci(n int) int {
	if n == 0 {
		return n
	} else if n <= 2 {
		return 1
	} else {
		cache := make([]int, n+1)
		for i := 1; i <= n; i++ {
			if i <= 2 {
				cache[i] = 1
			} else {
				cache[i] = cache[i-1] + cache[i-2] + cache[i-3]
			}
		}
		return cache[n]
	}
}