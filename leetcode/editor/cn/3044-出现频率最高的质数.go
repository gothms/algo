package main

import (
	"fmt"
	"slices"
)

func main() {
	mat := [][]int{{1, 1}, {9, 9}, {1, 1}}
	prime := mostFrequentPrime(mat)
	fmt.Println(prime)
}

//leetcode submit region begin(Prohibit modification and deletion)

var prime3044 []bool

func init() {
	const n = 1e6
	prime3044 = make([]bool, n)
	for i := 2; i < n; i++ {
		if prime3044[i] {
			continue
		}
		for j := i << 1; j < n; j += i {
			prime3044[j] = true
		}
	}
}
func mostFrequentPrime(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	memo := make(map[int]int)
	for _, row := range mat {
		cur := 0
		for _, v := range row {
			cur = cur*10 + v
			memo[cur]++
		}
		cur = 0
		for _, v := range slices.Backward(row) {
			cur = cur*10 + v
			memo[cur]++
		}
	}
	for j := 0; j < n; j++ {
		cur := 0
		for i := 0; i < m; i++ {
			cur = cur*10 + mat[i][j]
			memo[cur]++
		}
		cur = 0
		for i := m - 1; i >= 0; i-- {
			cur = cur*10 + mat[i][j]
			memo[cur]++
		}
	}
	for k := 0; k < m+n-1; k++ {
		i0 := max(0, k-n+1)
		// 捺
		j0 := max(0, m-k-1)
		cur := 0
		for j := j0; j < min(n, j0+m-i0); j++ {
			fmt.Println(k, i0, j0, j, i0+j-j0)
			cur = cur*10 + mat[i0+j-j0][j]
			memo[cur]++
		}
		// k 的顺序倒过来了
		i0 = m - i0 - 1
		j0 = n - j0 - 1
		cur = 0
		for j := j0; j >= max(0, j0-i0); j-- {
			cur = cur*10 + mat[i0-j+j0][j]
			memo[cur]++
		}
		//	 map[1:12 9:4 11:4 19:5 91:2 99:2 191:4]
		// 撇
		j0 = min(k, n-1)
		cur = 0
		for j := j0; j >= max(0, j0-m+1); j-- {
			fmt.Println(i0, j0, j0-m+1, j, i0+j0-j)
			cur = cur*10 + mat[i0+j0-j][j]
			memo[cur]++
		}
		i0 = m - i0 - 1
		j0 = n - j0 - 1
		cur = 0
		for j := j0; j < min(n, j0+i0+1); j++ {
			cur = cur*10 + mat[i0-j0+j][j]
			memo[cur]++
		}
	}

	fmt.Println(memo)
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
