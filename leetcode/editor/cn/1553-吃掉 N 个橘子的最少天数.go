package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	var n int = 2e9
	days := minDays(n)
	fmt.Println(days)

	fmt.Println(math.Log(1))
}

//leetcode submit region begin(Prohibit modification and deletion)

func minDays(n int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

// const n1553 = 1e6
//
// var md1553 [n1553]uint8
//
//	func init() {
//		md1553[1] = 1
//		for i := 2; i < n1553; i++ {
//			//j := i % 3
//			//md1553[i] = min(uint8(i&1)+md1553[i>>1], uint8(j)+md1553[(i-j)/3]) + 1
//			md1553[i] = min(uint8(i&1)+md1553[i>>1], uint8(i%3)+md1553[i/3]) + 1
//		}
//	}
func minDays_(n int) int {
	// 启发式搜索
	cost := func(rest int) int {
		if rest == 0 {
			return 0
		}
		return int(math.Log(float64(rest))/math.Log(3)) + 1
	}
	h := &hp1553{{cost(n), 0, n}}
	memo := make(map[int]struct{})
	var ans int
	for {
		t := heap.Pop(h).(star1553)
		day, rest := t.days, t.rest
		if _, ok := memo[rest]; ok {
			continue
		}
		if rest == 1 {
			ans = day + 1
			break
		}
		cur := day + rest&1 + 1
		heap.Push(h, star1553{cur + cost(rest>>1), cur, rest >> 1})
		cur = day + rest%3 + 1
		heap.Push(h, star1553{cur + cost(rest/3), cur, rest / 3})
	}
	return ans

	// Dijkstra

	// 记忆化搜索：最快
	//memo := make(map[int]uint8) // 使用 uint8 优化空间
	//var dfs func(int) uint8
	//dfs = func(i int) uint8 {
	//	if i < n1553 {
	//		return md1553[i]
	//	} else if memo[i] > 0 {
	//		return memo[i]
	//	}
	//	v := min(uint8(i&1)+dfs(i>>1), uint8(i%3)+dfs(i/3)) + 1
	//	memo[i] = v
	//	return v
	//}
	//return int(dfs(n))
}

type star1553 struct {
	cost, days, rest int
}
type hp1553 []star1553

func (h hp1553) Len() int           { return len(h) }
func (h hp1553) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h hp1553) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp1553) Push(x any)        { *h = append(*h, x.(star1553)) }
func (h *hp1553) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}
