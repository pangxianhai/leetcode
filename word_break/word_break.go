package word_break

import (
    "strings"
)

func wordBreak(s string, wordDict []string) bool {
    dp := make([][]int, len(s)+1)
    for i := 0; i < len(s)+1; i++ {
        dp[i] = make([]int, len(s)+1)
    }
    wordBreakB(dp, s, 0, len(s), wordDict)
    //fmt.Println(dp)
    //fmt.Println(dp[0][len(s)])
    return dp[0][len(s)] == 1
}

func wordBreakB(dp [][]int, source string, start int, end int, wordDict []string) {
    if start > end {
        dp[start][end] = -1
        return
    }
    if start == end {
        dp[start][end] = 1
        return
    }
    if dp[start][end] != 0 {
        return
    }
    s := source[start:end]
    for i := 0; i < len(wordDict); i++ {
        w := wordDict[i]
        arr := strings.Split(s, w)
        if len(arr) == 0 {
            dp[start][end] = 1
            return
        }
        empty := true
        for j := 0; j < len(arr); j++ {
            if len(arr[j]) > 0 {
                empty = false
                break
            }
        }
        if empty {
            dp[start][end] = 1
            return
        }
        for j := 0; j < len(arr)-1; j++ {
            ts := strings.Join(arr[0:j+1], w)
            sp := len(ts) + len(w)
            //fmt.Println(s+":"+source[start:start+sp-len(w)], " "+w+" "+source[start+sp:end])
            if sp <= end {
                if dp[start][start+sp-len(w)] == 0 {
                    wordBreakB(dp, source, start, start+sp-len(w), wordDict)
                }
                if dp[start+sp][end] == 0 {
                    wordBreakB(dp, source, start+sp, end, wordDict)
                }
                if dp[start][start+sp-len(w)] == 1 && dp[start+sp][end] == 1 {
                    dp[start][end] = 1
                    return
                }
            } else {
                //fmt.Println("sp=", sp)
            }
        }
        dp[start][end] = -1
    }
}

func wordBreakII(s string, wordDict []string) []string {
    resList := wordBreakA("", s, wordDict)
    for i := 0; i < len(resList); i++ {
        resList[i] = strings.TrimSpace(resList[i])
    }
    return resList
}

func wordBreakA(source string, s string, wordDict []string) []string {
    resList := make([]string, 0)

    for _, w := range wordDict {
        if strings.HasPrefix(s, w) {
            ts := source + " " + w
            ls := strings.Replace(s, w, "", 1)
            if len(ls) > 0 {
                resList1 := wordBreakA(ts, ls, wordDict)
                if len(resList1) > 0 {
                    for _, r := range resList1 {
                        resList = append(resList, r)
                    }
                }
            } else {
                resList = append(resList, ts)
            }
        }
    }

    return resList
}
