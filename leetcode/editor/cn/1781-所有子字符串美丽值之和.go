//一个字符串的 美丽值 定义为：出现频率最高字符与出现频率最低字符的出现次数之差。
//
//
// 比方说，"abaacc" 的美丽值为 3 - 1 = 2 。
//
//
// 给你一个字符串 s ，请你返回它所有子字符串的 美丽值 之和。
//
//
//
// 示例 1：
//
//
//输入：s = "aabcb"
//输出：5
//解释：美丽值不为零的字符串包括 ["aab","aabc","aabcb","abcb","bcb"] ，每一个字符串的美丽值都为 1 。
//
// 示例 2：
//
//
//输入：s = "aabcbaa"
//输出：17
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 500
// s 只包含小写英文字母。
//
//
// Related Topics 哈希表 字符串 计数 👍 92 👎 0

package main

import "fmt"

func main() {
	s := "aabcb"
	s = "aabcbaa"
	s = "xzvfsppsjfbxdwkqe" // 64
	i := beautySum(s)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func beautySum(s string) int {
	// 暴力
	ret := 0
	for i := range s {
		memo := [26]int{}
		//ma, mi, ch := 0, len(s), int32(0)	// 不能通过字符再次出现来判断，因为最小字符可能有多个
		ma, mi := 0, len(s)
		for _, c := range s[i:] {
			idx := c - 'a'
			memo[idx]++
			ma = max(ma, memo[idx]) // 最大
			if memo[idx] == 1 {     // 新增的字符
				mi = 1
			} else if memo[idx]-1 == mi { // 最小次数的字符，被更新了
				mi = len(s)
				for _, v := range memo {
					if v > 0 {
						mi = min(mi, v) // 最小
					}
				}
			}
			//if j >= 2 { // 至少三个字符
			ret += ma - mi
			//}
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
