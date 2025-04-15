package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	words := []string{"aa", "bc"}
	operations := maxPalindromesAfterOperations(words)
	fmt.Println(operations)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxPalindromesAfterOperations(words []string) int {
	// lc：以及逆向思维方法
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	tot, mask := 0, 0
	for _, w := range words {
		tot += len(w)
		for _, c := range w {
			mask ^= 1 << (c - 'a')
		}
	}
	tot -= bits.OnesCount(uint(mask))
	ans := 0
	for _, w := range words {
		if tot -= len(w) &^ 1; tot < 0 {
			break
		}
		ans++
	}
	return ans

	// 个人
	//sort.Slice(words, func(i, j int) bool {
	//	return len(words[i]) < len(words[j])
	//})
	//memo := make([]int, 26)
	//for _, w := range words {
	//	for _, c := range w {
	//		memo[c-'a']++
	//	}
	//}
	//sort.Ints(memo)
	//ans, i := 0, len(memo)-1
	//for _, w := range words {
	//	k := len(w) &^ 1
	//	if k == 0 {
	//		ans++
	//		continue
	//	}
	//	for k > 0 && i >= 0 {
	//		k -= memo[i] &^ 1
	//		i--
	//	}
	//	if k > 0 { // 间接保证了 i 的有效性
	//		break
	//	}
	//	ans++
	//	i++
	//	memo[i] = -k // k 已经 <=0
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
