package sub_sort_lcci

import "testing"

func TestSubSort(t *testing.T) {
    t.Log("开始测试部分排序")
    var arr []int
    var result []int
    // arr = []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
    // result = subSort(arr)
    // t.Log("返回结果:", result)

    // arr = []int{1, 3, 9, 7, 5}
    // result = subSort(arr)
    // t.Log("返回结果:", result)

    // arr = []int{5, 3, 1, 7, 9}
    // result = subSort(arr)
    // t.Log("返回结果:", result)

    arr = []int{1, 3, 5, 7, 9}
    result = subSort(arr)
    t.Log("返回结果:", result)
}
