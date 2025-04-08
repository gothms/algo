package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	// lc：将 [26]int 作为 key

	memo := make(map[string][]string)
	for _, s := range strs {
		buf := []byte(s)
		sort.Slice(buf, func(i, j int) bool {
			return buf[i] < buf[j]
		})
		key := string(buf)
		memo[key] = append(memo[key], s)
	}
	ans := make([][]string, 0, len(memo))
	for _, s := range memo {
		ans = append(ans, s)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
