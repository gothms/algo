package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func reverseDegree(s string) int {
	ans := 0
	for i, c := range s {
		ans += int(26-c+'a') * (i + 1)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
