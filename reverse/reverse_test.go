package reverse

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Log("开始测试整数翻转")
	res := reverse(123)
	fmt.Println(res)

}
