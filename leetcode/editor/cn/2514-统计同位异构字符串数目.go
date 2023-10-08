//给你一个字符串 s ，它包含一个或者多个单词。单词之间用单个空格 ' ' 隔开。
//
// 如果字符串 t 中第 i 个单词是 s 中第 i 个单词的一个 排列 ，那么我们称字符串 t 是字符串 s 的同位异构字符串。
//
//
// 比方说，"acb dfe" 是 "abc def" 的同位异构字符串，但是 "def cab" 和 "adc bef" 不是。
//
//
// 请你返回 s 的同位异构字符串的数目，由于答案可能很大，请你将它对 10⁹ + 7 取余 后返回。
//
//
//
// 示例 1：
//
// 输入：s = "too hot"
//输出：18
//解释：输入字符串的一些同位异构字符串为 "too hot" ，"oot hot" ，"oto toh" ，"too toh" 以及 "too oht" 。
//
//
// 示例 2：
//
// 输入：s = "aa"
//输出：1
//解释：输入字符串只有一个同位异构字符串。
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁵
// s 只包含小写英文字母和空格 ' ' 。
// 相邻单词之间由单个空格隔开。
//
//
// Related Topics 哈希表 数学 字符串 组合数学 计数 👍 14 👎 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "too hot"
	anagrams := countAnagrams(s)
	fmt.Println(anagrams)
}

/*
组合计数
	单词长度为 n，全排列个数为 n!
	有重复元素 x 个字符 a，则需除以 a 的全排列个数 x!

费马小定理：a, p 互素, 且 p 为素数 a^(p-1) mod p = 1 可以推导出：
	(ans / mul) mod p = (ans / mul * 1) mod p =
	(ans / mul * mul ^ (p - 1)) mod p = (ans * mul ^ (p - 2)) mod p
*/
// leetcode submit region begin(Prohibit modification and deletion)
func countAnagrams(s string) int {
	// lc
	const mod = 1_000_000_007
	pow := func(x, n int) int {
		ret := 1
		for ; n != 0; n >>= 1 {
			if n&1 == 1 {
				ret = ret * x % mod
			}
			x = x * x % mod
		}
		return ret
	}
	ret, mul := 1, 1 // 总数为 ret / mul
	for _, ss := range strings.Split(s, " ") {
		cnt := [26]int{}
		for i, c := range ss {
			idx := c - 'a'
			cnt[idx]++
			mul = mul * cnt[idx] % mod // 分母：重复元素的全排列
			ret = ret * (i + 1) % mod  // 分子：所有元素的全排列
		}
	}
	return ret * pow(mul, mod-2) % mod // 费马小定理：求 ret / mul % mod
}

//leetcode submit region end(Prohibit modification and deletion)
