func largestOddNumber(num string) string {
    endIndex := -1
    for index, char := range num {
        if char%2 == 1 {
            endIndex = index
        }
    }
    return num[:endIndex+1]
}