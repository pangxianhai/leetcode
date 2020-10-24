package string_to_integer_atoi

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	status := map[string][]string{
		"start":     {"start", "signed", "in_number", "end"},
		"signed":    {"end", "end", "in_number", "end"},
		"in_number": {"end", "end", "in_number", "end"},
		"end":       {"end", "end", "end", "end"},
	}
	state := "start"
	ans := 0
	signed := 1
	for i := 0; i < len(s); i++ {
		col := getCol(s[i])
		state = status[state][col]
		if state == "in_number" {
			tmp := ans*10 + int(s[i]-uint8('0'))
			if tmp > math.MaxInt32 {
				if signed == 1 {
					ans = math.MaxInt32

				} else {
					ans = math.MaxInt32 + 1
				}
				state = "end"
			} else {
				ans = tmp
			}
		} else if state == "signed" {
			if s[i] == uint8('-') {
				signed = -1
			}
		} else if state == "start" {
			continue
		} else if state == "end" {
			break
		}
	}
	return ans * signed
}

func getCol(c uint8) int {
	if c >= uint8('0') && c <= uint8('9') {
		return 2
	} else if c == uint8('-') || c == uint8('+') {
		return 1
	} else if c == uint8(' ') {
		return 0
	} else {
		return 3
	}
}

func myAtoi1(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	isNegaive := false
	var num int = 0
	for i := 0; i < len(s); i++ {
		if i == 0 && (s[i] == uint8('-') || s[i] == uint8('+')) {
			if s[i] == uint8('-') {
				isNegaive = true
			} else {
				isNegaive = false
			}
		} else if s[i] >= uint8('0') && s[i] <= uint8('9') {
			tmp := num*10 + int(s[i]-uint8('0'))
			if tmp > math.MaxInt32 {
				if isNegaive {
					num = math.MaxInt32 + 1
				} else {
					num = math.MaxInt32
				}
				break
			} else {
				num = tmp
			}
		} else {
			break
		}
	}
	if isNegaive {
		return 0 - num
	} else {
		return num
	}
}
