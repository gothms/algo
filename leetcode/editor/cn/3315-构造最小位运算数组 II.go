package main

import "fmt"

func main() {
	v := 5
	fmt.Println(v &^ ((v + 1) & (-v - 1) >> 1))
}

// leetcode submit region begin(Prohibit modification and deletion)
func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		if v&1 == 0 {
			ans[i] = -1
		} else {
			ans[i] = v &^ ((v + 1) & (-v - 1) >> 1) // 100111 变 100011
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
