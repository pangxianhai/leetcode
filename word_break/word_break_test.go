package word_break

import (
    "fmt"
    "testing"
)

func TestWordBreakII(t *testing.T) {
    //t.Log(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
    //t.Log(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
    //t.Log(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))

    t.Log(wordBreak("aaaaa", []string{"a", "aa", "aaa"}))
}

func TestWordBreak(t *testing.T) {
    //t.Log(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
    t.Log(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
    //t.Log(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))

    //t.Log(wordBreak("aaaa", []string{"a"}))
}

func TestA(t *testing.T) {
    s := []string{"a", "b"}
    fmt.Println(s[0:1])
}
