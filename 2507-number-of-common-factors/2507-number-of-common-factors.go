func countDivisors(x int) (divisors int) {
    divisors = 0
    for i := 1; i*i <= x; i++ {
        if x%i == 0 {
            divisors++
            if x/i != i {
                divisors++
            }
        }
    }
    return
}

func gcd(a int, b int) int {
    if a == 0 {
        return b
    } else {
        return gcd(b%a, a)
    }
}

func commonFactors(a int, b int) int {
    return countDivisors(gcd(a, b))
}