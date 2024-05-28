package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "a1b2"
	s = "abc"
	permutation := letterCasePermutation(s)
	fmt.Println(permutation)
}

// leetcode submit region begin(Prohibit modification and deletion)
func letterCasePermutation(s string) []string {
	// 二进制位图

	// bfs

	// 回溯：每个字符的大小写形式刚好差了 32

	// 回溯：写法二
	ans, buf := make([]string, 0), []rune(s)
	var dfs func(int)
	dfs = func(i int) {
		if i < 0 {
			return
		}
		dfs(i - 1)
		if unicode.IsDigit(buf[i]) {
			return
		} else if unicode.IsLower(buf[i]) {
			buf[i] = unicode.ToUpper(buf[i])
		} else {
			buf[i] = unicode.ToLower(buf[i])
		}
		ans = append(ans, string(buf))
		dfs(i - 1)
	}
	dfs(len(s) - 1)
	return append(ans, s)

	// 回溯
	//ans, buf := make([]string, 0), []rune(s)
	//var dfs func(int)
	//dfs = func(i int) {
	//	if i < 0 {
	//		ans = append(ans, string(buf))
	//		return
	//	}
	//	dfs(i - 1)
	//	if unicode.IsDigit(buf[i]) {
	//		return
	//	} else if unicode.IsLower(buf[i]) {
	//		buf[i] = unicode.ToUpper(buf[i])
	//		dfs(i - 1)
	//        buf[i] = unicode.ToLower(buf[i])
	//    } else {
	//        buf[i] = unicode.ToLower(buf[i])
	//        dfs(i - 1)
	//        buf[i] = unicode.ToUpper(buf[i])
	//	}
	//}
	//dfs(len(s) - 1)
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
