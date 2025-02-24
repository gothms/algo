package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func evenOddBit(n int) []int {
	ans := make([]int, 2)
	for i := 0; n != 0; n >>= 1 {
		ans[i] += n & 1
		i ^= 1 // 切换奇偶
	}
	return ans

	//eb := make([]int, 2)
	//for n > 0 {
	//	if n&1 == 1 {
	//		eb[0]++
	//	}
	//	n >>= 1
	//	if n&1 == 1 {
	//		eb[1]++
	//	}
	//	n >>= 1
	//}
	//return eb
}

//leetcode submit region end(Prohibit modification and deletion)
