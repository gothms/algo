//节点间通路。给定有向图，设计一个算法，找出两个节点之间是否存在一条路径。
//
// 示例1:
//
//  输入：n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]], start = 0, target = 2
// 输出：true
//
//
// 示例2:
//
//  输入：n = 5, graph = [[0, 1], [0, 2], [0, 4], [0, 4], [0, 1], [1, 3], [1, 4], [
//1, 3], [2, 3], [3, 4]], start = 0, target = 4
// 输出 true
//
//
// 提示：
//
//
// 节点数量n在[0, 1e5]范围内。
// 节点编号大于等于 0 小于 n。
// 图中可能存在自环和平行边。
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 👍 86 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	adj, visited := make([][]int, n), make([]bool, n)
	for _, g := range graph {
		adj[g[0]] = append(adj[g[0]], g[1])
	}
	visited[start] = true
	var dfs func(int, int) bool
	dfs = func(f, t int) bool {
		for _, i := range adj[t] {
			if i == target {
				return true
			}
			if !visited[i] {
				visited[i] = true
				if dfs(t, i) {
					return true
				}
			}
		}
		return false
	}
	return dfs(-1, start)
}

//leetcode submit region end(Prohibit modification and deletion)
