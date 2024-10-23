func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    lo, hi := 1, x/2+1
    for hi > lo+1 {
        mid := lo + (hi-lo)/2
        if mid <= x/mid {
            lo = mid
        } else {
            hi = mid
        }
    }
    return lo
}