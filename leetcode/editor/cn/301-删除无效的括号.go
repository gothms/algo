package main

import "fmt"

func main() {
	s := ")(f"
	parentheses := removeInvalidParentheses(s)
	fmt.Println(parentheses)
}

// leetcode submit region begin(Prohibit modification and deletion)
func removeInvalidParentheses(s string) []string {

}

//leetcode submit region end(Prohibit modification and deletion)

func removeInvalidParentheses_(s string) []string {
	n := len(s)
	memo, path := make(map[string]struct{}), make([]byte, 0, n)
	l, r := 0, 0 // 需要删除 ( 和 ) 各 l r 个
	for _, c := range s {
		switch c {
		case '(':
			l++
		case ')':
			if l == 0 {
				r++
			} else {
				l--
			}
		}
	}

	var dfs func(int, int, int, int)
	dfs = func(l, r, i, d int) {
		if i == n {
			if l|r|d == 0 { // 满足题意
				memo[string(path)] = struct{}{}
			}
			return
		}
		c := s[i]
		i++
		switch c {
		case '(':
			path = append(path, c)
			dfs(l, r, i, d+1)
			path = path[:len(path)-1]
			l-- // 不选
		case ')':
			if d > 0 {
				path = append(path, c)
				dfs(l, r, i, d-1)
				path = path[:len(path)-1]
			}
			r--
		default:
			path = append(path, c)
			dfs(l, r, i, d)
			path = path[:len(path)-1] // 回溯
			return
		}
		dfs(l, r, i, d) // 不选
	}
	dfs(l, r, 0, 0)

	ans := make([]string, 0, len(memo))
	for str := range memo {
		ans = append(ans, str)
	}
	return ans
}
