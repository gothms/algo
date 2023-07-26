//给你一个整数数组 nums，请你找出并返回能被三整除的元素最大和。
//
//
//
//
//
//
// 示例 1：
//
// 输入：nums = [3,6,5,1,8]
//输出：18
//解释：选出数字 3, 6, 1 和 8，它们的和是 18（可被 3 整除的最大和）。
//
// 示例 2：
//
// 输入：nums = [4]
//输出：0
//解释：4 不能被 3 整除，所以无法选出数字，返回 0。
//
//
// 示例 3：
//
// 输入：nums = [1,2,3,4,4]
//输出：12
//解释：选出数字 1, 3, 4 以及 4，它们的和是 12（可被 3 整除的最大和）。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 4 * 10^4
// 1 <= nums[i] <= 10^4
//
//
// Related Topics 贪心 数组 动态规划 排序 👍 215 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{3, 6, 5, 1, 8, 1}
	//nums = []int{1, 2, 3, 4, 4}
	three := maxSumDivThree(nums)
	fmt.Println(three)
}

/*
思路：排序
	1.所有元素的和除以3的余数只可能是 0 / 1 / 2，即 mod = 0/1/2
	2.最后数组中没有被”挑选“的数只可能有这几种情况：
		2.1.mod=0：没有数剩下
		2.2.mod=1：剩下两个 %3=2 的数，或者一个 %3=1 的数
		2.3.mod=2：剩下一个 %3=2 的数，或者两个 %3=1 的数
	3.当 2.2. 时，两个 %3=2 的数的和 < 一个 %3=1 的数，则挑前者，否则挑后者
	4.所以可能用两个集合，分别记录  %3=2 的数，和 %3=1 的数
		分别排序后，进行挑选
		未被挑选的数的和，即为所求
思路：dp
	对于任意 i，当前 元素最大和 %3 的状态 dp[i]，分3种情况：
		mod=0：
			dp[i][0] = max(dp[i][0], dp[i-1][0]+nums[i])
			dp[i][1] = max(dp[i][1], dp[i-1][1]+nums[i])
			dp[i][2] = max(dp[i][2], dp[i-1][2]+nums[i])
		mod=1：注意 dp[i-1][j]+nums[i]，(j+mod)%3 = 0/1/2
			dp[i][0] = max(dp[i][0], dp[i-1][2]+nums[i])
			dp[i][1] = max(dp[i][1], dp[i-1][0]+nums[i])
			dp[i][2] = max(dp[i][2], dp[i-1][1]+nums[i])
		mod=2：类似 mod=1
*/
// leetcode submit region begin(Prohibit modification and deletion)
func maxSumDivThree(nums []int) int {
	// dp
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp, n := [3]int{0, math.MinInt32, math.MinInt32}, len(nums)
	for i := 0; i < n; i++ {
		mod := nums[i] % 3
		dp[0], dp[1], dp[2] =
			maxVal(dp[0], dp[(3-mod)%3]+nums[i]),
			maxVal(dp[1], dp[(4-mod)%3]+nums[i]),
			maxVal(dp[2], dp[2-mod]+nums[i])
	}
	return dp[0]

	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//dp := [3]int{0, math.MinInt32, math.MinInt32}
	//for _, v := range nums {
	//	temp := [3]int{}
	//	for i := 0; i < 3; i++ {
	//		temp[(i+v)%3] = maxVal(dp[(i+v)%3], dp[i]+v)
	//	}
	//	dp = temp
	//}
	//return dp[0]

	// 排序
	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//sum, n, one, two := 0, len(nums), make([]int, 0), make([]int, 0)
	//for i := 0; i < n; i++ {
	//	sum += nums[i] // 总和
	//	switch nums[i] % 3 {
	//	case 1:
	//		one = append(one, nums[i])
	//	case 2:
	//		two = append(two, nums[i])
	//	}
	//}
	//sort.Ints(one)
	//sort.Ints(two)
	//if sum%3 == 0 {
	//	return sum
	//} else if sum%3 == 2 {
	//	one, two = two, one
	//}
	//if len(two) > 1 { // 挑出剩下的数，同时要注意两个集合的长度，避免越界
	//	if len(one) > 0 {
	//		sum -= minVal(two[0]+two[1], one[0])
	//	} else {
	//		sum -= two[0] + two[1]
	//	}
	//} else {
	//	sum -= one[0]
	//}
	//return sum

	// 错误
	//sort.Ints(nums)
	//sum, n := 0, len(nums)-1
	//f, s := 0, 0
	//div := func(t, v int) {
	//	if f == 0 {
	//		f = v
	//		return
	//	}
	//	if s == 0 && f%3 != t {
	//		s = v
	//		return
	//	}
	//	if f%3 == t {
	//		sum += v + f
	//		f, s = s, 0
	//	} else if s != 0 {
	//		if s%3 == t {
	//			sum += v + s
	//			s = 0
	//		} else {
	//			sum += v + f + s
	//			f, s = 0, 0
	//		}
	//	}
	//}
	//for i := n; i >= 0; i-- {
	//	switch nums[i] % 3 {
	//	case 0:
	//		sum += nums[i]
	//	case 1:
	//		div(2, nums[i])
	//	case 2:
	//		div(1, nums[i])
	//	}
	//}
	//return sum
}

//leetcode submit region end(Prohibit modification and deletion)
