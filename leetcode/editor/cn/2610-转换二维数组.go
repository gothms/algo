package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findMatrix(nums []int) [][]int {
	// lc
	memo := make(map[int]int)
	for _, v := range nums {
		memo[v]++
	}
	ans := make([][]int, 0)
	for len(memo) > 0 {
		arr := make([]int, 0, len(memo))
		for k, v := range memo {
			arr = append(arr, k)
			if v == 1 {
				delete(memo, k)
			} else {
				memo[k]--
			}
		}
		ans = append(ans, arr)
	}
	return ans

	// 个人
	//sort.Ints(nums)
	//n := len(nums)
	//ans := make([][]int, 1)
	//ans[0] = make([]int, 0)
	//for i := 0; i < n; {
	//	v := nums[i]
	//	ans[0] = append(ans[0], v)
	//	j := 1
	//	for ; i+j < n && nums[i+j] == v; j++ {
	//		if len(ans) == j {
	//			ans = append(ans, make([]int, 0))
	//		}
	//		ans[j] = append(ans[j], v)
	//	}
	//	i += j
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
