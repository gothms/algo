//如果字符串中的所有字符都相同，那么这个字符串是单字符重复的字符串。
//
// 给你一个字符串 text，你只能交换其中两个字符一次或者什么都不做，然后得到一些单字符重复的子串。返回其中最长的子串的长度。
//
//
//
// 示例 1：
//
// 输入：text = "ababa"
//输出：3
//
//
// 示例 2：
//
// 输入：text = "aaabaaa"
//输出：6
//
//
// 示例 3：
//
// 输入：text = "aaabbaaa"
//输出：4
//
//
// 示例 4：
//
// 输入：text = "aaaaa"
//输出：5
//
//
// 示例 5：
//
// 输入：text = "abcdef"
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= text.length <= 20000
// text 仅由小写英文字母组成。
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 109 👎 0

package main

import "fmt"

func main() {
	text := "a"
	opt1 := maxRepOpt1(text)
	fmt.Println(opt1)
}

/*
思路：双指针
	1.使用 i j 双指针，记录符合条件的单字符重复子串 text[i:j+1]，i 快 j 慢
		1.1.j 到 i 之间的字符，全部相同
		1.2.j 到 i 之间的字符，有一个不同
		1.3.使用 cache 记录各个字符的总数，以便查询是否有额外的字符用作交换
	2.单字符重复子串 判定规则
		2.1.当1.1.时，即 cache[text[j]-'a'] == i-j
			表示没有多余的字符 text[j] 去交换 text[i]
		2.2.当1.1.时，k 记录了一个不同的字符
			a)若 cache[text[j]-'a'] == i-j，表示 [j,i-1] 以外
				没有多余的 text[j] 去交换 text[k]，cnt=cache[text[j]-'a']
			b)若 cache[text[j]-'a'] > i-j，表示 [j,i-1] 以外
				有多余的 text[j] 去交换 text[k]，cnt=i-j
		2.3.下一次需要从 k 字符开始查询，以免漏掉 text[k] 字符符合 单字符重复子串 的情况
			此时必须满足 i == k+1 / i == k+2
			也可以据此判断，下一次查询要不要从 text[k] 开始，不过实现更加繁杂
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxRepOpt1(text string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	cnt, cache, n := 0, [26]int{}, len(text)
	for i := 0; i < n; i++ {
		cache[text[i]-'a']++
	}
	for i, j, k, idx := 0, 0, 0, 0; i < n; i++ {
		for i < n && text[i] == text[j] {
			i++
		}
		if idx = cache[text[j]-'a']; idx == i-j {
			cnt, j = max(cnt, i-j), i
			continue
		}
		k, i = i, i+1
		for i < n && text[i] == text[j] {
			i++
		}
		cnt, j, i = max(cnt, min(idx, i-j)), k, k
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
