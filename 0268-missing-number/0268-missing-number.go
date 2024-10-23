func missingNumber(nums []int) int {
    xor := len(nums)
    for index, value := range nums {
        xor ^= value
        xor ^= index
    }
    return xor
}