package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findClosestNumber(nums []int) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	res := nums[0]
	for _, v := range nums[1:] {
		if abs(v) < abs(res) || abs(v) == abs(res) && v > res {
			res = v
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
