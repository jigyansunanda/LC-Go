func isThree(n int) bool {
    divisors := 0
    for i := 1; i*i <= n; i++ {
        if n%i == 0 {
            divisors += 1
            other := n / i
            if other != i {
                divisors += 1
            }
        }
    }
    return (divisors == 3)
}