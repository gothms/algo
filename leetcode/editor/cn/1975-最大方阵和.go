package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{{-1, 0, -1}, {-2, 1, 3}, {3, 2, 2}} // 15
	//matrix = [][]int{{2, 9, 3}, {5, 4, -4}, {1, 7, 1}}    // 34
	i := maxMatrixSum(matrix)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxMatrixSum(matrix [][]int) int64 {

}

//leetcode submit region end(Prohibit modification and deletion)

func maxMatrixSum_(matrix [][]int) int64 {
	// 提示：奇数个负数，有 0，最小绝对值
	s, minV, cnt := 0, math.MaxInt32, 0
	for _, m := range matrix {
		for _, v := range m {
			if v < 0 {
				cnt++
				v = -v
			}
			s += v
			minV = min(minV, v)
		}
	}
	if cnt&1 == 0 {
		return int64(s)
	}
	return int64(s - minV<<1)
}
