//给你一个下标从 0 开始的字符串数组 words 以及一个二维整数数组 queries 。
//
// 每个查询 queries[i] = [li, ri] 会要求我们统计在 words 中下标在 li 到 ri 范围内（包含 这两个值）并且以元音开头和结尾
//的字符串的数目。
//
// 返回一个整数数组，其中数组的第 i 个元素对应第 i 个查询的答案。
//
// 注意：元音字母是 'a'、'e'、'i'、'o' 和 'u' 。
//
//
//
// 示例 1：
//
//
//输入：words = ["aba","bcb","ece","aa","e"], queries = [[0,2],[1,4],[1,1]]
//输出：[2,3,0]
//解释：以元音开头和结尾的字符串是 "aba"、"ece"、"aa" 和 "e" 。
//查询 [0,2] 结果为 2（字符串 "aba" 和 "ece"）。
//查询 [1,4] 结果为 3（字符串 "ece"、"aa"、"e"）。
//查询 [1,1] 结果为 0 。
//返回结果 [2,3,0] 。
//
//
// 示例 2：
//
//
//输入：words = ["a","e","i"], queries = [[0,2],[0,1],[2,2]]
//输出：[3,2,1]
//解释：每个字符串都满足这一条件，所以返回 [3,2,1] 。
//
//
//
// 提示：
//
//
// 1 <= words.length <= 10⁵
// 1 <= words[i].length <= 40
// words[i] 仅由小写英文字母组成
// sum(words[i].length) <= 3 * 10⁵
// 1 <= queries.length <= 10⁵
// 0 <= queries[j][0] <= queries[j][1] < words.length
//
//
// Related Topics 数组 字符串 前缀和 👍 15 👎 0

package main

func main() {

}

/*
思路：前缀和
	1.只需要知道 queries[i][0] 到 queries[i][1] 之间有多少个符合题意的元音字符串
		就可以查询任意区间的元音字符串数量
		1.1.将 words 中字符串遍历，判断queries[i]是否是元音字符串，记录到数组 prefix
		1.2.计算 prefix 的前缀和，就可以 O(1) 的查询了
	2.上面使用的是前缀和，而前缀和问题，还可以使用 树状数组 和 线段树
		2.1.后两者也能都解决问题，而且支持动态修改前缀和
		2.2.树状数组能解决的问题，线段树都能解决，反之则不然
			这里提供一种树状数组的实现
*/
//n, m := len(words), len(queries)
//prefix, ans := make([]int, n), make([]int, m)
//ok := func(c byte) bool {
//	switch c {
//	case 'a', 'e', 'i', 'o', 'u':
//		return true
//	}
//	return false
//}
//for i := 0; i < n; i++ {
//	if ok(words[i][0]) && ok(words[i][len(words[i])-1]) {
//		prefix[i] = 1
//	}
//}
//bit := BIT(prefix)
//for i := 0; i < m; i++ {
//	ans[i] = rangeSum(bit, queries[i][0], queries[i][1])
//}
//return ans
//leetcode submit region begin(Prohibit modification and deletion)
func vowelStrings(words []string, queries [][]int) []int {
	n, m := len(words), len(queries)
	prefix, ans := make([]int, n+1), make([]int, m)
	ok := func(c byte) bool {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}
		return false
	}
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i]
		if ok(words[i][0]) && ok(words[i][len(words[i])-1]) {
			prefix[i+1]++
		}
	}
	for i := 0; i < m; i++ {
		ans[i] = prefix[queries[i][1]+1] - prefix[queries[i][0]]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
