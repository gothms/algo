package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
}

//leetcode submit region end(Prohibit modification and deletion)
