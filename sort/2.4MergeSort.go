package sort

func MergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)>>1
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, m, r int) {
	temp := make([]int, r-l+1)
	j, k := m+1, 0
	for i := l; i <= m; i++ {
		for ; j <= r && arr[j] < arr[i]; j++ {
			temp[k], k = arr[j], k+1
		}
		temp[k], k = arr[i], k+1
	}
	copy(arr[l:], temp[:k])
}
