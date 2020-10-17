package total_nqueens

import (
	"fmt"
)

func totalNQueens(n int) int {
	num := make([]int, n)
	t := queens(num, 0)
	fmt.Println(num)
	return t
}

func queens(num []int, i int) int {
	n := len(num)
	t := 0
	for j := 1; j <= n; j++ {
		if !isInsertI(num, i, j) {
			continue
		}
		if i < n-1 {
			num[i] = j
			tmp := queens(num, i+1)
			if tmp > 0 {
				t += tmp
				continue
			} else {
				num[i] = 0
				continue
			}
		} else {
			num[i] = j
			t++
		}
	}
	return t
}

func isInsertI(num []int, i int, j int) bool {
	for k := 0; k < i; k++ {
		if num[k] == j {
			return false
		}
		tmp1 := num[k] - j
		tmp2 := k - i
		if abs(tmp1) == abs(tmp2) {
			return false
		}
	}
	return true
}

func abs(n int) int {
	if n > 0 {
		return n
	} else {
		return 0 - n
	}
}
