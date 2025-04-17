package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	cnt := func(edges [][]int, k int) func(int, int, int) int {
		g := make([][]int, len(edges)+1)
		for _, e := range edges {
			x, y := e[0], e[1]
			g[x], g[y] = append(g[x], y), append(g[y], x)
		}
		var dfs func(int, int, int) int
		dfs = func(s, p, f int) int {
			if s > k {
				return 0
			}
			s++
			cnt := 1
			for _, t := range g[f] {
				if t != p {
					cnt += dfs(s, f, t)
				}
			}
			return cnt
		}
		return dfs
	}
	mx := 0
	if k <= 1 {
		mx = k
	}
	f := cnt(edges2, k-1)
	for i := 0; i <= len(edges2); i++ {
		mx = max(mx, f(0, -1, i))
	}
	n := len(edges1) + 1
	ans := make([]int, n)
	f = cnt(edges1, k)
	for i := 0; i < n; i++ {
		ans[i] = f(0, -1, i) + mx
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
