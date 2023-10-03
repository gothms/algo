//一个数组的 最小乘积 定义为这个数组中 最小值 乘以 数组的 和 。
//
//
// 比方说，数组 [3,2,5] （最小值是 2）的最小乘积为 2 * (3+2+5) = 2 * 10 = 20 。
//
//
// 给你一个正整数数组 nums ，请你返回 nums 任意 非空子数组 的最小乘积 的 最大值 。由于答案可能很大，请你返回答案对 10⁹ + 7 取余 的
//结果。
//
// 请注意，最小乘积的最大值考虑的是取余操作 之前 的结果。题目保证最小乘积的最大值在 不取余 的情况下可以用 64 位有符号整数 保存。
//
// 子数组 定义为一个数组的 连续 部分。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,2]
//输出：14
//解释：最小乘积的最大值由子数组 [2,3,2] （最小值是 2）得到。
//2 * (2+3+2) = 2 * 7 = 14 。
//
//
// 示例 2：
//
//
//输入：nums = [2,3,3,1,2]
//输出：18
//解释：最小乘积的最大值由子数组 [3,3] （最小值是 3）得到。
//3 * (3+3) = 3 * 6 = 18 。
//
//
// 示例 3：
//
//
//输入：nums = [3,1,5,6,4,2]
//输出：60
//解释：最小乘积的最大值由子数组 [5,6,4] （最小值是 4）得到。
//4 * (5+6+4) = 4 * 15 = 60 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁷
//
//
// Related Topics 栈 数组 前缀和 单调栈 👍 121 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2}
	//nums = []int{3, 1, 5, 6, 4, 2}
	nums = []int{2, 5, 4, 2, 4, 5, 3, 1, 2, 4}
	nums = []int{4, 10, 6, 4, 8, 7, 8, 3, 5, 3, 4, 9, 9, 5, 10, 7, 10, 7, 6, 4} // 387
	product := maxSumMinProduct(nums)
	fmt.Println(product)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSumMinProduct(nums []int) int {
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	const mod = 1_000_000_007
	msmp, n := 0, len(nums)
	pre, stack := make([]int, n+1), make([]int, 0, n)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + nums[i] // pre[0] = 0
	}
	for i, v := range nums {
		right[i] = n // 没有被弹出
		last := len(stack) - 1
		for last >= 0 && v <= nums[stack[last]] { // 单调递增
			right[stack[last]] = i // stack[last] 的右边界，默认为 n（pre 的索引从 1 开始）
			last--
		}
		if last != -1 {
			left[i] = stack[last] + 1 // i 的左边界，默认为 0（即小于 nums[i] 的元素索引）
		}
		stack = stack[:last+1]
		stack = append(stack, i)
	}
	for i := 0; i < n; i++ {
		msmp = maxVal(msmp, (pre[right[i]]-pre[left[i]])*nums[i])
	}
	return msmp % mod

	// 个人写法
	//maxVal := func(a, b int) int {
	//	if b > a {
	//		return b
	//	}
	//	return a
	//}
	//const mod = 1_000_000_007
	//msmp, n := 0, len(nums)
	//pre, stack := make([]int, n+1), make([]int, 1, n+1)
	//stack[0] = -1
	//for i := 0; i < n; i++ {
	//	pre[i+1] = pre[i] + nums[i]
	//}
	////fmt.Println(pre)
	//for i, v := range nums {
	//	last := len(stack) - 1
	//	for last > 0 && v <= nums[stack[last]] {
	//		last--
	//	}
	//	msmp = maxVal(msmp, (pre[i+1]-pre[stack[last]+1])*v)
	//	for j := 1; j <= last; j++ { // TODO
	//		msmp = maxVal(msmp, (pre[i+1]-pre[stack[j-1]+1])*nums[stack[j]])
	//	}
	//	//fmt.Println(i, pre[i+1], pre[stack[last]+1], v)
	//	stack = stack[:last+1]
	//	stack = append(stack, i)
	//}
	//return msmp % mod
}

//leetcode submit region end(Prohibit modification and deletion)
