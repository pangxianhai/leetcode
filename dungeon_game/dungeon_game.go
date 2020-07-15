package dungeon_game

func calculateMinimumHP(dungeon [][]int) int {
	res := make([][]int, len(dungeon))
	for i := 0; i < len(dungeon); i++ {
		res[i] = make([]int, len(dungeon[i]))
	}
	n := len(dungeon)
	m := len(dungeon[0])
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			v := dungeon[i][j]
			hasV := res[i][j]
			nextMin := 0
			if i < n-1 {
				nextMin = res[i+1][j]
			}
			if j < m-1 && (nextMin > res[i][j+1] || nextMin <= 0) {
				nextMin = res[i][j+1]
			}
			if nextMin <= 0 {
				nextMin = 1
			}
			thisMin := nextMin - v
			if thisMin <= 0 {
				thisMin = 1
			}
			if hasV <= 0 || hasV > thisMin {
				res[i][j] = thisMin
			}
		}
	}
	//fmt.Println(res)
	return res[0][0]
}
