//给定一个由 n 个节点组成的网络，用 n x n 个邻接矩阵 graph 表示。在节点网络中，只有当 graph[i][j] = 1 时，节点 i 能够直接
//连接到另一个节点 j。
//
// 一些节点 initial 最初被恶意软件感染。只要两个节点直接连接，且其中至少一个节点受到恶意软件的感染，那么两个节点都将被恶意软件感染。这种恶意软件的传
//播将继续，直到没有更多的节点可以被这种方式感染。
//
// 假设 M(initial) 是在恶意软件停止传播之后，整个网络中感染恶意软件的最终节点数。
//
// 我们可以从 initial 中删除一个节点，并完全移除该节点以及从该节点到任何其他节点的任何连接。
//
// 请返回移除后能够使 M(initial) 最小化的节点。如果有多个节点满足条件，返回索引 最小的节点 。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：graph = [[1,1,0],[1,1,0],[0,0,1]], initial = [0,1]
//输出：0
//
//
// 示例 2：
//
//
//输入：graph = [[1,1,0],[1,1,1],[0,1,1]], initial = [0,1]
//输出：1
//
//
// 示例 3：
//
//
//输入：graph = [[1,1,0,0],[1,1,1,0],[0,1,1,1],[0,0,1,1]], initial = [0,1]
//输出：1
//
//
//
//
// 提示：
//
//
//
// n == graph.length
// n == graph[i].length
// 2 <= n <= 300
// graph[i][j] 是 0 或 1.
// graph[i][j] == graph[j][i]
// graph[i][i] == 1
// 1 <= initial.length < n
// 0 <= initial[i] <= n - 1
// initial 中每个整数都不同
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 哈希表 👍 56 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	graph := [][]int{{1, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 1, 1},
		{0, 0, 1, 1}}
	initial := []int{0, 1}
	graph = [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 1, 0, 1, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 1, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 1, 1}}
	initial = []int{3, 7} // 3
	graph = [][]int{
		{1, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 1}}
	initial = []int{4} // 4
	spread := minMalwareSpread(graph, initial)
	fmt.Println(spread)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minMalwareSpread(graph [][]int, initial []int) int {
	if len(initial) == 1 { // fast path
		return initial[0]
	}
	sort.Ints(initial)               // 节点可能不是有序的
	id, cnt, n := -1, -1, len(graph) // 初始化 id & cnt
	var (
		clean     = make([]int, n)   // 标记感染源 / clean 节点
		visited   = make([]int, n)   // 标记本次 dfs 节点是否已访问
		counter   = make([][]int, n) // clean 节点的感染源
		initGraph = make([]int, n)   // 感染源节点直接感染的节点（只被这个感染源感染了）的数目
	)
	for i := range visited { // 初始化 visited
		visited[i] = -1
	}
	for _, init := range initial { // 标记感染源
		clean[init] = 1
	}
	var dfs func(int, int)
	dfs = func(f, cur int) {
		for i := 0; i < n; i++ { // 遍历：感染源直接感染的节点
			if graph[cur][i] == 1 && clean[i] == 0 && visited[i] != f {
				visited[i] = f
				counter[i] = append(counter[i], f) // 被感染节点：感染源
				dfs(f, i)
			}
		}
	}
	for _, f := range initial { // 遍历感染源
		dfs(f, f)
	}
	for _, inits := range counter {
		if len(inits) == 1 { // clean 节点只有一个传染源
			initGraph[inits[0]]++ // 被感染的节点，传播给其他节点的数量
		}
	}
	for _, i := range initial {
		if initGraph[i] > cnt { // 寻找最大传染数的节点
			cnt, id = initGraph[i], i
		}
	}
	return id
}

//leetcode submit region end(Prohibit modification and deletion)
