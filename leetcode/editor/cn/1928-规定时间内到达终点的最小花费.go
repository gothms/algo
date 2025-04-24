package main

import (
	"fmt"
	"math"
)

func main() {
	edges := [][]int{{0, 1, 10},
		{1, 2, 10},
		{2, 5, 10},
		{0, 3, 1},
		{3, 4, 10},
		{4, 5, 15}}
	passingFees := []int{5, 1, 2, 20, 20, 3}
	maxTime := 29 // 48
	cost := minCost(maxTime, edges, passingFees)
	fmt.Println(cost)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minCost(maxTime int, edges [][]int, passingFees []int) int {
	// 优化
	// 但`new_time > maxTime`舍弃超时方案的判断，只有在用时接近maxTime时才会频繁生效，而如果maxTime设置得比较大，那么很多最终会超时的方案，在初期都会被当成合理方案处理
	// 如果能够在早期访问某个节点x时，就知道当前方案最终会不会超时，那就会减少大量无意义的遍历
	// 所以需要提前知晓每个节点x到n-1的最短时间。可以n-1为起点，不考虑花费问题，预处理最短时间
	//n := len(passingFees)
	//var (
	//	g       = make([][][2]int, n)
	//	minCost = make([]int, n)
	//	cost    = make([]int, n) // only time
	//	h       = &hp{}          // i fee time
	//)
	//for _, e := range edges {
	//	x, y, t := e[0], e[1], e[2]
	//	g[x] = append(g[x], [2]int{y, t})
	//	g[y] = append(g[y], [2]int{x, t})
	//}
	//for i := range cost {
	//	minCost[i], cost[i] = math.MaxInt32, math.MaxInt32
	//}
	//minCost[n-1], cost[0] = 0, 0
	//
	//heap.Push(h, []int{n - 1, 0})
	//for h.Len() > 0 {
	//	cur := heap.Pop(h).([]int)
	//	x, time := cur[0], cur[1]
	//	for _, p := range g[x] {
	//		y, t := p[0], p[1]
	//		if ty := time + t; ty < minCost[y] {
	//			minCost[y] = ty
	//			heap.Push(h, []int{y, ty})
	//		}
	//	}
	//}
	//
	//clear(*h)
	//heap.Push(h, []int{0, passingFees[0], 0})
	//for h.Len() > 0 {
	//	cur := heap.Pop(h).([]int)
	//	x, fee, time := cur[0], cur[1], cur[2]
	//	if x == n-1 {
	//		return fee
	//	}
	//	for _, p := range g[x] {
	//		y, t := p[0], p[1]
	//		if ty := time + t; ty+minCost[y] <= maxTime && ty < cost[y] {
	//			cost[y] = ty
	//			heap.Push(h, []int{y, fee + passingFees[y], ty})
	//		}
	//	}
	//}
	//return -1

	// lc：dijkstra
	// 假设策略是优先走用时最短的节点，那么每次经过n-1时，还要对当前花费做最小化处理，那么就要考虑到所有能够抵达n-1的方案，这样毫无疑问会非常慢
	// 假设策略是优先走花费最少的节点，那么当第一次抵达n-1时，只要没有超时，就是答案了，这样只需要考虑一种抵达n-1的方案即可，就会加速很多
	//n := len(passingFees)
	//var (
	//	g    = make([][][2]int, n)
	//	cost = make([]int, n) // only time
	//	h    = &hp{}          // i fee time
	//)
	//for _, e := range edges {
	//	x, y, t := e[0], e[1], e[2]
	//	g[x] = append(g[x], [2]int{y, t})
	//	g[y] = append(g[y], [2]int{x, t})
	//}
	//for i := range cost {
	//	cost[i] = math.MaxInt32
	//}
	//
	//// 1.以支付的通行费为优先策略，那么最先到达 n-1 的路线即为最优路线
	//// 2.以花费的时间为入队列的条件，那么假设有 a b 都能满足的两个方案，有没有可能是 a 先加入队列呢？不可能的
	//// a 方案：花费时间更少，但是通行费更多
	//// b 方案：花费时间更多，但是通行费更少，在条件 1 情况下 b 肯定优先级更高
	//// 3.不适用 visited 来判断
	//heap.Push(h, [3]int{0, passingFees[0], 0})
	//for h.Len() > 0 {
	//	cur := heap.Pop(h).([3]int)
	//	x, fee, time := cur[0], cur[1], cur[2]
	//	if x == n-1 {
	//		return fee
	//	}
	//	for _, p := range g[x] {
	//		y, t := p[0], p[1]
	//		if ty := time + t; ty <= maxTime && ty < cost[y] {
	//			cost[y] = ty
	//			heap.Push(h, [3]int{y, fee + passingFees[y], ty})
	//		}
	//	}
	//}
	//return -1

	// dp：优化
	n := len(passingFees)
	dp := make([][]int, maxTime+1)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = passingFees[0]
	for i := 1; i <= maxTime; i++ {
		for _, e := range edges {
			f, t, time := e[0], e[1], e[2]
			if i >= time {
				dp[i][t] = min(dp[i][t], dp[i-time][f]+passingFees[t])
				dp[i][f] = min(dp[i][f], dp[i-time][t]+passingFees[f])
			}
		}
	}
	ans := math.MaxInt32
	for _, d := range dp {
		ans = min(ans, d[n-1])
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans

	// dp
	//n := len(passingFees)
	//g := make([][][2]int, n)
	//for _, e := range edges {
	//	x, y, t := e[0], e[1], e[2]
	//	g[x] = append(g[x], [2]int{y, t})
	//	g[y] = append(g[y], [2]int{x, t})
	//}
	//dp := make([][]int, maxTime+1)
	//for i := range dp {
	//	dp[i] = make([]int, n)
	//	for j := range dp[i] {
	//		dp[i][j] = math.MaxInt32
	//	}
	//}
	//dp[0][0] = passingFees[0]
	//for i := 1; i <= maxTime; i++ {
	//	for j := 1; j < n; j++ {
	//		for _, p := range g[j] {
	//			k, t := p[0], p[1]
	//			if i >= t {
	//				dp[i][j] = min(dp[i][j], dp[i-t][k]+passingFees[j])
	//			}
	//		}
	//	}
	//}
	//ans := math.MaxInt32
	//for _, d := range dp {
	//	ans = min(ans, d[n-1])
	//}
	//if ans == math.MaxInt32 {
	//	return -1
	//}
	//return ans
}

// 优化
//type hp [][]int
//
//func (h hp) Len() int           { return len(h) }
//func (h hp) Less(i, j int) bool { return h[i][1] < h[j][1] }
//func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//func (h *hp) Push(x any)        { *h = append(*h, x.([]int)) }
//func (h *hp) Pop() any {
//	v := (*h)[h.Len()-1]
//	*h = (*h)[:h.Len()-1]
//	return v
//}

// lc
type hp [][3]int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.([3]int)) }
func (h *hp) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
