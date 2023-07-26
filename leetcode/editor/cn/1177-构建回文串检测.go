//给你一个字符串 s，请你对 s 的子串进行检测。
//
// 每次检测，待检子串都可以表示为 queries[i] = [left, right, k]。我们可以 重新排列 子串 s[left], ..., s[
//right]，并从中选择 最多 k 项替换成任何小写英文字母。
//
// 如果在上述检测过程中，子串可以变成回文形式的字符串，那么检测结果为 true，否则结果为 false。
//
// 返回答案数组 answer[]，其中 answer[i] 是第 i 个待检子串 queries[i] 的检测结果。
//
// 注意：在替换时，子串中的每个字母都必须作为 独立的 项进行计数，也就是说，如果 s[left..right] = "aaa" 且 k = 2，我们只能替换
//其中的两个字母。（另外，任何检测都不会修改原始字符串 s，可以认为每次检测都是独立的）
//
//
//
// 示例：
//
// 输入：s = "abcda", queries = [[3,3,0],[1,2,0],[0,3,1],[0,3,2],[0,4,1]]
//输出：[true,false,false,true,true]
//解释：
//queries[0] : 子串 = "d"，回文。
//queries[1] : 子串 = "bc"，不是回文。
//queries[2] : 子串 = "abcd"，只替换 1 个字符是变不成回文串的。
//queries[3] : 子串 = "abcd"，可以变成回文的 "abba"。 也可以变成 "baab"，先重新排序变成 "bacd"，然后把 "cd"
//替换为 "ab"。
//queries[4] : 子串 = "abcda"，可以变成回文的 "abcba"。
//
//
//
//
// 提示：
//
//
// 1 <= s.length, queries.length <= 10^5
// 0 <= queries[i][0] <= queries[i][1] < s.length
// 0 <= queries[i][2] <= s.length
// s 中只有小写英文字母
//
//
// Related Topics 位运算 数组 哈希表 字符串 前缀和 👍 80 👎 0

package main

import (
	"fmt"
	"math/bits"
)

func main() {
	s := "shezu"
	queries := [][]int{{3, 3, 1}, {3, 4, 2}, {2, 2, 1}, {3, 4, 2}}
	paliQueries := canMakePaliQueries(s, queries)
	fmt.Println(paliQueries)
}

/*
思路：位运算
	1.读懂题意：[left,right] 的子串可以是任意组合变成回文串，而不是原子串的字符位置都不变
	2.对于任意的 queries[i]，left=queries[i][0]，right=queries[i][1]，k=queries[i][2]
		s[left,right] 之间的 [a,z] 的字符数量，如果为偶数，则不用替换；如果为奇数，则最多替换一次
		如 s[left,right] 之间的子串中，字符'a' 的数量分别为：
			4：则不用替换，排列为 aa...aa，'a'字符就满足回文
			5：则最多替换一次，排列为 aaa...xaa 替换一次，排列为 aa...a...aa，替换0次
	3.异或运算恰好能使一个 数字x 的偶数次异或运算 =0，奇数次异或运算 =x
		3.1.而 a-z 的26个字符 <=32，能使用 int32 的每一位来标识 a-z 的每一个字符
		3.2.s[left,right] 子串的所有字符进行异或运算，得到一个数字 bitL2R，假设 'a'字符 出现：
			偶数次：则 bitL2R 的第0位（从低位往高位）的二进制为 0
			奇数次：bitL2R 的第0位的二进制为 1
		3.3.bitL2R 二进制的 1 的数量 oneCnt 就是需要考虑替换的字符的数量
			当 oneCnt 除以 2 后，小于等于 k，则表示子串 s[left,right] 经过 k 次替换可以变成回文串
	4.根据异或运算的特点，可以先将 s 字符串的所有字符都进行 异或运算，记为 异或运算“前缀和” 数组 cBits
		s[left,right] 子串的异或运算值 = cBits[left-1] ^ cBits[r]
*/
// leetcode submit region begin(Prohibit modification and deletion)
func canMakePaliQueries(s string, queries [][]int) []bool {
	// 位运算
	n, m := len(s), len(queries)
	can, cBits := make([]bool, m), make([]int, n+1) // 哨兵：避免 cBits[l-1]
	for i := 0; i < n; i++ {
		cBits[i+1] = cBits[i] ^ 1<<(s[i]-'a')
	} // 计算s串的二进制前缀和：异或运算
	ok := func(l, r, k int) bool { // 验证 [l,r] 的子串经过 k 次替换，能否变成 回文串
		//bitL2R := uint(cBits[l] ^ cBits[r+1])	// [l,r] 之间的字符 和
		//oneCnt := bits.OnesCount(bitL2R)	// bitL2R 二进制有多少个 1
		//return oneCnt>>1 <= k	// 字符数经过 k 次替换，能否变成 回文串
		return bits.OnesCount(uint(cBits[l]^cBits[r+1]))>>1 <= k // 上面三行
	}
	for i := 0; i < m; i++ {
		if l, r := queries[i][0], queries[i][1]; l == r || ok(l, r, queries[i][2]) {
			can[i] = true // l==r：一个字符的子串
		}
	}
	return can

	// 理解错题意：子串顺序可以不是原顺序
	//ok := func(l, r, k int) bool {
	//	for k > 0 && l < r {
	//		if s[l] != s[r] {
	//			k--
	//		}
	//		l++
	//		r--
	//	}
	//	return l >= r
	//}
	//n := len(queries)
	//can := make([]bool, n)
	//for i := 0; i < n; i++ {
	//	if l, r := queries[i][0], queries[i][1]; l == r ||
	//		queries[i][2] >= (r-l+1)>>1 || ok(l, r, queries[i][2]) {
	//		can[i] = true
	//	}
	//}
	//return can
}

//leetcode submit region end(Prohibit modification and deletion)
