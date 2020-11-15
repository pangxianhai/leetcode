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
    dp := make([][][]string, len(s)+1)
    for i := 0; i < len(s)+1; i++ {
        dp[i] = make([][]string, len(s)+1)
    }
    wordBreakIIB(dp, s, 0, len(s), wordDict)
    return dp[0][len(s)]
}

func wordBreakIIB(dp [][][]string, source string, start int, end int, wordDict []string) {
    //dp[i][j]==nil 表示还未计算
    if start >= end {
        dp[start][end] = make([]string, 0)
        return
    }
    if dp[start][end] != nil {
        return
    }
    s := source[start:end]
    for i := 0; i < len(wordDict); i++ {
        w := wordDict[i]
        if s == w {
            dp[start][end] = append(dp[start][end], w)
            AddValue(dp, start, end, w)
            continue
        }
        arr := strings.Split(s, w)
        if len(arr) == 1 {
            continue
        }
        for j := 0; j < len(arr)-1; j++ {
            ts := strings.Join(arr[0:j+1], w)
            sp := len(ts) + len(w)
            if sp > end || start+sp > end {
                continue
            }
            if start+sp-len(w) != end && dp[start][start+sp-len(w)] == nil {
                wordBreakIIB(dp, source, start, start+sp-len(w), wordDict)
            }
            if start+sp != start && dp[start+sp][end] == nil {
                wordBreakIIB(dp, source, start+sp, end, wordDict)
            }
            if start+sp-len(w) > start && start+sp < end {
                for t1 := 0; t1 < len(dp[start][start+sp-len(w)]); t1++ {
                    for t2 := 0; t2 < len(dp[start+sp][end]); t2++ {
                        tmp := dp[start][start+sp-len(w)][t1] + " " + w + " " + dp[start+sp][end][t2]
                        AddValue(dp, start, end, tmp)
                    }
                }
            } else if start+sp-len(w) > start {
                for t1 := 0; t1 < len(dp[start][start+sp-len(w)]); t1++ {
                    tmp := dp[start][start+sp-len(w)][t1] + " " + w
                    AddValue(dp, start, end, tmp)
                }
            } else if start+sp < end {
                for t2 := 0; t2 < len(dp[start+sp][end]); t2++ {
                    tmp := w + " " + dp[start+sp][end][t2]
                    AddValue(dp, start, end, tmp)
                }
            }
        }
    }
    if dp[start][end] == nil {
        dp[start][end] = make([]string, 0)
    }
}

func AddValue(dp [][][]string, start int, end int, w string) {
    if dp[start][end] == nil {
        dp[start][end] = make([]string, 0)
    }
    for i := 0; i < len(dp[start][end]); i++ {
        if dp[start][end][i] == w {
            return
        }
    }
    dp[start][end] = append(dp[start][end], w)
}

func wordBreakIII(s string, wordDict []string) (sentences []string) {
    wordSet := map[string]struct{}{}
    for _, word := range wordDict {
        wordSet[word] = struct{}{}
    }
    n := len(s)
    dp := make([][][]string, n)
    var backtrack func(index int) [][]string
    backtrack = func(index int) [][]string {
        if dp[index] != nil {
            return dp[index]
        }
        wordList := [][]string{}
        for i := index + 1; i < n; i++ {
            word := s[index:i]
            if _, has := wordSet[word]; has {
                for _, nextWord := range backtrack(i) {
                    wordList = append(wordList, append([]string{word}, nextWord...))
                }
            }
        }
        word := s[index:]
        if _, has := wordSet[word]; has {
            wordList = append(wordList, []string{word})
        }
        dp[index] = wordList
        return wordList
    }
    for _, wordList := range backtrack(0) {
        sentences = append(sentences, strings.Join(wordList, " "))
    }
    return
}
