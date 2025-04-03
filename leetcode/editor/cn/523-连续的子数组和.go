package main

import "fmt"

func main() {
	nums := []int{23, 2, 4, 6, 6} // true
	k := 7
	nums = []int{0} // false
	k = 1
	subarraySum := checkSubarraySum(nums, k)
	fmt.Println(subarraySum)
}

// leetcode submit region begin(Prohibit modification and deletion)
func checkSubarraySum(nums []int, k int) bool {
	memo := make(map[int]struct{})
	sum, pre := nums[0], 0
	for i, v := range nums[1:] {
		memo[pre] = struct{}{}
		sum = (sum + v) % k
		if _, ok := memo[sum]; ok {
			return true
		}
		pre = (pre + nums[i]) % k // currI - 1
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
