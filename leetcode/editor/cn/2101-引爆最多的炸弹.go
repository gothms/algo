package main

import "fmt"

func main() {
	bombs := [][]int{{1, 1, 5}, {10, 10, 5}}
	detonation := maximumDetonation(bombs)
	fmt.Println(detonation)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumDetonation(bombs [][]int) int {
	// 为什么不能用并查集？
	n := len(bombs)
	g := make([][]int, n)
	for i, f := range bombs {
		x, y, r := f[0], f[1], f[2]
		for j, t := range bombs {
			dx, dy := x-t[0], y-t[1]
			if r*r >= dx*dx+dy*dy {
				g[i] = append(g[i], j) // i 引爆 j
			}
		}
	}
	vis := make([]bool, n)
	var dfs func(int) int
	dfs = func(i int) int {
		vis[i] = true
		res := 1
		for _, j := range g[i] {
			if !vis[j] {
				res += dfs(j)
			}
		}
		return res
	}
	ans := 0
	for i := range g {
		clear(vis)
		ans = max(ans, dfs(i))
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
