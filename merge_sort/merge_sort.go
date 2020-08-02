package merge_sort

import "math"

func MergeSort(data []int, p int, r int) {
	if p >= r {
		return
	}
	q := p + (r-p)/2
	MergeSort(data, p, q)
	MergeSort(data, q+1, r)
	merge(data, p, q, r)
}

func merge(data []int, p int, q int, r int) {
	L := make([]int, q-p+1)
	for i := p; i <= q; i++ {
		L[i-p] = data[i]
	}
	L = append(L, math.MaxInt32)
	R := make([]int, r-q)
	for i := q + 1; i <= r; i++ {
		R[i-q-1] = data[i]
	}
	R = append(R, math.MaxInt32)
	i, j := 0, 0
	for k := p; k <= r; k++ {
		if L[i] < R[j] {
			data[k] = L[i]
			i++
		} else {
			data[k] = R[j]
			j++
		}
	}
}
