package main

import (
	"fmt"
)

func main() {
	s := "aab"
	s = "abbab" // [["a","b","b","a","b"],["a","b","bab"],["a","bb","a","b"],["abba","b"]]
	i := partition(s)
	fmt.Println(i)
}

/*
思路：dp+dfs
	1.
*/
// leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	// 记忆化搜索

}

//leetcode submit region end(Prohibit modification and deletion)

func partition_(s string) [][]string {
	// 记忆化搜索
	n := len(s)
	memo := make([][]bool, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]bool, n)
		memo[i][i] = true
	}
	for i := n - 2; i >= 0; i-- {
		memo[i][i+1] = s[i] == s[i+1]
		for j := i + 2; j < n; j++ {
			memo[i][j] = s[i] == s[j] && memo[i+1][j-1]
		}
	}
	ret, path := make([][]string, 0), make([]int, 1, n+1)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			pathStr := make([]string, len(path)-1)
			for k := 1; k < len(path); k++ {
				pathStr[k-1] = s[path[k-1]:path[k]]
			}
			ret = append(ret, pathStr)
			return
		}
		for j := i; j < n; j++ {
			if memo[i][j] {
				path = append(path, j+1)
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ret

	// dp+dfs：节约一半空间
	//n := len(s)
	//memo := make([][]bool, n)
	//for i := 0; i < n; i++ {
	//	memo[i] = make([]bool, i+1)
	//	memo[i][i] = true
	//}
	//for i := 1; i < n; i++ {
	//	memo[i][i-1] = s[i] == s[i-1]
	//	for j := i - 2; j >= 0; j-- {
	//		memo[i][j] = s[i] == s[j] && memo[i-1][j+1]
	//	} // [j,i] 是回文串
	//}
	//ret, path := make([][]string, 0), make([]int, 1, n+1)
	//var dfs func(int)
	//dfs = func(i int) {
	//	if i == n {
	//		pathStr := make([]string, len(path)-1)
	//		for k := 1; k < len(path); k++ {
	//			pathStr[k-1] = s[path[k-1]:path[k]]
	//		}
	//		ret = append(ret, pathStr)
	//		return
	//	}
	//	for j := i; j < n; j++ {
	//		if memo[j][i] { // [i,j] 是回文串
	//			path = append(path, j+1)
	//			dfs(j + 1)
	//			path = path[:len(path)-1]
	//		}
	//	}
	//}
	//dfs(0)
	//return ret
}
