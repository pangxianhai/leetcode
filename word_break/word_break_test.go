package word_break

import (
    "fmt"
    "strings"
    "testing"
)

func TestWordBreakII(t *testing.T) {
    //t.Log(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
    //t.Log(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
    //t.Log(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))

    t.Log(wordBreak("aaaaa", []string{"a", "aa", "aaa"}))
}

func TestWordBreak(t *testing.T) {
    var res []string

    res = wordBreakII("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
    t.Log(strings.Join(res, ","))

    res = wordBreakII("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})
    t.Log("res=", strings.Join(res, ","))

    res = wordBreak1("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"})
    t.Log("res=", strings.Join(res, ","))

    //t.Log(wordBreak("aaaa", []string{"a"}))
}

func wordBreak1(s string, wordDict []string) (sentences []string) {
    wordSet := map[string]struct{}{}
    for _, w := range wordDict {
        wordSet[w] = struct{}{}
    }

    n := len(s)
    dp := make([][][]string, n)
    var backtrack func(index int) [][]string
    backtrack = func(index int) [][]string {
        if dp[index] != nil {
            return dp[index]
        }
        wordsList := [][]string{}
        for i := index + 1; i < n; i++ {
            word := s[index:i]
            if _, has := wordSet[word]; has {
                for _, nextWords := range backtrack(i) {
                    wordsList = append(wordsList, append([]string{word}, nextWords...))
                }
            }
        }
        word := s[index:]
        if _, has := wordSet[word]; has {
            wordsList = append(wordsList, []string{word})
        }
        dp[index] = wordsList
        return wordsList
    }
    for _, words := range backtrack(0) {
        sentences = append(sentences, strings.Join(words, " "))
    }
    return
}

func TestA(t *testing.T) {
    wordSet := map[string]struct{}{}
    wordSet["aa"] = struct{}{}
    wordSet["ab"] = struct{}{}
    fmt.Println(wordSet)
}
