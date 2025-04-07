package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func exchangeBits(num int) int {
	odd := num & 0x55555555  // 偶数
	even := num & 0xaaaaaaaa // 奇数
	return odd<<1 | even>>1
}

//leetcode submit region end(Prohibit modification and deletion)
