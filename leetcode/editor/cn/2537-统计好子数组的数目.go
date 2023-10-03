//给你一个整数数组 nums 和一个整数 k ，请你返回 nums 中 好 子数组的数目。
//
// 一个子数组 arr 如果有 至少 k 对下标 (i, j) 满足 i < j 且 arr[i] == arr[j] ，那么称它是一个 好 子数组。
//
// 子数组 是原数组中一段连续 非空 的元素序列。
//
//
//
// 示例 1：
//
// 输入：nums = [1,1,1,1,1], k = 10
//输出：1
//解释：唯一的好子数组是这个数组本身。
//
//
// 示例 2：
//
// 输入：nums = [3,1,4,3,2,2,4], k = 2
//输出：4
//解释：总共有 4 个不同的好子数组：
//- [3,1,4,3,2,2] 有 2 对。
//- [3,1,4,3,2,2,4] 有 3 对。
//- [1,4,3,2,2,4] 有 2 对。
//- [4,3,2,2,4] 有 2 对。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i], k <= 10⁹
//
//
// Related Topics 数组 哈希表 滑动窗口 👍 41 👎 0

package main

import "fmt"

func main() {
	//fmt.Println(-1 & 1)
	nums := []int{1, 1, 1, 1, 1}
	nums = []int{3, 1, 4, 3, 2, 2, 4}
	k := 10
	k = 2
	good := countGood(nums, k)
	fmt.Println(good)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countGood(nums []int, k int) int64 {
	// 优化
	cg, cnt, l, n := int64(0), 0, 0, len(nums)
	memo := make(map[int]int)
	for r, v := range nums {
		// x 个相同的数，组成 (x-1)*x / 2 对
		cnt += memo[v] // 在 m[v]++ 之前，对数的增量为 m[v]
		memo[v]++
		for cnt >= k {
			memo[nums[l]]--
			cnt -= memo[nums[l]] // 对数的增量为 m[v]
			l++
			cg += int64(n - r) // r 及其后面的元素，都能成为一个好子数组
		}
	}
	return cg

	// 滑动窗体
	//cg, cnt, l, n := int64(0), 0, 0, len(nums)
	//memo := make(map[int]int)
	//for r, v := range nums {
	//	// x 个相同的数，组成 (x-1)*x / 2 对
	//	cnt += memo[v] // 在 m[v]++ 之前，对数的增量为 m[v]
	//	memo[v]++
	//	if cnt < k {
	//		continue
	//	}
	//	cg += int64(n - r) // r 及其后面的元素，都能成为一个好子数组
	//	for {
	//		memo[nums[l]]--
	//		cnt -= memo[nums[l]] // 对数的增量为 m[v]
	//		l++
	//		if cnt < k {
	//			break
	//		}
	//		cg += int64(n - r)
	//	}
	//}
	//return cg
}

//leetcode submit region end(Prohibit modification and deletion)
