package leetcode208

// 实现 Trie （前缀树)  TODO:未实现
type Trie struct {
	prefixMap map[string]bool
}

func Constructor() Trie {
	return Trie{
		prefixMap: map[string]bool{},
	}
}

func (this *Trie) Insert(word string) {
	for i := len(word); i > 0; i-- {
		this.prefixMap[word[:i]] = i == len(word)
	}
}

func (this *Trie) Search(word string) bool {
	full, ok := this.prefixMap[word]
	return ok && full
}

func (this *Trie) StartsWith(prefix string) bool {
	_, ok := this.prefixMap[prefix]
	return ok
}
