package main

import (
	"fmt"
)

func main() {
	requests := [][]int{{0, 1},
		{1, 0},
		{0, 1},
		{1, 2},
		{2, 0},
		{3, 4}}
	n := 5 // 5
	requests = [][]int{{2, 0},
		{0, 4},
		{3, 1},
		{2, 4},
		{1, 3},
		{4, 3},
		{0, 1},
		{4, 2}}
	n = 5 // 5
	i := maximumRequests(n, requests)
	fmt.Println(i)
}

/*
图
	将楼看作顶点，搬楼的请求看作有权有向边
	只有在环内的顶点才能搬楼
难点
	0 -> 1 -> 2 -> 0：权重分别是 100,1,99
	答案是 3，问题是 3 如何计算得来？
	又如：0 -> 1 -> 2 -> 0 且 0 -> 3 -> 4 -> 0，权重分别是 100,1,1 和 1,1,100，又如何计算？
	再如：0 -> 1 -> 0 且 0 -> 2 -> 3 -> 0，各种权重情况下，又如何计算？
	假如 n=5 且每两个顶点都互相连通，那又如何计算？
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumRequests(n int, requests [][]int) int {

	// 个人：没完成
	//g, vis := make([][]int, n), make([]int, n)
	//for i := range g {
	//	g[i] = make([]int, n)
	//}
	//for _, r := range requests { // 有权有向边
	//	g[r[0]][r[1]]++
	//}
	//uni := make([]int, n) // 并查集：合并环内的楼
	//for i := range uni {
	//	uni[i] = i
	//}
	//var find func(int) int
	//find = func(i int) int {
	//	if uni[i] != i {
	//		uni[i] = find(uni[i])
	//	}
	//	return uni[i]
	//}
	//union := func(p, q int) {
	//	uni[find(q)] = find(p)
	//}
	//var dfs func(int) int
	//dfs = func(i int) int {
	//	if vis[i] == 1 {
	//		return -1
	//	} else if vis[i] == -1 { // 环开始
	//		return i
	//	}
	//	vis[i] = -1 // 访问中
	//	c := -1
	//	for j, v := range g[i] {
	//		if v > 0 {
	//			c = dfs(j)
	//			if i == c { // 闭环
	//				c = -1
	//			}
	//			if c != -1 { // 标记环已闭
	//				union(c, i)
	//			}
	//		}
	//	}
	//	vis[i] = 1 // 已访问
	//	return c
	//}
	//for i := 0; i < n; i++ {
	//	if vis[i] == 0 {
	//		dfs(i)
	//	}
	//}
	//out, in := make([]int, n), make([]int, n)
	//for i := range g {
	//	for j, v := range g[i] {
	//		if v > 0 {
	//			if find(i) != find(j) {
	//				g[i][j] = 0 // 环外的请求全部作废
	//			} else {
	//				out[i] += v // 统计入/出度
	//				in[j] += v
	//			}
	//		}
	//	}
	//}
	////for i, d := range g {
	////	fmt.Println(i, d)
	////}
	////fmt.Println(uni)
	////fmt.Println(out)
	////fmt.Println(in)
	//arr := make([]*hp1601, n)
	//for i := range uni { // 将每个环内的顶点放到一个堆中
	//	idx := find(i)
	//	if arr[idx] == nil {
	//		arr[idx] = &hp1601{}
	//	}
	//	heap.Push(arr[idx], [2]int{i, out[i]})
	//}
	//for _, h := range arr {
	//	if h == nil {
	//		continue
	//	}
	//	// 写不下去了
	//}
	//var ans int
	//return ans
}

type hp1601 [][2]int

func (h hp1601) Len() int           { return len(h) }
func (h hp1601) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h hp1601) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp1601) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *hp1601) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)

//func maximumRequests(n int, requests [][]int) int {
//	//	var ans int
//	//out:
//	//	for mask := 1<<len(requests) - 1; mask > 0; mask-- { // 枚举
//	//		cnt := bits.OnesCount(uint(mask))
//	//		if cnt <= ans {
//	//			continue
//	//		}
//	//		delta := make([]int, n)
//	//		for bit := mask; bit > 0; bit &= bit - 1 { // 满足被标记的请求
//	//			idx := bits.TrailingZeros(uint(bit & -bit))
//	//			x, y := requests[idx][0], requests[idx][1]
//	//			delta[x]--
//	//			delta[y]++
//	//		}
//	//		for _, v := range delta { // check 是否成功
//	//			if v != 0 {
//	//				continue out
//	//			}
//	//		}
//	//		ans = cnt
//	//	}
//	//	return ans
//
//	m := len(requests)
//	var ans, cnt, zero int
//	delta := make([]int, n)
//	var dfs func(int)
//	dfs = func(i int) {
//		if i == m {
//			if zero == 0 {
//				ans = max(ans, cnt)
//			}
//			return
//		}
//		dfs(i + 1)  // 不选
//		tag := zero // zero 标记所有楼房是否“出入均衡”
//		cnt++
//		x, y := requests[i][0], requests[i][1]
//		if delta[x] == 0 {
//			zero--
//		}
//		delta[x]--
//		if delta[x] == 0 {
//			zero++
//		}
//		if delta[y] == 0 {
//			zero--
//		}
//		delta[y]++
//		if delta[y] == 0 {
//			zero++
//		}
//		dfs(i + 1) // 选
//		delta[x]++ // 回溯
//		delta[y]--
//		cnt--
//		zero = tag
//	}
//	dfs(0)
//	return ans
//}
