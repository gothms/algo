package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
