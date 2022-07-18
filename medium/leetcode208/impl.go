package leetcode208

// 实现 Trie （前缀树)
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (this *Trie) Search(word string) bool {
	if t := this.SearchPrefix(word); t != nil && t.isEnd {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	if t := this.SearchPrefix(prefix); t != nil {
		return true
	}
	return false
}
