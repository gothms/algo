//堆盘子。设想有一堆盘子，堆太高可能会倒下来。因此，在现实生活中，盘子堆到一定高度时，我们就会另外堆一堆盘子。请实现数据结构SetOfStacks，模拟这种行
//为。SetOfStacks应该由多个栈组成，并且在前一个栈填满时新建一个栈。此外，SetOfStacks.push()和SetOfStacks.pop()应该与
//普通栈的操作方法相同（也就是说，pop()返回的值，应该跟只有一个栈时的情况一样）。 进阶：实现一个popAt(int index)方法，根据指定的子栈，执行
//pop操作。
//
// 当某个栈为空时，应当删除该栈。当栈中没有元素或不存在该栈时，pop，popAt 应返回 -1.
//
// 示例1:
//
//  输入：
//["StackOfPlates", "push", "push", "popAt", "pop", "pop"]
//[[1], [1], [2], [1], [], []]
// 输出：
//[null, null, null, 2, 1, -1]
//
//
// 示例2:
//
//  输入：
//["StackOfPlates", "push", "push", "push", "popAt", "popAt", "popAt"]
//[[2], [1], [2], [3], [0], [0], [0]]
// 输出：
//[null, null, null, null, 2, 1, 3]
//
//
// Related Topics 栈 设计 链表 👍 58 👎 0

package main

import "fmt"

func main() {
	arr := [][]int{{1, 2}}
	fmt.Printf("%p\n", arr[0])
	t := arr[0]
	arr[0] = t[:len(t)-1]
	fmt.Println(arr)
	fmt.Printf("%p\n", arr[0])
}

// leetcode submit region begin(Prohibit modification and deletion)
type StackOfPlates struct {
	sop [][]int
	cap int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{cap: cap}
}
func (this *StackOfPlates) Push(val int) {
	if this.cap == 0 { // cap 可能被设置为 0
		return
	}
	last := len(this.sop) - 1
	if last == -1 || len(this.sop[last]) == this.cap { // 增加栈
		this.sop = append(this.sop, make([]int, 0, this.cap))
		last++
	}
	this.sop[last] = append(this.sop[last], val) // 增加元素
}
func (this *StackOfPlates) Pop() int {
	if this.cap == 0 || len(this.sop) == 0 { // 空
		return -1
	}
	return this.PopAt(len(this.sop) - 1)
}
func (this *StackOfPlates) PopAt(index int) int {
	if this.cap == 0 || index >= len(this.sop) { // 越界
		return -1
	}
	t := this.sop[index] // 目标栈
	val := t[len(t)-1]
	if len(t) == 1 {
		this.sop = append(this.sop[:index], this.sop[index+1:]...) // 删除栈
	} else {
		this.sop[index] = t[:len(t)-1] // 删除元素
	}
	return val
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
//leetcode submit region end(Prohibit modification and deletion)
