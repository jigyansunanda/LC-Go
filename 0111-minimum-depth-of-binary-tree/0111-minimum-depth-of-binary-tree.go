/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    } else if root.Left == nil && root.Right == nil {
        return 1
    } else {
        lh, rh := 10000000, 10000000
        if root.Left != nil {
            lh = minDepth(root.Left)
        }
        if root.Right != nil {
            rh = minDepth(root.Right)
        }
        if min(lh, rh) == 10000000 {
            return 1
        } else {
            return 1 + min(lh, rh)
        }
    }
}