/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    var totalSum int = 0

    var sumToLeaf func(*TreeNode, int)
    sumToLeaf = func(node *TreeNode, pre int) {
        if node == nil {
            return
        }
        pre = (pre * 10) + node.Val
        if node.Left == nil && node.Right == nil {
            totalSum += pre
            return
        }
        if node.Left != nil {
            sumToLeaf(node.Left, pre)
        }
        if node.Right != nil {
            sumToLeaf(node.Right, pre)
        }
    }

    sumToLeaf(root, 0)
    return totalSum
}