//给你一个长度为 n 的整数数组 nums ，下标从 0 开始。
//
// 如果在下标 i 处 分割 数组，其中 0 <= i <= n - 2 ，使前 i + 1 个元素的乘积和剩余元素的乘积互质，则认为该分割 有效 。
//
//
// 例如，如果 nums = [2, 3, 3] ，那么在下标 i = 0 处的分割有效，因为 2 和 9 互质，而在下标 i = 1 处的分割无效，因为 6
// 和 3 不互质。在下标 i = 2 处的分割也无效，因为 i == n - 1 。
//
//
// 返回可以有效分割数组的最小下标 i ，如果不存在有效分割，则返回 -1 。
//
// 当且仅当 gcd(val1, val2) == 1 成立时，val1 和 val2 这两个值才是互质的，其中 gcd(val1, val2) 表示
//val1 和 val2 的最大公约数。
//
//
//
// 示例 1：
//
//
//
//
//输入：nums = [4,7,8,15,3,5]
//输出：2
//解释：上表展示了每个下标 i 处的前 i + 1 个元素的乘积、剩余元素的乘积和它们的最大公约数的值。
//唯一一个有效分割位于下标 2 。
//
// 示例 2：
//
//
//
//
//输入：nums = [4,7,15,8,3,5]
//输出：-1
//解释：上表展示了每个下标 i 处的前 i + 1 个元素的乘积、剩余元素的乘积和它们的最大公约数的值。
//不存在有效分割。
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 10⁴
// 1 <= nums[i] <= 10⁶
//
//
// Related Topics 数组 哈希表 数学 数论 👍 34 👎 0

package main

import "fmt"

func main() {
	a, b := 42*13*2*6, 42*5*7*11*23
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Println(a, b)
	nums := []int{4, 7, 15, 8, 3, 5}
	split := findValidSplit(nums)
	fmt.Println(split)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findValidSplit(nums []int) int {
	n := len(nums)
	memo, far := make(map[int]int), make([]int, n)
	record := func(i, p int) { // 关键逻辑
		if f, ok := memo[p]; ok {
			far[f] = i // 质因子 p 第一次出现为 f，最后一次出现为 i
			// f 记录的最远 i 会更新并覆盖
			// 区间 [f, i) 不可能是答案
		} else {
			memo[p] = i // 质因子 p 第一次出现的位置为 i
		}
	}
	for i, val := range nums { // 求出质因子
		v := val
		for _, p := range prime2584 {
			if p*p > v {
				break
			}
			if v%p == 0 {
				record(i, p)
				for v /= p; v%p == 0; v /= p {
				}
			}
		}
		if v > 1 {
			record(i, v)
		}
	}
	maxFar := 0 // 重复质因子的最远位置
	for i, f := range far {
		if i > maxFar { // maxFar 为分割点
			return maxFar
		}
		maxFar = max(maxFar, f) // 更新 maxFar
	}
	return -1
}

var prime2584 []int

func init() { // 求出质数
	const N = 1e6
	prime := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		if !prime[i] {
			for j := i << 1; j <= N; j += i {
				prime[j] = true
			}
			prime2584 = append(prime2584, i)
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
