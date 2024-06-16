package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4}      // 23
	nums = []int{1, 2, 3, 4, 5, 6} // 201
	widths := sumSubseqWidths(nums)
	fmt.Println(widths)
}

// leetcode submit region begin(Prohibit modification and deletion)
func sumSubseqWidths(nums []int) int {
	// lc
	// 考虑 x=nums[i]，x 作为最大值的序列数为 2^i，x 作为最小值的序列数为 2^(n-1-i)
	// 则 x 的贡献值为 (2^i-2^(n-1-i)) * x

	// dp：个人
	// 排序后，对于包含最小值和最大值 nums[j] nums[i] 的子序列一共有 sum = 2^(i-j-1) 个
	// 分别计算最大值和最小值的总数，最大值总数量 fMax(i) = fMax(i-1)<<1 + 1，则 fMax(i) * nums[i] 为增加量
	// 对于位置 i 的最小值总和 sumMin(i) = sumMin(i-1)<<1 + nums[i-1]
	const mod = 1_000_000_007
	sort.Ints(nums)
	n := len(nums)
	addSum, subSum, powSum := 0, 0, 1
	for i := 1; i < n; i++ {
		addSum = (addSum + powSum*nums[i]%mod - subSum + mod) % mod // f(i) 为 i 位置的 sumSubseqWidths
		subSum = (subSum<<1 + nums[i-1]) % mod                      // i 位置需要减去的 最小值 的总和
		powSum = (powSum<<1 + 1) % mod
	}
	return (addSum - subSum + mod) % mod
}

//leetcode submit region end(Prohibit modification and deletion)
