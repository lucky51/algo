package leetcode676

type MagicDictionary struct {
	dict []string
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.dict = dictionary
}

func (this *MagicDictionary) Search(searchWord string) bool {
loop:
	for _, word := range this.dict {
		if len(searchWord) != len(word) {
			continue
		}
		diff := false
		for j := range word {
			if word[j] != searchWord[j] {
				if !diff {
					diff = true
				} else {
					continue loop
				}
			}
		}
		if diff {
			return true
		}
	}
	return false
}
