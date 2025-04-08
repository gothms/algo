package main

import "fmt"

func main() {
	words := []string{"DirNnILhARNS hOYIFB", "SM ", "YSPBaovrZBS", "evMMBOf", "mCrS", "oRJfjw gwuo", "xOpSEXvfI"}
	s := "mCrS"
	i := findString(words, s)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findString(words []string, s string) int {
	// 个人
	l, r := 0, len(words)-1
	for l <= r {
		m := (l + r) >> 1
		ml := m
		for ml >= l && words[ml] == "" {
			ml--
		}
		//fmt.Println(l, ml, words[ml])
		if ml < l || words[ml] < s {
			l = m + 1 // TODO
		} else if words[ml] > s {
			r = ml - 1
		} else {
			return ml
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
