package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	ans, cnt := 0, 0
	for _, v := range nums {
		if v == ans {
			cnt++
		} else if cnt == 0 {
			ans = v
			cnt = 1
		} else {
			cnt--
		}
	}
	cnt = 0
	for _, v := range nums {
		if v == ans {
			cnt++
		}
	}
	if cnt <= len(nums)>>1 {
		return -1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
