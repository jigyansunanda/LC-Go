/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    var oddTail, evenTail, evenHead *ListNode
    isOddNode := true
    currNode := head
    for currNode != nil {
        if isOddNode {
            if oddTail == nil {
                oddTail = currNode
            } else {
                oddTail.Next = currNode
                oddTail = oddTail.Next
            }
        } else {
            if evenHead == nil {
                evenHead = currNode
            }
            if evenTail == nil {
                evenTail = evenHead
            } else {
                evenTail.Next = currNode
                evenTail = evenTail.Next
            }
        }
        isOddNode = !isOddNode
        currNode = currNode.Next
    }
    if evenTail != nil {
        evenTail.Next = nil
    }
    oddTail.Next = evenHead
    return head
}