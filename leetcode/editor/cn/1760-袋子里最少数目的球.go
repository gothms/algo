//给你一个整数数组 nums ，其中 nums[i] 表示第 i 个袋子里球的数目。同时给你一个整数 maxOperations 。
//
// 你可以进行如下操作至多 maxOperations 次：
//
//
// 选择任意一个袋子，并将袋子里的球分到 2 个新的袋子中，每个袋子里都有 正整数 个球。
//
//
//
// 比方说，一个袋子里有 5 个球，你可以把它们分到两个新袋子里，分别有 1 个和 4 个球，或者分别有 2 个和 3 个球。
//
//
//
//
// 你的开销是单个袋子里球数目的 最大值 ，你想要 最小化 开销。
//
// 请你返回进行上述操作后的最小开销。
//
//
//
// 示例 1：
//
//
//输入：nums = [9], maxOperations = 2
//输出：3
//解释：
//- 将装有 9 个球的袋子分成装有 6 个和 3 个球的袋子。[9] -> [6,3] 。
//- 将装有 6 个球的袋子分成装有 3 个和 3 个球的袋子。[6,3] -> [3,3,3] 。
//装有最多球的袋子里装有 3 个球，所以开销为 3 并返回 3 。
//
//
// 示例 2：
//
//
//输入：nums = [2,4,8,2], maxOperations = 4
//输出：2
//解释：
//- 将装有 8 个球的袋子分成装有 4 个和 4 个球的袋子。[2,4,8,2] -> [2,4,4,4,2] 。
//- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,4,4,4,2] -> [2,2,2,4,4,2] 。
//- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,2,2,4,4,2] -> [2,2,2,2,2,4,2] 。
//- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,2,2,2,2,4,2] -> [2,2,2,2,2,2,2,2] 。
//装有最多球的袋子里装有 2 个球，所以开销为 2 并返回 2 。
//
//
// 示例 3：
//
//
//输入：nums = [7,17], maxOperations = 2
//输出：7
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= maxOperations, nums[i] <= 10⁹
//
//
// Related Topics 数组 二分查找 👍 253 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{9}
	maxOperations := 2
	nums = []int{2, 4, 8, 2}
	maxOperations = 4
	nums = []int{7, 17}
	maxOperations = 2
	nums = []int{1, 1, 1}
	maxOperations = 2
	nums = []int{431, 922, 158, 60, 192, 14, 788, 146, 788, 775, 772, 792, 68, 143, 376, 375, 877, 516, 595, 82, 56, 704, 160, 403, 713, 504, 67, 332, 26}
	maxOperations = 80 // 129
	//sort.Ints(nums)
	//cnt := 0
	//for _, v := range nums[7:] {
	//	//cnt += (v - 1) / 128 // 81
	//	cnt += v / (128 + 1) //
	//}
	//fmt.Println(cnt)

	//nums = []int{99, 99, 99, 99}
	//maxOperations = 3
	size := minimumSize(nums, maxOperations)
	fmt.Println(size)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumSize(nums []int, maxOperations int) int {
	// 二分查找：个人
	//sort.Ints(nums) // 为了二分：2
	//s, maxV, n := 0, 0, len(nums)
	//if nums[0] == nums[n-1] { // fast path
	//	return max(1, (nums[0]-1)/(maxOperations/n+1)+1)
	//}
	//for _, v := range nums {
	//	s += v              // 和
	//	maxV = max(maxV, v) // 最大值
	//}
	//minV := max(1, (s+1)/(n+maxOperations))
	//v := sort.Search(maxV-minV, func(i int) bool { // 二分：1
	//	op, t := 0, i+minV                     // t 可能是 0，防止 /0，则初始化 minV >= 1
	//	j := sort.Search(n, func(j int) bool { // 二分：2
	//		return nums[j] > t
	//	})
	//	for _, v := range nums[j:] {
	//		op += (v - 1) / t
	//		if op > maxOperations {
	//			return false
	//		}
	//	}
	//	//return op <= maxOperations  // 操作次数 <= maxOperations 为合法
	//	return true
	//}) + minV
	//return v

	// 二分
	s, maxV, n := 0, 0, len(nums)
	for _, v := range nums {
		s += v
		maxV = max(maxV, v)
	}
	minV := max(1, (s+1)/(n+maxOperations))
	v := sort.Search(maxV-minV, func(i int) bool {
		op, t := 0, i+minV
		for _, v := range nums {
			if op += (v - 1) / t; op > maxOperations {
				return false
			}
		}
		return true
	}) + minV
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
