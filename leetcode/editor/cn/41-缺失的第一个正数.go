//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。 请你实现时间复杂度为
//O(n) 并且只使用常数级别额外空间的解决方案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,0]
//输出：3
//
//
// 示例 2：
//
//
//输入：nums = [3,4,-1,1]
//输出：2
//
//
// 示例 3：
//
//
//输入：nums = [7,8,9,11,12]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
// Related Topics 数组 哈希表 👍 1834 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 2, 0}
	nums = []int{3, 4, -1, 1}
	nums = []int{7, 8, 9, 11, 12}
	nums = []int{1}
	//nums = []int{1, 1}
	positive := firstMissingPositive(nums)
	fmt.Println(positive)
}

// leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

//func firstMissingPositive(nums []int) int {
//	abs := func(v int) int {
//		if v < 0 {
//			return -v
//		}
//		return v
//	}
//	n := len(nums) + 1
//	for i, v := range nums { // 只保留 [1,n) 的值
//		if v <= 0 {
//			nums[i] = n
//		}
//	}
//	for _, v := range nums { // 将 [1,n) 的值对号入座 [0,len(nums)-1]
//		v = abs(v)
//		if v < n {
//			nums[v-1] = -abs(nums[v-1])
//		}
//	}
//	for i, v := range nums { // 第一个没有被入座的位置
//		if v > 0 {
//			return i + 1
//		}
//	}
//	return n
//}
