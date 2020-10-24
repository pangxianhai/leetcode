package video_stitching

import (
	"fmt"
	"math/rand"
)

func videoStitching(clips [][]int, T int) int {
	clipMaxDict := make(map[int]int)
	starts := make([]int, 0)
	for _, clip := range clips {
		end, exist := clipMaxDict[clip[0]]
		if exist {
			if clip[1] > end {
				clipMaxDict[clip[0]] = clip[1]
			}
		} else {
			clipMaxDict[clip[0]] = clip[1]
			starts = append(starts, clip[0])
		}
	}
	if len(starts) <= 0 {
		return -1
	}
	fmt.Println(clipMaxDict)
	sortStart(starts)
	if starts[0] > 0 {
		return -1
	}
	packages := make([]int, T+1)
	choices := make([]string, T+1)
	for i := 0; i <= T; i++ {
		packages[i] = -1
	}
	for index, start := range starts {
		end := clipMaxDict[start]
		for i := start + 1; i <= min(end, T); i++ {
			if index == 0 {
				packages[i] = 1
				choices[i] = fmt.Sprintf("%d-%d", start, end)
			} else {
				if packages[start] == -1 {
					continue
				}
				if packages[i] == -1 || packages[start]+1 < packages[i] {
					packages[i] = packages[start] + 1
					choices[i] = fmt.Sprintf("%d-%d", start, end)
				}
			}
		}
	}
	fmt.Println(choices)
	return packages[T]
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func sortStart(starts []int) {
	quickSort(starts, 0, len(starts)-1)
}

func quickSort(data []int, p int, r int) {
	if p >= r {
		return
	}
	q := patten(data, p, r)
	quickSort(data, p, q-1)
	quickSort(data, q+1, r)
}

func patten(data []int, p int, r int) int {
	t := rand.Intn(r - p + 1)
	exchange(data, p+t, r)
	q := p - 1
	for i := p; i < r; i++ {
		if data[i] < data[r] {
			exchange(data, q+1, i)
			q++
		}
	}
	exchange(data, q+1, r)
	return q + 1
}

func exchange(data []int, x int, y int) {
	if x == y {
		return
	}
	t := data[x]
	data[x] = data[y]
	data[y] = t
}
