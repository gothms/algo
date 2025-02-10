package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func incremovableSubarrayCount(nums []int) int64 {
	i, n := 0, len(nums)
	for i < n-1 && nums[i] < nums[i+1] {
		i++
	}
	if i == n-1 {
		return int64(n * (n + 1) >> 1)
	}
	ans := i + 2
	for j := n - 1; j == n-1 || nums[j] < nums[j+1]; j-- {
		for i >= 0 && nums[i] >= nums[j] {
			i--
		}
		ans += i + 2
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
