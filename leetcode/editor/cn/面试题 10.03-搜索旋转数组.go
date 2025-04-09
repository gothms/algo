package main

import "fmt"

func main() {
	arr := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	target := 5 // 8
	arr = []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	target = 11 // -1
	//arr = []int{1, 1, 1, 1, 1, 2, 1, 1, 1}
	//target = 2	// 5
	i := search(arr, target)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func search(arr []int, target int) int {
	// 区分 81
	l, r := 0, len(arr)-1
	for l < r {
		m := (l + r) >> 1
		if arr[l] == arr[r] {
			r--
		} else if arr[l] <= arr[m] { // m 在左半部分
			if arr[l] <= target && arr[m] >= target {
				r = m
			} else {
				l = m + 1
			}
		} else {
			if arr[l] > target && arr[m] < target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	if arr[l] == target {
		return l
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
