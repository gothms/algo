//给你一个下标从 0 开始的正整数数组 nums 。请你找出并统计满足下述条件的三元组 (i, j, k) 的数目：
//
//
// 0 <= i < j < k < nums.length
// nums[i]、nums[j] 和 nums[k] 两两不同 。
//
// 换句话说：nums[i] != nums[j]、nums[i] != nums[k] 且 nums[j] != nums[k] 。
//
//
//
// 返回满足上述条件三元组的数目。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,4,2,4,3]
//输出：3
//解释：下面列出的三元组均满足题目条件：
//- (0, 2, 4) 因为 4 != 2 != 3
//- (1, 2, 4) 因为 4 != 2 != 3
//- (2, 3, 4) 因为 2 != 4 != 3
//共计 3 个三元组，返回 3 。
//注意 (2, 0, 4) 不是有效的三元组，因为 2 > 0 。
//
//
// 示例 2：
//
//
//输入：nums = [1,1,1,1,1]
//输出：0
//解释：不存在满足条件的三元组，所以返回 0 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 100
// 1 <= nums[i] <= 1000
//
//
// Related Topics 数组 哈希表 👍 22 👎 0

package main

import "fmt"

func main() {
	nums := []int{4, 4, 2, 4, 3}
	triplets := unequalTriplets(nums)
	fmt.Println(triplets)
}

/*
思路：
	1.如果数字不会重复，就是一个组合问题
		那么数字重复时，怎么计算三元组的数目呢？
	2.考虑 0<=i<j<k<nums.length，nums[j] 组成的三元组数目计算方式
		nums[j] 值唯一：f(j) = j * (nums.length-j-1)
			nums[j] 左边有 j 个数，右边有 nums.length-j-1 个数
		nums[j] 值唯一，有 v 个：f(j) = j * v * nums.length-j-v
			这里假设 nums[j] 第一次出现，则左边有j个数，右边有nums.length-j-v
	3.用 map 统计每个数值出现的次数，遍历 map 并按照 2. 的思路计算三元组数目
*/
// leetcode submit region begin(Prohibit modification and deletion)
func unequalTriplets(nums []int) int {
	cache, n := make(map[int]int), len(nums)
	for i := 0; i < n; i++ {
		cache[nums[i]]++
	}
	if len(cache) < 3 {
		return 0
	}
	cnt, l, r := 0, 0, n
	for _, v := range cache {
		r -= v
		cnt, l = cnt+l*v*r, l+v
	}
	return cnt

	//cache, n := make(map[int]int), len(nums)
	//for i := 0; i < n; i++ {
	//	cache[nums[i]]++
	//}
	//if len(cache) < 3 {
	//	return 0
	//}
	//cnt, pre := 0, 0
	//for _, v := range cache {
	//	cnt, pre = cnt+pre*v*(n-pre-v), pre+v
	//}
	//return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
