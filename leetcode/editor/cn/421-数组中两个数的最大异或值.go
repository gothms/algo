//给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入：nums = [3,10,5,25,2,8]
//输出：28
//解释：最大运算结果是 5 XOR 25 = 28.
//
// 示例 2：
//
//
//输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
//输出：127
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2 * 10⁵
// 0 <= nums[i] <= 2³¹ - 1
//
//
// Related Topics 位运算 字典树 数组 哈希表 👍 558 👎 0

package main

import (
	"fmt"
	"math/bits"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findMaximumXOR(nums []int) int {
	// 个人：找出第二个数的集合
	// lc
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	max, maxV := 0, 0
	for _, v := range nums {
		maxV = maxVal(maxV, v) // maxV 最高位可视为第一个数的集合
	}
	k := bits.Len(uint(maxV))
	for i := k - 1; i >= 0; i-- { // 第 i 位是否 ok
		temp := make(map[int]bool)
		for _, v := range nums {
			temp[v>>i] = true
		}
		max = max<<1 + 1
		found := false
		for _, v := range nums {
			t := v >> i
			if temp[t^max] { // 关键逻辑，见 lc
				found = true // 有第 i 位
				break
			}
		}
		fmt.Println(i, max, temp)
		if !found { // 没有第 i 位
			max--
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
