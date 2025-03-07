package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
