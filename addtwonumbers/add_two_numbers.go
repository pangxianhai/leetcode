package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	carry := 0
	for {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		curr.Val = carry % 10
		carry = carry / 10

		if l1 != nil || l2 != nil || carry > 0 {
			curr.Next = new(ListNode)
			curr = curr.Next
		} else {
			break
		}

	}
	return dummy
}
