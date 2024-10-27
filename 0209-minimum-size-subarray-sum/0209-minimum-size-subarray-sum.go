func minSubArrayLen(target int, nums []int) int {
    var (
        l   int = 0
        sum int = 0
        ans int = len(nums) + 1
    )
    for r := 0; r < len(nums); r++ {
        sum += nums[r]
        for sum >= target {
            ans = min(ans, r-l+1)
            sum -= nums[l]
            l++
        }
    }
    if ans > len(nums) {
        return 0
    }
    return ans
}