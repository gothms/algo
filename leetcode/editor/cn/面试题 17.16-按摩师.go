package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func massage(nums []int) int {
	// dp：个人
	pre, cur := 0, 0
	for _, v := range nums {
		pre, cur = cur, max(cur, pre+v)
	}
	return cur
}

//leetcode submit region end(Prohibit modification and deletion)
