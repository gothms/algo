package sorts

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		var bubble bool
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				bubble = true // 最好时间复杂度：O(n)
			}
		}
		if !bubble {
			break
		}
	}
}

// InsertionSort 插入排序
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j, cur := i-1, arr[i]
		for ; j >= 0 && arr[j] > cur; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = cur
	}
}

// SelectionSort 选择排序
func SelectionSort(arr []int) {
	n := len(arr) - 1
	for i, idx := 0, 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if arr[j] < arr[idx] {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i] // arr[idx] 值已更新
	}
}
