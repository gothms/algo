package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "leet2code3"
	k := 20 // t
	index := decodeAtIndex(s, k)
	fmt.Printf("%s", index)
}

// leetcode submit region begin(Prohibit modification and deletion)
func decodeAtIndex(s string, k int) string {
	// lc：逆向工作法
	//n := len(s)
	//size := 0
	//for _, c := range s {
	//	if unicode.IsLetter(c) {
	//		size++
	//	} else {
	//		size *= int(c - '0')
	//	}
	//}
	//for i := n - 1; i >= 0; i-- {
	//	k %= size // 巧妙之处：
	//	// 1.k>size
	//	// 2.k<size，但“答案”不在这轮的“重复”中，k值会一直不变
	//	// 3.k<size，“答案”在这轮的“重复”中，k 会突然变为 0
	//	if k == 0 && unicode.IsLetter(rune(s[i])) {
	//		return string(s[i])
	//	}
	//	if unicode.IsNumber(rune(s[i])) {
	//		size /= int(s[i] - '0')
	//	} else {
	//		size--
	//	}
	//}
	//return ""

	// 个人
	n := len(s)
	ans := ""
	var dfs func(int, int) int
	dfs = func(i, c int) int {
		j, pre := i, c
		for j < n && unicode.IsLetter(rune(s[j])) { // 过滤区间 [i,j) 的连续字符，寻找数字
			j++
		}
		if j < n { // 数字
			c += (c + j - i) * int(s[j]-'1') // 重复 “c + j - i”
		}
		c += j - i // 补上区间 [i,j) 的连续字符

		var back int // 标记逆向找第几个字符，从 0 开始计数
		if c >= k {  // 1.已超过目标 k
			back = c - k
		} else {
			back = dfs(j+1, c) // 2.还未找到目标 k
		}
		if back == -1 { // 已完成，终止递归
			return -1
		}
		
		back %= pre + j - i // 排除掉重复的，其中 pre + j - i 即为 “c + j - i”
		if back < j-i {     // 目标 k 就在 [i,j) 区间中
			ans = string(s[j-back-1])
			return -1
		} else {
			back -= j - i // 目标 k 不在 [i,j) 区间中，继续逆向找
		}
		return back
	}
	dfs(0, 0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
