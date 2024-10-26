func judgeSquareSum(c int) bool {
    var lo, hi, C int64 = 0, int64(math.Sqrt(float64(c))), int64(c)
    for hi >= lo {
        var squareSum int64 = (lo * lo) + (hi * hi)
        if squareSum == C {
            return true
        } else if squareSum < C {
            lo++
        } else {
            hi--
        }
    }
    return false
}