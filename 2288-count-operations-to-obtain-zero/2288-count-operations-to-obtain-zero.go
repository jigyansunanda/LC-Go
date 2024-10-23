func countOperations(a int, b int) int {
    count := 0
    for true {
        if a == 0 || b == 0 {
            break
        }
        if b > a {
            a, b = b, a
        }
        count += a / b
        a = a % b
    }
    return count
}