package main

import "slices"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func plusOne(digits []int) []int {
	idx := -1
	for i, v := range slices.Backward(digits) {
		if v != 9 {
			idx = i
			break
		}
		digits[i] = 0
	}
	if idx == -1 {
		return append([]int{1}, digits...)
	}
	digits[idx]++
	return digits
}

//leetcode submit region end(Prohibit modification and deletion)
