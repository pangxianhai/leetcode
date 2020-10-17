package four_sum

func fourSum(nums []int, target int) [][]int {
	arr := make([][]int, 0)
	mySort(nums)
	L := len(nums)
	if L < 4 {
		return arr
	}
	for i1 := 0; i1 < L; i1++ {
		if i1 > 0 && nums[i1] == nums[i1-1] {
			continue
		}
		if nums[i1]+nums[L-3]+nums[L-2]+nums[L-1] < target {
			continue
		}
		for i2 := i1 + 1; i2 < L; i2++ {
			if i2 > i1+1 && nums[i2] == nums[i2-1] {
				continue
			}
			if nums[i1]+nums[i2]+nums[L-2]+nums[L-1] < target {
				continue
			}
			left, right := i2+1, L-1
			t := target - nums[i1] - nums[i2]
			for left < right {
				sum := nums[left] + nums[right]
				if sum > t {
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
					// right--
				} else if sum < t {
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					// left++
				} else {
					ri := []int{nums[i1], nums[i2], nums[left], nums[right]}
					arr = append(arr, ri)
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				}
			}
		}
	}
	return arr
}

func mySort(nums []int) {
	mySort1(nums, 0, len(nums)-1)
}

func mySort1(nums []int, p int, r int) {
	if p >= r {
		return
	}
	q := pattern(nums, p, r)
	mySort1(nums, p, q-1)
	mySort1(nums, q+1, r)
}

func pattern(nums []int, p int, r int) int {
	i, j := p-1, p
	for k := p; k < r; k++ {
		if nums[k] > nums[r] {
			j++
			continue
		} else {
			i++
			j++
			exchange(nums, i, k)
		}
	}
	exchange(nums, i+1, r)
	return i + 1
}

func exchange(nums []int, i int, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
	return
}
