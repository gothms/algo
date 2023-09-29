//给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
//
//
//
// 示例 1：
//
//
//输入：nums = [3,2,3]
//输出：[3]
//
// 示例 2：
//
//
//输入：nums = [1]
//输出：[1]
//
//
// 示例 3：
//
//
//输入：nums = [1,2]
//输出：[1,2]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
//
// Related Topics 数组 哈希表 计数 排序 👍 698 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 9, 2, 9, 3, 9}
	element := majorityElement(nums)
	fmt.Println(element)
}

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}
	v1, v2, c1, c2 := math.MinInt32, math.MinInt32, 0, 0
	for _, v := range nums {
		switch v {
		case v1:
			c1++
		case v2:
			c2++
		default:
			switch {
			case c1 == 0:
				v1, c1 = v, 1
			case c2 == 0:
				v2, c2 = v, 1
			default:
				c1--
				c2--
			}
		}
	}
	c1, c2 = 0, 0
	for _, v := range nums {
		switch v {
		case v1:
			c1++
		case v2:
			c2++
		}
	}
	ret := make([]int, 0, 2)
	if c1 > (n / 3) {
		ret = append(ret, v1)
	}
	if c2 > (n / 3) {
		ret = append(ret, v2)
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
