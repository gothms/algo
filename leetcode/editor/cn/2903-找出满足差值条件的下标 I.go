package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	// lc：双指针
	n := len(nums)
	var maxIdx, minIdx int
	for i := indexDifference; i < n; i++ {
		j := i - indexDifference
		if nums[j] > nums[maxIdx] {
			maxIdx = j
		} else if nums[j] < nums[minIdx] {
			minIdx = j
		}
		if nums[maxIdx]-nums[i] >= valueDifference {
			return []int{maxIdx, i}
		}
		if nums[i]-nums[minIdx] >= valueDifference {
			return []int{minIdx, i}
		}
	}
	return []int{-1, -1}

	//n := len(nums)
	//for i := indexDifference; i < n; i++ {
	//	for j := i - indexDifference; j >= 0; j-- {
	//		d := nums[i] - nums[j]
	//		if d < 0 {
	//			d = -d
	//		}
	//		if d >= valueDifference {
	//			return []int{j, i}
	//		}
	//	}
	//}
	//return []int{-1, -1}
}

//leetcode submit region end(Prohibit modification and deletion)
