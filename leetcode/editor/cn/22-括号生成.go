package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	// 栈
	ans := make([]string, 0)
	st := make([]byte, 0)
	var dfs func(int, int)
	dfs = func(l, r int) {
		if l|r == 0 {
			ans = append(ans, string(st))
		}
		if l > 0 {
			st = append(st, '(')
			dfs(l-1, r)
			st = st[:len(st)-1]
		}
		if r > l {
			st = append(st, ')')
			dfs(l, r-1)
			st = st[:len(st)-1]
		}
	}
	dfs(n, n)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
