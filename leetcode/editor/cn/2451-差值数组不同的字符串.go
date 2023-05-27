//给你一个字符串数组 words ，每一个字符串长度都相同，令所有字符串的长度都为 n 。
//
// 每个字符串 words[i] 可以被转化为一个长度为 n - 1 的 差值整数数组 difference[i] ，其中对于 0 <= j <= n - 2
// 有 difference[i][j] = words[i][j+1] - words[i][j] 。注意两个字母的差值定义为它们在字母表中 位置 之差，也就是
//说 'a' 的位置是 0 ，'b' 的位置是 1 ，'z' 的位置是 25 。
//
//
// 比方说，字符串 "acb" 的差值整数数组是 [2 - 0, 1 - 2] = [2, -1] 。
//
//
// words 中所有字符串 除了一个字符串以外 ，其他字符串的差值整数数组都相同。你需要找到那个不同的字符串。
//
// 请你返回 words中 差值整数数组 不同的字符串。
//
//
//
// 示例 1：
//
//
//输入：words = ["adc","wzy","abc"]
//输出："abc"
//解释：
//- "adc" 的差值整数数组是 [3 - 0, 2 - 3] = [3, -1] 。
//- "wzy" 的差值整数数组是 [25 - 22, 24 - 25]= [3, -1] 。
//- "abc" 的差值整数数组是 [1 - 0, 2 - 1] = [1, 1] 。
//不同的数组是 [1, 1]，所以返回对应的字符串，"abc"。
//
//
// 示例 2：
//
//
//输入：words = ["aaa","bob","ccc","ddd"]
//输出："bob"
//解释：除了 "bob" 的差值整数数组是 [13, -13] 以外，其他字符串的差值整数数组都是 [0, 0] 。
//
//
//
//
// 提示：
//
//
// 3 <= words.length <= 100
// n == words[i].length
// 2 <= n <= 20
// words[i] 只含有小写英文字母。
//
//
// Related Topics 数组 哈希表 字符串 👍 15 👎 0

package main

import "fmt"

func main() {
	words := []string{"adc", "wzy", "abc"}
	words = []string{"aaa", "bob", "ccc", "ddd"}
	words = []string{"ddd", "poo", "baa", "onn"}
	s := oddString(words)
	fmt.Println(s)
}

/*
思路：
	1.差值概念优化：对于任意差值 difference[i][j] = words[i][j+1] - words[i][j]
		与其用 words[i][j+1] - words[i][j]，作为判断不同字符串的差值是否相等
		不如用 words[i][j] - words[i][0]，也就是字符串的第一个 字符 作为标杆
	2.字符串之间差值比较：
		与其两个字符串之间进行比较，或者统计字符串
		不如用 words[0] 作为标杆：与 words[0] 相同记为 w0，不相同记为 w，2者出现次数：
			w0>=2，w==1：返回 words[i-1]/words[i]
			w0==1，w>=2：返回 words[0]
*/
//leetcode submit region begin(Prohibit modification and deletion)
func oddString(words []string) string {
	n, m := len(words), len(words[0])
	for i := 1; i < m; i++ {
		v0, w0, w := words[0][i]-words[0][0], false, false
		for j := 1; j < n; j++ {
			switch words[j][i] - words[j][0] {
			case v0:
				if w {
					return words[j-1]
				}
				w0 = true
			default:
				if w0 {
					return words[j]
				} else if w {
					return words[0]
				}
				w = true
			}
		}
	}
	return ""
}
	//n, m := len(words), len(words[0])
	//cd := make([]byte, n)
	//for i := 0; i < n; i++ {
	//	cd[i] = words[i][0] - 'a'
	//}
	//for i := 1; i < m; i++ {
	//	v0, c0, c := words[0][i]-cd[0], 1, 0
	//	for j := 1; j < n; j++ {
	//		if words[j][i]-cd[j] == v0 {
	//			c0++
	//			if c == 1 {
	//				return words[j-1]
	//			}
	//		} else {
	//			c++
	//			if c >= 2 {
	//				return words[0]
	//			} else if c0 >= 2 {
	//				return words[j]
	//			}
	//		}
	//	}
	//}
	//return "?"
}

//leetcode submit region end(Prohibit modification and deletion)
