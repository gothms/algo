//给你一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
//
// 你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,1,3,2,5]
//输出：[3,5]
//解释：[5, 3] 也是有效的答案。
//
//
// 示例 2：
//
//
//输入：nums = [-1,0]
//输出：[-1,0]
//
//
// 示例 3：
//
//
//输入：nums = [0,1]
//输出：[1,0]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 3 * 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
// 除两个只出现一次的整数外，nums 中的其他数字都出现两次
//
//
// Related Topics 位运算 数组 👍 746 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	number := singleNumber(nums)
	fmt.Println(number)
}

// leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) []int {
	// 异或
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	one, bit := 0, xor&-xor
	for _, v := range nums {
		if bit&v != 0 {
			one ^= v
		}
	}
	return []int{one, one ^ xor}

	//xor := 0
	//for _, v := range nums {
	//	xor ^= v
	//}
	//lowBit, x := xor&-xor, xor
	//for _, v := range nums {
	//	if v&lowBit != 0 {
	//		xor ^= v
	//	}
	//}
	//return []int{xor, xor ^ x}
}

//leetcode submit region end(Prohibit modification and deletion)
