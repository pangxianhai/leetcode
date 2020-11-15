/* 给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。

num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。

输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。

输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。 */

package remove_k_digits

import (
    "strconv"
    "strings"
)

func removeKdigits(num string, k int) string {
    n := len(num)
    if n <= k {
        return "0"
    }
    res := "0"
    for i := 0; i < n; i++ {
        ni := num[i : i+1]
        nii, _ := strconv.Atoi(ni)
        for k > 0 && len(res) > 0 {
            last, _ := strconv.Atoi(res[len(res)-1:])
            if last > nii {
                res = res[0 : len(res)-1]
                k--
            } else {
                break
            }
        }
        res += ni
        if k <= 0 {
            res += num[i+1:]
            break
        }
    }
    if k != 0 {
        res = res[:len(res)-k]
    }
    ans := strings.TrimLeft(res, "0")
    if len(ans) <= 0 {
        return "0"
    }
    return ans
}
