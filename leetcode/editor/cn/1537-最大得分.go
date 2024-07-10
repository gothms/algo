package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)

//leetcode submit region end(Prohibit modification and deletion)

func maxSum(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	var i, j, s1, s2, ans int
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			s1 += nums1[i]
			i++
		} else if nums1[i] > nums2[j] {
			s2 += nums2[j]
			j++
		} else {
			ans += max(s1, s2) + nums1[i]
			s1, s2 = 0, 0
			i++
			j++
		}
	}
	for ; i < m; i++ {
		s1 += nums1[i]
	}
	for ; j < n; j++ {
		s2 += nums2[j]
	}
	return (ans + max(s1, s2)) % 1_000_000_007

	// 个人
	// 1.找出值相等的索引对
	//const mod = 1_000_000_007
	//m, n := len(nums1), len(nums2)
	//memo := make(map[int]int, m)
	//for i, v := range nums1 {
	//	memo[v] = i
	//}
	//path := make([][2]int, 0)
	//for i, v := range nums2 {
	//	if idx, ok := memo[v]; ok {
	//		path = append(path, [2]int{idx, i})
	//	}
	//}
	//path = append(path, [2]int{m, n})
	//// 2.前缀和
	//preSum1, preSum2 := make([]int, m+1), make([]int, n+1)
	//for i, v := range nums1 {
	//	preSum1[i+1] = preSum1[i] + v
	//}
	//for i, v := range nums2 {
	//	preSum2[i+1] = preSum2[i] + v
	//}
	//// 3.获取最长的“片段”子数组
	//ans, lx, ly := 0, 0, 0
	//for _, p := range path {
	//	x, y := p[0], p[1]
	//	ans = (ans + max(preSum1[x]-preSum1[lx], preSum2[y]-preSum2[ly])) % mod
	//	lx, ly = x, y
	//}
	//return ans
}
