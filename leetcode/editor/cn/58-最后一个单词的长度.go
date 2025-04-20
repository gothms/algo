package main

import "slices"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLastWord(s string) int {
	n := len(s)
	j := n
	for i, c := range slices.Backward([]byte(s)) {
		switch c {
		case ' ':
			if j < n {
				return j - i
			}
		default:
			if j == n {
				j = i
			}
		}
	}
	return j+1
}

//leetcode submit region end(Prohibit modification and deletion)
