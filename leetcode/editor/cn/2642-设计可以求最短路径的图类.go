//给你一个有 n 个节点的 有向带权 图，节点编号为 0 到 n - 1 。图中的初始边用数组 edges 表示，其中 edges[i] = [fromi,
//toi, edgeCosti] 表示从 fromi 到 toi 有一条代价为 edgeCosti 的边。
//
// 请你实现一个 Graph 类：
//
//
// Graph(int n, int[][] edges) 初始化图有 n 个节点，并输入初始边。
// addEdge(int[] edge) 向边集中添加一条边，其中 edge = [from, to, edgeCost] 。数据保证添加这条边之前对应的两
//个节点之间没有有向边。
// int shortestPath(int node1, int node2) 返回从节点 node1 到 node2 的路径 最小 代价。如果路径不存在，
//返回 -1 。一条路径的代价是路径中所有边代价之和。
//
//
//
//
// 示例 1：
//
//
//
// 输入：
//["Graph", "shortestPath", "shortestPath", "addEdge", "shortestPath"]
//[[4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]], [3, 2], [0, 3], [[1, 3, 4]
//], [0, 3]]
//输出：
//[null, 6, -1, null, 6]
//
//解释：
//Graph g = new Graph(4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]);
//g.shortestPath(3, 2); // 返回 6 。从 3 到 2 的最短路径如第一幅图所示：3 -> 0 -> 1 -> 2 ，总代价为 3 +
// 2 + 1 = 6 。
//g.shortestPath(0, 3); // 返回 -1 。没有从 0 到 3 的路径。
//g.addEdge([1, 3, 4]); // 添加一条节点 1 到节点 3 的边，得到第二幅图。
//g.shortestPath(0, 3); // 返回 6 。从 0 到 3 的最短路径为 0 -> 1 -> 3 ，总代价为 2 + 4 = 6 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 100
// 0 <= edges.length <= n * (n - 1)
// edges[i].length == edge.length == 3
// 0 <= fromi, toi, from, to, node1, node2 <= n - 1
// 1 <= edgeCosti, edgeCost <= 10⁶
// 图中任何时候都不会有重边和自环。
// 调用 addEdge 至多 100 次。
// 调用 shortestPath 至多 100 次。
//
//
// Related Topics 图 设计 最短路 堆（优先队列） 👍 18 👎 0

package main

import (
	"container/heap"
	"math"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type Graph struct {
	adj [][][2]int
	n   int
}

func Constructor(n int, edges [][]int) Graph {
	adj := make([][][2]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], [2]int{e[1], e[2]})
	}
	return Graph{adj, n}
}

func (this *Graph) AddEdge(edge []int) {
	this.adj[edge[0]] = append(this.adj[edge[0]], [2]int{edge[1], edge[2]})
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	// Dijkstra
	d, h := make([]int, this.n), &hp2642{[2]int{node1, 0}}
	for i := range d {
		d[i] = math.MaxInt // 后面比较 v < d[e[0]]
	}
	d[node1] = 0
	for h.Len() > 0 {
		cur := heap.Pop(h).([2]int)
		if cur[0] == node2 { // 终点
			return cur[1]
		}
		if cur[1] > d[cur[0]] { // 已遍历，且没有 =
			continue
		}
		for _, e := range this.adj[cur[0]] {
			if v := cur[1] + e[1]; v < d[e[0]] { // 更新距离
				d[e[0]] = v
				heap.Push(h, [2]int{e[0], v})
			}
		}
	}
	return -1
}

type hp2642 [][2]int

func (h hp2642) Len() int           { return len(h) }
func (h hp2642) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h hp2642) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp2642) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *hp2642) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
//leetcode submit region end(Prohibit modification and deletion)
