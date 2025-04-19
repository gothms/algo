package main

import (
	"fmt"
)

func main() {
	mat := [][]int{{1, 1}, {9, 9}, {1, 1}}
	//mat = [][]int{{4, 5, 3}} // 53
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
	// 暴力
	dxy := [8][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	m, n := len(mat), len(mat[0])
	memo := make(map[int]int)
	for i, row := range mat {
		for j, v := range row {
			for _, p := range dxy {
				x, y, v := i+p[0], j+p[1], v
				for x >= 0 && x < m && y >= 0 && y < n {
					v = v*10 + mat[x][y]
					memo[v]++
					x += p[0]
					y += p[1]
				}
			}
		}
	}

	ans, cnt := -1, 0
	for k, v := range memo {
		if prime3044[k] {
			continue
		}
		if v > cnt || v == cnt && k > ans {
			ans, cnt = k, v
		}
	}
	return ans

	// 个人
	//m, n := len(mat), len(mat[0])
	//memo := make(map[int]int)
	//ans, cnt := 0, 0
	//for k, v := range memo {
	//	if k < 10 || prime3044[k] {
	//		continue
	//	}
	//	if v > cnt || v == cnt && k > ans {
	//		ans, cnt = k, v
	//	}
	//}
	//fmt.Println(memo)
	//if ans == 0 {
	//	return -1
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
