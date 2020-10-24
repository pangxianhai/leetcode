package string_to_integer_atoi

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	t.Log(myAtoi("42"))
	t.Log(myAtoi("   -42"))
	t.Log(myAtoi("4193 with words"))
	t.Log(myAtoi("words and 987"))
	t.Log(myAtoi("-91283472332"))
	t.Log(myAtoi("+1"))
	t.Log(myAtoi("21474836460"))
}

func TestA(t *testing.T) {

	status := map[string][]string{
		"start":     {"start", "signed", "in_number", "end"},
		"signed":    {"end", "end", "in_number", "end"},
		"in_number": {"end", "end", "in_number", "end"},
		"end":       {"end", "end", "end", "end"},
	}




	fmt.Println(status)
}
