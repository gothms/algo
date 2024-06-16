package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}

//leetcode submit region end(Prohibit modification and deletion)
