package main

import "fmt"

func main() {
	strs := []string{"aa", "aaa", "aaa"}                        // -1
	strs = []string{"aabbcc", "aabbcc", "bc", "bcc", "aabbccc"} // 7
	slength := findLUSlength(strs)
	fmt.Println(slength)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findLUSlength(strs []string) int {
	isSubSeq := func(s string, str string) bool { // s 是否为 str 的子序列
		if len(str) < len(s) {
			return false
		}
		j := 0
		for i := range str {
			if str[i] == s[j] {
				j++
				if j == len(s) {
					return true
				}
			}
		}
		return false
	}
	const k, b = 6, 1<<6 - 1
	memo := make(map[string]int)
	for i, s := range strs { // 去重，且记录次数
		memo[s] = i<<k | memo[s]&b + 1 // 高位是索引，低位是次数
	}
	ans := -1
out:
	for s, v := range memo {
		if v&b == 1 { // 重复的字符串不合法
			i := v >> k
			for str, val := range memo {
				if val>>k != i && isSubSeq(s, str) {
					continue out
				}
			}
			ans = max(ans, len(s))
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func findLUSlength_(strs []string) int {
	isSubSeq := func(s1, s2 string) bool { //  s1 是否为 s2 的子序列
		n := len(s1)
		if n > len(s2) {
			return false
		}
		i := 0
		for j := range s2 {
			if s2[j] == s1[i] {
				i++
				if i == n {
					return true
				}
			}
		}
		return false
	}
	ans := -1
out:
	for i, s1 := range strs {
		for j, s2 := range strs {
			if i != j && isSubSeq(s1, s2) { // s1 满足题意
				continue out
			}
		}
		ans = max(ans, len(s1))
	}
	return ans
}
