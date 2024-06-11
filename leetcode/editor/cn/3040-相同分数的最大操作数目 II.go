package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 2, 2, 1, 3, 3} // 4
	operations := maxOperations(nums)
	fmt.Println(operations)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxOperations(nums []int) int {
	// dp

}

// leetcode submit region end(Prohibit modification and deletion)
func maxOperations_(nums []int) int {
	// dp：记忆化搜索更优，https://leetcode.cn/problems/maximum-number-of-operations-with-the-same-score-ii/solutions/2643756/qu-jian-dp-de-tao-lu-pythonjavacgo-by-en-nynz/

	// 记忆化搜索：否则超时
	n := len(nums)
	if n == 2 { // fast path
		return 1
	}
	memo := make(map[int]map[int]int)
	memo[nums[0]+nums[1]] = make(map[int]int)
	if k := nums[n-1] + nums[n-2]; memo[k] == nil { // 去重
		memo[k] = make(map[int]int)
	}
	if k := nums[0] + nums[n-1]; memo[k] == nil {
		memo[k] = make(map[int]int)
	}

	var cur map[int]int
	var done bool // 优化：最多有 n/2 次操作
	ans := 0
	var dfs func(int, int, int) int
	dfs = func(l, r, s int) int { // 记忆化搜索
		if done {
			return 0
		}
		if l >= r { // 已最大化答案
			done = true
			return 0
		}
		key := l*n + r
		if v, ok := cur[key]; ok {
			return v
		}
		ret := 0
		if nums[l]+nums[l+1] == s {
			ret = dfs(l+2, r, s) + 1
		}
		if nums[l]+nums[r] == s {
			ret = max(ret, dfs(l+1, r-1, s)+1)
		}
		if nums[r]+nums[r-1] == s {
			ret = max(ret, dfs(l, r-2, s)+1)
		}
		cur[key] = ret
		return ret
	}
	for k, v := range memo {
		cur, done = v, false
		ans = max(ans, dfs(0, n-1, k))
	}
	return ans
}
