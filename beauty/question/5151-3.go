package main

import "fmt"

func main() {
	arr := []int{3, 2, 5, 1, 6}
	//bubbleSort(arr, len(arr))
	//insertSort(arr, len(arr))
	selectSort(arr, len(arr))
	fmt.Println(arr)
}
func bubbleSort(arr []int, n int) {
	for i := n - 1; i > 0; i-- {
		var bubble bool
		for j := 1; j <= i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				bubble = true
			}
		}
		if !bubble {
			break
		}
	}
}
func insertSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		v, j := arr[i], i-1
		for ; j >= 0 && arr[j] > v; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = v
	}
}
func selectSort(arr []int, n int) {
	for i := 0; i < n; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idx] {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}
