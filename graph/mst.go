package graph

import (
	"container/heap"
	"fmt"
	"math"
)

/*
https://oi-wiki.org/graph/mst/

最小生成树
	连通图：
		若无向图 G = (V, E)，满足其中任意两个顶点均连通，则称 G 是 连通图 (connected graph)
		若一张有向图的节点两两互相可达，则称这张图是 强连通的 (strongly connected)
		若一张有向图的边替换为无向边后可以得到一张连通图，则称原来这张有向图是 弱连通的 (weakly connected)
	连通分量：
		若 H 是 G 的一个连通子图，且不存在 F 满足 H 不包含于 F 包含于 G 且 F 为连通图
		则 H 是 G 的一个 连通块/连通分量 (connected component)（极大连通子图）
	生成子图：对一张图 G = (V, E)，若存在另一张图 H = (V', E') 满足 V' 包含于 V 且 E' 包含于 E，则称 H 是 G 的 子图 (subgraph)
		若 H 包含于 G 满足 V' = V，则称 H 为 G 的 生成子图/支撑子图 (spanning subgraph)
	生成树：一个连通无向图的生成子图，同时要求是树。也即在图的边集中选择 n - 1 条，将所有顶点连通
Kruskal 算法
	Kruskal 算法是一种常见并且好写的最小生成树算法，由 Kruskal 发明
	该算法的基本思想是从小到大加入边，是个贪心算法

	具体来说，维护一个森林，查询两个结点是否在同一棵树中，连接两棵树
	抽象一点地说，维护一堆 集合，查询两个元素是否属于同一集合，合并两个集合
	其中，查询两点是否连通和连接两点可以使用并查集维护
	如果使用 O(m\log m) 的排序算法，并且使用 O(m*max(m, n)) 或 O(m*log n) 的并查集，就可以得到时间复杂度为 O(m*log m) 的 Kruskal 算法
	// TODO
Prim 算法
	Prim 算法是另一种常见并且好写的最小生成树算法
	该算法的基本思想是从一个结点开始，不断加点（而不是 Kruskal 算法的加边）
	// TODO
Boruvka 算法
	该算法的思想是前两种算法的结合
	它可以用于求解无向图的最小生成森林（无向连通图就是最小生成树）
	// TODO

最小生成树的唯一性
次小生成树
瓶颈生成树
最小瓶颈路
Kruskal 重构树

lc
	1584
	1489
*/

// Kruskal 算法，参考自 ChatGPT
func Kruskal(n int, edges [][]int) int {
	uf := newUnionFind(n)
	totalWeight := 0
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if uf.union(u, v) {
			totalWeight += w
		}
	}
	return totalWeight
}

type unionFind struct {
	parent []int
	rank   []int
}

func newUnionFind(n int) *unionFind {
	uf := &unionFind{
		make([]int, n),
		make([]int, n),
	}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
}
func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}
func (uf *unionFind) union(x, y int) bool {
	rootX, rootY := uf.find(x), uf.find(y)
	if rootX == rootY {
		return false
	}
	if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	return true
}

// Prim 算法
func Prim(n int, edges [][]int) {
	var (
		adj = make([][][2]int, n)
		dis = make([]int, n)
		vis = make([]bool, n)
		hp  = &hpPrim{{}}
	)
	var (
		cnt int
		res int
	)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u], adj[v] = append(adj[u], [2]int{v, w}), append(adj[v], [2]int{u, w})
	} // 必须双向
	for i := range dis {
		dis[i] = math.MaxInt32
	}
	dis[0] = 0

	for hp.Len() > 0 {
		if cnt >= n {
			break
		}
		u := heap.Pop(hp).([2]int)[0]
		if vis[u] {
			continue
		}
		vis[u] = true
		cnt++ // 重要
		res += dis[u]
		for _, e := range adj[u] {
			v, w := e[0], e[1]
			if w < dis[v] { // 贪心
				dis[v] = w
				heap.Push(hp, [2]int{v, w})
			}
		}
	}
	if cnt == n {
		fmt.Println(res)
	} else {
		fmt.Println("No MST.")
	}
}

type hpPrim [][2]int

func (h hpPrim) Len() int           { return len(h) }
func (h hpPrim) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h hpPrim) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hpPrim) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *hpPrim) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}
