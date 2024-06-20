package main

import (
	"math/rand"
	"slices"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type Solution struct {
	arr, temp []int
}

func Constructor(nums []int) Solution {
	return Solution{nums, slices.Clone(nums)}
}

func (this *Solution) Reset() []int {
	return this.arr
}

func (this *Solution) Shuffle() []int {
	// 洗牌算法
	n := len(this.temp)
	for i := range this.temp {
		j := rand.Intn(n-i) + i
		if i != j {
			this.temp[i], this.temp[j] = this.temp[j], this.temp[i]
		}
	}
	return this.temp
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
//leetcode submit region end(Prohibit modification and deletion)
