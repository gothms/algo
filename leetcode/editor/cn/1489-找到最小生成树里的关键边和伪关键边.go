package main

import (
	"fmt"
	"sort"
)

func main() {
	edges := [][]int{{0, 1, 1},
		{1, 2, 1},
		{0, 2, 1},
		{2, 3, 4},
		{3, 4, 2},
		{3, 5, 2},
		{4, 5, 2}}
	n := 6 // [[3],[0,1,2,4,5,6]]
	criticalEdges := findCriticalAndPseudoCriticalEdges(n, edges)
	fmt.Println(criticalEdges)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	// Kruskal：涉及 tarjan 算法，还没懂
	// https://leetcode.cn/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/
	m := len(edges)
	edgeType := make([]int, m) // -1：不在最小生成树中；0：伪关键边；1：关键边

	for i, e := range edges {
		edges[i] = append(e, i)
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

	type neighbor struct{ to, edgeID int }
	graph := make([][]neighbor, n)
	dfn := make([]int, n) // 遍历到该顶点时的时间戳
	timestamp := 0
	var tarjan func(int, int) int
	tarjan = func(v, pid int) int {
		timestamp++
		dfn[v] = timestamp
		lowV := timestamp
		for _, e := range graph[v] {
			if w := e.to; dfn[w] == 0 {
				lowW := tarjan(w, e.edgeID)
				if lowW > dfn[v] {
					edgeType[e.edgeID] = 1
				}
				lowV = min(lowV, lowW)
			} else if e.edgeID != pid {
				lowV = min(lowV, dfn[w])
			}
		}
		return lowV
	}

	uf := newUnionFind(n)
	for i := 0; i < m; {
		vs := []int{}
		// 将权值相同的边分为一组，建图，然后用 Tarjan 算法找桥边
		for weight := edges[i][2]; i < m && edges[i][2] == weight; i++ {
			e := edges[i]
			v, w, edgeID := uf.find(e[0]), uf.find(e[1]), e[3]
			if v != w {
				graph[v] = append(graph[v], neighbor{w, edgeID})
				graph[w] = append(graph[w], neighbor{v, edgeID})
				vs = append(vs, v, w) // 记录图中顶点
			} else {
				edgeType[edgeID] = -1
			}
		}
		for _, v := range vs {
			if dfn[v] == 0 {
				tarjan(v, -1)
			}
		}
		// 合并顶点、重置数据
		for j := 0; j < len(vs); j += 2 {
			v, w := vs[j], vs[j+1]
			uf.union(v, w)
			graph[v] = nil
			graph[w] = nil
			dfn[v] = 0
			dfn[w] = 0
		}
	}

	var keyEdges, pseudokeyEdges []int
	for i, tp := range edgeType {
		if tp == 0 {
			pseudokeyEdges = append(pseudokeyEdges, i)
		} else if tp == 1 {
			keyEdges = append(keyEdges, i)
		}
	}
	return [][]int{keyEdges, pseudokeyEdges}

	// prim
	//var (
	//	adj = make([][][3]int, n)
	//	dis []int
	//	vis []bool
	//	hp  *hp1489
	//)
	//for i, e := range edges {
	//	u, v, w := e[0], e[1], e[2]
	//	adj[u], adj[v] = append(adj[u], [3]int{v, w, i}), append(adj[v], [3]int{u, w, i})
	//}
	//prim := func(idx int) (int, int) { // prim 算法
	//	var (
	//		cnt int
	//		res int
	//	)
	//	dis = make([]int, n)
	//	vis = make([]bool, n)
	//	hp = &hp1489{{}}
	//	for i := range dis {
	//		dis[i] = math.MaxInt32
	//	}
	//	dis[0] = 0
	//
	//	for hp.Len() > 0 {
	//		if cnt >= n {
	//			break
	//		}
	//		u := heap.Pop(hp).([2]int)[0]
	//		if vis[u] {
	//			continue
	//		}
	//		vis[u] = true
	//		cnt++ // 重要
	//		res += dis[u]
	//		for _, e := range adj[u] {
	//			v, w, i := e[0], e[1], e[2]
	//			if i != idx && w < dis[v] { // 贪心
	//				dis[v] = w
	//				heap.Push(hp, [2]int{v, w})
	//			}
	//		}
	//	}
	//	return cnt, res
	//}
	//_, mst := prim(-1) // 最小生成树
	//ans := make([][]int, 2)
	//for i, e := range edges {
	//	c, r := prim(i)       // 不选择边 i
	//	if c < n || r > mst { // 没有最小生成树 | 关键边
	//		ans[0] = append(ans[0], i)
	//	} else {
	//		u, v := e[0], e[1]
	//		adj[u], adj[v] = append(adj[u], [3]int{v, 0, i}), append(adj[v], [3]int{u, 0, i}) // 模拟 0 权的边
	//		c, r = prim(-1)
	//		if r+e[2] == mst { // 必须选择边 i，如果 == mst 则是伪关键边，否则不是
	//			ans[1] = append(ans[1], i)
	//		}
	//		adj[u], adj[v] = adj[u][:len(adj[u])-1], adj[v][:len(adj[v])-1] // 将模拟边删掉
	//	}
	//}
	//return ans
}

type unionFind struct {
	parent, size []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent, size}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
	return true
}

type hp1489 [][2]int

func (h hp1489) Len() int           { return len(h) }
func (h hp1489) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h hp1489) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp1489) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *hp1489) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
