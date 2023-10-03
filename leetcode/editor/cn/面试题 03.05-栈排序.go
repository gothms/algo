//栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。该栈支持如下操作：
//push、pop、peek 和 isEmpty。当栈为空时，peek 返回 -1。
//
// 示例1:
//
//  输入：
//["SortedStack", "push", "push", "peek", "pop", "peek"]
//[[], [1], [2], [], [], []]
// 输出：
//[null,null,null,1,null,2]
//
//
// 示例2:
//
//  输入：
//["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
//[[], [], [], [1], [], []]
// 输出：
//[null,null,null,null,null,true]
//
//
// 说明:
//
//
// 栈中的元素数目在[0, 5000]范围内。
//
//
// Related Topics 栈 设计 单调栈 👍 102 👎 0

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
	down []int // 降序
	up   []int // 升序
}

func Constructor() SortedStack {
	return SortedStack{}
}
func (this *SortedStack) Push(val int) {
	move(&this.down, &this.up, func(v int) bool { // 移掉小值
		return val > v
	})
	move(&this.up, &this.down, func(v int) bool { // 移掉大值
		return val < v
	})
	this.down = append(this.down, val)
}
func (this *SortedStack) Pop() { // 查并删
	if this.IsEmpty() {
		return
	}
	move(&this.up, &this.down, func(v int) bool {
		return true
	})
	this.down = this.down[:len(this.down)-1]
}
func (this *SortedStack) Peek() int { // 只查不删
	if this.IsEmpty() {
		return -1
	}
	move(&this.up, &this.down, func(v int) bool {
		return true
	})
	return this.down[len(this.down)-1]
}
func (this *SortedStack) IsEmpty() bool {
	return len(this.down)+len(this.up) == 0
}
func move(f, t *[]int, fn func(v int) bool) {
	for i := len(*f) - 1; i >= 0 && fn((*f)[i]); i-- {
		*t = append(*t, (*f)[i])
		*f = (*f)[:i]
	}
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
