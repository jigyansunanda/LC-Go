/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var addNumbers func(*ListNode, *ListNode, int) *ListNode

    addNumbers = func(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
        if l1 == nil && l2 == nil {
            if carry == 0 {
                return nil
            } else {
                return &ListNode{carry, nil}
            }
        } else if l1 == nil {
            if carry == 0 {
                return l2
            } else {
                sum := l2.Val + carry
                carry = sum / 10
                data := sum % 10
                currHead := &ListNode{data, nil}
                currTail := addNumbers(l1, l2.Next, carry)
                currHead.Next = currTail
                return currHead
            }
        } else if l2 == nil {
            if carry == 0 {
                return l1
            } else {
                sum := l1.Val + carry
                carry = sum / 10
                data := sum % 10
                currHead := &ListNode{data, nil}
                currTail := addNumbers(l1.Next, l2, carry)
                currHead.Next = currTail
                return currHead
            }
        } else {
            sum := l1.Val + l2.Val + carry
            carry = sum / 10
            data := sum % 10
            currHead := &ListNode{data, nil}
            currTail := addNumbers(l1.Next, l2.Next, carry)
            currHead.Next = currTail
            return currHead
        }
    }

    return addNumbers(l1, l2, 0)
}