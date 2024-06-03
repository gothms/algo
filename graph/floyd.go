package graph

import (
	"fmt"
	"math"
)

/*
https://oi-wiki.org/graph/shortest-path/#floyd-%E7%AE%97%E6%B3%95

Floyd 算法
	是用来求任意两个结点之间的最短路的
	复杂度比较高，但是常数小，容易实现（只有三个 for）
	适用于任何图，不管有向无向，边权正负，但是最短路必须存在。（不能有个负环）
	定义
		f[k][x][y]，表示只允许经过结点 1 到 k（也就是说，在子图 V'={1, 2, ..., k} 中的路径，注意，x 与 y 不一定在这个子图中），结点 x 到结点 y 的最短路长度
		很显然，f[n][x][y] 就是结点 x 到结点 y 的最短路长度（因为 V'={1, 2, ..., n} 即为 V 本身，其表示的最短路径就是所求路径）
		f[0][x][y]：x 与 y 的边权，或者 0，或者 +∞
			当 x 与 y 间有直接相连的边的时候，为它们的边权
			当 x = y 的时候为零，因为到本身的距离为零
			当 x 与 y 没有直接相连的边的时候，为 +∞
	f[k][x][y] = min(f[k-1][x][y], f[k-1][x][k]+f[k-1][k][y])
		（f[k-1][x][y]，为不经过 k 点的最短路径，而 f[k-1][x][k]+f[k-1][k][y]，为经过了 k 点的最短路）
BellmanFord 算法
	一种基于松弛（relax）操作的最短路算法，可以求出有负权的图的最短路，并可以对最短路不存在的情况进行判断
	SPFA（Shortest Path Faster Algorithm），就是 Bellman–Ford 算法的一种实现
	Bellman–Ford 算法要用到的松弛操作（Dijkstra 算法也会用到松弛操作）
		对于边 (u,v)，松弛操作对应的式子：dis(v) = min(dis(v), dis(u) + w(u, v))
	每次循环是 O(m) 的，那么最多会循环多少次呢？
		在最短路存在的情况下，由于一次松弛操作会使最短路的边数至少 +1，而最短路的边数最多为 n-1
		因此整个算法最多执行 n-1 轮松弛操作。故总时间复杂度为 O(nm)
	负环
		如果从 S 点出发，抵达一个负环时，松弛操作会无休止地进行下去
		对于最短路存在的图，松弛操作最多只会执行 n-1 轮，因此如果第 n 轮循环时仍然存在能松弛的边，说明从 S 点出发，能够抵达一个负环
	负环判断中存在的常见误区
		需要注意的是，以 S 点为源点跑 Bellman–Ford 算法时，如果没有给出存在负环的结果，只能说明从 S 点出发不能抵达一个负环，而不能说明图上不存在负环
		因此如果需要判断整个图上是否存在负环，最严谨的做法是建立一个超级源点，向图上每个节点连一条权值为 0 的边，然后以超级源点为起点执行 Bellman–Ford 算法
	队列优化：SPFA
		func BellmanFord(n int, edges [][]int)
	Bellman–Ford 的其他优化
		堆优化：将队列换成堆，与 Dijkstra 的区别是允许一个点多次入队。在有负权边的图可能被卡成指数级复杂度
		栈优化：将队列换成栈（即将原来的 BFS 过程变成 DFS），在寻找负环时可能具有更高效率，但最坏时间复杂度仍然为指数级
		LLL 优化：将普通队列换成双端队列，每次将入队结点距离和队内距离平均值比较，如果更大则插入至队尾，否则插入队首
		SLF 优化：将普通队列换成双端队列，每次将入队结点距离和队首比较，如果更大则插入至队尾，否则插入队首
		D´Esopo–Pape 算法：将普通队列换成双端队列，如果一个节点之前没有入队，则将其插入队尾，否则插入队首
		更多优化以及针对这些优化的 Hack 方法，参考：https://www.zhihu.com/question/292283275/answer/484871888
Dijkstra 算法
	// TODO

Johnson 全源最短路径算法
	// TODO

不同方法的比较
	最短路算法	Floyd				Bellman–Ford	Dijkstra	Johnson
	最短路类型	每对结点之间的最短路	单源最短路		单源最短路	每对结点之间的最短路
	作用于		任意图				任意图			非负权图		任意图
	能否检测负环？	能					能				不能			能
	时间复杂度	O(N^3)				O(N*M)			O(M*log M)	O(NM*log M)

	总结
		在没有负权边时最好使用 Dijkstra 算法
		在有负权边且题目中的图没有特殊性质时，若 SPFA 是标算的一部分，题目不应当给出 Bellman–Ford 算法无法通过的数据范围

Floyd 应用
	给一个正权无向图，找一个最小权值和的环
		考虑环上编号最大的结点 u
		f[u-1][x][y] 和 (u,x),(u,y) 共同构成了环
	已知一个有向图中任意两点之间是否有连边，要求判断任意两点是否连通
		该问题即是求 图的传递闭包
		只需要按照 Floyd 的过程，逐个加入点判断一下
		只是此时的边的边权变为 1/0，而取 \min 变成了 或 运算
		再进一步用 bitset 优化，复杂度可以到 O(n^3/w)
			// std::bitset<SIZE> f[SIZE];
			for (k = 1; k <= n; k++)
			  for (i = 1; i <= n; i++)
				if (f[i][k]) f[i] = f[i] | f[k];

lc
	847
*/

// BellmanFord 算法
func BellmanFord(n int, edges [][]int) {
	const inf = math.MaxInt32
	var (
		adj   = make([][][2]int, n)
		dis   = make([]int, n)
		cnt   = make([]int, n)
		vis   = make([]bool, n)
		queue []int
	)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2] // w 为权值
		adj[x], adj[y] = append(adj[x], [2]int{y, w}), append(adj[y], [2]int{x, w})
	}
	for i := range dis {
		dis[i] = inf
	}
	spfa := func(s int) bool {
		queue = []int{s}
		dis[s] = 0
		vis[s] = true
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, e := range adj[u] {
				v, w := e[0], e[1]
				if dis[v] > dis[u]+w { // spfa：更新距离 & 检测负权环
					dis[v] = dis[u] + w
					cnt[v] = cnt[u] + 1 // 记录最短路经过的边数
					if cnt[v] >= n {    // 在不经过负环的情况下，最短路至多经过 n-1 条边
						return false // 因此如果经过了多于 n 条边，一定说明经过了负环
					}
					if !vis[v] {
						queue = append(queue, v)
						vis[v] = true
					}
				}
			}
		}
		return true
	}
	// test
	if spfa(0) {
		fmt.Println("No negative weight cycle detected")
		fmt.Println(dis)
	} else {
		fmt.Println("Negative weight cycle detected")
	}
}

// floydON2 Floyd 算法
func floydON2(n int, graph [][]int) {
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	// ...
	for k := 1; k <= n; k++ {
		for x := 1; x <= n; x++ {
			for y := 1; y <= n; y++ {
				f[x][y] = min(f[x][y], f[x][k]+f[k][y])
			}
		}
	}
}

func floydON3(n int, graph [][]int) {
	f := make([][][]int, n) // O(n^3)
	for i := range f {
		f[i] = make([][]int, n)
		for j := range f[i] {
			f[i][j] = make([]int, n)
		}
	}
	// ...
	for k := 1; k <= n; k++ { // 因为第一维对结果无影响，可以发现数组的第一维是可以省略的
		for x := 1; x <= n; x++ {
			for y := 1; y <= n; y++ {
				f[k][x][y] = min(f[k-1][x][y], f[k-1][x][k]+f[k-1][k][y])
			}
		}
	}
}
