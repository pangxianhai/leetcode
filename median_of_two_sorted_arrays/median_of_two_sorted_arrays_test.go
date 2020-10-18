package median_of_two_sorted_arrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	var nums1 []int
	var nums2 []int

	nums1 = []int{1, 3}
	nums2 = []int{2}
	t.Log(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	t.Log(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{0, 0}
	nums2 = []int{0, 0}
	t.Log(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 1}
	nums2 = []int{1, 1}
	t.Log(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{}
	nums2 = []int{1}
	t.Log(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{2}
	nums2 = []int{}
	t.Log(findMedianSortedArrays(nums1, nums2))
}
