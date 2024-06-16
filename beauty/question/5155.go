package main

import "fmt"

func main() {
	arr := []int{3, 2, 5, 1, 6, 7, 1, 0, 3, 9, 8, 4}
	mergeSort(arr, len(arr))
	fmt.Println(arr)
}
func mergeSort(arr []int, n int) {
	var mSort func(int, int)
	mSort = func(l, r int) {
		if l >= r {
			return
		} else if l+1 == r {
			if arr[l] > arr[r] {
				arr[l], arr[r] = arr[r], arr[l]
			}
			return
		}
		m := (l + r) >> 1
		mSort(l, m)
		mSort(m+1, r)
		temp := make([]int, r-l+1)
		copy(temp, arr[l:r+1])
		for i, j, k := 0, m+1-l, l; i <= m-l; k++ {
			if j < len(temp) && temp[j] < temp[i] {
				arr[k] = temp[j]
				j++
			} else {
				arr[k] = temp[i]
				i++
			}
		}
	}
	mSort(0, n-1)
}
