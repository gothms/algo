package main

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{{0, 0},
		{1, 1},
		{1, 0},
		{-1, 1}} // 4
	points = [][]int{{-8, -3},
		{-16, -17},
		{11, 14},
		{-11, 5},
		{-7, -20},
		{-12, -11},
		{-18, -2},
		{8, -19},
		{9, 14},
		{10, -15},
		{-17, 15},
		{9, 5},
		{1, -18},
		{-20, -2},
		{0, 12},
		{12, -3},
		{3, -8}} // 157
	connectPoints := minCostConnectPoints(points)
	fmt.Println(connectPoints)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minCostConnectPoints(points [][]int) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	// 建图优化的 Kruskal：https://leetcode.cn/problems/min-cost-to-connect-all-points/solutions/565801/lian-jie-suo-you-dian-de-zui-xiao-fei-yo-kcx7/

	// kruskal
	n := len(points)
	edges := make([][3]int, 0, n*(n-1)>>1)
	for i, l := range points {
		for j, r := range points[i+1:] {
			edges = append(edges, [3]int{i, j + i + 1, abs(r[0]-l[0]) + abs(r[1]-l[1])})
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })
	uf := newUnionFind1574(n)
	ans := 0
	for _, e := range edges {
		if uf.unite(e[0], e[1]) {
			ans += e[2]
			if uf.cnt == 1 {
				break
			}
		}
	}
	return ans

	// prim
	//n := len(points)
	//adj := make([][][2]int, n)
	//for i := range adj {
	//	adj[i] = make([][2]int, 0, n-1)
	//}
	//for i, l := range points {
	//	for j, r := range points[i+1:] {
	//		k := j + i + 1
	//		manhattan := abs(r[0]-l[0]) + abs(r[1]-l[1])
	//		adj[i], adj[k] = append(adj[i], [2]int{k, manhattan}), append(adj[k], [2]int{i, manhattan})
	//	}
	//}
	//dis, vis := make([]int, n), make([]bool, n)
	//for i := range dis {
	//	dis[i] = math.MaxInt32
	//}
	//dis[0] = 0
	//
	//h := &hp1574{{}}
	//for ans, cnt := 0, 0; ; {
	//	if cnt >= n {
	//		return ans
	//	}
	//	u := heap.Pop(h).([2]int)[0]
	//	if vis[u] {
	//		continue
	//	}
	//	vis[u] = true
	//	ans += dis[u]
	//	cnt++
	//	for _, e := range adj[u] {
	//		v, w := e[0], e[1]
	//		if w < dis[v] {
	//			dis[v] = w
	//			heap.Push(h, [2]int{v, w})
	//		}
	//	}
	//}
}

type UnionFind1574 struct {
	p    []int
	size []int
	cnt  int
}

func newUnionFind1574(n int) *UnionFind1574 {
	p, size := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		size[i] = 1
	}
	return &UnionFind1574{p, size, n}
}

func (uf *UnionFind1574) find(x int) int {
	if uf.p[x] != x {
		uf.p[x] = uf.find(uf.p[x])
	}
	return uf.p[x]
}
func (uf *UnionFind1574) unite(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.p[fy] = fx
	uf.size[fx] += uf.size[fy]
	uf.cnt--
	return true
}

type hp1574 [][2]int

func (h hp1574) Len() int           { return len(h) }
func (h hp1574) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h hp1574) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp1574) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *hp1574) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
