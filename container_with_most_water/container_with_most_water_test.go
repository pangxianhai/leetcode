package container_with_most_water

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	var res int
	var height []int

	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res = maxArea(height)
	fmt.Println("49", res)
}
