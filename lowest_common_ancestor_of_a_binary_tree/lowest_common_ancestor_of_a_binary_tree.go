package lowest_common_ancestor_of_a_binary_tree

/**
 * 求二叉树的公共最深祖先
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//存储节点方式 不如递归快
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := make(map[int]*TreeNode)
	pParent := make(map[int]bool)

	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
		}
		_, pe := parent[p.Val]
		_, qe := parent[q.Val]
		if pe && qe {
			return
		}
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	pp := p
	for ; ; {
		pParent[pp.Val] = true
		pp = parent[pp.Val]
		if pp == nil {
			break
		}
	}
	pq := q
	for ; ; {
		_, qe := pParent[pq.Val]
		if qe {
			return pq
		}
		pq = parent[pq.Val]
		if pq == nil {
			break
		}
	}
	return nil
}

//递归方式 更快
//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	if root == nil {
//		return root
//	}
//	left := lowestCommonAncestor(root.Left, p, q)
//	right := lowestCommonAncestor(root.Right, p, q)
//
//	if root.Val == p.Val || root.Val == q.Val {
//		return root
//	}
//	if left != nil && right != nil {
//		return root
//	}
//	if left != nil {
//		return left
//	}
//	return right
//}
