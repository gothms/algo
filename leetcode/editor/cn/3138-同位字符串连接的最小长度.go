//给你一个字符串 s ，它由某个字符串 t 和若干 t 的 同位字符串 连接而成。
//
// 请你返回字符串 t 的 最小 可能长度。
//
// 同位字符串 指的是重新排列一个单词得到的另外一个字符串，原来字符串中的每个字符在新字符串中都恰好只使用一次。
//
//
//
// 示例 1：
//
//
// 输入：s = "abba"
//
//
// 输出：2
//
// 解释：
//
// 一个可能的字符串 t 为 "ba" 。
//
// 示例 2：
//
//
// 输入：s = "cdef"
//
//
// 输出：4
//
// 解释：
//
// 一个可能的字符串 t 为 "cdef" ，注意 t 可能等于 s 。
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁵
// s 只包含小写英文字母。
//
//
// Related Topics 哈希表 字符串 计数 👍 7 👎 0

package main

import (
	"fmt"
)

func main() {
	s := "abba"
	s = "abbabbaab"
	s = "cdef"
	s = "abbcbcac"
	s = "aabbbb"
	s = "aabbabab"
	s = "jjj"
	length := minAnagramLength(s)
	fmt.Println(length)
}

// leetcode submit region begin(Prohibit modification and deletion)

//func init() {
//	prime := make([]bool, n3138)
//	for i := 2; i < n3138; i++ {
//		if !prime[i] {
//			for j := i << 1; j < n3138; j += i {
//				prime[j] = true
//			}
//			prime3138 = append(prime3138, i)
//		}
//	}
//	//fmt.Println(prime3138)
//}
//
//const n3138 = 1e5 + 1
//
//var prime3138 []int

func minAnagramLength(s string) int {
	// lc
	// 题意：t = aabb，则同位字符串可以是 abab
	n := len(s)
out:
	for i := 1; i <= n>>1; i++ {
		if n%i != 0 {
			continue
		}
		cache := [26]int{}
		for j := 0; j < i; j++ {
			cache[s[j]-'a']++
		}
		for j := i; j < n; j += i {
			temp := [26]int{}
			for k := j; k < j+i; k++ {
				x := s[k] - 'a'
				if temp[x] == cache[x] {
					continue out
				}
				temp[x]++
			}
		}
		return i
	}
	return n

	// 个人：失败 s = "aabbabab"
	// 题意都读错了
	//memo := [26]int{}
	//for _, c := range s {
	//	memo[c-'a']++
	//}
	//count := 0
	//for _, c := range memo {
	//	if c > 0 {
	//		count++
	//	}
	//}
	//ans, last, start, cnt := count, 0, 0, count
	//cache := [26]int{}
	//for i, c := range s {
	//	idx := c - 'a'
	//	if cache[idx] == 0 {
	//		cnt--
	//	}
	//	cache[c-'a']++
	//	if cnt > 0 {
	//		continue
	//	}
	//	if i-start+1 != count {
	//		j := start
	//		for ; cache[s[j]-'a'] > 1; j++ {
	//			cache[s[j]-'a']--
	//		}
	//		if i-j+1 == count {
	//			if j-last < count {
	//				ans = max(ans, j-last+count)
	//			} else {
	//				ans = max(ans, j-last)
	//			}
	//		} else {
	//			ans = max(ans, i-last+1)
	//		}
	//	}
	//	start, cnt = i+1, count
	//	if ans == count {
	//		last = start
	//	}
	//	cache = [26]int{}
	//}
	//fmt.Println(ans, last, cnt)
	//if cnt < count && cnt > 0 {
	//	ans = max(ans, len(s)-last)
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
