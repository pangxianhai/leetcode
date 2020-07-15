package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen <= 1 {
		return sLen
	}
	r := 0
	maxLen := 0
	for i := 1; i < sLen; i++ {
		for j := r; j < i; j++ {
			if s[j] == s[i] {
				if i-r > maxLen {
					maxLen = i - r
				}
				r = j + 1
				break
			}
		}
	}
	if sLen-r > maxLen {
		return sLen - r
	} else {
		return maxLen
	}
}
