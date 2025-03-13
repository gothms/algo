package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	largest := findKthLargest(nums, k)
	fmt.Println(largest)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	return heapSortK(nums, k)
}
func heapSortK(arr []int, k int) int {
	n := len(arr)
	for i := n>>1 - 1; i >= 0; i-- {
		heapify(arr, i, n)
	}
	for i := n - 1; i > n-k; i-- { // 第一大的数，则 k =1
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, 0, i)
	}
	return arr[0]
}
func heapify(arr []int, i, n int) {
	for idx := i<<1 + 1; idx < n; i, idx = idx, idx<<1+1 {
		if idx+1 < n && arr[idx] < arr[idx+1] {
			idx++
		}
		if arr[i] >= arr[idx] {
			break
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
