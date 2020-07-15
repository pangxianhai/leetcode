package towsum

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target - v]; ok {
			return []int {i, k}
		}
		m[v] = i
	}
	return nil
}
