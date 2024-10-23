func countEven(num int) int {
    count := 0
    for i := 1; i <= num; i++ {
        digitSum := 0
        D := i
        for D > 0 {
            digitSum += D % 10
            D /= 10
        }
        if digitSum%2 == 0 {
            count++
        }
    }
    return count
}