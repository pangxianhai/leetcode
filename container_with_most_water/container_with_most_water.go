package container_with_most_water

func maxArea(height []int) int {
	L := len(height)
	//if L < 2 {
	//	return 0
	//}
	max := 0
	for i, j := 0, L-1; i < j; {
		var area int
		if height[i] > height[j] {
			area = (j - i) * height[j]
			j--
		} else {
			area = (j - i) * height[i]
			i++
		}
		if area > max {
			max = area
		}
	}
	return max
}

//func maxArea(height []int) int {
//	L := len(height)
//	if L < 2 {
//		return 0
//	}
//	max := 0
//	for i := 1; i < L; i++ {
//		for j := 0; j < i; j++ {
//			area := calArea(height, i, j)
//			if area > max {
//				max = area
//			}
//		}
//	}
//	return max
//}

func calArea(height []int, i int, j int) int {
	h := height[i]
	if height[j] < h {
		h = height[j]
	}
	return h * (i - j)
}
