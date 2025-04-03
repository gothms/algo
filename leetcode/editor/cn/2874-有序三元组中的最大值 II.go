package main

import "fmt"

func main() {
	nums := []int{15, 3, 3, 18, 19, 13, 7, 5, 18, 1, 8, 5} // 252
	value := maximumTripletValue(nums)
	fmt.Println(value)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumTripletValue(nums []int) int64 {
	// 枚举 k
	ans, maxDif, mx := 0, 0, 0
	for _, v := range nums {
		ans = max(ans, maxDif*v)
		maxDif = max(maxDif, mx-v)
		mx = max(mx, v)
	}
	return int64(ans)
}

// leetcode submit region end(Prohibit modification and deletion)

func maximumTripletValue_(nums []int) int64 {
	// O(n) 做法，参见 2874

	// 枚举 k：优化
	//ans, maxDif, mx := 0, 0, 0
	//for _, v := range nums {
	//	// k j i 三者的记录顺序为倒序
	//	ans = max(ans, maxDif*v)   // 枚举 nums[k]=v
	//	maxDif = max(maxDif, mx-v) // 记录 maxDif
	//	mx = max(mx, v)            // 记录 mx
	//}
	//return int64(ans)

	// 枚举 k：stack
	stack := make([]int, 0, len(nums))
	ans, maxDif := 0, 0
	for _, v := range nums {
		for len(stack) > 0 && stack[len(stack)-1] <= v {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
		ans = max(ans, maxDif*v)
		maxDif = max(maxDif, stack[0]-v)
	}
	return int64(ans)

	// 枚举 j
	//n := len(nums)
	//suf := make([]int, n+1)
	//for i := n - 1; i > 1; i-- {
	//	suf[i] = max(suf[i+1], nums[i])
	//}
	//ans := 0
	//for j, mx := 1, 0; j < n-1; j++ {
	//	mx = max(mx, nums[j-1])
	//	ans = max(ans, (mx-nums[j])*suf[j+1])
	//}
	//return int64(ans)
}
