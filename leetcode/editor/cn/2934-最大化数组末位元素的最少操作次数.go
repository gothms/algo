package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums1 []int, nums2 []int) int {
	ans, all, n := 0, 0, len(nums1)
	v1, v2 := nums1[n-1], nums2[n-1]
	m := min(v1, v2)
	for i := 0; i < n-1; i++ {
		if nums1[i] > v1 || nums2[i] > v2 { // 不换 n-1 的情况下，ans 记录的都需要换
			if nums1[i] > v2 || nums2[i] > v1 {
				return -1
			}
			ans++
		} else if max(nums1[i], nums2[i]) <= m { // 不管 n-1 换不换，all 记录的都不用换
			all++
		}
	}
	return min(ans, n-ans-all) // 不换 n-1 / 换 n-1
}

//leetcode submit region end(Prohibit modification and deletion)
