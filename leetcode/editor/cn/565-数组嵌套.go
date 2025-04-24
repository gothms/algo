package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func arrayNesting(nums []int) int {
	ans := 0
	for i, v := range nums {
		nums[i] = -1
		c := 0
		for ; v >= 0; v, nums[v] = nums[v], -1 {
			c++
		}
		ans = max(ans, c)
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

func arrayNesting_(nums []int) int {
	// lc
	ans := 0
	for i := range nums {
		c := 0
		for nums[i] >= 0 {
			i, nums[i] = nums[i], -1
			c++
		}
		ans = max(ans, c)
	}
	return ans

	// 个人
	//ans := 0
	//for i, v := range nums {
	//	c := 1
	//	for v >= 0 && v != i {
	//		v, nums[v] = nums[v], -1
	//		c++
	//	}
	//	nums[i] = -1
	//	if v == i {
	//		ans = max(ans, c)
	//	}
	//}
	//return ans
}
