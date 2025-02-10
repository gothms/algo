package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sortedSquares(nums []int) []int {
	n := len(nums)
	i := sort.SearchInts(nums, 0)
	j := i - 1
	ans := make([]int, 0, n)
	for j >= 0 && i < n {
		if nums[i] > -nums[j] {
			ans = append(ans, nums[j]*nums[j])
			j--
		} else {
			ans = append(ans, nums[i]*nums[i])
			i++
		}
	}
	if i < n {
		for ; i < n; i++ {
			ans = append(ans, nums[i]*nums[i])
		}
	} else {
		for ; j >= 0; j-- {
			ans = append(ans, nums[j]*nums[j])
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
