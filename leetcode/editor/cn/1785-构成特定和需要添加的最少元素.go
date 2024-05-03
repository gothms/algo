//给你一个整数数组 nums ，和两个整数 limit 与 goal 。数组 nums 有一条重要属性：abs(nums[i]) <= limit 。
//
// 返回使数组元素总和等于 goal 所需要向数组中添加的 最少元素数量 ，添加元素 不应改变 数组中 abs(nums[i]) <= limit 这一属性。
//
//
// 注意，如果 x >= 0 ，那么 abs(x) 等于 x ；否则，等于 -x 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,-1,1], limit = 3, goal = -4
//输出：2
//解释：可以将 -2 和 -3 添加到数组中，数组的元素总和变为 1 - 1 + 1 - 2 - 3 = -4 。
//
//
// 示例 2：
//
//
//输入：nums = [1,-10,9,1], limit = 100, goal = 0
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= limit <= 10⁶
// -limit <= nums[i] <= limit
// -10⁹ <= goal <= 10⁹
//
//
// Related Topics 贪心 数组 👍 56 👎 0

package main

import "fmt"

func main() {
	nums := []int{2, 2, 2, 5, 1, -2}
	limit, goal := 300, 126614243 // 25322847
	nums = []int{1, 2, 3}
	limit, goal = 5, 6 // 25322847
	elements := minElements(nums, limit, goal)
	fmt.Println(elements)
	//fmt.Println(-298 % 99)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minElements(nums []int, limit int, goal int) int {
	s := goal
	for _, v := range nums {
		s -= v
	}
	if s == 0 { // 刚好
		return 0
	} else if s < 0 { // 转为正数
		s = -s
	}
	return (s-1)/limit + 1 // 向上取整
}

//leetcode submit region end(Prohibit modification and deletion)
