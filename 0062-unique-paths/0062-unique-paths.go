func uniquePaths(m int, n int) int {
    cache := make([][]int, m)
    for row := 0; row < m; row++ {
        cache[row] = make([]int, n)
        for col := 0; col < n; col++ {
            if row == 0 || col == 0 {
                cache[row][col] = 1
            } else {
                cache[row][col] = cache[row-1][col] + cache[row][col-1]
            }
        }
    }
    return cache[m-1][n-1]
}