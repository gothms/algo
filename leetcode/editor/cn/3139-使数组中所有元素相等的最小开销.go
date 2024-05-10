//给你一个整数数组 nums 和两个整数 cost1 和 cost2 。你可以执行以下 任一 操作 任意 次：
//
//
// 从 nums 中选择下标 i 并且将 nums[i] 增加 1 ，开销为 cost1。
// 选择 nums 中两个 不同 下标 i 和 j ，并且将 nums[i] 和 nums[j] 都 增加 1 ，开销为 cost2 。
//
//
// 你的目标是使数组中所有元素都 相等 ，请你返回需要的 最小开销 之和。
//
// 由于答案可能会很大，请你将它对 10⁹ + 7 取余 后返回。
//
//
//
// 示例 1：
//
//
// 输入：nums = [4,1], cost1 = 5, cost2 = 2
//
//
// 输出：15
//
// 解释：
//
// 执行以下操作可以使数组中所有元素相等：
//
//
// 将 nums[1] 增加 1 ，开销为 5 ，nums 变为 [4,2] 。
// 将 nums[1] 增加 1 ，开销为 5 ，nums 变为 [4,3] 。
// 将 nums[1] 增加 1 ，开销为 5 ，nums 变为 [4,4] 。
//
//
// 总开销为 15 。
//
// 示例 2：
//
//
// 输入：nums = [2,3,3,3,5], cost1 = 2, cost2 = 1
//
//
// 输出：6
//
// 解释：
//
// 执行以下操作可以使数组中所有元素相等：
//
//
// 将 nums[0] 和 nums[1] 同时增加 1 ，开销为 1 ，nums 变为 [3,4,3,3,5] 。
// 将 nums[0] 和 nums[2] 同时增加 1 ，开销为 1 ，nums 变为 [4,4,4,3,5] 。
// 将 nums[0] 和 nums[3] 同时增加 1 ，开销为 1 ，nums 变为 [5,4,4,4,5] 。
// 将 nums[1] 和 nums[2] 同时增加 1 ，开销为 1 ，nums 变为 [5,5,5,4,5] 。
// 将 nums[3] 增加 1 ，开销为 2 ，nums 变为 [5,5,5,5,5] 。
//
//
// 总开销为 6 。
//
// 示例 3：
//
//
// 输入：nums = [3,5,3], cost1 = 1, cost2 = 3
//
//
// 输出：4
//
// 解释：
//
// 执行以下操作可以使数组中所有元素相等：
//
//
// 将 nums[0] 增加 1 ，开销为 1 ，nums 变为 [4,5,3] 。
// 将 nums[0] 增加 1 ，开销为 1 ，nums 变为 [5,5,3] 。
// 将 nums[2] 增加 1 ，开销为 1 ，nums 变为 [5,5,4] 。
// 将 nums[2] 增加 1 ，开销为 1 ，nums 变为 [5,5,5] 。
//
//
// 总开销为 4 。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁶
// 1 <= cost1 <= 10⁶
// 1 <= cost2 <= 10⁶
//
//
// Related Topics 贪心 数组 枚举 👍 4 👎 0

package main

import (
	"fmt"
)

func main() {
	nums := []int{55, 52, 29, 11}
	c1, c2 := 18, 2
	array := minCostToEqualizeArray(nums, c1, c2)
	fmt.Println(array)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minCostToEqualizeArray(nums []int, cost1 int, cost2 int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

//func minCostToEqualizeArray(nums []int, cost1 int, cost2 int) int {
//	// lc：https://leetcode.cn/problems/minimum-cost-to-equalize-array/solutions/2766600/fen-lei-tao-lun-on-zuo-fa-pythonjavacgo-9bsb4/
//	const mod = 1_000_000_007
//	minV, maxV, s, n := math.MaxInt32, 0, 0, len(nums)
//	for _, v := range nums {
//		s += v
//		minV = min(minV, v)
//		maxV = max(maxV, v)
//	}
//	s = n*maxV - s
//	if n <= 2 || cost1<<1 <= cost2 {
//		return s * cost1 % mod
//	}
//	f := func(x int) int {
//		t, d := s+(x-maxV)*n, x-minV
//		if d<<1 <= t {
//			return (t>>1)*cost2 + (t&1)*cost1
//		}
//		return (t-d)*cost2 + (d<<1-t)*cost1
//	}
//	i := (n*maxV - minV<<1 - s + n - 3) / (n - 2)
//	if i <= maxV {
//		return min(f(maxV), f(maxV+1)) % mod
//	}
//	return min(f(maxV), f(i-1), f(i), f(i+1)) % mod
//}
