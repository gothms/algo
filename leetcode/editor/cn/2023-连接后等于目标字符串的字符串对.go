package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numOfPairs(nums []string, target string) int {
	// lc
	
	// 个人
	memo := make(map[string]int)
	ans, n := 0, len(target)
	for _, v := range nums {
		m := len(v)
		if m >= n {
			continue
		}
		if v == target[:m] { // 前缀
			ans += memo[target[m:]]
		}
		if v == target[n-m:] {
			ans += memo[target[:n-m]]
		}
		memo[v]++
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
