package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	// lc：枚举因子/枚举倍数

	// 个人：超时
	//memo := make(map[int]int)
	//for _, v := range nums1 {
	//	if v%k == 0 {
	//		memo[v/k]++
	//	}
	//}
	//ans := 0
	//for v, c := range memo {
	//	for _, w := range nums2 {
	//		if v%w == 0 {
	//			ans += c
	//		}
	//	}
	//}
	//return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
