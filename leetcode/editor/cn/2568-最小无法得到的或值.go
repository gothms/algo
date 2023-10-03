//给你一个下标从 0 开始的整数数组 nums 。
//
// 如果存在一些整数满足 0 <= index1 < index2 < ... < indexk < nums.length ，得到 nums[index1]
// | nums[index2] | ... | nums[indexk] = x ，那么我们说 x 是 可表达的 。换言之，如果一个整数能由 nums 的某个子
//序列的或运算得到，那么它就是可表达的。
//
// 请你返回 nums 不可表达的 最小非零整数 。
//
//
//
// 示例 1：
//
// 输入：nums = [2,1]
//输出：4
//解释：1 和 2 已经在数组中，因为 nums[0] | nums[1] = 2 | 1 = 3 ，所以 3 是可表达的。由于 4 是不可表达的，所以我们返
//回 4 。
//
//
// 示例 2：
//
// 输入：nums = [5,3,2]
//输出：1
//解释：1 是最小不可表达的数字。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁹
//
//
// Related Topics 位运算 脑筋急转弯 数组 👍 16 👎 0

package main

import (
	"fmt"
	"math/bits"
)

func main() {
	l := bits.Len(1e9)
	fmt.Println(l)
	fmt.Println(1 << 30)
	nums := []int{1, 2, 4, 8}
	or := minImpossibleOR(nums)
	fmt.Println(or)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minImpossibleOR(nums []int) int {
	// 2^n 覆盖的范围内的整数，都能被 或运算 达到
	//const L = 31
	memo := make(map[int]bool, 31)
	for _, v := range nums {
		if bits.OnesCount(uint(v)) == 1 {
			memo[v] = true
		}
	}
	//for i := 0; i < L; i++ {
	//	if !memo[1<<i] {
	//		return 1 << i
	//	}
	//}
	for i := 1; ; i <<= 1 {
		if !memo[i] {
			return i
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
