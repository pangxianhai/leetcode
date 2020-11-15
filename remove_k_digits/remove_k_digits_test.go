package remove_k_digits

import (
    "strconv"
    "testing"
)

func TestRemoveKdigits(t *testing.T) {
    var res string

    res = removeKdigits("1432219", 3)
    t.Log("1219", res)

    res = removeKdigits("10200", 1)
    t.Log("200", res)

    res = removeKdigits("10", 2)
    t.Log("0", res)

    res = removeKdigits("112", 1)
    t.Log("11", res)

    res = removeKdigits("1234567890", 9)
    t.Log("0", res)
}

func TestA(t *testing.T) {
    t.Log(strconv.Atoi("0200"))
    s := "1439991"

    t.Log(s[0 : len(s)-2])

}
