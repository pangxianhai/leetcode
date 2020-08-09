package recover_binary_search_tree

import "testing"

func TestRecoverTree(test *testing.T) {
	//t0 := TreeNode{
	//	Val: 0,
	//}
	t1 := TreeNode{
		Val: 1,
	}
	t2 := TreeNode{
		Val: 2,
	}
	t3 := TreeNode{
		Val: 3,
	}
	//t4 := TreeNode{
	//	Val: 4,
	//}
	//t5 := TreeNode{
	//	Val: 5,
	//}
	//t6 := TreeNode{
	//	Val: 6,
	//}

	//t3.Left = &t2
	//t3.Right = &t5
	//
	//t2.Left = &t0
	//t2.Right = &t1
	//
	//t5.Left = &t4
	//t5.Right = &t6
	//
	//recoverTree(&t3)
	//
	//t3.Left = &t1
	//t3.Right = &t2
	//
	//t1.Left = &t0
	//t1.Right = &t5
	//
	//t2.Left = &t4
	//t2.Right = &t6
	//
	//recoverTree(&t3)

	t1.Left = &t3
	t3.Right = &t2
	recoverTree(&t1)
}
