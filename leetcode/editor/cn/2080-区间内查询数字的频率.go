package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type RangeFreqQuery struct {
	m map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	m := make(map[int][]int)
	for i, v := range arr {
		m[v] = append(m[v], i)
	}
	return RangeFreqQuery{m}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	return sort.SearchInts(this.m[value], right+1) - sort.SearchInts(this.m[value], left)
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
