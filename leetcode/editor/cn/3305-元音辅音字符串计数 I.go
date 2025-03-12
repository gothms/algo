package main

import (
	"fmt"
)

func main() {
	word := "ieaouqqieaouqq" // 3
	k := 1
	word = "iqeaouqi" // 3
	k = 2
	word = "ieiaoud" // 2
	k = 0
	word = "aaeuoiouee" // 10
	k = 0
	substrings := countOfSubstrings(word, k)
	fmt.Println(substrings)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countOfSubstrings(word string, k int) int {
	y := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	f := func(k int) int {
		memo := make(map[byte]int)
		ans, l, cnt := 0, 0, 0
		for i := range word {
			c := word[i]
			if y[c] {
				memo[c]++
			} else {
				cnt++
			}
			for len(memo) == 5 && cnt >= k {
				c = word[l]
				if y[c] {
					if memo[c] == 1 {
						delete(memo, c)
					} else {
						memo[c]--
					}
				} else {
					cnt--
				}
				l++
			}
			ans += l
		}
		return ans
	}
	return f(k) - f(k+1)

	// 滑动窗体：失败 word = "aaeuoiouee"
	//y := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	//memo := make(map[byte]int)
	//ans, l, cnt := 0, 0, 0
	//for i := range word {
	//	c := word[i]
	//	if y[c] {
	//		memo[c]++
	//	} else {
	//		cnt++
	//	}
	//	if len(memo) == 5 && cnt == k {
	//		ans++
	//		continue
	//	}
	//	for cnt > k {
	//		c = word[l]
	//		if y[c] {
	//			if memo[c] == 1 {
	//				delete(memo, c)
	//			} else {
	//				memo[c]--
	//				if len(memo) == 5 {
	//					ans++
	//				}
	//			}
	//		} else {
	//			cnt--
	//			if len(memo) == 5 {
	//				ans++
	//			}
	//		}
	//		l++
	//	}
	//}
	//if len(memo) == 5 {
	//	for c := word[l]; y[c] && cnt == k; c = word[l] {
	//		if memo[c] == 1 {
	//			break
	//		} else {
	//			memo[c]--
	//		}
	//		ans++
	//		l++
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
