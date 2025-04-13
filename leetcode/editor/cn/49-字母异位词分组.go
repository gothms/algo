package main

import (
	"slices"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	memo := make(map[string][]string)
	for _, s := range strs {
		buf := []byte(s)
		slices.Sort(buf)
		key := string(buf)
		memo[key] = append(memo[key], s)
	}
	ans := make([][]string, 0, len(memo))
	for _, v := range memo {
		ans = append(ans, v)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
