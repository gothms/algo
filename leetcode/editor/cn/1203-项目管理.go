package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	// dfs

	// khan
	//var (
	//	degree = make([]int, n)
	//	g      = make([][]int, n)
	//	q      = make([]int, 0, n)
	//	ans    = make([]int, 0, n)
	//)
	//for i, bi := range beforeItems {
	//	degree[i] = len(bi)
	//	if degree[i] == 0 {
	//		q = append(q, i)
	//	}
	//	for _, f := range bi {
	//		g[f] = append(g[f], i)
	//	}
	//}
	//cnt := 0
	//for len(q) > 0 {
	//	f := q[0]
	//	q = q[1:]
	//	for _, t := range g[f] {
	//		degree[t]--
	//		if degree[t] == 0 {
	//			q = append(q, t)
	//		}
	//	}
	//}
	//if cnt < n {
	//	return nil
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
