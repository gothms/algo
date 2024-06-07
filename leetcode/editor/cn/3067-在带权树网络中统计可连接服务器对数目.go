package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {

}

//leetcode submit region end(Prohibit modification and deletion)

func countPairsOfConnectableServers_(edges [][]int, signalSpeed int) []int {
	// dfs 2
	n := len(edges) + 1
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u], adj[v] = append(adj[u], [2]int{v, w}), append(adj[v], [2]int{u, w})
	}
	var dfs func(int, int, int) int
	dfs = func(f, t, s int) int {
		ret := 0
		if s == 0 {
			ret++
		}
		for _, e := range adj[t] {
			if e[0] != f {
				ret += dfs(t, e[0], (s+e[1])%signalSpeed)
			}
		}
		return ret
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if len(adj[i]) == 1 {
			continue
		}
		cnt := 0
		for _, e := range adj[i] {
			c := dfs(i, e[0], e[1]%signalSpeed)
			ans[i] += cnt * c
			cnt += c
		}
	}
	return ans

	// dfs
	//n := len(edges) + 1
	//adj := make([][][2]int, n)
	//for _, e := range edges {
	//	u, v, w := e[0], e[1], e[2]
	//	adj[u], adj[v] = append(adj[u], [2]int{v, w}), append(adj[v], [2]int{u, w})
	//}
	//var cnt int
	//var dfs func(int, int, int)
	//dfs = func(f, t, s int) {
	//	if s%signalSpeed == 0 {
	//		cnt++ // 一个解
	//	}
	//	for _, e := range adj[t] {
	//		if e[0] != f {
	//			dfs(t, e[0], s+e[1])
	//		}
	//	}
	//}
	//ans := make([]int, n)
	//for i := 0; i < n; i++ { // 误区：”归“有多少节点可以整除
	//	sum := 0
	//	for _, e := range adj[i] {
	//		dfs(i, e[0], e[1])  // 正解：“递”有多少节点可以整除
	//		ans[i] += sum * cnt // 关键计算
	//		sum += cnt          // 总数
	//		cnt = 0             // 重置
	//	}
	//}
	//return ans
}
