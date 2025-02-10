package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isArraySpecial(nums []int) bool {
	t := -1
	for _, v := range nums {
		if cur := v & 1; cur == t {
			return false
		} else {
			t = cur
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
