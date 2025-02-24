package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type OrderedStream struct {
	m   []string
	ptr int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{make([]string, n+1), 1}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.m[idKey] = value
	i := this.ptr
	for this.ptr < len(this.m) && this.m[this.ptr] != "" {
		this.ptr++
	}
	return this.m[i:this.ptr]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
