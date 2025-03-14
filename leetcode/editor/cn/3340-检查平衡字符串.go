package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isBalanced(num string) bool {
	v := 0
	for i := len(num) - 1; i >= 0; i-- {
		v += int(num[i]-'0') * (i&1<<1 - 1)
	}
	return v == 0
}

//leetcode submit region end(Prohibit modification and deletion)
