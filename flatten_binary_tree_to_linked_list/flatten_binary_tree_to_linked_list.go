package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	flatten0(root)
}

func flatten0(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	left1, left2 := flatten0(root.Left)
	right1, right2 := flatten0(root.Right)
	root.Left = nil
	root.Right = left1
	if left2 != nil {
		left2.Right = right1
	} else {
		if left1 == nil {
			root.Right = right1
		} else {
			left1.Right = right1
		}
	}
	if right2 == nil {
		right2 = right1
	}
	if right2 == nil {
		right2 = left2
	}
	if right2 == nil {
		right2 = left1
	}
	return root, right2
}
