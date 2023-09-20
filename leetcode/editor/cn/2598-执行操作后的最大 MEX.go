//给你一个下标从 0 开始的整数数组 nums 和一个整数 value 。
//
// 在一步操作中，你可以对 nums 中的任一元素加上或减去 value 。
//
//
// 例如，如果 nums = [1,2,3] 且 value = 2 ，你可以选择 nums[0] 减去 value ，得到 nums = [-1,2,3]
//。
//
//
// 数组的 MEX (minimum excluded) 是指其中数组中缺失的最小非负整数。
//
//
// 例如，[-1,2,3] 的 MEX 是 0 ，而 [1,0,3] 的 MEX 是 2 。
//
//
// 返回在执行上述操作 任意次 后，nums 的最大 MEX 。
//
//
//
// 示例 1：
//
// 输入：nums = [1,-10,7,13,6,8], value = 5
//输出：4
//解释：执行下述操作可以得到这一结果：
//- nums[1] 加上 value 两次，nums = [1,0,7,13,6,8]
//- nums[2] 减去 value 一次，nums = [1,0,2,13,6,8]
//- nums[3] 减去 value 两次，nums = [1,0,2,3,6,8]
//nums 的 MEX 是 4 。可以证明 4 是可以取到的最大 MEX 。
//
//
// 示例 2：
//
// 输入：nums = [1,-10,7,13,6,8], value = 7
//输出：2
//解释：执行下述操作可以得到这一结果：
//- nums[2] 减去 value 一次，nums = [1,-10,0,13,6,8]
//nums 的 MEX 是 2 。可以证明 2 是可以取到的最大 MEX 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length, value <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 贪心 数组 哈希表 数学 👍 16 👎 0

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, -10, 7, 13, 6, 8}
	value := 5
	nums = []int{3, 0, 3, 2, 4, 2, 1, 1, 0, 4}
	nums = []int{0, -3}
	value = 4
	integer := findSmallestInteger(nums, value)
	fmt.Println(integer)
}

/*
思路：贪心
	1.操作次数可以是任意次，所以只需要计算有没有数能变成目标值 v
		如果有，则跳过
		没有，则返回 v 值
	2.示例：nums = [1,-10000,7,13,6,8], value = 5
		-10000 经过任意次操作，总能变成 0
		1 经过任意次操作，总能变成 1
		...
	3.以 value 为一个周期，比如 value = 5
		第一个周期为 0 1 2 3 4
		第二个周期为 5 6 7 8 9
		...
		所以模运算相同的值，对于每个周期是等价的
		只需要考虑模运算分组后，每组的元素个数是否够当前周期使用
*/
// leetcode submit region begin(Prohibit modification and deletion)
func findSmallestInteger(nums []int, value int) int {
	memo := make(map[int]int, value)
	for _, v := range nums {
		memo[(v%value+value)%value]++ // 模运算分组，并统计每组的元素个数
	}
	for i := 0; ; i++ {
		for v := 0; v < value; v++ { // 以 value 为一轮，检查分组的元素个数是否够
			if memo[v] == i { // 不够，则此处不能连续了
				return v + i*value // 返回断的地方
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
