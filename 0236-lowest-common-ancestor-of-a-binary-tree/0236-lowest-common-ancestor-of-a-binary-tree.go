/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    if root == p || root == q {
        return root
    } else {
        l, r := lowestCommonAncestor(root.Left,p,q), lowestCommonAncestor(root.Right,p,q)
        if l != nil && r != nil {
            return root
        } else if l != nil {
            return l
        } else if r != nil {
            return r
        } else {
            return nil
        }
    }
}