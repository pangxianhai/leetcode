package heap_sort

func HeapSort(arr []int) {
	L := len(arr)
	BuildMaxHeap(arr)
	heapSize := L
	for i := L - 1; i > 0; i-- {
		tmp := arr[0]
		arr[0] = arr[i]
		arr[i] = tmp
		heapSize--
		MaxHeapify(arr, 0, heapSize)
	}
}

func BuildMaxHeap(arr []int) {
	L := len(arr)
	for i := 2 / L; i >= 0; i-- {
		MaxHeapify(arr, i, L)
	}
}

//以 i 为根节点 构建最大堆
func MaxHeapify(arr []int, i int, heapSize int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if i >= heapSize {
		return
	}
	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		tmp := arr[i]
		arr[i] = arr[largest]
		arr[largest] = tmp
		MaxHeapify(arr, largest, heapSize)
	}
}
