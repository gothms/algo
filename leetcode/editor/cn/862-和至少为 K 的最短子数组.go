//给你一个整数数组 nums 和一个整数 k ，找出 nums 中和至少为 k 的 最短非空子数组 ，并返回该子数组的长度。如果不存在这样的 子数组 ，返回
//-1 。
//
// 子数组 是数组中 连续 的一部分。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [1], k = 1
//输出：1
//
//
// 示例 2：
//
//
//输入：nums = [1,2], k = 4
//输出：-1
//
//
// 示例 3：
//
//
//输入：nums = [2,-1,2], k = 3
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁵ <= nums[i] <= 10⁵
// 1 <= k <= 10⁹
//
//
// Related Topics 队列 数组 二分查找 前缀和 滑动窗口 单调队列 堆（优先队列） 👍 691 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-36, 10, -28, -42, 17, 83, 87, 79, 51, -26, 33, 53, 63, 61, 76, 34, 57, 68, 1, -30}
	k := 484 // 10
	nums = []int{27, 20, 79, 87, -36, 78, 76, 72, 50, -26}
	k = 453
	nums = []int{-23, 51, -14, -6, 36, 33, 76, -26, -6, 58, -16, 1, 98, 2, -20, 48, -19, -41, -34, 62}
	k = 221
	nums = []int{45, 11, -47, 46, 43, 75, 22, 82, 64, 22}
	k = 22
	nums = []int{84, -37, 32, 40, 95}
	k = 167 // 3
	nums = []int{1}
	k = 1
	nums = []int{1, 2}
	k = 4
	nums = []int{2, -1, 2}
	k = 3 // 3
	//nums = []int{-34, 37, 51, 3, -12, -50, 51, 100, -47}
	//k = 151 // 2
	subarray := shortestSubarray(nums, k)
	fmt.Println(subarray)
}

// leetcode submit region begin(Prohibit modification and deletion)
func shortestSubarray(nums []int, k int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func shortestSubarray_(nums []int, k int) int {
	// 双端队列
	ans, n := math.MaxInt32, len(nums)
	pre, deque := make([]int, n+1), make([]int, 0, n)
	for i, v := range nums {
		pre[i+1] = pre[i] + v
	}
	for i, v := range pre {
		for len(deque) > 0 && v-pre[deque[0]] >= k { // 找到最短子数组
			ans = min(ans, i-deque[0])
			deque = deque[1:]
		}
		for len(deque) > 0 && v <= pre[deque[len(deque)-1]] { // 最重要的逻辑：说明从 i 倒序查 nums 的值 <=0
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans

	// 滑动窗体：不能 AC，根本原因在于，没有跳过“递减”的情况
	// 例如：nums = []int{84, -37, 32, 40, 95}，k = 167，pre = [0 84 47 79 119 214]
	// 此时应该跳过 84 而计算 47
	//n := len(nums)
	//minVal := func(a, b int) int {
	//	if b < a {
	//		return b
	//	}
	//	return a
	//}
	//pre := make([]int, n+1)
	//for i, v := range nums {
	//	pre[i+1] = v + pre[i]
	//}
	//fmt.Println(pre)
	//ss, i := n+1, 0
	//for i < n && nums[i] < 0 {
	//	i++
	//}
	//for j := 0; i <= n; i++ {
	//	for ; j < i && pre[i]-pre[j] >= k; j++ {
	//		ss = minVal(ss, i-j)
	//	}
	//}
	//if ss > n {
	//	return -1
	//}
	//return ss
}
