package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	for i := 0; i < 10; i++ {
		v := rand.Intn(4)
		fmt.Println(v)
	}
	w := []int{1, 3}
	con := Constructor(w)
	con.PickIndex()
}

// leetcode submit region begin(Prohibit modification and deletion)
type Solution struct {
	pre []int
	sum int
}

func Constructor(w []int) Solution {
	pre := make([]int, len(w)+1)
	for i, v := range w {
		pre[i+1] = pre[i] + v
	}
	return Solution{pre, pre[len(w)]}
}

func (this *Solution) PickIndex() int {
	//v := rand.Intn(this.sum)
	//return sort.Search(len(this.pre), func(i int) bool { return this.pre[i] > v }) - 1
	return sort.SearchInts(this.pre, rand.Intn(this.sum)+1) - 1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
//leetcode submit region end(Prohibit modification and deletion)
