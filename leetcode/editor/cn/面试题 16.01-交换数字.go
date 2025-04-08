package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func swapNumbers(numbers []int) []int {
	numbers[0] ^= numbers[1]
	numbers[1] ^= numbers[0]
	numbers[0] ^= numbers[1]
	return numbers
}

//leetcode submit region end(Prohibit modification and deletion)
