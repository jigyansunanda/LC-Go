func maxSubArray(nums []int) int {
    maxSumEndingAtCurrentIndex, maxSubArraySum := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        maxSumEndingAtCurrentIndex = max(maxSumEndingAtCurrentIndex+nums[i], nums[i])
        maxSubArraySum = max(maxSubArraySum, maxSumEndingAtCurrentIndex)
    }
    return maxSubArraySum
}