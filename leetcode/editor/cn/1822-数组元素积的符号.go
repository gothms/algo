package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func arraySign(nums []int) int {
	xor := 0
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		xor ^= v
	}
	if xor >= 0 {
		return 1
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
