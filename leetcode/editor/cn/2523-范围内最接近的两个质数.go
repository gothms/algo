//给你两个正整数 left 和 right ，请你找到两个整数 num1 和 num2 ，它们满足：
//
//
// left <= nums1 < nums2 <= right 。
// nums1 和 nums2 都是 质数 。
// nums2 - nums1 是满足上述条件的质数对中的 最小值 。
//
//
// 请你返回正整数数组 ans = [nums1, nums2] 。如果有多个整数对满足上述条件，请你返回 nums1 最小的质数对。如果不存在符合题意的质数
//对，请你返回 [-1, -1] 。
//
// 如果一个整数大于 1 ，且只能被 1 和它自己整除，那么它是一个质数。
//
//
//
// 示例 1：
//
//
//输入：left = 10, right = 19
//输出：[11,13]
//解释：10 到 19 之间的质数为 11 ，13 ，17 和 19 。
//质数对的最小差值是 2 ，[11,13] 和 [17,19] 都可以得到最小差值。
//由于 11 比 17 小，我们返回第一个质数对。
//
//
// 示例 2：
//
//
//输入：left = 4, right = 6
//输出：[-1,-1]
//解释：给定范围内只有一个质数，所以题目条件无法被满足。
//
//
//
//
// 提示：
//
//
// 1 <= left <= right <= 10⁶
//
//
// Related Topics 数学 数论 👍 19 👎 0

package main

import (
	"sort"
)

func main() {
	closestPrimes(1, 1_000_000)
}

// leetcode submit region begin(Prohibit modification and deletion)
func closestPrimes(left int, right int) []int {
	// 预处理
	idx, d := 0, CPN
	for i := sort.SearchInts(primes, left) + 1; primes[i] <= right; i++ { // 从第二个质数开始
		if dc := primes[i] - primes[i-1]; dc < d {
			d, idx = dc, i // 更新最小值和索引
		}
	}
	if idx == 0 {
		return []int{-1, -1}
	}
	return []int{primes[idx-1], primes[idx]}

	// 超时
	//	if left == 1 || left != 2 && left&1 == 0 {
	//		left++
	//	}
	//	prime, last, d := make([]int, 2), 0, math.MaxInt32
	//out:
	//	for i := left; i <= right; i++ {
	//		sqrt := int(math.Sqrt(float64(i)))
	//		for j := 2; j <= sqrt; j++ {
	//			if i%j == 0 {
	//				continue out
	//			}
	//		}
	//		if last != 0 && i-last < d {
	//			d = i - last
	//			prime[0], prime[1] = last, i
	//		}
	//		last = i
	//	}
	//	if d == math.MaxInt32 {
	//		return []int{-1, -1}
	//	}
	//	return prime
}

const CPN = 1_000_001

var primes = make([]int, 1, 78500) // 实际有 78498 个

func init() {
	primes[0] = 2
	memo := [CPN]bool{}
	for i := 3; i < CPN; i += 2 { // 找出质数
		if memo[i] {
			continue
		}
		primes = append(primes, i)             // 质数
		for j := i * i; j < CPN; j += i << 1 { // 非质数
			memo[j] = true
		}
	}
	primes = append(primes, CPN, CPN) // 方便二分查找
}

//leetcode submit region end(Prohibit modification and deletion)
