package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	path := make([]byte, 0, n<<1)
	ans := make([]string, 0)
	var dfs func(int, int)
	dfs = func(l, r int) {
		if r == 0 {
			ans = append(ans, string(path))
			return
		}
		if l > 0 {
			path = append(path, '(')
			dfs(l-1, r)
			path = path[:len(path)-1]
		}
		if r > l {
			path = append(path, ')')
			dfs(l, r-1)
			path = path[:len(path)-1]
		}
	}
	dfs(n, n)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
