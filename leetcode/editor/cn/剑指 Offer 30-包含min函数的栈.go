//定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
//
//
//
// 示例:
//
// MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.min();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.min();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// 各函数的调用总次数不超过 20000 次
//
//
//
//
// 注意：本题与主站 155 题相同：https://leetcode-cn.com/problems/min-stack/
//
// Related Topics 栈 设计 👍 495 👎 0

package main

func main() {

}

/*
两个栈：
	一个按顺序记录进栈的数
	一个记录当前最小的数（数量和第一个栈相等）
*/

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	stack []int
	min   []int
	last  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{last: -1}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if this.last >= 0 && x > this.min[this.last] {
		x = this.min[this.last]
	}
	this.min = append(this.min, x)
	this.last++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.last]
	this.min = this.min[:this.last]
	this.last--
}

func (this *MinStack) Top() int {
	return this.stack[this.last]
}

func (this *MinStack) Min() int {
	return this.min[this.last]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
//leetcode submit region end(Prohibit modification and deletion)
