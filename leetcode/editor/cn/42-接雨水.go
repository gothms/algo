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

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	i := trap(height)
	fmt.Println(i)
}

/*
思路：dp
	1.l/r 为从 左/右 遍历，当前最高的柱子
	2.dp
		2.1.假设总能接到雨水 >= 0
			假设 n 位置有个无限高的柱子，从左开始计算每个位置的最大积水，记为 lv
			假设 -1 位置有个无限高的柱子，从右开始计算每个位置的最大积水，记为 rv
		2.2.每个位置能接到的雨水为：min(lv[i], rv[i])
		2.3.dp 方程请见代码
思路：栈
	1.遍历所有柱子，当前柱子 height[i] 与已遍历柱子能接雨水的情况是：
		存在一个左边界柱子和 height[i] 中间有凹地，这个凹地就能接雨水
		所以需要至少3根柱子来形成这个凹地
	2.考虑维护一个单调栈（非递增） stack
		2.1.height[i] <= stack[last]：不能形成凹地，直接将 height[i] 加入 stack
		2.2.height[i] > stack[last]：寻找可能形成的凹地，计算凹地接的雨水
			maxV += 宽度 * 高度
			宽度：height[i] - 左边界 - 1
			高度：min(height[i], 左边界) - 凹地柱子的高度
		2.3.重复 2.2. 直到没有左边界，或者 左边界 >= height[i]
思路：双指针
	1.lMax/rMax 分别记录 左/右 最高柱子
	2.左/右指针 i/j 分别从左/右遍历柱子
		2.1.height[i] < height[j]：右边界 height[j]
			if height[i] < lMax：以 lMax 为左边界，更新雨水
			else：左边界更新为 height[i]
		2.2.同上
*/
// leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	// 双指针
	//maxV, lMax, rMax, n := 0, 0, 0, len(height)
	//for i, j := 0, n-1; i < j; {
	//	if height[i] < height[j] { // 右边界 height[j]
	//		if height[i] < lMax { // 更新雨水
	//			maxV += lMax - height[i]
	//		} else {
	//			lMax = height[i] // 更新左边界
	//		}
	//		i++
	//	} else { // 左边界 height[i]
	//		if height[j] < rMax { // 更新雨水
	//			maxV += rMax - height[j]
	//		} else {
	//			rMax = height[j] // 更新右边界
	//		}
	//		j--
	//	}
	//}
	//return maxV

	// 栈
	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//maxV, n := 0, len(height)
	//stack := []int{0} // 单调栈（非递增）
	//for i := 1; i < n; i++ {
	//	l := len(stack) - 1
	//	for last := height[stack[l]]; height[i] > last; last = height[stack[l]] {
	//		l--
	//		if l < 0 { // 两个柱子没雨水
	//			break
	//		}
	//		// 三个柱子：宽度 * (minVal(左高, 右高) - 中高)
	//		maxV += (i - stack[l] - 1) * (minVal(height[i], height[stack[l]]) - last)
	//	}
	//	stack = append(stack[:l+1], i) // 添加当前柱子
	//}
	//return maxV

	// dp
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	maxV, l, r, n := 0, 0, 0, len(height) // l/r 为 左/右 最高柱子
	lv, rv := make([]int, n), make([]int, n)
	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		l, r = maxVal(l, height[i]), maxVal(r, height[j]) // 更新 l r
		lv[i], rv[j] = l-height[i], r-height[j]           // 更新 左/右 雨水量
	}
	for i := 0; i < n; i++ {
		maxV += minVal(lv[i], rv[i])
	}
	return maxV
}

//leetcode submit region end(Prohibit modification and deletion)
