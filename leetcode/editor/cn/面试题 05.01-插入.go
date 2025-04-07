package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func insertBits(N int, M int, i int, j int) int {
	//x := (1<<(j-i+1) - 1) << i
	//return N&^x | M<<i&x // 不 &x 也可以，题已经“保证”了
	return N&^((1<<(j-i+1)-1)<<i) | M<<i
}

//leetcode submit region end(Prohibit modification and deletion)
