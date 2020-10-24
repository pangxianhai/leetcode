package video_stitching

import (
	"fmt"
	"testing"
)

func TestVideoStitching(t *testing.T) {
	var clips [][]int
	var T int

	//clips = [][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}
	//T = 9
	//t.Log(videoStitching(clips, T))

	//clips = [][]int{{0, 4}, {2, 8}}
	//T = 5
	//t.Log(videoStitching(clips, T))

	clips = [][]int{{8,10},{17,39},{18,19},{8,16},{13,35},{33,39},{11,19},{18,35}}
	T = 20
	t.Log(videoStitching(clips, T))
}

func TestA(t *testing.T) {
	s := fmt.Sprintf("%d-%d", 2, 3)
	t.Log(s)
}
