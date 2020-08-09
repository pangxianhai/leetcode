package recover_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var list []*TreeNode

func recoverTree(root *TreeNode) {
	list = make([]*TreeNode, 0)
	inorder(root)

	L := len(list)
	var n1 *TreeNode
	var n2 *TreeNode
	for i, v := range list {
		if i == L-1 {
			break
		}
		if v.Val > list[i+1].Val {
			if n1 == nil {
				n1 = v
				n2 = list[i+1]
			} else {
				n2 = list[i+1]
			}
		}
	}
	if n1 != nil && n2 != nil {
		n1.Val, n2.Val = n2.Val, n1.Val
	}
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorder(root.Left)
	}
	list = append(list, root)
	if root.Right != nil {
		inorder(root.Right)
	}
}
