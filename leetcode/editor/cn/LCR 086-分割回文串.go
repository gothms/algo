package main

import (
	"fmt"
	"slices"
)

func main() {
	s := "google"
	s = "baaaaac"
	i := partition(s)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range s {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		for j := i - 1; j >= 0; j-- {
			dp[j][i] = s[j] == s[i] && (j+1 == i || dp[j+1][i-1])
		}
	}
	ans, path := make([][]string, 0), make([]string, 0, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func partition_(s string) [][]string {
	n := len(s)
	palindrome := make([][]bool, n) // 统计
	for i := range palindrome {
		palindrome[i] = make([]bool, n)
		palindrome[i][i] = true
	}
	for i := n - 2; i >= 0; i-- { // 统计
		if s[i] == s[i+1] {
			palindrome[i][i+1] = true
		}
		for j := i + 2; j < n; j++ {
			palindrome[i][j] = s[i] == s[j] && palindrome[i+1][j-1]
		}
	}
	ans, path := make([][]string, 0), make([]string, 0, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		for j := i; j < n; j++ {
			if palindrome[i][j] { // 查表
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1] // 回溯
			}
		}
	}
	dfs(0)
	return ans
}
