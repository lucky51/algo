package leetcode1206

import "math/rand"

const maxLevel = 32
const pFactor = 0.25

type SkiplistNode struct {
	val     int
	forward []*SkiplistNode
}

// Skiplist 设计跳跃表
type Skiplist struct {
	head  *SkiplistNode
	level int
}

func Constructor() Skiplist {
	return Skiplist{&SkiplistNode{-1, make([]*SkiplistNode, maxLevel)}, 0}
}

// randomLevel 随机生成 1 ~ maxlevel 之间的数
//  1/2 的概率返回 2
//  1/4 的概率返回 3
//  1/8 分之一的概率返回 4
// https://leetcode.cn/problems/design-skiplist/solution/she-ji-tiao-biao-by-capital-worker-3vqk/
func (Skiplist) randomLevel() int {
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}

// Search 搜索元素
func (this *Skiplist) Search(target int) bool {
	curr := this.head
	for i := this.level - 1; i >= 0; i-- {
		// 同新增，找到第i层小于且接近target的元素
		for curr.forward[i] != nil && curr.forward[i].val < target {
			curr = curr.forward[i]
		}
	}
	// 已经在第一层  ? 这个位置为什么要查到最后一层才结束可不可以提前终止
	curr = curr.forward[0]
	// 当前元素的值是否等于 target
	return curr != nil && curr.val == target
}

// Add 添加元素
func (this *Skiplist) Add(num int) {
	// 用来记录更新的位置
	update := make([]*SkiplistNode, maxLevel)
	for i := range update {
		update[i] = this.head
	}
	curr := this.head
	// 遍历每一层，找到小于且接近num的元素
	for i := this.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	lv := this.randomLevel()
	// 更新level的最大值
	this.level = max(lv, this.level)
	// 创建新节点
	newNode := &SkiplistNode{num, make([]*SkiplistNode, lv)}
	// 更新小于lv的所有层
	for i, node := range update[:lv] {
		// 对于第 i层的状态进行更新，将当前元素的 forward指向新的节点
		// 如官方图示在插入 23数字的时候， update 先记录了所有小于23的每一层前驱元素
		// 然后将新节点23的下一个元素指向这个前驱元素的下一个节点
		// 最后把前驱元素的下一个节点指向新节点，那么这样每一层的新节点就添加完毕
		newNode.forward[i] = node.forward[i]
		node.forward[i] = newNode
	}
}
func (this *Skiplist) Erase(num int) bool {
	// 同 add方法
	update := make([]*SkiplistNode, maxLevel)
	curr := this.head
	for i := this.level - 1; i >= 0; i-- {
		// 找到 i层小于且最接近num的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	curr = curr.forward[0]
	// 如果值不存在则返回false
	if curr == nil || curr.val != num {
		return false
	}
	// 找到目标值，遍历目标值对应的所有层
	for i := 0; i < this.level && update[i].forward[i] == curr; i++ {
		// 对第 i 层的状态进行更新，将forward 指向被删除节点的下一跳
		update[i].forward[i] = curr.forward[i]
	}
	// 更新当前的level
	for this.level > 1 && this.head.forward[this.level-1] == nil {
		this.level--
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
