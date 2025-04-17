package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findLonely(nums []int) []int {
	// lc
	memo := make(map[int]int)
	for _, v := range nums {
		memo[v]++
	}
	ans := make([]int, 0)
	for k, v := range memo {
		if v == 1 && memo[k-1]|memo[k+1] == 0 {
			ans = append(ans, k)
		}
	}
	return ans

	// 个人
	//memo := make(map[int]bool)
	//for _, v := range nums {
	//	lonely := true
	//	for i := v - 1; i <= v+1; i++ {
	//		if _, ok := memo[i]; ok { // 必须判断是否存在，而不是 true
	//			memo[i] = false
	//			lonely = false
	//		}
	//	}
	//	memo[v] = lonely
	//}
	//ans := make([]int, 0, len(memo))
	//for k, v := range memo {
	//	if v {
	//		ans = append(ans, k)
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
