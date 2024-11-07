/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}

func isBalanced(root *TreeNode) bool {
    var height func(*TreeNode) int

    height = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        lh, rh := height(node.Left), height(node.Right)
        if lh == -1 || rh == -1 || abs(lh-rh) > 1 {
            return -1
        }
        return max(lh, rh) + 1
    }

    h := height(root)
    if h == -1 {
        return false
    } else {
        return true
    }
}