package main

func main() {

}
func maximumTripletValue(nums []int) int64 {
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	// lc
	ans, maxDiff, preMax := 0, 0, 0
	for _, x := range nums {
		ans = maxVal(ans, maxDiff*x)        // 最大结果
		maxDiff = maxVal(maxDiff, preMax-x) // 左侧最大差值
		preMax = maxVal(preMax, x)          // 左侧最大值
	}
	return int64(ans)

	// 单调栈
	//mtv, mv, n := int64(0), int64(0), len(nums)
	//stack := make([]int, 0, n)
	//for i, v := range nums {
	//	last := len(stack) - 1
	//	for last >= 0 && v >= nums[stack[last]] { // 单调递减
	//		last--
	//	}
	//	stack = stack[:last+1]
	//	stack = append(stack, i)
	//	if i > 1 {
	//		mtv = maxVal(mtv, mv*int64(v)) // 左侧的最大差值 * nums[i]
	//	}
	//	if last >= 0 {
	//		mv = maxVal(mv, int64(nums[stack[0]]-v)) // 左侧最大差值
	//	}
	//}
	//return mtv
}
