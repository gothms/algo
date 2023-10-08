//给你一个字符串 text ，请你返回满足下述条件的 不同 非空子字符串的数目：
//
//
// 可以写成某个字符串与其自身相连接的形式（即，可以写为 a + a，其中 a 是某个字符串）。
//
//
// 例如，abcabc 就是 abc 和它自身连接形成的。
//
//
//
// 示例 1：
//
// 输入：text = "abcabcabc"
//输出：3
//解释：3 个子字符串分别为 "abcabc"，"bcabca" 和 "cabcab" 。
//
//
// 示例 2：
//
// 输入：text = "leetcodeleetcode"
//输出：2
//解释：2 个子字符串为 "ee" 和 "leetcodeleetcode" 。
//
//
//
//
// 提示：
//
//
// 1 <= text.length <= 2000
// text 只包含小写英文字母。
//
//
// Related Topics 字典树 字符串 哈希函数 滚动哈希 👍 53 👎 0

package main

import "fmt"

func main() {
	text := "abcabcabc"
	text = "leetcodeleetcode"
	substrings := distinctEchoSubstrings(text)
	fmt.Println(substrings)
}

// leetcode submit region begin(Prohibit modification and deletion)
func distinctEchoSubstrings(text string) int {
	// Rabin-Karp
	const PrimeRK = 16777619
	cnt, n := 0, len(text)
	hashes, pows, matched := make([]uint32, n+1), make([]uint32, n+1), make(map[uint32]bool)
	pows[0] = 1
	for i := 0; i < n; i++ {
		hashes[i+1] = hashes[i]*PrimeRK + uint32(text[i]) // hash 集合
		pows[i+1] = pows[i] * PrimeRK                     // pow 集合
	}
	get := func(l, r int) uint32 { // [l,r)
		return hashes[r] - hashes[l]*pows[r-l]
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if j<<1-i > n { // 防止越界
				continue
			}
			if hash := get(i, j); !matched[hash] && hash == get(j, j<<1-i) { // 连续的子串匹配
				cnt++
				matched[hash] = true // 记录已匹配的子串
			}
		}
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
