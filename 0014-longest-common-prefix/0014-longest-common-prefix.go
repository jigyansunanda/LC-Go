func longestCommonPrefix(strs []string) string {
    slices.Sort(strs)
    s1, s2, index := strs[0], strs[len(strs)-1], 0
    for index < len(s1) && index < len(s2) {
        if s1[index] == s2[index] {
            index++
        } else {
            break
        }
    }
    return s1[0:index]
}