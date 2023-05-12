package sort

func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	pivot := partition(arr, l, r)
	QuickSort(arr, l, pivot-1)
	QuickSort(arr, pivot+1, r)
}
func partition(arr []int, l, r int) int {
	//pivot, counter := l+rand.Intn(r-l+1), l+1	// 随机数
	//arr[l], arr[pivot] = arr[pivot], arr[l]
	//pivot = l
	pivot, counter := l, l+1
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[pivot] {
			arr[counter], arr[i] = arr[i], arr[counter]
			counter++
		}
	}
	counter--
	arr[counter], arr[pivot] = arr[pivot], arr[counter]
	return counter
}
