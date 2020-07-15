package wildcard_matching

import "fmt"

func isMatch(s string, p string) bool {
	for len(s) > 0 && len(p) > 0 && p[len(p)-1] != '*' {
		if charMatch(s[len(s)-1], p[len(p)-1]) {
			s = s[:len(s)-1]
			p = p[:len(p)-1]
		} else {
			return false
		}
	}
	if len(p) == 0 {
		return len(s) == 0
	}
	for len(s) > 0 && len(p) > 0 && p[0] != '*' {
		if charMatch(s[0], p[0]) {
			s = s[1:]
			p = p[1:]
		} else {
			return false
		}
	}
	if len(p) == 0 {
		return len(s) == 0
	}

	dp := [][]bool{}
	for r := 0; r <= len(p); r++ {
		dp = append(dp, []bool{})
		for l := 0; l <= len(s); l++ {
			dp[r] = append(dp[r], false)
		}
	}
	dp[0][0] = true
	for r := 1; r <= len(p); r++ {
		pi := fmt.Sprintf("%c", p[r-1])
		if pi == "*" {
			for l := 0; l < len(dp[r]); l++ {
				dp[r][l] = true
			}
		} else {
			break
		}
	}

	for r := 1; r <= len(p); r++ {
		pi := fmt.Sprintf("%c", p[r-1])
		mated := false
		for l := 1; l <= len(s); l++ {
			last := false
			if pi == "*" {
				last = dp[r-1][l]
			} else {
				last = dp[r-1][l-1]
			}
			if pi == "*" {
				if last || mated {
					dp[r][l] = true
					mated = true
				}
			} else if pi == "?" {
				if last {
					dp[r][l] = true
				}
			} else {
				if last && s[l-1] == p[r-1] {
					dp[r][l] = true
				}
			}
		}
	}
	return dp[len(p)][len(s)]
}

func charMatch(u, v byte) bool {
	return u == v || v == '?'
}
