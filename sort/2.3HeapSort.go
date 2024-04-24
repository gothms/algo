package sort

func HeapSort(arr []int) {
	n := len(arr)
	for i := n>>1 - 1; i >= 0; i-- {
		heapify(arr, i, n)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i) // i=1 时，无效 heapify
	}
}

func heapify(arr []int, i, n int) {
	for idx := i<<1 + 1; idx < n; idx, i = idx<<1+1, idx {
		if idx+1 < n && arr[idx+1] > arr[idx] {
			idx++
		}
		if arr[i] >= arr[idx] {
			break
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}
