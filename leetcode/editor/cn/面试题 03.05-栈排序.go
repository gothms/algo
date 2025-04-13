package main

func main() {
	ss := Constructor()
	ss.Push(1)
	ss.Push(2)
	ss.Peek()
	ss.Pop()
	ss.Peek()
}

// leetcode submit region begin(Prohibit modification and deletion)
type SortedStack struct {
}

func Constructor() SortedStack {
}

func (this *SortedStack) Push(val int) {

}

func (this *SortedStack) Pop() {
}

func (this *SortedStack) Peek() int {
}

func (this *SortedStack) IsEmpty() bool {
}

/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */
//leetcode submit region end(Prohibit modification and deletion)

//type SortedStack struct {
//	down []int // 降序
//	up   []int // 升序
//}
//
//func Constructor() SortedStack {
//	return SortedStack{}
//}
//func (this *SortedStack) Push(val int) {
//	move(&this.down, &this.up, func(v int) bool { // 移掉小值
//		return val > v
//	})
//	move(&this.up, &this.down, func(v int) bool { // 移掉大值
//		return val < v
//	})
//	this.down = append(this.down, val)
//}
//func (this *SortedStack) Pop() { // 查并删
//	if this.IsEmpty() {
//		return
//	}
//	move(&this.up, &this.down, func(v int) bool {
//		return true
//	})
//	this.down = this.down[:len(this.down)-1]
//}
//func (this *SortedStack) Peek() int { // 只查不删
//	if this.IsEmpty() {
//		return -1
//	}
//	move(&this.up, &this.down, func(v int) bool {
//		return true
//	})
//	return this.down[len(this.down)-1]
//}
//func (this *SortedStack) IsEmpty() bool {
//	return len(this.down)+len(this.up) == 0
//}
//func move(f, t *[]int, fn func(v int) bool) {
//	for i := len(*f) - 1; i >= 0 && fn((*f)[i]); i-- {
//		*t = append(*t, (*f)[i])
//		*f = (*f)[:i]
//	}
//}
