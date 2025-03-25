package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 2, 3},
		{3, 1, 5},
		{3, 2, 1}}
	values := differenceOfDistinctValues(grid)
	fmt.Println(values)
}

// leetcode submit region begin(Prohibit modification and deletion)
func differenceOfDistinctValues(grid [][]int) [][]int {
	// lc
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	memo := make(map[int]struct{})
	for k := 1; k < m+n; k++ {
		mnJ, mxJ := max(n-k, 0), min(m+n-k-1, n-1)
		clear(memo)
		for j := mnJ; j <= mxJ; j++ {
			i := k + j - n
			ans[i][j] = len(memo)
			memo[grid[i][j]] = struct{}{}
		}
		clear(memo)
		for j := mxJ; j >= mnJ; j-- {
			i := k + j - n
			ans[i][j] = abs(ans[i][j] - len(memo))
			memo[grid[i][j]] = struct{}{}
		}
	}
	return ans

	// 个人
	//m, n := len(grid), len(grid[0])
	//ans := make([][]int, m)
	//for i := range ans {
	//	ans[i] = make([]int, n)
	//}
	//memo := make(map[int]struct{})
	//for j := 0; j < n; j++ {
	//	clear(memo)
	//	for i := 0; i < min(n-j, m); i++ {
	//		ans[i][i+j] = len(memo)
	//		memo[grid[i][i+j]] = struct{}{}
	//	}
	//}
	//for i := 1; i < m; i++ {
	//	clear(memo)
	//	for j := 0; j < min(m-i, n); j++ {
	//		ans[i+j][j] = len(memo)
	//		memo[grid[i+j][j]] = struct{}{}
	//	}
	//}
	//for j := n - 1; j >= 0; j-- {
	//	clear(memo)
	//	for i := min(n-j-1, m-1); i >= 0; i-- {
	//		ans[i][i+j] = int(math.Abs(float64(ans[i][i+j] - len(memo))))
	//		memo[grid[i][i+j]] = struct{}{}
	//	}
	//}
	//for i := m - 1; i > 0; i-- {
	//	clear(memo)
	//	for j := min(m-i-1, n-1); j >= 0; j-- {
	//		ans[i+j][j] = int(math.Abs(float64(ans[i+j][j] - len(memo))))
	//		memo[grid[i+j][j]] = struct{}{}
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
