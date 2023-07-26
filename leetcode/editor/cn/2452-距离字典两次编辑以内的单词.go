//给你两个字符串数组 queries 和 dictionary 。数组中所有单词都只包含小写英文字母，且长度都相同。
//
// 一次 编辑 中，你可以从 queries 中选择一个单词，将任意一个字母修改成任何其他字母。从 queries 中找到所有满足以下条件的字符串：不超过 两
//次编辑内，字符串与 dictionary 中某个字符串相同。
//
// 请你返回 queries 中的单词列表，这些单词距离 dictionary 中的单词 编辑次数 不超过 两次 。单词返回的顺序需要与 queries 中原
//本顺序相同。
//
//
//
// 示例 1：
//
// 输入：queries = ["word","note","ants","wood"], dictionary = ["wood","joke",
//"moat"]
//输出：["word","note","wood"]
//解释：
//- 将 "word" 中的 'r' 换成 'o' ，得到 dictionary 中的单词 "wood" 。
//- 将 "note" 中的 'n' 换成 'j' 且将 't' 换成 'k' ，得到 "joke" 。
//- "ants" 需要超过 2 次编辑才能得到 dictionary 中的单词。
//- "wood" 不需要修改（0 次编辑），就得到 dictionary 中相同的单词。
//所以我们返回 ["word","note","wood"] 。
//
//
// 示例 2：
//
// 输入：queries = ["yes"], dictionary = ["not"]
//输出：[]
//解释：
//"yes" 需要超过 2 次编辑才能得到 "not" 。
//所以我们返回空数组。
//
//
//
//
// 提示：
//
//
// 1 <= queries.length, dictionary.length <= 100
// n == queries[i].length == dictionary[j].length
// 1 <= n <= 100
// 所有 queries[i] 和 dictionary[j] 都只包含小写英文字母。
//
//
// Related Topics 数组 字符串 👍 4 👎 0

package main

func main() {

}

/*
思路：暴力枚举
	对于任意两个字符串，判断它们同位置的字符是否相等
		当不等的字符数 > 2 时，说明两者不符合题意
		当不等的字符数 <= 2 时，将 queries[i] 添加到结果集中
*/
// leetcode submit region begin(Prohibit modification and deletion)
func twoEditWords(queries []string, dictionary []string) []string {
	ans, idx, n := make([]string, len(queries)), 0, len(queries[0])
	match := func(a, b *string) bool { // 两个字符串的编辑次数是否 <= 2
		for i, cnt := 0, 0; i < n; i++ {
			if (*a)[i] != (*b)[i] {
				if cnt == 2 {
					return false
				}
				cnt++
			}
		}
		return true
	}
	for _, q := range queries { // 暴力枚举
		for _, d := range dictionary {
			if match(&q, &d) {
				ans[idx] = q
				idx++
				break // 终止 q 的匹配
			}
		}
	}
	return ans[:idx]
}

//leetcode submit region end(Prohibit modification and deletion)
