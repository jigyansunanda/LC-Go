func isPalindrome(x int) bool {
    if (x < 0) || (x!=0 && x%10 == 0) {
        return false
    }
    var rev int = 0
    for true {
        if x <= rev {
            break
        }
        rev = (rev * 10) + (x % 10)
        x = x / 10
    }
    return (x == rev) || (x == (rev / 10))
}