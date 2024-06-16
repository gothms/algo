package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 7, 9}
	value := 9
	search := binarySearch(arr, len(arr), value)
	fmt.Println(search)
}
func binarySearch(arr []int, n int, value int) int {
	//l, r, m := 0, n-1, 0
	//for l <= r {
	//	m = int(uint(l+r) >> 1)
	//	if arr[m] == value {
	//		return m
	//	} else if arr[m] > value {
	//		r = m - 1
	//	} else {
	//		l = m + 1
	//	}
	//}
	//return -1

	var dfs func(int, int) int
	dfs = func(l, r int) int {
		if l > r {
			return -1
		}
		m := int(uint(l+r) >> 1)
		if arr[m] == value {
			return m
		} else if arr[m] > value {
			return dfs(l, m-1)
		} else {
			return dfs(m+1, r)
		}
	}
	return dfs(0, n-1)
}
