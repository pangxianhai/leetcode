package unique_binary_search_trees

/*
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
  这竟然是个数学题  卡塔兰数

题目要求是计算不同二叉搜索树的个数。为此，我们可以定义两个函数：
G(n): 长度为  的序列能构成的不同二叉搜索树的个数。

F(i, n)F(i,n): 以 i 为根、序列长度为 n 的不同二叉搜索树个数 (1≤i≤n)。

可见，G(n) 是我们求解需要的函数。

稍后我们将看到，G(n) 可以从 F(i, n)F(i,n) 得到，而 F(i, n)F(i,n) 又会递归地依赖于 G(n)G(n)。

首先，根据上一节中的思路，不同的二叉搜索树的总数 G(n)G(n)，是对遍历所有 ii (1 \le i \le n)(1≤i≤n) 的 F(i, n)F(i,n) 之和。换言之：

G(n)= sum(F(i,n)) i从1到n
F(i,n)=G(i−1)⋅G(n−i)
所有 G(n) = sum(G(i-1)*G(n-i)) i从1到n
边界 G(0)=1 G(1)=1

由卡塔兰数

C(n+1) = [2(2n+1)] / [n+2] * C(n)
解法如下
​
*/

func numTrees(n int) int {
	c := 1
	for i := 1; i < n; i++ {
		c = c * 2 * (2*i + 1) / (i + 2)
	}

	return c
}
