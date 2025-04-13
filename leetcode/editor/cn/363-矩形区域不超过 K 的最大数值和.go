package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"math"
)

func main() {
	matrix := [][]int{{2, 2, -1}}
	k := 3 // 3
	k = 0  // -1
	matrix = [][]int{
		{5, -4, -3, 4},
		{-3, -4, 4, 5},
		{5, 1, 5, -4},
	}
	k = 3 // 2
	submatrix := maxSumSubmatrix(matrix, k)
	fmt.Println(submatrix)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSumSubmatrix(matrix [][]int, k int) int {
	// 有序集合
	n := len(matrix[0])
	ans := math.MinInt32
	for i := range matrix { // 上边界
		colSum := make([]int, n)         // 各列的和
		for _, row := range matrix[i:] { // 下边界
			for j, v := range row {
				colSum[j] += v
			}
			rbt := redblacktree.NewWithIntComparator()
			rbt.Put(0, 0)
			cur := 0
			for _, v := range colSum {
				cur += v
				if ceil, ok := rbt.Ceiling(cur - k); ok {
					ans = max(ans, cur-ceil.Key.(int))
				}
				rbt.Put(cur, 0)
			}
		}
	}
	return ans

	// 暴力：其中之一的遍历方式
	//m, n := len(matrix), len(matrix[0])
	//ans := math.MinInt32
	//for r := range matrix {	// 枚举上边界
	//	col := make([]int, n)
	//	for i := r; i < m; i++ {	// 枚举下边界
	//		for j := 0; j < n; j++ {
	//			col[j] += matrix[i][j]
	//		}
	//		for f := 0; f < n; f++ {	// 枚举左边界
	//			cur := 0
	//			for t := f; t < n; t++ {	// 枚举右边界
	//				cur += col[t]
	//				if cur <= k && cur > ans {
	//					ans = cur
	//				}
	//			}
	//		}
	//	}
	//}
	//return ans

	// Error：矩形的面积计算公式错误
	//rbt := redblacktree.NewWithIntComparator()
	//rbt.Put(0, 0) // 用于当前的完整矩形 [0,0] [i,j]
	//ans, sum := math.MinInt32, 0
	//for _, m := range matrix {
	//	for _, v := range m {
	//		sum += v
	//		if f, ok := rbt.Ceiling(sum - k); ok {
	//			fmt.Println(sum, f)
	//			if area := sum - f.Key.(int); area <= k {
	//				ans = max(ans, area)
	//			}
	//		}
	//		rbt.Put(sum, 0)
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
