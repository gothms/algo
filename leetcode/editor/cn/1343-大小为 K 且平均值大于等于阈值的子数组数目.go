package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numOfSubarrays(arr []int, k int, threshold int) int {
	// lc
	s, t := 0, k*threshold
	ans := 0
	for i, v := range arr {
		s += v
		if i < k-1 {
			continue
		}
		if s >= t {
			ans++
		}
		s -= arr[i-k+1]
	}
	return ans

	// 滑动窗口
	//ans := 0
	//var s, n, t float64 = 0, float64(k), float64(threshold)
	//for i, v := range arr {
	//	if i >= k {
	//		if s/n >= t {
	//			ans++
	//		}
	//		s -= float64(arr[i-k])
	//	}
	//	s += float64(v)
	//}
	//if s/n >= t {
	//	ans++
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
