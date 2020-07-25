package burst_balloons

import (
	"fmt"
	"testing"
)

func TestMaxCoins(t *testing.T) {
	var nums []int
	var coins int

	nums = []int{3, 1, 5, 8}
	coins = maxCoins(nums)
	fmt.Println(167, coins)

	nums = []int{9, 76, 64, 21, 97, 60, 5}
	coins = maxCoins(nums)
	fmt.Println(1088290, coins)
}
