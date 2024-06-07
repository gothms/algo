package main

import (
	"fmt"
	"slices"
)

func main() {
	graph := [][]int{{4, 3, 1}, {3, 2, 4}, {}, {4}, {}}
	target := allPathsSourceTarget(graph)
	fmt.Println(target)
}

// leetcode submit region begin(Prohibit modification and deletion)
func allPathsSourceTarget(graph [][]int) [][]int {

}

//leetcode submit region end(Prohibit modification and deletion)

func allPathsSourceTarget_(graph [][]int) [][]int {
	// dfs
	n := len(graph) - 1
	ans := make([][]int, 0)
	var dfs func(int, []int)
	dfs = func(i int, path []int) {
		if i == n { // 注意是 0 -> n-1
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
