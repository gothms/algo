package main

import "slices"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func allPathsSourceTarget(graph [][]int) [][]int {
	// dfs
	n := len(graph) - 1
	ans := make([][]int, 0)
	var dfs func(int, []int)
	dfs = func(i int, path []int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		for _, j := range graph[i] {
			dfs(j, append(path, j))
		}
	}
	dfs(0, make([]int, 1, n))
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
