func fib(n int) int {
	if n <= 1 {
		return n
	} else {
		cache := make([]int, n+1)
		for i := 0; i <= n; i++ {
			if i <= 1 {
				cache[i] = i
			} else {
				cache[i] = cache[i-1] + cache[i-2]
			}
		}
		return cache[n]
	}
}