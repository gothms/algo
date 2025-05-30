package main

import "math/rand"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	m := make(map[int][]int)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}
	return Solution{m}
}

func (this *Solution) Pick(target int) int {
	idx := this.m[target]
	i := rand.Intn(len(idx))
	return idx[i]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
//leetcode submit region end(Prohibit modification and deletion)
