package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func percentageLetter(s string, letter byte) int {
	cnt := 0
	for _, c := range s {
		if byte(c) == letter {
			cnt++
		}
	}
	return cnt * 100 / len(s)
}

//leetcode submit region end(Prohibit modification and deletion)
