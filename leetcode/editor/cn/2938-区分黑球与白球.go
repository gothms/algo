package main

import "fmt"

func main() {
	s := "abc"
	fmt.Println(len(s))
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumSteps(s string) int64 {
	// 双指针
	ans := 0
	for i, j := 0, len(s)-1; i < j; {
		if s[i] == '0' {
			i++
		} else if s[j] == '1' {
			j--
		} else if s[i] == '1' && s[j] == '0' {
			ans += j - i
			i++
			j--
		}
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
