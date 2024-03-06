//给你一个长度为 n 下标从 0 开始的整数数组 maxHeights 。
//
// 你的任务是在坐标轴上建 n 座塔。第 i 座塔的下标为 i ，高度为 heights[i] 。
//
// 如果以下条件满足，我们称这些塔是 美丽 的：
//
//
// 1 <= heights[i] <= maxHeights[i]
// heights 是一个 山脉 数组。
//
//
// 如果存在下标 i 满足以下条件，那么我们称数组 heights 是一个 山脉 数组：
//
//
// 对于所有 0 < j <= i ，都有 heights[j - 1] <= heights[j]
// 对于所有 i <= k < n - 1 ，都有 heights[k + 1] <= heights[k]
//
//
// 请你返回满足 美丽塔 要求的方案中，高度和的最大值 。
//
//
//
// 示例 1：
//
//
//输入：maxHeights = [5,3,4,1,1]
//输出：13
//解释：和最大的美丽塔方案为 heights = [5,3,3,1,1] ，这是一个美丽塔方案，因为：
//- 1 <= heights[i] <= maxHeights[i]
//- heights 是个山脉数组，峰值在 i = 0 处。
//13 是所有美丽塔方案中的最大高度和。
//
// 示例 2：
//
//
//输入：maxHeights = [6,5,3,9,2,7]
//输出：22
//解释： 和最大的美丽塔方案为 heights = [3,3,3,9,2,2] ，这是一个美丽塔方案，因为：
//- 1 <= heights[i] <= maxHeights[i]
//- heights 是个山脉数组，峰值在 i = 3 处。
//22 是所有美丽塔方案中的最大高度和。
//
// 示例 3：
//
//
//输入：maxHeights = [3,2,5,5,2,3]
//输出：18
//解释：和最大的美丽塔方案为 heights = [2,2,5,5,2,2] ，这是一个美丽塔方案，因为：
//- 1 <= heights[i] <= maxHeights[i]
//- heights 是个山脉数组，最大值在 i = 2 处。
//注意，在这个方案中，i = 3 也是一个峰值。
//18 是所有美丽塔方案中的最大高度和。
//
//
//
//
// 提示：
//
//
// 1 <= n == maxHeights <= 10⁵
// 1 <= maxHeights[i] <= 10⁹
//
//
// Related Topics 栈 数组 单调栈 👍 88 👎 0

package main

import (
	"fmt"
)

func main() {
	arr := []int{6, 5, 3, 9, 2, 7}
	arr = []int{3, 5, 3, 5, 1, 5, 4, 4, 4} // 22
	heights := maximumSumOfHeights(arr)
	fmt.Println(heights)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumSumOfHeights(maxHeights []int) int64 {
	// 单调栈
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	ret, n := 0, len(maxHeights)
	// 以 maxHeights[i] 为塔尖的最大值
	left, right := make([]int, n+1), make([]int, n+1)
	stack := []int{-1}
	for i, v := range maxHeights { // 从左往右
		last := len(stack) - 1
		for last > 0 && maxHeights[stack[last]] >= v {
			last--
		}
		left[i+1] = left[stack[last]+1] + (i-stack[last])*v // i 为塔尖
		stack = append(stack[:last+1], i)
	}
	stack = stack[:1]
	stack[0] = n
	for i := n - 1; i >= 0; i-- { // 从右往左
		last := len(stack) - 1
		for last > 0 && maxHeights[stack[last]] >= maxHeights[i] {
			last--
		}
		right[i] = right[stack[last]] + (stack[last]-i)*maxHeights[i]
		stack = append(stack[:last+1], i)
	}
	for i := 0; i <= n; i++ {
		ret = maxVal(ret, left[i]+right[i])
	}
	return int64(ret)

	// 单调栈：错误
	//minVal := func(a, b int) int {
	//	if b < a {
	//		return b
	//	}
	//	return a
	//}
	//ret, cur, n := 0, math.MaxInt32, len(maxHeights)
	//stack := []int{-1}     // 单调递增
	//dp := make([]int, n+1) // 以 maxHeights[i] 为塔尖
	//for i, v := range maxHeights {
	//	last := len(stack) - 1
	//	for last > 0 && maxHeights[stack[last]] >= v {
	//		last--
	//	}
	//	dp[i+1] = dp[stack[last]+1] + (i-stack[last])*v // 以 v 为塔尖
	//	stack = append(stack[:last+1], i)
	//
	//	cur = minVal(cur, v) // 记录塔尖
	//	ret += cur
	//	fmt.Println(i, ret)
	//	if ret < dp[i+1] {
	//		ret, cur = dp[i+1], v
	//	}
	//}
	//fmt.Println(dp)
	//return int64(ret)
}

//leetcode submit region end(Prohibit modification and deletion)
