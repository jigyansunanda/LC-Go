func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    cache := make([][]int, m)
    for row := 0; row < m; row++ {
        cache[row] = make([]int, n)
        for col := 0; col < n; col++ {
            if row+col == 0 {
                cache[row][col] = grid[0][0]
            } else if row == 0 {
                cache[row][col] = cache[row][col-1] + grid[row][col]
            } else if col == 0 {
                cache[row][col] = cache[row-1][col] + grid[row][col]
            } else {
                cache[row][col] = min(cache[row-1][col], cache[row][col-1]) + grid[row][col]
            }
        }
    }
    return cache[m-1][n-1]
}