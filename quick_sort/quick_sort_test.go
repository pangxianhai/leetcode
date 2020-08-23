package quick_sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{16, 14, 10, 8, 7, 5, 9, 3, 2, 4, 1, 5}
	QuickSortS(arr)
	fmt.Print(arr)
}
