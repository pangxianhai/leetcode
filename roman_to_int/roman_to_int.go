package roman_to_int

func romanToInt(s string) int {
	numDct := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	numTwoDct := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}

	L := len(s)
	num := 0
	for i := 0; i < L; i++ {
		if i < L-1 {
			tmp := s[i : i+2]
			n, e := numTwoDct[tmp]
			if e {
				num += n
				i++
				continue
			}
		}
		n, _ := numDct[s[i:i+1]]
		num += n
	}
	return num
}
