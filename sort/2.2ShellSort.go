package sort

func ShellSort(arr []int) {
	for t := len(arr) >> 1; t > 0; t >>= 1 {
		for i := t; i < len(arr); i++ {
			j, curr := i, arr[i]
			for ; j >= t && curr < arr[j-t]; j -= t {
				arr[j] = arr[j-t]
			}
			arr[j] = curr
		}
	}
}
