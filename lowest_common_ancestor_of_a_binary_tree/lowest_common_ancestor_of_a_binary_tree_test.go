package lowest_common_ancestor_of_a_binary_tree

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	qi := TreeNode{
		Val: 7,
	}
	si := TreeNode{
		Val: 4,
	}
	liu := TreeNode{
		Val: 6,
	}
	er := TreeNode{
		Val:   2,
		Left:  &qi,
		Right: &si,
	}
	lin := TreeNode{
		Val: 0,
	}
	ba := TreeNode{
		Val: 8,
	}
	wu := TreeNode{
		Val:   5,
		Left:  &liu,
		Right: &er,
	}
	yi := TreeNode{
		Val:   1,
		Left:  &lin,
		Right: &ba,
	}
	san := TreeNode{
		Val:   3,
		Left:  &wu,
		Right: &yi,
	}

	var comm *TreeNode
	comm = lowestCommonAncestor(&san, &wu, &yi)
	fmt.Println("3", comm.Val)

	comm = lowestCommonAncestor(&san, &wu, &si)
	fmt.Println("5", comm.Val)

}
