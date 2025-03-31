package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println('a', 'A')
}

// leetcode submit region begin(Prohibit modification and deletion)
func frequencySort(s string) string {
	memo := make(map[rune]int)
	for _, c := range s {
		memo[c]++
	}
	idx := make([]rune, len(memo))
	i := 0
	for k := range memo {
		idx[i] = k
		i++
	}
	sort.Slice(idx, func(i, j int) bool {
		return memo[idx[i]] > memo[idx[j]]
	})
	var ans strings.Builder
	for _, c := range idx {
		ans.Write(bytes.Repeat([]byte{byte(c)}, memo[c]))
	}
	return ans.String()
}

//leetcode submit region end(Prohibit modification and deletion)
