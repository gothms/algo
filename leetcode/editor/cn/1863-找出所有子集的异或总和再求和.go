package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func subsetXORSum(nums []int) int {
	xor := 0
	for _, v := range nums {
		xor |= v
	}
	return xor << (len(nums) - 1)
}

//leetcode submit region end(Prohibit modification and deletion)
