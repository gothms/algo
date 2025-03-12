package main

import "slices"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countSubarrays(nums []int, k int) int64 {
	maxV := slices.Max(nums)
	ans, l, cnt := 0, 0, 0
	for _, v := range nums {
		if v == maxV {
			cnt++
		}
		for cnt == k {
			if nums[l] == maxV {
				cnt--
			}
			l++
		}
		ans += l
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
