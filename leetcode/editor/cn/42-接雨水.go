//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 4347 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	// dp
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	v, n := 0, len(height)
	dpL, dpR, left, right := make([]int, n), make([]int, n), 0, 0
	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		left, right = max(left, height[i]), max(right, height[j])
		dpL[i], dpR[j] = left-height[i], right-height[j]
	}
	for i := 0; i < n; i++ {
		v += min(dpL[i], dpR[i])
	}
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
