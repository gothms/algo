//给你一个二进制字符串 s 和一个正整数 k 。
//
// 如果 s 的某个子字符串中 1 的个数恰好等于 k ，则称这个子字符串是一个 美丽子字符串 。
//
// 令 len 等于 最短 美丽子字符串的长度。
//
// 返回长度等于 len 且字典序 最小 的美丽子字符串。如果 s 中不含美丽子字符串，则返回一个 空 字符串。
//
// 对于相同长度的两个字符串 a 和 b ，如果在 a 和 b 出现不同的第一个位置上，a 中该位置上的字符严格大于 b 中的对应字符，则认为字符串 a 字典
//序 大于 字符串 b 。
//
//
// 例如，"abcd" 的字典序大于 "abcc" ，因为两个字符串出现不同的第一个位置对应第四个字符，而 d 大于 c 。
//
//
//
//
// 示例 1：
//
//
//输入：s = "100011001", k = 3
//输出："11001"
//解释：示例中共有 7 个美丽子字符串：
//1. 子字符串 "100011001" 。
//2. 子字符串 "100011001" 。
//3. 子字符串 "100011001" 。
//4. 子字符串 "100011001" 。
//5. 子字符串 "100011001" 。
//6. 子字符串 "100011001" 。
//7. 子字符串 "100011001" 。
//最短美丽子字符串的长度是 5 。
//长度为 5 且字典序最小的美丽子字符串是子字符串 "11001" 。
//
//
// 示例 2：
//
//
//输入：s = "1011", k = 2
//输出："11"
//解释：示例中共有 3 个美丽子字符串：
//1. 子字符串 "1011" 。
//2. 子字符串 "1011" 。
//3. 子字符串 "1011" 。
//最短美丽子字符串的长度是 2 。
//长度为 2 且字典序最小的美丽子字符串是子字符串 "11" 。
//
//
// 示例 3：
//
//
//输入：s = "000", k = 1
//输出：""
//解释：示例中不存在美丽子字符串。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 100
// 1 <= k <= s.length
//
//
// Related Topics 字符串 滑动窗口 👍 7 👎 0

package main

import "fmt"

func main() {
	s := "001011111010001001101101110" // 10111110100010011011
	k := 12
	substring := shortestBeautifulSubstring(s, k)
	fmt.Println(substring)
}

// leetcode submit region begin(Prohibit modification and deletion)
func shortestBeautifulSubstring(s string, k int) string {
	n := len(s)
	cache := make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			cache = append(cache, i)
		}
	}
	m := len(cache)
	if k == 1 && m > 0 {
		return "1"
	}
	l, f := n+1, -1
	for i := 0; i <= m-k; i++ { // 滑动窗体
		length := cache[i+k-1] - cache[i]
		if length < l {
			l, f = length, i
		} else if length == l {
			for j := 1; j < k-1; j++ {
				if d := cache[i+j] - cache[i] - cache[f+j] + cache[f]; d == 0 {
					continue
				} else if d > 0 {
					f = i
				}
				break
			}
		}
	}
	if f >= 0 {
		return s[cache[f] : cache[f+k-1]+1]
	}
	return ""
}

//leetcode submit region end(Prohibit modification and deletion)
