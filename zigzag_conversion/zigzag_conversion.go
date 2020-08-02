package zigzag_conversion

func convert(s string, numRows int) string {
	N := numRows*2 - 2
	if numRows == 1 {
		N = 1
	}
	var res []byte
	L := len(s)
	for line := 0; line < numRows; line++ {
		for i := 0; i < L; i++ {
			p := N*i + line
			if p > L-1 {
				break
			}
			res = append(res, s[p])
			if line != 0 && line != numRows-1 {
				q := N*i + N - line
				if q > L-1 {
					break
				}
				res = append(res, s[q])
			}
		}
	}

	return string(res)
}
