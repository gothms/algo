package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)

// lc
type MinStack struct {
	st  []int
	mst []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]int, 0), []int{math.MaxInt64}}
}

func (this *MinStack) Push(x int) {
	this.st = append(this.st, x)
	this.mst = append(this.mst, min(this.mst[len(this.mst)-1], x))
}

func (this *MinStack) Pop() {
	this.st = this.st[:len(this.st)-1]
	this.mst = this.mst[:len(this.mst)-1]
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
	return this.mst[len(this.mst)-1]
}

// 个人
//type MinStack struct {
//	st  []int
//	mSt []int // 存的是 st 的索引
//}
//
///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{make([]int, 0), make([]int, 0)}
//}
//
//func (this *MinStack) Push(x int) {
//	i := len(this.st)
//	this.st = append(this.st, x)
//	if i == 0 || x < this.st[this.mSt[len(this.mSt)-1]] {
//		this.mSt = append(this.mSt, i)
//	}
//}
//
//func (this *MinStack) Pop() {
//	i := len(this.st) - 1
//	//if i >= 0 {
//	this.st = this.st[:i]
//	if this.mSt[len(this.mSt)-1] == i {
//		this.mSt = this.mSt[:len(this.mSt)-1]
//	}
//	//}
//}
//
//func (this *MinStack) Top() int {
//	return this.st[len(this.st)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	//if len(this.mSt) > 0 {
//	return this.st[this.mSt[len(this.mSt)-1]]
//	//}
//	//return 0
//}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
