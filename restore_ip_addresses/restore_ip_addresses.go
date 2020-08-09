package restore_ip_addresses

import (
	"strconv"
)

var ans []string
var segments [4]int

func restoreIpAddresses(s string) []string {
	ans = make([]string, 0)
	defs(s, 0, 0)
	return ans
}

func defs(s string, sid int, start int) {
	if sid == 4 {
		if start == len(s) {
			ipAddr := ""
			for i := 0; i < 4; i++ {
				ipAddr += strconv.Itoa(segments[i])
				if i != 3 {
					ipAddr += "."
				}
			}
			ans = append(ans, ipAddr)
		}
		return
	}
	if start == len(s) {
		return
	}
	if s[start] == '0' {
		segments[sid] = 0
		defs(s, sid+1, start+1)
		return
	}

	for i := start + 1; i <= len(s); i++ {
		sv, _ := strconv.Atoi(s[start:i])
		if sv > 0 && sv <= 255 {
			segments[sid] = sv
			defs(s, sid+1, i)
		} else {
			break
		}
	}
}
