package insert_sort

import (
	"fmt"
	"testing"
)

func TestInsertOrder(t *testing.T) {
	data := []int{3, 4, 1, 5, 9, 5, 7}
	afterOrder := InsertSort(data)
	fmt.Println(afterOrder)
}

func Test(t *testing.T) {
	s := "abcd"
	fmt.Println(string(s[1]))
}
