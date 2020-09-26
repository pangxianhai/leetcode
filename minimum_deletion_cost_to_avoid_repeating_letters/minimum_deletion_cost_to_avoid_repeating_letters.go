package minimum_deletion_cost_to_avoid_repeating_letters

import "fmt"

// 题目描述
// 给你一个字符串 s 和一个整数数组 cost ，其中 cost[i] 是从 s 中删除字符 i 的代价。
// 返回使字符串任意相邻两个字母不相同的最小删除成本。
// 请注意，删除一个字符后，删除其他字符的成本不会改变。
func MinCost(s string, cost []int) int {
	totalCost := 0
	maxCost := 0
	sameChar := s[0]
	L := len(s)
	for i := 0; i < L; i++ {
		if s[i] != sameChar {
			sameChar = s[i]
			totalCost -= maxCost
			fmt.Println("减去的最大值：", maxCost)
			maxCost = cost[i]
		}
		totalCost += cost[i]
		if cost[i] > maxCost {
			maxCost = cost[i]
		}
		if i == L-1 {
			totalCost -= maxCost
			fmt.Println("减去的最大值：", maxCost)
		}
	}
	return totalCost
}
