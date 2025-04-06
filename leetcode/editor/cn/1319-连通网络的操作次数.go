package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func makeConnected(n int, connections [][]int) int {
	// dfs
	m := len(connections)
	if m+1 < n {
		return -1
	}
	g := make([][]int, n)
	for _, c := range connections {
		x, y := c[0], c[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	ans := -1
	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}
	return ans

	// 并查集
	//m := len(connections)
	//if m+1 < n {
	//	return -1
	//}
	//uni := make([]int, n)
	//for i := range uni {
	//	uni[i] = i
	//}
	//ans := n
	//var find func(int) int
	//find = func(i int) int {
	//	for uni[i] != i {
	//		i, uni[i] = uni[i], uni[uni[i]]
	//	}
	//	return uni[i]
	//}
	//union := func(x, y int) bool {
	//	xr, yr := find(x), find(y)
	//	if xr != yr {
	//		uni[xr] = yr
	//		return true
	//	}
	//	return false
	//}
	//for _, c := range connections {
	//	x, y := c[0], c[1]
	//	if union(x, y) {
	//		ans--
	//	}
	//}
	//return ans - 1
}

//leetcode submit region end(Prohibit modification and deletion)
