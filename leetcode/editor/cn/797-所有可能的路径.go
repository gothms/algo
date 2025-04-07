package main

import (
	"fmt"
	"slices"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	//vis := make([]bool, n)	// 无环，不需要 vis
	ans, path := make([][]int, 0, n), make([]int, 1, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n-1 {
			ans = append(ans, slices.Clone(path))
			return
		}
		m := len(path)
		path = append(path, 0)
		for _, j := range graph[i] {
			path[m] = j
			dfs(j)
		}
		path = path[:m] // 回溯
	}
	dfs(0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
