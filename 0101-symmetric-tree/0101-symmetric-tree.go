/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    var isSameTree func(*TreeNode, *TreeNode) bool

    isSameTree = func(p *TreeNode, q *TreeNode) bool {
        if p == nil && q == nil {
            return true
        }
        if p == nil || q == nil {
            return false
        }
        return (p.Val == q.Val) && isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
    }

    return isSameTree(root, root)
}