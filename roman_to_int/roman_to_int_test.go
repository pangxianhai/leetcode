package roman_to_int

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	t.Log("III=", romanToInt("III"))
	t.Log("IV=", romanToInt("IV"))
	t.Log("IX=", romanToInt("IX"))
	t.Log("LVIII=", romanToInt("LVIII"))
	t.Log("MCMXCIV=", romanToInt("MCMXCIV"))
}
