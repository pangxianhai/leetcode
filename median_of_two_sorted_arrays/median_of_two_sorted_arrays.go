package median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	i1, i2 := 0, 0
	for {
		if i1 == len(nums1) {
			return nums2[i2+k-1]
		} else if i2 == len(nums2) {
			return nums1[i1+k-1]
		}
		if k == 1 {
			return min(nums1[i1], nums2[i2])
		}
		h := k / 2
		ni1 := min(i1+h, len(nums1)) - 1
		ni2 := min(i2+h, len(nums2)) - 1
		p1, p2 := nums1[ni1], nums2[ni2]
		if p1 < p2 {
			k -= (ni1 - i1 + 1)
			i1 = ni1 + 1
		} else {
			k -= (ni2 - i2 + 1)
			i2 = ni2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	L1, L2 := len(nums1), len(nums2)
// 	TL := L1 + L2
// 	c1 := 0
// 	c2 := 0
// 	if TL%2 == 0 {
// 		c1 = TL / 2
// 		c2 = c1 + 1
// 	} else {
// 		c1 = TL/2 + 1
// 		c2 = c1
// 	}
// 	i1 := 0
// 	i2 := 0
// 	for i := 0; i < TL; i++ {
// 		if i+1 == c1 {
// 			if c1 == c2 {
// 				if choiceNums1(i1, L1, i2, L2, nums1, nums2) {
// 					return float64(nums1[i1])
// 				} else {
// 					return float64(nums2[i2])
// 				}
// 			} else {
// 				v1, v2 := 0, 0
// 				if choiceNums1(i1, L1, i2, L2, nums1, nums2) {
// 					v1 = nums1[i1]
// 					i1++
// 				} else {
// 					v1 = nums2[i2]
// 					i2++
// 				}
// 				if choiceNums1(i1, L1, i2, L2, nums1, nums2) {
// 					v2 = nums1[i1]
// 					i1++
// 				} else {
// 					v2 = nums2[i2]
// 					i2++
// 				}
// 				return (float64(v1) + float64(v2)) / float64(2)
// 			}
// 		} else {
// 			if choiceNums1(i1, L1, i2, L2, nums1, nums2) {
// 				i1++
// 			} else {
// 				i2++
// 			}
// 		}
// 	}
// 	return 0
// }

// func choiceNums1(i1 int, L1 int, i2 int, L2 int, nums1 []int, nums2 []int) bool {
// 	if i1 > L1-1 {
// 		return false
// 	} else if i2 > L2-1 {
// 		return true
// 	} else {
// 		return nums1[i1] < nums2[i2]
// 	}
// }
