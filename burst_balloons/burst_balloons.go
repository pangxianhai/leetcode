package burst_balloons

/*
有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。

求所能获得硬币的最大数量。

说明:

你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
*/

func maxCoins(nums []int) int {
	numN := len(nums)
	val := make([]int, numN+2)
	for i := 0; i < numN; i++ {
		val[i+1] = nums[i]
	}
	val[0], val[numN+1] = 1, 1
	rec := make([][]int, numN+2)
	for i := 0; i < len(rec); i++ {
		rec[i] = make([]int, numN+2)
	}

	for i := numN - 1; i >= 0; i-- {
		for j := i + 2; j < numN+2; j++ {
			for k := i + 1; k < j; k++ {
				rec[i][j] = max(val[i]*val[k]*val[j]+rec[i][k]+rec[k][j], rec[i][j])
			}
		}
	}
	return rec[0][numN+1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
