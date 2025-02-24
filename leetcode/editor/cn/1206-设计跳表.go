package main

import "math/rand"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)

const (
	maxLevel         = 32
	p        float64 = 0.5
)

type Skiplist struct {
	level int
	head  *SkiplistNode
}
type SkiplistNode struct {
	val  int
	next []*SkiplistNode
}

func Constructor() Skiplist {
	return Skiplist{0, &SkiplistNode{-1, make([]*SkiplistNode, maxLevel)}}
}
func (this *Skiplist) Search(target int) bool {
	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < target {
			cur = cur.next[i]
		}
	}
	cur = cur.next[0]
	return cur != nil && cur.val == target
}
func (this *Skiplist) Add(num int) {
	update := make([]*SkiplistNode, maxLevel)
	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		update[i] = cur
	}
	lv := this.randomLevel()
	if lv > this.level {
		for i := this.level; i < lv; i++ {
			update[i] = this.head
		}
		this.level = lv
	}
	newNode := &SkiplistNode{num, make([]*SkiplistNode, maxLevel)}
	for i := 0; i < this.level; i++ {
		update[i].next[i], newNode.next[i] = newNode, update[i].next[i]
	}
}
func (this *Skiplist) Erase(num int) bool {
	update := make([]*SkiplistNode, maxLevel)
	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		update[i] = cur
	}
	cur = cur.next[0]
	if cur == nil || cur.val != num {
		return false
	}
	for i := 0; i < this.level && update[i].next[i] == cur; i++ {
		update[i].next[i] = cur.next[i]
	}
	for this.level > 1 && this.head.next[this.level-1] == nil {
		this.level--
	}
	return true
}

func (this *Skiplist) randomLevel() int {
	level := 1
	for level < maxLevel && rand.Float64() < p {
		level++
	}
	return level
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
//leetcode submit region end(Prohibit modification and deletion)

//type SkiplistNode struct {
//	value int
//	next  []*SkiplistNode
//}
//
//type Skiplist struct {
//	head  *SkiplistNode
//	level int
//}
//
//func Constructor() Skiplist {
//	return Skiplist{
//		level: 0,
//		head: &SkiplistNode{
//			value: -1 << 31, // 头节点存储最小整数值
//			next:  make([]*SkiplistNode, maxLevel),
//		},
//	}
//}
//
//func (s *Skiplist) randomLevel() int {
//	level := 1
//	for rand.Float64() < p && level < s.level {
//		level++
//	}
//	return level
//}
//
//func (s *Skiplist) Add(num int) {
//	update := make([]*SkiplistNode, maxLevel)
//	current := s.head
//
//	// 从最高层开始查找插入位置
//	for i := s.level - 1; i >= 0; i-- {
//		for current.next[i] != nil && current.next[i].value < num {
//			current = current.next[i]
//		}
//		update[i] = current
//	}
//
//	// 检查是否已存在
//	//if current.next[0] != nil && current.next[0].value == num {
//	//	return
//	//}
//
//	// 生成新节点层数
//	level := s.randomLevel()
//	if level > s.level {
//		for i := s.level; i < level; i++ {
//			update[i] = s.head
//		}
//		s.level = level
//	}
//
//	// 创建新节点
//	newNode := &SkiplistNode{
//		value: num,
//		next:  make([]*SkiplistNode, level),
//	}
//
//	// 更新各层指针
//	for i := 0; i < level; i++ {
//		newNode.next[i], update[i].next[i] = update[i].next[i], newNode
//	}
//}
//
//func (s *Skiplist) Search(num int) bool {
//	current := s.head
//	for i := s.level - 1; i >= 0; i-- {
//		for current.next[i] != nil && current.next[i].value < num {
//			current = current.next[i]
//		}
//	}
//	current = current.next[0]
//	return current != nil && current.value == num
//}
//
//func (s *Skiplist) Erase(num int) bool {
//	update := make([]*SkiplistNode, s.level)
//	current := s.head
//
//	// 查找需要更新的节点
//	for i := s.level - 1; i >= 0; i-- {
//		for current.next[i] != nil && current.next[i].value < num {
//			current = current.next[i]
//		}
//		update[i] = current
//	}
//
//	// 定位目标节点
//	target := current.next[0]
//	if target == nil || target.value != num {
//		return false
//	}
//
//	// 更新各层指针
//	for i := 0; i < s.level; i++ {
//		if update[i].next[i] != target {
//			break
//		}
//		update[i].next[i] = target.next[i]
//	}
//
//	// 调整当前层数
//	for s.level > 1 && s.head.next[s.level-1] == nil {
//		s.level--
//	}
//	return true
//}
