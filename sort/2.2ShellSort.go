package sort

// copyInsertionSort 插入排序
func copyInsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j, cur := i, arr[i]
		for ; j > 0 && arr[j-1] > cur; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = cur
	}
}

// ShellSort 希尔排序
func ShellSort(arr []int) {
	for t := len(arr) >> 1; t > 0; t >>= 1 {
		for i := t; i < len(arr); i++ {
			j, cur := i, arr[i]
			for ; j >= t && arr[j-t] < cur; j -= t {
				arr[j] = arr[j-t]
			}
			arr[j] = cur
		}
	}
}
