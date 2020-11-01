package string_to_integer_atoi

import (
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
