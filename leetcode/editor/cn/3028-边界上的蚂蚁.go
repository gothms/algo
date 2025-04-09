package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func returnToBoundaryCount(nums []int) int {
	ans, s := 0, 0
	for _, v := range nums {
		if s += v; s == 0 {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
