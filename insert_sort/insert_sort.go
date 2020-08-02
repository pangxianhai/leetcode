package insert_sort

// 插入排序

func InsertSort(data []int) []int {
	N := len(data)
	for i := 1; i < N; i++ {
		v := data[i]
		j := i - 1
		for ; j >= 0 && data[j] > v; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = v
	}

	return data
}
