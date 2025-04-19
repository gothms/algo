package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	for l, r := 0, len(nums)-1; l <= r; {
		m := (l + r) >> 1
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
