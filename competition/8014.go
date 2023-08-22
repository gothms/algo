package main

import (
	"fmt"
)

func main() {
	str1 := "abc"
	str2 := "ad"
	str1 = "ab"
	str2 = "d"
	subsequence := canMakeSubsequence(str1, str2)
	fmt.Println(subsequence)
}

/*
思路：双指针
	1.贪心
		在选择子序列匹配中，我们可以总是优先选择先出现的匹配元素
	2.str1 str2 匹配原则
		字符相等：str1[i] == str2[j]
			i++
			j++
		字符不等：
			相差一个字符，即：str1[i]+1 == str2[j] || str1[i] == 'z' && str2[j] == 'a'
				i++
				j++
			否则
				i++
*/

func canMakeSubsequence(str1 string, str2 string) bool {
	// 双指针
	n, m := len(str1), len(str2)
	for i, j := 0, 0; i < n; i++ {
		// 贪心：字符匹配时，优先选择
		if str1[i] == str2[j] || str1[i]+1 == str2[j] || str1[i] == 'z' && str2[j] == 'a' {
			j++
			if j == m { // 终止条件：子序列匹配
				return true
			}
		}
	}
	return false

	// 双指针
	//dp, j, n, m := false, 0, len(str1), len(str2)
	//for i := 0; i < n && j < m; i++ {
	//	if str1[i] == str2[j] || str1[i]+1 == str2[j] || str1[i] == 'z' && str2[j] == 'a' {
	//		j++
	//		dp = true
	//	} else {
	//		dp = false
	//	}
	//}
	//return j == m && dp

	// dfs：超时
	//n, m := len(str1), len(str2)
	//if n < m {
	//	return false
	//}
	//var dfs func(int, int) bool
	//dfs = func(i, j int) bool {
	//	if j == m {
	//		return true
	//	} else if i == n {
	//		return false
	//	}
	//	if str1[i] == str2[j] {
	//		return dfs(i+1, j+1)
	//	} else if str1[i]+1 == str2[j] || str1[i] == 'z' && str2[j] == 'a' {
	//		return dfs(i+1, j) || dfs(i+1, j+1)
	//	}
	//	return dfs(i+1, j)
	//}
	//return dfs(0, 0)
}
