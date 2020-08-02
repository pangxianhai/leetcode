package merge_sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	data := []int{3, 4, 1, 5, 9, 5, 7}
	MergeSort(data, 0, len(data)-1)
	fmt.Println(data)
}
