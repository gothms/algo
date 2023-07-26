//有一个只含有 'Q', 'W', 'E', 'R' 四种字符，且长度为 n 的字符串。
//
// 假如在该字符串中，这四个字符都恰好出现 n/4 次，那么它就是一个「平衡字符串」。
//
//
//
// 给你一个这样的字符串 s，请通过「替换一个子串」的方式，使原字符串 s 变成一个「平衡字符串」。
//
// 你可以用和「待替换子串」长度相同的 任何 其他字符串来完成替换。
//
// 请返回待替换子串的最小可能长度。
//
// 如果原字符串自身就是一个平衡字符串，则返回 0。
//
//
//
// 示例 1：
//
//
//输入：s = "QWER"
//输出：0
//解释：s 已经是平衡的了。
//
// 示例 2：
//
//
//输入：s = "QQWE"
//输出：1
//解释：我们需要把一个 'Q' 替换成 'R'，这样得到的 "RQWE" (或 "QRWE") 是平衡的。
//
//
// 示例 3：
//
//
//输入：s = "QQQW"
//输出：2
//解释：我们可以把前面的 "QQ" 替换成 "ER"。
//
//
// 示例 4：
//
//
//输入：s = "QQQQ"
//输出：3
//解释：我们可以替换后 3 个 'Q'，使 s = "QWER"。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10^5
// s.length 是 4 的倍数
// s 中只含有 'Q', 'W', 'E', 'R' 四种字符
//
//
// Related Topics 字符串 滑动窗口 👍 243 👎 0

package main

import "fmt"

func main() {
	s := "WWEQERQWQWWRWWERQWEQ"
	s = "WQWRQQQW"
	i := balancedString(s)
	fmt.Println(i)
}

/*
思路：滑动窗体
	1.counter 记录 QWER 出现的总次数
	2.滑动窗体，当每个字符出现的次数 <= len(s)/4 时，则变更窗体中的子串，满足要求
*/
// leetcode submit region begin(Prohibit modification and deletion)
func balancedString(s string) int {
	// 滑动窗体
	counter, n := [26]int{}, len(s)
	minSub, m := n, n>>2
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	goon := func() bool { // 是否继续滑动快指针，以变更数量超出 m 的字符
		return counter['Q'-'A'] > m || counter['W'-'A'] > m ||
			counter['E'-'A'] > m || counter['R'-'A'] > m
	}
	for i := 0; i < n; i++ { // 统计字符数量
		counter[s[i]-'A']++
	}
	fast := 0             // 快指针
	for slow := range s { // 慢指针
		for ; fast < n && goon(); fast++ {
			counter[s[fast]-'A']-- // 滑动快指针，直到满足要求
		}
		if goon() { // 不能使用 fast==n 来判断，后缀子串也可能是最小子串
			break
		}
		minSub = minVal(minSub, fast-slow) // 满足要求时，记录子串长度
		counter[s[slow]-'A']++             // 滑动慢指针
	}
	return minSub
}
	// 题意是替换子串，这里是随意替换
	//counter, n := [4]int{}, len(s)
	//for i := 0; i < n; i++ {
	//	switch s[i] {
	//	case 'Q':
	//		counter[0]++
	//	case 'W':
	//		counter[1]++
	//	case 'E':
	//		counter[2]++
	//	case 'R':
	//		counter[3]++
	//	}
	//}
	//b, m := 0, n>>2
	//fmt.Println(counter, m)
	//for i := 0; i < 4; i++ {
	//	if counter[i] > m {
	//		b += counter[i] - m
	//	}
	//}
	//return b
}

//leetcode submit region end(Prohibit modification and deletion)
