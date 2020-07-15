package towsum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Log("开始测试两数之和")
	num1 := []int{2, 7, 11, 15}
	res := twoSum(num1, 9)
	fmt.Println(res)

}
