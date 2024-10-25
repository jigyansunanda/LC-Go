func findAnagrams(s string, p string) []int {
    m, n := len(s), len(p)
    var ans []int
    if n > m {
        return ans
    }
    charCount := [26]int{}
    patternCharCount := [26]int{}
    for i := 0; i < n; i++ {
        patternCharCount[p[i]-'a']++
    }
    for i := 0; i+n-1 < m; i++ {
        startIndex, endIndex := i, i+n-1
        if startIndex == 0 {
            for index := startIndex; index <= endIndex; index++ {
                charCount[s[index]-'a']++
            }
        } else {
            charCount[s[startIndex-1]-'a']--
            charCount[s[endIndex]-'a']++
        }
        if charCount == patternCharCount {
            ans = append(ans, i)
        }
    }
    return ans
}