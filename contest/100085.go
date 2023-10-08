package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println()
}
func minProcessingTime(processorTime []int, tasks []int) int {
	// 贪心
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	sort.Ints(processorTime)
	sort.Ints(tasks)
	mpt := 0
	for i, j := 0, len(tasks)-1; i < len(processorTime); i, j = i+1, j-4 {
		mpt = maxVal(mpt, processorTime[i]+tasks[j])
	}
	return mpt
}
