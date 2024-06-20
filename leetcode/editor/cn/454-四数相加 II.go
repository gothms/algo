package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 分组
	ans, n := 0, len(nums1)
	memo := make(map[int]int, n*n)
	for _, u := range nums3 {
		for _, v := range nums4 {
			memo[-u-v]++
		}
	}
	for _, u := range nums1 {
		for _, v := range nums2 {
			ans += memo[u+v]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
