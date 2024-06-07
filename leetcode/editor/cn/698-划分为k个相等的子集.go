package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	k := 4 // true
	nums = []int{2, 2, 2, 2, 3, 4, 5}
	k = 4 // false
	nums = []int{3522, 181, 521, 515, 304, 123, 2512, 312, 922, 407, 146, 1932, 4037, 2646, 3871, 269}
	k = 5 // true
	nums = []int{1, 2, 2, 2, 2}
	k = 3 // false
	kSubsets := canPartitionKSubsets(nums, k)
	fmt.Println(kSubsets)
}

// leetcode submit region begin(Prohibit modification and deletion)
func canPartitionKSubsets(nums []int, k int) bool {
	// dp：lc
	s, n := 0, len(nums)
	for _, v := range nums {
		s += v
	}
	if s%k != 0 {
		return false
	}
	t := s / k
	sort.Ints(nums)
	if nums[n-1] > t {
		return false
	}

	m := 1 << n
	dp, ss := make([]bool, m), make([]int, m) // 记录当前的和
	dp[m-1] = true
	for mask := m - 1; mask > 0; mask-- {
		if !dp[mask] { // dp[mask]==true，则递推
			continue
		}
		for i, v := range nums {
			if v+ss[mask] > t { // 已排序
				break
			}
			if mask>>i&1 > 0 {
				j := mask &^ (1 << i)
				dp[j] = true
				ss[j] = (ss[mask] + v) % t // 要用于后面的判断 if v+ss[mask] > t
			}
		}
	}
	return dp[0]

	// 记忆化搜索
	//s, maxV, n := 0, 0, len(nums) // 预处理
	//for _, v := range nums {
	//	maxV = max(maxV, v)
	//	s += v
	//}
	//if s%k != 0 {
	//	return false
	//}
	//target := s / k
	//if maxV > target {
	//	return false
	//}
	//m := 1 << n // 记忆化搜索
	//memo := make([]uint8, m)
	//memo[0] = 1
	//var dfs func(int, int) bool
	//dfs = func(t, mask int) bool {
	//	if t == 0 {
	//		t = target
	//	}
	//	v := &memo[mask]
	//	if *v != 0 {
	//		return *v == 1
	//	}
	//	for i := 0; i < n; i++ {
	//		if nums[i] <= t && 1<<i&mask > 0 && dfs(t-nums[i], mask&^(1<<i)) {
	//			*v = 1 // yes
	//			return true
	//		}
	//	}
	//	*v = 2 // no
	//	return false
	//}
	//return dfs(target, m-1)
}

//leetcode submit region end(Prohibit modification and deletion)
