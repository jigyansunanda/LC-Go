func oddLengthPalindromeStartingAt(s string, l int, r int) (startIndex int, length int) {
    startIndex = l
    length = 0
    if s[l] != s[r] {
        return
    }
    for l >= 0 && r < len(s) && s[l] == s[r] {
        if l == r {
            length += 1
        } else {
            length += 2
        }
        startIndex = l
        l--
        r++
    }
    return
}

func longestPalindrome(s string) string {
    maxLength, palindromeIndex := 0, 0
    for i := 0; i < len(s); i++ {
        startIndex, length := oddLengthPalindromeStartingAt(s, i, i)
        if length > maxLength {
            maxLength = length
            palindromeIndex = startIndex
        }
    }
    for i := 0; i < len(s)-1; i++ {
        startIndex, length := oddLengthPalindromeStartingAt(s, i, i+1)
        if length > maxLength {
            maxLength = length
            palindromeIndex = startIndex
        }
    }
    return s[palindromeIndex : palindromeIndex+maxLength]
}