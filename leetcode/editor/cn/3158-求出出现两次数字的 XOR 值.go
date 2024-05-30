package main

import "fmt"

func main() {
	fmt.Println(1 << 31)
	fmt.Println(1 << 62)
}

// leetcode submit region begin(Prohibit modification and deletion)
func duplicateNumbersXOR(nums []int) int {
	ans, t := 0, 0
	for _, v := range nums {
		if t>>v&1 != 0 {
			ans ^= v
		} else {
			t |= 1 << v
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
