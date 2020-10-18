package remove_nth_node_from_end_of_list

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	l := ListNode{
		Val:  1,
		Next: nil,
	}
	l1 := ListNode{
		Val:  2,
		Next: nil,
	}
	l.Next = &l1

	h := removeNthFromEnd(&l, 1)
	if h == nil {
		fmt.Println("无结果")
	}
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}
