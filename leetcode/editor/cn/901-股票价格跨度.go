package main

import (
	"math"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type StockSpanner struct {
	st [][2]int
	i  int
}

func Constructor() StockSpanner {
	return StockSpanner{[][2]int{{-1, math.MaxInt32}}, 0}
}
func (this *StockSpanner) Next(price int) int {
	st := this.st
	for price >= st[len(st)-1][1] {
		st = st[:len(st)-1]
	}
	ans := this.i - st[len(st)-1][0]
	this.st = append(st, [2]int{this.i, price})
	this.i++
	return ans
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
//leetcode submit region end(Prohibit modification and deletion)
