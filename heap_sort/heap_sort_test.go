package heap_sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	HeapSort(arr)

	fmt.Print(arr)
}
