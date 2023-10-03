package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{1, 2}
	//nums = []int{1, 1, 1, 2, 3}
	target := 72
	subarray := minSizeSubarray(nums, target)
	fmt.Println(subarray)
}
func minSizeSubarray(nums []int, target int) int {
	// 滑动窗体：1 <= nums[i] <= 105
	//minVal := func(a, b int) int {
	//	if b < a {
	//		return b
	//	}
	//	return a
	//}
	//mss, cnt, sum, n := math.MaxInt32, 0, 0, len(nums)
	//for i := 0; i < n; i++ {
	//	sum += nums[i] // 总和
	//}
	//cnt, target, sum = target/sum, target%sum, 0 // 需要加上 nums 的次数为 cnt
	//if target == 0 {
	//	return n * cnt
	//}
	//for i, j := 0, 0; i < n<<1; i++ { // 拼接两个 nums
	//	sum += nums[i%n]   // 右边界
	//	for sum > target { // 左边界
	//		sum -= nums[j%n]
	//		j++
	//	}
	//	if sum == target {
	//		mss = minVal(mss, i-j+1) // 最短子数组
	//	}
	//}
	//if mss == math.MaxInt32 {
	//	return -1
	//}
	//return mss + cnt*n

	// 笨方法
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	n := len(nums)
	sum := make([]int, n+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v // 前缀和
	}
	pre := func(p int) int {
		return sum[n]*(p/n) + sum[p%n] // 计算和
	}
	query := func(l, r int) int {
		return pre(r) - pre(l)
	}
	ans := math.MaxInt32
	for l := range nums {
		r := sort.Search(2e9, func(r int) bool {
			return query(l, r) >= target // 二分查找前缀差
		})
		if query(l, r) == target {
			ans = minVal(ans, r-l)
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
