package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func intersect(nums1 []int, nums2 []int) []int {
	ans, memo := make([]int, 0), make(map[int]int)
	for _, v := range nums1 {
		memo[v]++
	}
	for _, v := range nums2 {
		if memo[v] > 0 {
			memo[v]--
			ans = append(ans, v)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
