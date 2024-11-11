func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    cache := make([][]int, m+1)
    for row := 0; row <= m; row++ {
        cache[row] = make([]int, n+1)
        for col := 0; col <= n; col++ {
            if row == 0 || col == 0 {
                cache[row][col] = row + col
            } else {
                if word1[row-1] == word2[col-1] {
                    cache[row][col] = cache[row-1][col-1]
                } else {
                    cache[row][col] = min(cache[row-1][col], cache[row][col-1], cache[row-1][col-1]) + 1
                }
            }
        }
    }
    return cache[m][n]
}