package main

func main() {

}
func maxSum(nums []int, m int, k int) int64 {
	maxVal := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	max, sum, i, n := int64(0), int64(0), 0, len(nums)
	window := make(map[int]int)
	for ; i < k; i++ {
		window[nums[i]]++
		sum += int64(nums[i])
	}
	if len(window) >= m {
		max = sum
	}
	for j := 0; i < n; i, j = i+1, j+1 {
		window[nums[i]]++
		window[nums[j]]--
		if window[nums[j]] == 0 {
			delete(window, nums[j])
		}
		sum += int64(nums[i] - nums[j])
		if len(window) >= m {
			max = maxVal(max, sum)
		}
	}
	return max
}
