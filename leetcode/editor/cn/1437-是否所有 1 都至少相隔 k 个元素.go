package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func kLengthApart(nums []int, k int) bool {
	last := -k - 1
	for i, v := range nums {
		if v == 1 {
			if i-last <= k {
				return false
			}
			last = i
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
