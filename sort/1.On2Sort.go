package sort

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// InsertionSort 插入排序
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j, curr := i, arr[i]
		for ; j > 0 && arr[j-1] > curr; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = curr
	}
}

// SelectionSort 选择排序
func SelectionSort(arr []int) {
	n := len(arr) - 1
	for i, mi := 0, 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if arr[j] < arr[mi] {
				mi = j
			}
		}
		arr[i], arr[mi] = arr[mi], arr[i]
	}
}
