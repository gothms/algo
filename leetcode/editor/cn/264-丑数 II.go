package main

import (
	"fmt"
	"math"
)

func main() {
	n := 10
	number := nthUglyNumber(n)
	fmt.Println(number)
}

// leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int {
	nums := [3]int{2, 3, 5}
	idx := make([]int, len(nums))
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
		for j, k := range idx {
			dp[i] = min(dp[i], dp[k]*nums[j])
		}
		for j, k := range idx {
			if dp[k]*nums[j] == dp[i] {
				idx[j]++
			}
		}
	}
	return dp[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func nthUglyNumber_(n int) int {
	// dp
	nums := []int{2, 3, 5}
	dp := make([]int, n)
	dp[0] = 1
	ids := make([]int, len(nums))
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
		for j, k := range ids {
			dp[i] = min(dp[i], nums[j]*dp[k])
		}
		for j, k := range ids {
			if dp[i] == nums[j]*dp[k] {
				ids[j]++
			}
		}
	}
	return dp[n-1]

	// 堆
	//nums, memo := []int{2, 3, 5}, make(map[int]bool)
	//h := &hp{1}
	//for k := 0; k < n-1; k++ {
	//	top := heap.Pop(h).(int)
	//	for _, v := range nums {
	//		if val := v * top; !memo[val] {
	//			heap.Push(h, val)
	//			memo[val] = true
	//		}
	//	}
	//}
	//return heap.Pop(h).(int)
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] < h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.(int)) }
func (h *hp) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}
