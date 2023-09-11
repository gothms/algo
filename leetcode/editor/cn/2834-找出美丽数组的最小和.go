//给你两个正整数：n 和 target 。
//
// 如果数组 nums 满足下述条件，则称其为 美丽数组 。
//
//
// nums.length == n.
// nums 由两两互不相同的正整数组成。
// 在范围 [0, n-1] 内，不存在 两个 不同 下标 i 和 j ，使得 nums[i] + nums[j] == target 。
//
//
// 返回符合条件的美丽数组所可能具备的 最小 和。
//
//
//
// 示例 1：
//
//
//输入：n = 2, target = 3
//输出：4
//解释：nums = [1,3] 是美丽数组。
//- nums 的长度为 n = 2 。
//- nums 由两两互不相同的正整数组成。
//- 不存在两个不同下标 i 和 j ，使得 nums[i] + nums[j] == 3 。
//可以证明 4 是符合条件的美丽数组所可能具备的最小和。
//
// 示例 2：
//
//
//输入：n = 3, target = 3
//输出：8
//解释：
//nums = [1,3,4] 是美丽数组。
//- nums 的长度为 n = 3 。
//- nums 由两两互不相同的正整数组成。
//- 不存在两个不同下标 i 和 j ，使得 nums[i] + nums[j] == 3 。
//可以证明 8 是符合条件的美丽数组所可能具备的最小和。
//
// 示例 3：
//
//
//输入：n = 1, target = 1
//输出：1
//解释：nums = [1] 是美丽数组。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁵
// 1 <= target <= 10⁵
//
//
// Related Topics 贪心 数学 👍 9 👎 0

package main

import "fmt"

func main() {
	n := 39636
	target := 49035 // 156198582
	possibleSum := minimumPossibleSum(n, target)
	fmt.Println(possibleSum) // 601,107,806

	var v32 int32 = 156198582
	var v64 int64 = 1156198589
	fmt.Printf("%d, %[1]b\n", v32)
	fmt.Printf("%d, %[1]b\n", v64)

	/*
		   1001010011110110011010110110
		1000100111010100011000010111101
	*/
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumPossibleSum(n int, target int) int {
	// lc-2829
	//m := (target - 1) >> 1
	//if target-m > n {
	//	return (n + 1) * n >> 1
	//}
	//return (n+1)*n>>1 + (n-target+m+1)*m

	const mod = 1_000_000_007 // 官方改了题，但是没写要 % 1_000_000_007
	sum := func(f, t int) int {
		return (f + t) * (t - f + 1) >> 1
	}
	if n <= target>>1 {
		return sum(1, n)
	}
	return (sum(1, target>>1) + sum(target, target+n-target>>1-1)) % mod

	//lc
	//minVal := func(a, b int) int {
	//	if b < a {
	//		return b
	//	}
	//	return a
	//}
	//m := minVal(target/2, n)
	//return int((int64(m)*int64(m+1) + int64(target*2+n-m-1)*int64(n-m)) / 2)
}

// leetcode submit region end(Prohibit modification and deletion)
