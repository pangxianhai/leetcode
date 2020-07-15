package diving_board_lcci

func divingBoard(shorter int, longer int, k int) []int {
	if k <= 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	if k == 1 {
		return []int{shorter, longer}
	}
	ks := make([]int, k+1)
	for i := 0; i <= k; i++ {
		s := shorter*(k-i) + longer*i
		ks[i] = s
	}
	return ks
}
