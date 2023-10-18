//堆箱子。给你一堆n个箱子，箱子宽 wi、深 di、高 hi。箱子不能翻转，将箱子堆起来时，下面箱子的宽度、高度和深度必须大于上面的箱子。实现一种方法，搭出最
//高的一堆箱子。箱堆的高度为每个箱子高度的总和。
//
// 输入使用数组[wi, di, hi]表示每个箱子。
//
// 示例1:
//
//  输入：box = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
// 输出：6
//
//
// 示例2:
//
//  输入：box = [[1, 1, 1], [2, 3, 4], [2, 6, 7], [3, 4, 5]]
// 输出：10
//
//
// 提示:
//
//
// 箱子的数目不大于3000个。
//
//
// Related Topics 数组 动态规划 排序 👍 90 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	box := [][]int{{20, 14, 9},
		{6, 3, 4},
		{6, 5, 8},
		{5, 3, 3},
		{3, 9, 20},
		{2, 16, 13},
		{2, 6, 15},
		{3, 11, 7},
		{2, 2, 11},
		{3, 12, 4}}
	i := pileBox(box) // 126
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func pileBox(box [][]int) int {
	// dp
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	sort.Slice(box, func(i, j int) bool { // 排序
		return box[i][0] < box[j][0] || box[i][0] == box[j][0] &&
			(box[i][1] < box[j][1] || box[i][1] == box[j][1] && box[i][2] < box[j][2]) // 注意
	})
	max, n := 0, len(box)
	dp := make([]int, n)
	for i := 0; i < n; i++ { // dp
		dp[i] = box[i][2]        // 初始值
		for j := 0; j < i; j++ { // O(n^2)
			if box[i][0] > box[j][0] && box[i][1] > box[j][1] && box[i][2] > box[j][2] {
				dp[i] = maxVal(dp[i], dp[j]+box[i][2])
			}
		}
		max = maxVal(max, dp[i])
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
