package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type TripleInOne struct {
	stack []int
	idx   [3]int
	cap   int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{stack: make([]int, stackSize*3),
		idx: [3]int{0, stackSize, stackSize << 1},
		cap: stackSize,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	i := this.idx[stackNum]
	if i == this.cap*(stackNum+1) {
		return
	}
	this.stack[i] = value
	this.idx[stackNum]++
}

func (this *TripleInOne) Pop(stackNum int) int {
	v := this.Peek(stackNum)
	if v != -1 {
		this.idx[stackNum]--
	}
	return v
}

func (this *TripleInOne) Peek(stackNum int) int {
	i := this.idx[stackNum]
	if i == this.cap*(stackNum) {
		return -1
	}
	return this.stack[i-1]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.idx[stackNum] == this.cap*(stackNum)
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
//leetcode submit region end(Prohibit modification and deletion)
