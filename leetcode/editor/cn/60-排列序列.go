package main

import (
	"fmt"
	"strings"
)

func main() {
	n, k := 3, 3
	n, k = 4, 9
	n, k = 2, 2
	n, k = 3, 5 // 312
	permutation := getPermutation(n, k)
	fmt.Println(permutation)
}

// leetcode submit region begin(Prohibit modification and deletion)
func getPermutation(n int, k int) string {
	// 个人
	cache := make([]int, n+1)
	for i := range cache {
		cache[i] = i + 1
	}
	var sb strings.Builder
	cnt := func(i int) int {
		cnt := 1
		for ; i > 1; i-- {
			cnt *= i
		}
		return cnt
	}(n - 1)
	for i, c := n, k-1; i > 1; i-- {
		idx := c / cnt
		sb.WriteRune(rune(cache[idx] + '0'))
		cache = append(cache[:idx], cache[idx+1:]...)
		c %= cnt
		cnt /= i - 1
	}
	sb.WriteRune(rune(cache[0] + '0'))
	return sb.String()

	// lc

}

//leetcode submit region end(Prohibit modification and deletion)
