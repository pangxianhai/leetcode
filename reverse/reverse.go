package reverse

func reverse(x int) int {
	MaxInt32 := 1<<31 - 1
	MinInt32 := -1 << 31

	r := 0
	for x != 0 {
		p := x % 10
		x = x / 10
		r = r*10 + p
		if r > MaxInt32 || r < MinInt32 {
			return 0
		}
	}

	return r
}
