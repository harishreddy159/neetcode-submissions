/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
    result := dummy
	carry := 0
	for l1 != nil || l2 != nil{
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + carry
		carry = 0
		if sum > 9 {
			carry = sum / 10
			sum = sum % 10
		}
		result.Next = &ListNode{Val: sum}
		result = result.Next
	}
	if carry > 0 {
		result.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
