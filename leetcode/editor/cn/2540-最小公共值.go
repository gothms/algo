package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func getCommon(nums1 []int, nums2 []int) int {
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
