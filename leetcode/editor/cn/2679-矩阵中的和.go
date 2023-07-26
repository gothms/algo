//给你一个下标从 0 开始的二维整数数组 nums 。一开始你的分数为 0 。你需要执行以下操作直到矩阵变为空：
//
//
// 矩阵中每一行选取最大的一个数，并删除它。如果一行中有多个最大的数，选择任意一个并删除。
// 在步骤 1 删除的所有数字中找到最大的一个数字，将它添加到你的 分数 中。
//
//
// 请你返回最后的 分数 。
//
//
//
// 示例 1：
//
//
//输入：nums = [[7,2,1],[6,4,2],[6,5,3],[3,2,1]]
//输出：15
//解释：第一步操作中，我们删除 7 ，6 ，6 和 3 ，将分数增加 7 。下一步操作中，删除 2 ，4 ，5 和 2 ，将分数增加 5 。最后删除 1 ，2
// ，3 和 1 ，将分数增加 3 。所以总得分为 7 + 5 + 3 = 15 。
//
//
// 示例 2：
//
//
//输入：nums = [[1]]
//输出：1
//解释：我们删除 1 并将分数增加 1 ，所以返回 1 。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 300
// 1 <= nums[i].length <= 500
// 0 <= nums[i][j] <= 10³
//
//
// Related Topics 数组 矩阵 排序 模拟 堆（优先队列） 👍 7 👎 0

package main

import (
	"sort"
)

func main() {

}

/*
思路：排序
	1.按每行进行排序
	2.从每列中选出最大值，累加
*/
// leetcode submit region begin(Prohibit modification and deletion)
func matrixSum(nums [][]int) int {
	sum, n, m := 0, len(nums), len(nums[0])
	for i := 0; i < n; i++ {
		sort.Ints(nums[i]) // 行排序
	}
	for j := 0; j < m; j++ { // 选出列的最大值
		max := nums[0][j]
		for i := 1; i < n; i++ {
			if nums[i][j] > max {
				max = nums[i][j]
			}
		}
		sum += max // 累加
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
