package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minAddToMakeValid(s string) int {
	ans, l := 0, 0
	for _, c := range s {
		switch c {
		case '(':
			l++
		case ')':
			if l == 0 {
				ans++
			} else {
				l--
			}
		}
	}
	return ans + l
}

//leetcode submit region end(Prohibit modification and deletion)
