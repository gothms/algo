package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numSplits(s string) int {
	// 或者用 [26] 数组
	l, r := make(map[uint8]int), make(map[uint8]int)
	for i := range s {
		r[s[i]]++
	}
	ans := 0
	for i := range s[:len(s)-1] {
		c := s[i]
		l[c]++
		r[c]--
		if r[c] == 0 {
			delete(r, c)
		}
		if len(l) == len(r) {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
