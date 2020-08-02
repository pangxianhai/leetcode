package largest_1_bordered_square

func largest1BorderedSquare(grid [][]int) int {
	var maxLen int

	for i := 0; i < len(grid); i++ {
		s := -1
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				for k := s; k <= j; k++ {
					if j-k > maxLen {
						is := checkout1Border(grid, i, k, j)
						if is {
							maxLen = j - k
						}
					}
				}

			} else {
				s = j
			}
		}
	}
	return maxLen * maxLen
}

func checkout1Border(grid [][]int, line int, rows int, rowe int) bool {
	borderLen := rowe - rows
	if len(grid)-line < borderLen {
		return false
	}
	for i := line + 1; i < line+borderLen; i++ {
		if grid[i][rows+1] == 0 {
			return false
		}
		if grid[i][rowe] == 0 {
			return false
		}
	}
	lastLine := line + borderLen - 1
	for i := rows + 1; i < rowe; i++ {
		if grid[lastLine][i] == 0 {
			return false
		}
	}
	return true
}
