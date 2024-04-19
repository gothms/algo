//给你一个正整数数组 arr，请你找出一个长度为 m 且在数组中至少重复 k 次的模式。
//
// 模式 是由一个或多个值组成的子数组（连续的子序列），连续 重复多次但 不重叠 。 模式由其长度和重复次数定义。
//
// 如果数组中存在至少重复 k 次且长度为 m 的模式，则返回 true ，否则返回 false 。
//
//
//
// 示例 1：
//
// 输入：arr = [1,2,4,4,4,4], m = 1, k = 3
//输出：true
//解释：模式 (4) 的长度为 1 ，且连续重复 4 次。注意，模式可以重复 k 次或更多次，但不能少于 k 次。
//
//
// 示例 2：
//
// 输入：arr = [1,2,1,2,1,1,1,3], m = 2, k = 2
//输出：true
//解释：模式 (1,2) 长度为 2 ，且连续重复 2 次。另一个符合题意的模式是 (2,1) ，同样重复 2 次。
//
//
// 示例 3：
//
// 输入：arr = [1,2,1,2,1,3], m = 2, k = 3
//输出：false
//解释：模式 (1,2) 长度为 2 ，但是只连续重复 2 次。不存在长度为 2 且至少重复 3 次的模式。
//
//
// 示例 4：
//
// 输入：arr = [1,2,3,1,2], m = 2, k = 2
//输出：false
//解释：模式 (1,2) 出现 2 次但并不连续，所以不能算作连续重复 2 次。
//
//
// 示例 5：
//
// 输入：arr = [2,2,2,2], m = 2, k = 3
//输出：false
//解释：长度为 2 的模式只有 (2,2) ，但是只连续重复 2 次。注意，不能计算重叠的重复次数。
//
//
//
//
// 提示：
//
//
// 2 <= arr.length <= 100
// 1 <= arr[i] <= 100
// 1 <= m <= 100
// 2 <= k <= 100
//
//
// Related Topics 数组 枚举 👍 66 👎 0

package main

import "fmt"

func main() {
	arr := []int{1, 2, 4, 4, 4, 4}
	m, k := 1, 3
	//arr = []int{1, 2, 1, 2, 1, 1, 1, 3}
	//arr = []int{1, 2, 3, 1, 2}
	//m, k = 2, 2
	pattern := containsPattern(arr, m, k)
	fmt.Println(pattern)
}

// leetcode submit region begin(Prohibit modification and deletion)
func containsPattern(arr []int, m int, k int) bool {
	// Rabin-Karp
	n := len(arr)
	if m*k > n { // fast fail
		return false
	}
	const prime = 16777619
	hashes, times := make([]uint32, n), make([]int, n)
	pow := func(i int) uint32 {
		var ret, base uint32 = 1, prime
		for ; i > 0; i >>= 1 {
			if i&1 == 1 {
				ret *= base
			}
			base *= base
		}
		return ret
	}
	p := pow(m)
	var hash uint32
	for i := 0; i < m; i++ {
		hash = hash*prime + uint32(arr[i]) // Rabin-Karp：1
	}
	hashes[m-1], times[m-1] = hash, 1
	for i := m; i < n; i++ {
		hash = hash*prime + uint32(arr[i]) - uint32(arr[i-m])*p // Rabin-Karp：2
		if hash == hashes[i-m] {                                // 重点：连续 重复多次但 不重叠
			times[i] = times[i-m] + 1 // 连续
			if times[i] == k {
				return true
			}
		} else { // 无匹配 / 不连续
			times[i] = 1
		}
		hashes[i] = hash // 记录 hash
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
