package main

import "fmt"

func main() {
	word := "ieaouqqieaouqq" // 3
	k := 1
	//word = "iqeaouqi" // 3
	//k = 2
	//word = "ieiaoud" // 2
	//k = 0
	//word = "aaeuoiouee" // 10
	//k = 0
	substrings := countOfSubstrings(word, k)
	fmt.Println(substrings)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countOfSubstrings(word string, k int) int64 {
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
			for len(memo) == 5 && cnt >= k { // 找到最短的符合 cnt >= k 的子串
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
	return int64(f(k) - f(k+1))
}

//leetcode submit region end(Prohibit modification and deletion)
