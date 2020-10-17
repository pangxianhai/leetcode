package four_sum

import (
	"fmt"
	"testing"
)

func TestFourSum(t *testing.T) {
	var nums []int
	var target int
	var result [][]int

	// nums = []int{1, 0, -1, 0, -2, 2}
	// target = 0
	// result = fourSum(nums, target)
	// fmt.Print(result)

	// nums = []int{-2, -1, -1, 1, 1, 2, 2}
	// target = 0
	// result = fourSum(nums, target)
	// fmt.Print(result)

	nums = []int{1, 0, -1, 0, -2, 2}
	target = 0
	result = fourSum(nums, target)
	fmt.Print(result)
}

func TestA(t *testing.T) {
	var nums []int
	nums = []int{-2, -1, 1, 1, 2, -1, 2}
	mySort(nums)
	fmt.Print(nums)
}
