func findMax(arr []int) int {
    maxValue := 0
    for _, value := range arr {
        maxValue = max(maxValue, value)
    }
    return maxValue
}

func lengthOfLongestSubstring(s string) int {
    count, l, ans := make([]int, 256), 0, 0
    for r := 0; r < len(s); r++ {
        count[s[r]]++
        for findMax(count) > 1 {
            count[s[l]]--
            l++
        }
        if findMax(count) == 1 {
            ans = max(ans, r-l+1)
        }
    }
    return ans
}