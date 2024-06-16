package main

import "fmt"

func main() {
	strs := []string{"aa", "aaa", "aaa"}
	slength := findLUSlength(strs)
	fmt.Println(slength)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findLUSlength(strs []string) int {
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

//leetcode submit region end(Prohibit modification and deletion)
