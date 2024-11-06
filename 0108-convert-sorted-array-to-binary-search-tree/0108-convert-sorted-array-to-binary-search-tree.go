/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    var sortedArrayToBSTRec func([]int, int, int) *TreeNode

    sortedArrayToBSTRec = func(nums []int, l int, r int) *TreeNode {
        if l > r {
            return nil
        }
        if l == r {
            return &TreeNode{nums[l], nil, nil}
        }
        m := (l + r) / 2
        head := &TreeNode{nums[m], nil, nil}
        head.Left = sortedArrayToBSTRec(nums, l, m-1)
        head.Right = sortedArrayToBSTRec(nums, m+1, r)
        return head
    }

    return sortedArrayToBSTRec(nums, 0, len(nums)-1)
}
