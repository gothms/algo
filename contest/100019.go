package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 0, 1, 2}
	nums = []int{0, 0}
	nums = []int{8, 10, 23,
		26, 21, 28, 21, 14,
		21, 14, 9,
		16, 24, 29, 7, 26} // 4
	//nums = []int{22, 21, 29, 22}	// 1
	subarrays := maxSubarrays(nums)
	fmt.Println(subarrays)
	//v := 1<<30 - 1
	//for i := 0; i < 3; i++ {
	//	v &= nums[i]
	//}
	//fmt.Println(v)
	fmt.Println(1<<31 - 1)
}
func maxSubarrays(nums []int) int {
	and, n := nums[0], len(nums)
	for i := 1; i < n; i++ {
		and &= nums[i]
	}
	if and > 0 { // 总与值 >0，则不分割
		return 1
	}
	cnt, m := 0, 1<<31-1
	for i, v := 0, m; i < n; i++ {
		v &= nums[i]
		if v == 0 { // 与值为 0，分割出一个子数组
			v = m // 收尾的与值可能不是 0
			cnt++
		}
	}
	return cnt
}
