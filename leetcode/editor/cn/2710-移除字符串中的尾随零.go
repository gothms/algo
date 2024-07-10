package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func removeTrailingZeros(num string) string {
	i := len(num) - 1
	for i >= 0 && num[i] == '0' {
		i--
	}
	return num[:i+1]
}

//leetcode submit region end(Prohibit modification and deletion)
