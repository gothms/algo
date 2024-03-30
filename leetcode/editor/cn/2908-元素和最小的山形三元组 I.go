//给你一个下标从 0 开始的整数数组 nums 。
//
// 如果下标三元组 (i, j, k) 满足下述全部条件，则认为它是一个 山形三元组 ：
//
//
// i < j < k
// nums[i] < nums[j] 且 nums[k] < nums[j]
//
//
// 请你找出 nums 中 元素和最小 的山形三元组，并返回其 元素和 。如果不存在满足条件的三元组，返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：nums = [8,6,1,5,3]
//输出：9
//解释：三元组 (2, 3, 4) 是一个元素和等于 9 的山形三元组，因为：
//- 2 < 3 < 4
//- nums[2] < nums[3] 且 nums[4] < nums[3]
//这个三元组的元素和等于 nums[2] + nums[3] + nums[4] = 9 。可以证明不存在元素和小于 9 的山形三元组。
//
//
// 示例 2：
//
//
//输入：nums = [5,4,8,7,10,2]
//输出：13
//解释：三元组 (1, 3, 5) 是一个元素和等于 13 的山形三元组，因为：
//- 1 < 3 < 5
//- nums[1] < nums[3] 且 nums[5] < nums[3]
//这个三元组的元素和等于 nums[1] + nums[3] + nums[5] = 13 。可以证明不存在元素和小于 13 的山形三元组。
//
//
// 示例 3：
//
//
//输入：nums = [6,5,4,3,4,5]
//输出：-1
//解释：可以证明 nums 中不存在山形三元组。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 50
// 1 <= nums[i] <= 50
//
//
// Related Topics 数组 👍 11 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{5, 4, 8, 7, 10, 2}
	//nums = []int{1, 2, 3, 2}
	//nums = []int{2, 2, 1}
	nums = []int{2, 3, 2, 1}
	//nums = []int{3, 4, 2, 1}
	i := minimumSum(nums)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumSum(nums []int) int {
	// 前后缀
	ret, m, n := math.MaxInt32, math.MaxInt32, len(nums)
	pre := make([]int, n)
	for i, v := range nums { // 前缀最小值
		m = min(m, v)
		pre[i] = m
	}
	m = nums[n-1]
	for i := n - 2; i > 0; i-- {
		if nums[i] > pre[i-1] && nums[i] > m { // 山形
			ret = min(ret, pre[i-1]+nums[i]+m)
		}
		m = min(m, nums[i]) // 更新后缀最小值
	}
	if ret == math.MaxInt32 {
		return -1
	}
	return ret

	// stack：失败
	//ret, st := math.MaxInt32, []int{nums[0]}
	//for _, v := range nums[1:] {
	//	n := len(st)
	//	idx := sort.Search(n, func(i int) bool {
	//		return st[i] > v
	//	})
	//	if n >= 2 && idx < n {
	//		if idx == 0 {
	//			ret = min(ret, st[0]+st[1]+v)
	//		} else {
	//			ret = min(ret, st[0]+st[idx]+v)
	//		}
	//	}
	//	if idx == 0 && n == 1 || idx > 0 && st[idx-1] != v {
	//		st = append(st[:idx], v)
	//	}
	//	//fmt.Println(idx, st, ret)
	//}
	//if ret == math.MaxInt32 {
	//	return -1
	//}
	//return ret
}

//leetcode submit region end(Prohibit modification and deletion)
