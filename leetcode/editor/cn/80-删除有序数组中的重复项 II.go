package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 3 {
		return n
	}
	k := 2
	for i := 2; i < n; i++ {
		if nums[i] != nums[i-2] {
			nums[k] = nums[i]
			k++
		} else if nums[k] > nums[k-2] {
			k++
		}
	}
	return k
}

//leetcode submit region end(Prohibit modification and deletion)
