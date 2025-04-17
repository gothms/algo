package main

import "fmt"

func main() {
	s := "babab"
	palindrome := smallestPalindrome(s)
	fmt.Println(palindrome)
}

// leetcode submit region begin(Prohibit modification and deletion)
func smallestPalindrome(s string) string {
	// lc：先排序 s 左半部分，再反转相加

	// 优化：只计数排序左半部分，且使用 bytes.Repeat 函数

	// 个人：计数排序
	n := len(s)
	//var m int
	memo := [26]int{}
	for _, c := range s {
		memo[c-'a']++
	}
	buf := make([]byte, n)
	j := 0
	for i, c := range memo {
		var v = byte(i) + 'a'
		for ; c > 1; c -= 2 {
			buf[j], buf[n-1-j] = v, v
			j++
		}
		//if c&1 == 1 {
		//	m = i + 'a'
		//}
	}
	if n&1 == 1 {
		buf[n>>1] = s[n>>1]
	}
	return string(buf)
}

//leetcode submit region end(Prohibit modification and deletion)
