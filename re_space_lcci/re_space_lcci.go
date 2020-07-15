package re_space_lcci

func respace(dictionary []string, sentence string) int {
	if len(dictionary) <= 0 {
		return len(sentence)
	}
	dictSearch := DictSearch{
		d: dictionary,
	}
	dictSearch.init()

	n := len(sentence)
	dp := []int{0}
	for i := 1; i <= n; i++ {
		dp = append(dp, dp[i-1]+1)
		for j := 0; j < i; j++ {
			ok := dictSearch.Contain(sentence[j:i])
			if ok {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[n]
}

type DictSearch struct {
	d     []string
	rTrie TrieTree
}

type TrieTree struct {
	d      uint8
	c      []*TrieTree
	isWord bool
}

func (this *DictSearch) init() {
	root := TrieTree{
		d:      0,
		c:      []*TrieTree{},
		isWord: false,
	}
	for _, v := range this.d {
		cNot := &root
		for i := 0; i < len(v); i++ {
			hasNode := false
			if len(cNot.c) > 0 {
				for _, n := range cNot.c {
					if n.d == v[i] {
						cNot = n
						hasNode = true
						break
					}
				}
			}
			if !hasNode {
				newNode := TrieTree{
					d: v[i],
					c: []*TrieTree{},
				}
				cNot.c = append(cNot.c, &newNode)
				cNot = &newNode
			}
		}
		cNot.isWord = true
	}
	this.rTrie = root
}

func (this *DictSearch) Contain(text string) bool {
	cNot := &this.rTrie
	for i := 0; i < len(text); i++ {
		if cNot == nil || len(cNot.c) <= 0 {
			return false
		}
		hasSearch := false
		for j := 0; j < len(cNot.c); j++ {
			if cNot.c[j].d == text[i] {
				cNot = cNot.c[j]
				hasSearch = true
				break
			}
		}
		if !hasSearch {
			return false
		}
	}
	return cNot == nil || cNot.isWord
}

func min(a, b int) (m int) {
	if a < b {
		m = a
	} else {
		m = b
	}
	return
}
