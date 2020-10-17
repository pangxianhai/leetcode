package longest_palindromic_substring

func longestPalindrome(s string) string {
	L := len(s)
	dp := make([][]int, L)
	for i := 0; i < L; i++ {
		dp[i] = make([]int, L)
	}
	ans := ""
	for l := 0; l < L; l++ {
		for i := 0; i < L; i++ {
			j := i + l
			if j >= L {
				break
			}
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if len(ans) < l+1 && dp[i][j] == 1 {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}
