package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func leastBricks(wall [][]int) int {
	m := len(wall)
	memo := make(map[int]int)
	n := 0
	for _, v := range wall[0] {
		n += v
		memo[n]++
	}
	for _, w := range wall[1:] {
		s := 0
		for _, v := range w {
			s += v
			memo[s]++
		}
	}
	delete(memo, n)
	ans := m
	for _, v := range memo {
		ans = min(ans, m-v)
	}
	return ans

	// fatal error: runtime: cannot allocate memory
	// 测试用例:[[100000000],[100000000],[100000000]]
	//m, n := len(wall), 0
	//for _, v := range wall[0] {
	//	n += v
	//}
	//memo := make([]int, n+1) // n-1 条有效垂线
	//for _, w := range wall {
	//	s := 0
	//	for _, v := range w {
	//		s += v
	//		memo[s]++
	//	}
	//}
	//ans := m
	//for _, v := range memo[1:n] {
	//	ans = min(ans, m-v)
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
