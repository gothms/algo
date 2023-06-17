package sort

func HeapSort(arr []int) {
	n := len(arr)
	for i := n>>1 - 1; i >= 0; i-- {
		heapify(arr, i, n)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}
func heapify(arr []int, i, n int) {
	for di := i<<1 + 1; di < n; di, i = di<<1+1, di {
		if di+1 < n && arr[di+1] > arr[di] {
			di++
		}
		if arr[i] >= arr[di] {
			break
		}
		arr[i], arr[di] = arr[di], arr[i]
	}
}
