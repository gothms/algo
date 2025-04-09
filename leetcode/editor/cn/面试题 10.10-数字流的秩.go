package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"math"
)

func main() {
	fmt.Println(math.Floor(3))
	fmt.Println(math.Ceil(3))
}

// leetcode submit region begin(Prohibit modification and deletion)
type StreamRank struct {
	b []int
}

func (this *StreamRank) update(v int) {
	for i := v + 1; i < len(this.b); i += i & -i {
		this.b[i]++
	}
}
func (this *StreamRank) prefix(v int) int {
	sum := 0
	for i := v + 1; i > 0; i &= i - 1 {
		sum += this.b[i]
	}
	return sum
}

func Constructor() StreamRank {
	return StreamRank{make([]int, 50002)}
}

func (this *StreamRank) Track(x int) {
	this.update(x)
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	return this.prefix(x)
}

/**
 * Your StreamRank object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Track(x);
 * param_2 := obj.GetRankOfNumber(x);
 */
//leetcode submit region end(Prohibit modification and deletion)

//type StreamRank struct {
//	rbt *redblacktree.Tree
//}
//
//func Constructor() StreamRank {
//	return StreamRank{rbt: redblacktree.NewWithIntComparator()}
//}
//
//func (this *StreamRank) Track(x int) {
//	this.rbt.Put(x, struct{}{})
//}
//
//func (this *StreamRank) GetRankOfNumber(x int) int {
//	right, _ := this.rbt.Ceiling(x )
//	fmt.Println(this.rbt.Size(), right.Size())
//	return this.rbt.Size() - right.Size()
//}
