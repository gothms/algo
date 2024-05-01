//给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//
//
//
//
//
// 示例 1:
//
//
//输入: rowIndex = 3
//输出: [1,3,3,1]
//
//
// 示例 2:
//
//
//输入: rowIndex = 0
//输出: [1]
//
//
// 示例 3:
//
//
//输入: rowIndex = 1
//输出: [1,1]
//
//
//
//
// 提示:
//
//
// 0 <= rowIndex <= 33
//
//
//
//
// 进阶：
//
// 你可以优化你的算法到 O(rowIndex) 空间复杂度吗？
//
// Related Topics 数组 动态规划 👍 540 👎 0

package main

import "fmt"

func main() {
	rowIndex := 4
	row := getRow(rowIndex)
	fmt.Println(row)
}

// leetcode submit region begin(Prohibit modification and deletion)
func getRow(rowIndex int) []int {
	// dp
	dp := make([]int, rowIndex+1)
	dp[0] = 1
	for i := 1; i <= rowIndex; i++ {
		dp[i] = 1                     // 处理尾巴
		for j := i >> 1; j > 0; j-- { // 倒序
			dp[j] += dp[j-1] // 中点
			dp[i-j] = dp[j]  // 也可能是中点
		}
	}
	return dp
}

//leetcode submit region end(Prohibit modification and deletion)
