/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    nodeIndex := 1
    var curr, prev, newHead, newTail *ListNode
    curr = head
    newHead = &ListNode{-1, nil}
    newTail = newHead
    for curr != nil {
        isInsideRange := false
        currHead := curr
        for nodeIndex >= left && nodeIndex <= right {
            next := curr.Next
            curr.Next = prev
            prev = curr
            curr = next
            nodeIndex++
            isInsideRange = true
        }
        if isInsideRange {
            newTail.Next = prev
            newTail = currHead
        } else {
            nodeIndex++
            newTail.Next = currHead
            newTail = newTail.Next
            curr = curr.Next
        }

    }
    return newHead.Next
}