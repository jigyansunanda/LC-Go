func arraySign(nums []int) int {
    neg := 0
    for _, val := range nums {
        if val == 0 {
            return 0
        } else if val < 0 {
            neg++
        }
    }
    if neg%2 == 0 {
        return 1
    } else {
        return -1
    }
}