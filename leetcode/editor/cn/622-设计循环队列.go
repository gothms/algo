package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type MyCircularQueue struct {
	cq   []int
	h, t int
	k    int
}

func Constructor(k int) MyCircularQueue {
	k++
	return MyCircularQueue{make([]int, k), 0, 0, k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.cq[this.t] = value
	this.t = (this.t + 1) % this.k
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.h = (this.h + 1) % this.k
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.cq[this.h]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.cq[(this.t-1+this.k)%this.k]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.h == this.t
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.t+1)%this.k == this.h
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
//leetcode submit region end(Prohibit modification and deletion)
