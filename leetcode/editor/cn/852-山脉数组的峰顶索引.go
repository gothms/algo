package main

import "fmt"

func main() {
	arr := []int{0, 1, 0}
	arr = []int{3, 4, 5, 1}
	array := peakIndexInMountainArray(arr)
	fmt.Println(array)
}

// leetcode submit region begin(Prohibit modification and deletion)
func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := (l + r) >> 1
		if arr[m] < arr[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

//leetcode submit region end(Prohibit modification and deletion)
