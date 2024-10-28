func search(nums []int, target int) int {
    lo, hi := -1, len(nums)
    for hi > lo+1 {
        mid := (lo + hi) / 2
        if nums[mid] < target {
            lo = mid
        } else if nums[mid] > target {
            hi = mid
        } else {
            return mid
        }
    }
    return -1
}