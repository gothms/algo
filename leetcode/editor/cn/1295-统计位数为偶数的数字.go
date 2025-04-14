package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findNumbers(nums []int) int {
	ans := 0
	for _, v := range nums {
		k := -1
		for ; v != 0; v /= 10 {
			k++
		}
		ans += k & 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
