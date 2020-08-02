package zigzag_conversion

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	var s string
	var res string
	var right string

	s = "LEETCODEISHIRING"
	res = convert(s, 3)
	right = "LCIRETOESIIGEDHN"
	fmt.Println(right, res, right == res)

	s = "AB"
	res = convert(s, 1)
	right = "AB"
	fmt.Println(right, res, right == res)
}
