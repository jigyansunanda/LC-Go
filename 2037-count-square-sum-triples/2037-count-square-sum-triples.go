func countTriples(n int) int {
    triples := 0
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            for k := max(i, j); k <= n; k++ {
                if (i*i)+(j*j) == (k * k) {
                    triples++
                }
            }
        }
    }
    return triples
}