package quick_sort

func QuickSortS(arr []int) {
	QuickSort(arr, 0, len(arr)-1)
}

func QuickSort(arr []int, p int, r int) {
	if p >= r {
		return
	}
	q := Partition(arr, p, r)
	QuickSort(arr, p, q-1)
	QuickSort(arr, q+1, r)
}

func Partition(arr []int, p int, r int) int {
	i, j := p-1, p
	for ; j < r; j++ {
		if arr[j] <= arr[r] {
			i++
			exchange(arr, i, j)
		}
	}
	exchange(arr, i+1, r)
	return i + 1
}

func exchange(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
