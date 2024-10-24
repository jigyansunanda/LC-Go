func addDigits(num int) int {
    for num > 9 {
        sumOfDigits := 0
        for num > 0 {
            sumOfDigits += num % 10
            num /= 10
        }
        num = sumOfDigits
    }
    return num
}