package remove_nth_node_from_end_of_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sen := &ListNode{0, head}
	first, second, t := sen, head, 0
	for second != nil {
		second = second.Next
		t++
		if t >= n+1 {
			first = first.Next
		}
	}
	var next *ListNode
	if first.Next != nil {
		next = first.Next.Next
	}
	first.Next = next
	return sen.Next
}
