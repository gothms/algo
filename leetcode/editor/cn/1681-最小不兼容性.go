//给你一个整数数组 nums 和一个整数 k 。你需要将这个数组划分到 k 个相同大小的子集中，使得同一个子集里面没有两个相同的元素。
//
// 一个子集的 不兼容性 是该子集里面最大值和最小值的差。
//
// 请你返回将数组分成 k 个子集后，各子集 不兼容性 的 和 的 最小值 ，如果无法分成分成 k 个子集，返回 -1 。
//
// 子集的定义是数组中一些数字的集合，对数字顺序没有要求。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,1,4], k = 2
//输出：4
//解释：最优的分配是 [1,2] 和 [1,4] 。
//不兼容性和为 (2-1) + (4-1) = 4 。
//注意到 [1,1] 和 [2,4] 可以得到更小的和，但是第一个集合有 2 个相同的元素，所以不可行。
//
// 示例 2：
//
//
//输入：nums = [6,3,8,1,3,1,2,2], k = 4
//输出：6
//解释：最优的子集分配为 [1,2]，[2,3]，[6,8] 和 [1,3] 。
//不兼容性和为 (2-1) + (3-2) + (8-6) + (3-1) = 6 。
//
//
// 示例 3：
//
//
//输入：nums = [5,3,3,6,3,3], k = 3
//输出：-1
//解释：没办法将这些数字分配到 3 个子集且满足每个子集里没有相同数字。
//
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 16
// nums.length 能被 k 整除。
// 1 <= nums[i] <= nums.length
//
//
// Related Topics 位运算 数组 动态规划 状态压缩 👍 57 👎 0

package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	nums := []int{1, 2, 1, 4}
	k := 2
	//nums = []int{5, 3, 3, 6, 3, 3}6,3,8,1,3,1,2,2
	//k = 3
	//nums = []int{6, 3, 8, 1, 3, 1, 2, 2}
	//k = 4 // 6
	incompatibility := minimumIncompatibility(nums, k)
	fmt.Println(incompatibility)
}

/*
思路：状压dp
    1.由提示可知 1 <= nums[i] <= nums.length <= 16
		因为有重复的元素，所以相对于对元素进行状态压缩，考虑对唯一标识索引状态压缩
	2.对于任意元素个数为 k 的索引组合（子集 i），先求出其 不兼容性，并记录到 memo 映射中
		memo[i] = max - min，表示 索引子集i 的不兼容性，max min 为子集 i 中最大和最小元素值
	3.对于任意 [0, 1<<n) 区间的值 i，都是一个子集
		因为元素的无规律性，所以对于任意子集 i，不一定存在 f(i) 的解，即 f(i)=-1
		所以根据每个已求得的 f(i) 推出 子集i 的超集 f(i | idx) 的复杂度，要小于依次试图求出每个 f(i) 的复杂度
		其中 idx 表示 i 的补集的所有子集，如果 memo 中查询得 idx 合法，则可求出 f(i | idx) 的值
*/
// leetcode submit region begin(Prohibit modification and deletion)
func minimumIncompatibility(nums []int, k int) int {
	// 状压dp
	n := len(nums)
	if n == k { // 一个元素一个子集
		return 0
	}
	b, m, maxV := 1<<n, n/k, 17 // 1 <= k <= nums.length <= 16
	dp, memo := make([]int, b), make(map[int]int)
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
out:
	for i := 1; i < b; i++ { // 寻找所有 k 长度的子集
		dp[i] = math.MaxInt32 // 初始化，特别的 dp[0]=0，表示空数组的不兼容性的和为 0
		if bits.OnesCount(uint(i)) != m {
			continue
		}
		min, max, cache := maxV, 0, make([]int, maxV)
		for j := 0; j < n; j++ { // 子集映射为索引子集
			if i&(1<<j) > 0 {
				if cache[nums[j]] > 0 {
					continue out // 重复元素
				}
				cache[nums[j]] = 1
				min, max = minVal(min, nums[j]), maxVal(max, nums[j])
			}
		}
		memo[i] = max - min // 记录不兼容性
	}
	for i := 0; i < b; i++ {
		if dp[i] == math.MaxInt32 { // 不合法的子集 i
			continue
		}
		canUse, sub := make(map[int]int), 0
		for j := 0; j < n; j++ { // 寻找子集
			if i&(1<<j) == 0 {
				canUse[nums[j]] = j // 状压的是 索引
			}
		}
		if len(canUse) < m { // 长度不够
			continue
		}
		for _, v := range canUse {
			sub |= 1 << v // 无重复元素的子集
		}
		for idx := sub; idx > 0; idx = (idx - 1) & sub {
			if memo[idx] > 0 { // idx为合法的子集
				dp[i|idx] = minVal(dp[i|idx], dp[i]+memo[idx]) // 更新 dp[i|idx] 的最小值
			}
		}
	}
	if dp[b-1] < math.MaxInt32 { // 全集有解
		return dp[b-1]
	}
	return -1

	// 状压dp：对元素值状态压缩
	//counter, n := 0, len(nums)
	//if n == k { // 每个子集长度为 1
	//	return 0
	//}
	//l := 1 << n
	//arr, dp, m := make([]int, n+1), make([]int, l), n/k
	//for i := range nums {
	//	arr[nums[i]]++
	//}
	//for _, v := range arr {
	//	if v > counter { // 相同元素最多的个数
	//		counter = v
	//	}
	//}
	//if counter > k { // 不能组合
	//	return -1
	//}
	//for i := 1; i < l; i++ {
	//	dp[i] = math.MaxInt32
	//	if bits.OnesCount(uint(i)) < m {
	//		continue
	//	}
	//}
	//
	//cache := make(map[int]int, 16)
	//for i := 0; i < 16; i++ {
	//	cache[1<<i] = i
	//}
	//fmt.Println(cache)
	//tot, n, m := 0, len(nums), len(nums)/k
	//for i := 0; i < n; i++ {
	//	tot |= 1 << nums[i]
	//}
	//arr2 := make([][]int, 0)
	//for nxt := tot; nxt > 0; {
	//	if bits.OnesCount(uint(nxt)) == m {
	//		arr2 = append(arr2, []int{cache[GetHighestOneBit(nxt)], cache[nxt&-nxt]})
	//	}
	//	nxt = (nxt - 1) & tot
	//}
	//fmt.Println(arr)
	//return 0

	// lc
	//n := len(nums)
	//group := n / k // 每个子集的长度
	//inf := math.MaxInt32
	//dp := make([]int, 1<<n)
	//for i := range dp {
	//	dp[i] = inf
	//}
	//dp[0] = 0 // nums最大值为0
	//values := make(map[int]int)
	//
	//for mask := 1; mask < (1 << n); mask++ { // 对索引状态压缩
	//	if bits.OnesCount(uint(mask)) != group {
	//		continue
	//	} // 长度 != group
	//	minVal := 20 // 最大值为 16
	//	maxVal := 0
	//	cur := make(map[int]bool)
	//	for i := 0; i < n; i++ {
	//		if mask&(1<<i) != 0 {
	//			if cur[nums[i]] { // 元素重复
	//				break
	//			}
	//			cur[nums[i]] = true
	//			minVal = min(minVal, nums[i]) // 记录最值
	//			maxVal = max(maxVal, nums[i])
	//		}
	//	} // 可以组合成一个子集，key=组合为子集的索引
	//	if len(cur) == group {
	//		values[mask] = maxVal - minVal // value=差值
	//	}
	//}
	//fmt.Println(values)
	//for mask := 0; mask < (1 << n); mask++ {
	//	if dp[mask] == inf { // 没有这个组合
	//		continue
	//	}
	//	seen := make(map[int]int) // 可挑选元素的子集 key=nums[i] value=i
	//	for i := 0; i < n; i++ {
	//		if (mask & (1 << i)) == 0 { // 除去 mask，剩下可选的数
	//			seen[nums[i]] = i
	//		}
	//	}
	//	if len(seen) < group { // 不能组合为一个 group 长度的子集
	//		continue
	//	}
	//	sub := 0
	//	for _, v := range seen { // 子集包含哪些索引
	//		sub |= (1 << v)
	//	}
	//	nxt := sub
	//	for nxt > 0 {
	//		if val, ok := values[nxt]; ok { // 能组合成子集
	//			dp[mask|nxt] = min(dp[mask|nxt], dp[mask]+val)
	//		} // mask|nxt 的子集可组合，求 dp[mask|nxt] 最小值
	//		nxt = (nxt - 1) & sub // 遍历所有组合
	//	}
	//}
	//fmt.Println(dp)
	//if dp[(1<<n)-1] < inf { // 能否组合所有元素
	//	return dp[(1<<n)-1]
	//}
	//return -1
}

func GetHighestOneBit(c int) int {
	c = c | (c >> 1)
	c = c | (c >> 2)
	c = c | (c >> 4)
	c = c | (c >> 8)
	c = c | (c >> 16)
	return c>>1 + 1
}

//leetcode submit region end(Prohibit modification and deletion)
