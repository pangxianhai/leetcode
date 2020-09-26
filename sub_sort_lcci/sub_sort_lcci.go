package sub_sort_lcci

//给定一个整数数组，编写一个函数，
//找出索引m和n，只要将索引区间[m,n]的元素排好序，整个数组就是有序的
//。注意：n-m尽量最小，也就是说，找出符合条件的最短序列。函数返回值为[m,n]，
//若不存在这样的m和n（例如整个数组是有序的），请返回[-1,-1]。

func subSort(array []int) []int {
    last, first := -1, -1
    L := len(array)
    if L <= 1 {
        return []int{first, last}
    }
    max, min := array[0], array[L-1]
    for i := 0; i < L; i++ {
        if array[i] >= max {
            max = array[i]
        } else {
            last = i
        }
        if array[L-i-1] <= min {
            min = array[L-i-1]
        } else {
            first = L - i - 1
        }
    }
    return []int{first, last}
}
