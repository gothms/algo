package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	// dp

	// 滑动窗体：lc
	var s1, s2, s3, m1, m2, m3, i1, i2, idx1 int
	ret := make([]int, 3)
	for i := k << 1; i < len(nums); i++ {
		s1 += nums[i-k<<1] // 向前滑动，求和
		s2 += nums[i-k]
		s3 += nums[i]
		if i >= k*3-1 {
			if s1 > m1 { // m1,m2,m3：分别对应 1-3 个长度为 k 的子数组，的最大和
				m1 = s1
				i1 = i - k*3 + 1 // i1,i2：分别对应 m1,m2 情况下的索引
			}
			if m1+s2 > m2 {
				m2 = m1 + s2
				i2 = i - k<<1 + 1
				idx1 = i1 // 更新第一个索引
			}
			if m2+s3 > m3 {
				m3 = m2 + s3
				ret[0], ret[1], ret[2] = idx1, i2, i-k+1 // 更新
			}
			s1 -= nums[i-k*3+1] // 滑出
			s2 -= nums[i-k<<1+1]
			s3 -= nums[i-k+1]
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
