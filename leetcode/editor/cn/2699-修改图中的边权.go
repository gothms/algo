//给你一个 n 个节点的 无向带权连通 图，节点编号为 0 到 n - 1 ，再给你一个整数数组 edges ，其中 edges[i] = [ai, bi,
//wi] 表示节点 ai 和 bi 之间有一条边权为 wi 的边。
//
// 部分边的边权为 -1（wi = -1），其他边的边权都为 正 数（wi > 0）。
//
// 你需要将所有边权为 -1 的边都修改为范围 [1, 2 * 10⁹] 中的 正整数 ，使得从节点 source 到节点 destination 的 最短距
//离 为整数 target 。如果有 多种 修改方案可以使 source 和 destination 之间的最短距离等于 target ，你可以返回任意一种方案。
//
//
// 如果存在使 source 到 destination 最短距离为 target 的方案，请你按任意顺序返回包含所有边的数组（包括未修改边权的边）。如果不存
//在这样的方案，请你返回一个 空数组 。
//
// 注意：你不能修改一开始边权为正数的边。
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 5, edges = [[4,1,-1],[2,0,-1],[0,3,-1],[4,3,-1]], source = 0,
//destination = 1, target = 5
//输出：[[4,1,1],[2,0,1],[0,3,3],[4,3,1]]
//解释：上图展示了一个满足题意的修改方案，从 0 到 1 的最短距离为 5 。
//
//
// 示例 2：
//
//
//
//
//输入：n = 3, edges = [[0,1,-1],[0,2,5]], source = 0, destination = 2, target = 6
//输出：[]
//解释：上图是一开始的图。没有办法通过修改边权为 -1 的边，使得 0 到 2 的最短距离等于 6 ，所以返回一个空数组。
//
//
// 示例 3：
//
//
//
//
//输入：n = 4, edges = [[1,0,4],[1,2,3],[2,3,5],[0,3,-1]], source = 0, destination
//= 2, target = 6
//输出：[[1,0,4],[1,2,3],[2,3,5],[0,3,1]]
//解释：上图展示了一个满足题意的修改方案，从 0 到 2 的最短距离为 6 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 100
// 1 <= edges.length <= n * (n - 1) / 2
// edges[i].length == 3
// 0 <= ai, bi < n
// wi = -1 或者 1 <= wi <= 107
// ai != bi
// 0 <= source, destination < n
// source != destination
// 1 <= target <= 10⁹
// 输入的图是连通图，且没有自环和重边。
//
//
// Related Topics 图 最短路 堆（优先队列） 👍 29 👎 0

package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	edges := [][]int{{0, 1, -1}, {0, 2, 5}}
	n, source, destination, target := 3, 0, 2, 6
	edges = [][]int{{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1}}
	n, source, destination, target = 5, 0, 1, 5
	//edges = [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, 1}}
	//n, source, destination, target = 4, 0, 2, 6
	//edges = [][]int{{0, 1, -1}, {1, 2, -1}, {3, 1, -1}, {3, 0, 2}, {0, 2, 5}}
	//n, source, destination, target = 4, 2, 3, 8
	//edges = [][]int{{0, 1, 2}, {3, 0, -1}, {2, 3, -1}, {2, 1, 5}}
	//n, source, destination, target = 4, 0, 2, 8 // []
	graphEdges := modifiedGraphEdges(n, edges, source, destination, target)
	fmt.Println(graphEdges)

	//ed := &Ed{}
	//heap.Push(ed, [2]int{0, 3})
	//heap.Push(ed, [2]int{1, 2})
	//heap.Push(ed, [2]int{4, 1})
	//heap.Push(ed, [2]int{5, 8})
	//heap.Push(ed, [2]int{6, 0})
	//fmt.Println(*ed)
}

/*
思路：Dijkstra + 小顶堆
	1.单源最短路径算法，使用 Dijkstra 算法
		找到最短路径后还要修改 -1 的权值，需要进行再次 Dijkstra 验证路径值与 target 的关系
	2.当前顶点为 curr，考察 curr 可达的所有顶点，选出下一个离 起始点source 最近的顶点 next
		2.1.curr 离起点的最短路径 + curr和next边的权值 = next 离起点的路径距离
			那么可以将 curr 可达的所有顶点（和source到next的距离），存入小顶堆
		2.2.堆顶就是下一个离起点最近的顶点，也是该顶点离起点的最短路径
代码注释：
	1.初始化数据
		adj：邻接表（存入 edges 的索引，方便定位 -1 的权值）
		dist distC：顶点 离 source 的距离
		expand：第一次 Dijkstra 后，source 到 destination 的最短路径 小于 target，需要至少增加 expand
		ed：小顶堆
		visit：标记已经进入队列，并已出列的顶点
	2.查找 source 到 destination 的最短路径
		2.1.起点
		2.2.取堆顶元素（当前离起点最近的顶点）
		2.3.找到最短路径
		2.4.出列的顶点不能继续访问
		2.5.当前顶点 curr 可达的所有顶点
			a)过滤已经出列过的顶点
			b)记录 next 到 起点 的距离，并将 next和距离 入堆
			c)为 -1 的边，最小值要置为 1 来计算距离
		2.6.source 到 destination 的最短路径 > target，不符合题意
	3.expand>0，修改最短路径上的第一个 -1 权值，主要逻辑同2.
		3.1.修改第一个 -1：curr离起点的距离 + 要修改的值 = next离起点的距离 + 需要至少增加的值expand
			ew := dist[next[0]] + expand - distC[curr[0]]
		3.2.ew == 1：表示这条路径上，已经遇到过 -1，并已修改
		3.3.修改后的最短路径仍然 < target，不符合题意
	4.补全未修改过的 -1 权值
*/
//leetcode submit region begin(Prohibit modification and deletion)
func modifiedGraphEdges(n int, edges [][]int, source, destination, target int) [][]int {
	// Dijkstra + 小顶堆：二刷错误点
	adj, dist, dist1 := make([][][2]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([][2]int, 0)
	}
	for i, e := range edges {
		adj[e[0]] = append(adj[e[0]], [2]int{e[1], i})
		adj[e[1]] = append(adj[e[1]], [2]int{e[0], i}) // 记录的 i
	}
	for i := 0; i < n; i++ {
		dist[i], dist1[i] = math.MaxInt32, math.MaxInt32 // 初始化
	}
	dist[source], dist1[source] = 0, 0
	dijkstra := func(check bool, ex int) {
		ed, visit := Ed{{source, 0}}, make([]bool, n)
		for ed.Len() > 0 {
			curr := heap.Pop(&ed).([2]int)
			if visit[curr[0]] {
				continue
			}
			if curr[0] == destination { // 索引别搞错
				break
			}
			visit[curr[0]] = true
			for _, next := range adj[curr[0]] {
				if visit[next[0]] {
					continue
				}
				//fmt.Println(curr, next)
				w := edges[next[1]][2]
				if w < 0 {
					w = 1
					if check {
						wEx := ex + dist[next[0]] - dist1[curr[0]]
						if wEx > w {
							w, edges[next[1]][2] = wEx, wEx
						}
					}
				}
				if check { // 索引别搞错
					dist1[next[0]] = min(dist1[next[0]], curr[1]+w)
					heap.Push(&ed, [2]int{next[0], dist1[next[0]]})
				} else {
					dist[next[0]] = min(dist[next[0]], curr[1]+w)
					heap.Push(&ed, [2]int{next[0], dist[next[0]]})
				}
			}
		}
	}
	dijkstra(false, 0)
	if dist[destination] > target {
		return nil
	}
	if ex := target - dist[destination]; ex > 0 {
		dijkstra(true, ex)
		if dist1[destination] < target { // < 没有成功
			return nil
		}
	}
	for _, edge := range edges {
		if edge[2] == -1 {
			edge[2] = 1
		}
	}
	return edges

	//adj, dist, distC := make([][][2]int, n), make([]int, n), make([]int, n) // 1
	//for i, edge := range edges {
	//	adj[edge[0]] = append(adj[edge[0]], [2]int{edge[1], i})
	//	adj[edge[1]] = append(adj[edge[1]], [2]int{edge[0], i})
	//}
	//for i := 0; i < n; i++ {
	//	dist[i], distC[i] = math.MaxInt32, math.MaxInt32
	//}
	//dist[source], distC[source] = 0, 0
	//expand := 0
	//dijkstra := func(check bool) {
	//	ed, visit := Ed{}, make([]bool, n)
	//	heap.Push(&ed, [2]int{source, 0}) // 2.1
	//	for ed.Len() > 0 {
	//		curr := heap.Pop(&ed).([2]int) // 2.2
	//		if visit[curr[0]] {
	//			continue
	//		}
	//		if curr[0] == destination { // 2.3
	//			break
	//		}
	//		visit[curr[0]] = true                          // 2.4
	//		for i, m := 0, len(adj[curr[0]]); i < m; i++ { // 2.5
	//			next := adj[curr[0]][i]
	//			if visit[next[0]] { // 2.5.a
	//				continue
	//			}
	//			w := edges[next[1]][2]
	//			if w < 0 { // 2.5.c
	//				w = 1
	//				if check { // 3.1
	//					ew := dist[next[0]] + expand - distC[curr[0]]
	//					if ew > w { // 3.2
	//						edges[next[1]][2], w = ew, ew
	//					}
	//				}
	//			}
	//			if check {
	//				distC[next[0]] = min(distC[next[0]], distC[curr[0]]+w)
	//				heap.Push(&ed, [2]int{next[0], distC[next[0]]})
	//			} else { // 2.5.b
	//				dist[next[0]] = min(dist[next[0]], dist[curr[0]]+w)
	//				heap.Push(&ed, [2]int{next[0], dist[next[0]]})
	//			}
	//		}
	//	}
	//}
	//dijkstra(false)                 // 2
	//if dist[destination] > target { // 2.6
	//	return nil
	//}
	//if expand = target - dist[destination]; expand > 0 { // 3
	//	dijkstra(true)
	//	if distC[destination] < target { // 3.3
	//		return nil
	//	}
	//}
	//for i := range edges { // 4
	//	if edges[i][2] == -1 {
	//		edges[i][2] = 1
	//	}
	//}
	//return edges
}

//func modifiedGraphEdges_(n int, edges [][]int, source, destination, target int) [][]int {
//	// 一个 dist
//	minVal := func(a, b int) int {
//		if a < b {
//			return a
//		}
//		return b
//	}
//	adj, dist := make([][][2]int, n), make([]int, n)
//	for i, edge := range edges {
//		adj[edge[0]] = append(adj[edge[0]], [2]int{edge[1], i})
//		adj[edge[1]] = append(adj[edge[1]], [2]int{edge[0], i})
//	}
//	for i := 0; i < n; i++ {
//		dist[i] = math.MaxInt32
//	}
//	dist[source] = 0
//	expand := 0
//	dijkstra := func(visit []bool, check bool) {
//		ed := Ed{}
//		heap.Push(&ed, [2]int{source, 0})
//		for ed.Len() > 0 {
//			curr := heap.Pop(&ed).([2]int)
//			if visit[curr[0]] {
//				continue
//			}
//			if curr[0] == destination {
//				break
//			}
//			visit[curr[0]] = true
//			for i, m := 0, len(adj[curr[0]]); i < m; i++ {
//				next := adj[curr[0]][i]
//				if visit[next[0]] {
//					continue
//				}
//				w := edges[next[1]][2]
//				if w < 0 {
//					//fmt.Println(edges[next[1]], w)
//					w = 1
//					//w, edges[next[1]][2] = 1, 1
//					if check {
//						ew := dist[next[0]] + expand - dist[curr[0]]
//						if ew > w {
//							edges[next[1]][2], w = ew, ew
//						}
//					}
//					dist[next[0]] = dist[curr[0]] + w // 这种写法没法判断 dist[next[0]] 值是否更改
//				}
//				//if !check {
//				dist[next[0]] = minVal(dist[next[0]], dist[curr[0]]+w)
//				//}
//				heap.Push(&ed, [2]int{next[0], dist[next[0]]})
//			}
//		}
//	}
//	dijkstra(make([]bool, n), false)
//	if dist[destination] > target {
//		return nil
//	}
//	if dist[destination] < target {
//		expand = target - dist[destination]
//		dijkstra(make([]bool, n), true)
//		if dist[destination] < target {
//			return nil
//		}
//	}
//	for i := range edges {
//		if edges[i][2] == -1 {
//			edges[i][2] = 1
//		}
//	}
//	return edges
//}

//func modifiedGraphEdges(n int, edges [][]int, source, destination, target int) [][]int {
// 尝试一次 Dijkstra：不行
////adj, ed, dist, visit, path :=
////	make([][][2]int, n), Ed{{source, 0}}, make([]int, n), make([]bool, n), make([]int, n)
//adj, ed, dist, visit, path :=
//	make([][][2]int, n), Ed{{source, 0}}, make([]int, n), make([]bool, len(edges)), make([]int, n)
//for i := 0; i < n; i++ {
//	adj[i] = make([][2]int, n)
//	dist[i] = math.MaxInt32
//}
//dist[source] = 0
//for i, edge := range edges {
//	adj[edge[0]][edge[1]][0], adj[edge[1]][edge[0]][0] = edge[2], edge[2] // 只存不改
//	adj[edge[0]][edge[1]][1], adj[edge[1]][edge[0]][1] = i, i             // edges 索引
//}
//check := func(dis int) bool {
//	fmt.Println("dis", dis)
//	for t, f := destination, path[destination]; f != source; t, f = f, path[f] {
//		if adj[f][t][0] < 0 {
//			if edges[adj[f][t][1]][2] < target-dis {
//				edges[adj[f][t][1]][2] += target - dis
//			}
//			return false
//		}
//	}
//	return true
//}
////out:
//for len(ed) > 0 {
//	curr := heap.Pop(&ed).([2]int)
//	//if curr[1] >= target {
//	//	break
//	//}
//	//for curr[1] > dist[curr[0]] || visit[curr[0]] {
//	//	if len(ed) == 0 {
//	//		//fmt.Println(ed)
//	//		break out
//	//	}
//	//	curr = heap.Pop(&ed).([2]int)
//	//}
//	//if curr[0] == destination {
//	//	// TODO
//	//	check(destination, dist[])
//	//}
//	//visit[curr[0]] = true
//	for i := 0; i < n; i++ {
//		w := adj[i][curr[0]][0]
//		fmt.Println(curr[1], w, ed, adj[curr[0]][i][1], visit[adj[curr[0]][i][1]])
//		if w == 0 {
//			continue
//		}
//		if w == -1 {
//			edges[adj[i][curr[0]][1]][2], w = 1, 1
//			//w = 1
//		}
//		//if curr[1]+w >= target || visit[adj[curr[0]][i][1]] {
//		if curr[1]+w >= target {
//			//visit[i] = true
//			continue
//		}
//		visit[adj[curr[0]][i][1]] = true
//		//if dist[i] > curr[1]+w || i == destination { // 更小的 dist 肯定先访问到，这里更新的是 math.MaxInt32
//		dist[i], path[i] = curr[1]+w, curr[0]
//		//if !visit[i] {
//		//	visit[i] = true
//		//}
//		//fmt.Println(visit, ed)
//		//fmt.Println(i, destination)
//		if i == destination {
//			fmt.Println(dist, ed, visit)
//			if check(dist[i]) {
//				return nil
//			}
//			continue
//		}
//		heap.Push(&ed, [2]int{i, dist[i]})
//		//}
//	}
//}
//for i := range edges {
//	if edges[i][2] < 0 {
//		edges[i][2] = 1
//	}
//}
//return edges
//}

type Ed [][2]int

func (ed *Ed) Push(x interface{}) {
	*ed = append(*ed, x.([2]int))
}
func (ed *Ed) Pop() interface{} {
	v := (*ed)[len(*ed)-1]
	*ed = (*ed)[:len(*ed)-1]
	return v
}
func (ed Ed) Len() int {
	return len(ed)
}
func (ed Ed) Less(i, j int) bool {
	return ed[i][1] < ed[j][1]
}
func (ed Ed) Swap(i, j int) {
	ed[i], ed[j] = ed[j], ed[i]
}
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

//func modifiedGraphEdgesLC(n int, edges [][]int, source, destination, target int) [][]int {
//	type edge struct{ to, eid int }
//	g := make([][]edge, n)
//	for i, e := range edges {
//		x, y := e[0], e[1]
//		g[x] = append(g[x], edge{y, i})
//		g[y] = append(g[y], edge{x, i}) // 建图，额外记录边的编号
//	}
//
//	var delta int
//	dis := make([][2]int, n)
//	for i := range dis {
//		dis[i] = [2]int{math.MaxInt32, math.MaxInt32}
//	}
//	dis[source] = [2]int{}
//	dijkstra := func(k int) { // 这里 k 表示第一次/第二次
//		vis := make([]bool, n)
//		for {
//			// 找到当前最短路，去更新它的邻居的最短路
//			// 根据数学归纳法，dis[x][k] 一定是最短路长度
//			x := -1
//			for y, b := range vis {
//				if !b && (x < 0 || dis[y][k] < dis[x][k]) {
//					x = y
//				}
//			}
//			if x == destination { // 起点 source 到终点 destination 的最短路已确定
//				return
//			}
//			vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
//			for _, e := range g[x] {
//				y, wt := e.to, edges[e.eid][2]
//				if wt == -1 {
//					wt = 1 // -1 改成 1
//				}
//				if k == 1 && edges[e.eid][2] == -1 {
//					// 第二次 Dijkstra，改成 w
//					w := delta + dis[y][0] - dis[x][1]
//					if w > wt {
//						wt = w
//						edges[e.eid][2] = w // 直接在 edges 上修改
//					}
//				}
//				// 更新最短路
//				dis[y][k] = min(dis[y][k], dis[x][k]+wt)
//			}
//		}
//	}
//
//	dijkstra(0)
//	fmt.Println(edges)
//	delta = target - dis[destination][0]
//	if delta < 0 { // -1 全改为 1 时，最短路比 target 还大
//		return nil
//	}
//
//	dijkstra(1)
//	fmt.Println(edges)
//	if dis[destination][1] < target { // 最短路无法再变大，无法达到 target
//		return nil
//	}
//
//	for _, e := range edges {
//		if e[2] == -1 { // 剩余没修改的边全部改成 1
//			e[2] = 1
//		}
//	}
//	return edges
//}

//leetcode submit region end(Prohibit modification and deletion)
