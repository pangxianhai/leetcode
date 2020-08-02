package largest_1_bordered_square

import (
	"fmt"
	"testing"
)

func TestLargest1BorderedSquare(t *testing.T) {
	var grid [][]int
	var maxLen int

	//grid = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	//maxLen = largest1BorderedSquare(grid)
	//fmt.Println("9", maxLen)
	//
	//grid = [][]int{{1, 1, 0, 0}}
	//maxLen = largest1BorderedSquare(grid)
	//fmt.Println("1", maxLen)

	grid = [][]int{{0, 1, 0, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 0, 1, 1}, {1, 1, 1, 1, 0}, {0, 1, 1, 1, 1}, {1, 1, 1, 0, 1}, {0, 1, 1, 1, 1}, {1, 1, 1, 0, 1}}
	maxLen = largest1BorderedSquare(grid)
	fmt.Println("9", maxLen)
}
