package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxDistance(nums1 []int, nums2 []int) int {
	// 优化
	ans, i, n := 0, 0, len(nums1)
	for j, v := range nums2 {
		for i < n && nums1[i] > v {
			i++
		}
		if i == n {
			break
		}
		ans = max(ans, j-i)
	}
	return ans

	// 个人
	//ans, i, n := 0, 0, len(nums1)
	//for j, v := range nums2 {
	//	for i <= j && i < n && nums1[i] > v {
	//		i++
	//	}
	//	if i == n {
	//		break
	//	}
	//	if i > j {
	//		continue
	//	}
	//	ans = max(ans, j-i)
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
