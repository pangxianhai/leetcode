package minimum_deletion_cost_to_avoid_repeating_letters

import (
    "fmt"
    "testing"
)

func TestMinCost(t *testing.T) {
    s := "aabbcc"
    c := []int{0,1,2,3,4,5}
    r := MinCost(s, c)
    fmt.Println("结果:", r)
}

