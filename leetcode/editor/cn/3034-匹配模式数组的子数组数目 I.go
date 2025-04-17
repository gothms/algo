package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countMatchingSubarrays(nums []int, pattern []int) int {
	ans := 0
out:
	for i := range nums[:len(nums)-len(pattern)] {
		for j, v := range pattern {
			switch v {
			case 1:
				if nums[i+j] >= nums[i+j+1] {
					continue out
				}
			case 0:
				if nums[i+j] != nums[i+j+1] {
					continue out
				}
			case -1:
				if nums[i+j] <= nums[i+j+1] {
					continue out
				}
			}
		}
		ans++
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
