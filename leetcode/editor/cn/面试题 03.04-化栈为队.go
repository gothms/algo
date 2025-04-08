package main

import "slices"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type MyQueue struct {
	in  []int
	out []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{[]int{}, []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	v := this.Peek()
	if v >= 0 {
		this.out = this.out[:len(this.out)-1]
	}
	return v
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		this.move()
	}
	if len(this.out) == 0 {
		return -1
	}
	return this.out[len(this.out)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.out) == 0 && len(this.in) == 0
}

func (this *MyQueue) move() {
	for _, v := range slices.Backward(this.in) {
		this.out = append(this.out, v)
	}
	this.in = this.in[:0] // 清空
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)
