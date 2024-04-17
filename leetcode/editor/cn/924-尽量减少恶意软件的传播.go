//给出了一个由 n 个节点组成的网络，用 n × n 个邻接矩阵图
// graph 表示。在节点网络中，当 graph[i][j] = 1 时，表示节点 i 能够直接连接到另一个节点 j。
//
// 一些节点 initial 最初被恶意软件感染。只要两个节点直接连接，且其中至少一个节点受到恶意软件的感染，那么两个节点都将被恶意软件感染。这种恶意软件的传
//播将继续，直到没有更多的节点可以被这种方式感染。
//
// 假设 M(initial) 是在恶意软件停止传播之后，整个网络中感染恶意软件的最终节点数。
//
// 如果从 initial 中移除某一节点能够最小化 M(initial)， 返回该节点。如果有多个节点满足条件，就返回索引最小的节点。
//
// 请注意，如果某个节点已从受感染节点的列表 initial 中删除，它以后仍有可能因恶意软件传播而受到感染。
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
//输入：graph = [[1,0,0],[0,1,0],[0,0,1]], initial = [0,2]
//输出：0
//
//
// 示例 3：
//
//
//输入：graph = [[1,1,1],[1,1,1],[1,1,1]], initial = [1,2]
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
// graph[i][j] == 0 或 1.
// graph[i][j] == graph[j][i]
// graph[i][i] == 1
// 1 <= initial.length <= n
// 0 <= initial[i] <= n - 1
// initial 中所有整数均不重复
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 哈希表 👍 90 👎 0

package main

import "fmt"

func main() {
	graph := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	initial := []int{0, 2}
	spread := minMalwareSpread(graph, initial)
	fmt.Println(spread)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minMalwareSpread(graph [][]int, initial []int) int {
	n := len(graph)
	ids, size := make([]int, n), make([]int, 1)
	for i, id, cnt := 0, 0, 0; i < n; i++ {
		if ids[i] == 0 {
			id++ // 记录连通分量的编号
			ids[i] = id
			cnt = 1
			q := []int{i}
			for ; len(q) > 0; q = q[1:] {
				cur := q[0]
				for j, v := range graph[cur] {
					if ids[j] == 0 && v == 1 {
						ids[j] = id // 每个结点属于的连通分量
						cnt++
						q = append(q, j)
					}
				}
			}
			size = append(size, cnt) // 每个连通分量的大小
		}
	}
	initialSize := make([]int, len(size))
	ret, remove := n+1, 0
	for _, v := range initial {
		initialSize[ids[v]]++ // 每个连通分量中，被感染的结点数量
	}
	for _, v := range initial {
		var cur int
		if initialSize[ids[v]] == 1 { // 只有感染数为 1，才有意义
			cur = size[ids[v]] // 减少的感染数
		}
		if cur > remove || cur == remove && v < ret { // 更新
			ret, remove = v, cur
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
