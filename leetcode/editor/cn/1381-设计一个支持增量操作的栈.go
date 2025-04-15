package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type CustomStack struct {
	st  []int
	inc []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{make([]int, 0, maxSize), make([]int, maxSize)}
}

func (this *CustomStack) Push(x int) {
	if len(this.st) < cap(this.st) {
		this.st = append(this.st, x)
	}
}

func (this *CustomStack) Pop() int {
	if len(this.st) == 0 {
		return -1
	}
	i := len(this.st) - 1
	v := this.st[i]
	this.st = this.st[:i]

	incV := this.inc[i]
	this.inc[i] = 0
	if i > 0 {
		this.inc[i-1] += incV
	}
	return v + incV
}

func (this *CustomStack) Increment(k int, val int) {
	if i := min(k, len(this.st)); i > 0 {
		this.inc[i-1] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
//leetcode submit region end(Prohibit modification and deletion)
