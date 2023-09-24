package main

import "fmt"

func main() {
	arr := []int{6, 5, 3, 9, 2, 7}
	heights := maximumSumOfHeights(arr)
	fmt.Println(heights)
}
func maximumSumOfHeights(maxHeights []int) int64 {
	// dp+单调栈
	maxVal := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}
	msoh, n := int64(0), len(maxHeights)
	dpUp, dpDown := make([]int, n+1), make([]int, n+1)
	stack := []int{-1}             // dpUp 第一个元素为 0
	for i, h := range maxHeights { // 从左往右
		last := len(stack) - 1
		for last > 0 && maxHeights[stack[last]] >= h { // 递增单调栈
			last--
		}
		dpUp[i+1] = dpUp[stack[last]+1] + (i-stack[last])*h
		stack = stack[:last+1]   // 维持单调递增
		stack = append(stack, i) // 添加索引
	}
	stack = stack[:1]
	stack[0] = n                  // dpDown 最后一个元素为 0
	for i := n - 1; i >= 0; i-- { // 从右往左
		last := len(stack) - 1
		for last > 0 && maxHeights[stack[last]] >= maxHeights[i] { // 递增单调栈
			last--
		}
		dpDown[i] = dpDown[stack[last]] + (stack[last]-i)*maxHeights[i]
		stack = stack[:last+1]
		stack = append(stack, i)
	}
	for i := 0; i <= n; i++ {
		msoh = maxVal(msoh, int64(dpUp[i]+dpDown[i]))
	}
	return msoh
}
