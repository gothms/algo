package main

import "fmt"

func main() {
	nums := []int{12, 6, 1, 2, 7}
	//nums = []int{2, 3, 1}
	value := maximumTripletValue(nums)
	fmt.Println(value)
}
func maximumTripletValue(nums []int) int64 {
	maxVal := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}
	mtv, mv, n := int64(0), int64(0), len(nums)
	stack := make([]int, 0, n)
	for i, v := range nums {
		last := len(stack) - 1
		for last >= 0 && v >= nums[stack[last]] {
			last--
		}

		//fmt.Println(i, mv, v)
		//fmt.Println(stack)
		stack = stack[:last+1]
		stack = append(stack, i)
		if i > 1 {
			mtv = maxVal(mtv, mv*int64(v))
		}
		if last >= 0 {
			mv = maxVal(mv, int64(nums[stack[0]]-v))
		}
	}
	return mtv
}
