package sort

func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	pivot := partition(arr, l, r)
	QuickSort(arr, l, pivot-1)
	QuickSort(arr, pivot+1, r)
}

// partition 与桶排序 Bucket Sort
// 已知 arr[i] 的区间范围 [x,y]，n 的区间范围 [i,j]
// 则可根据两个区间范围，预设 k+1 个分区点 pivot，取值为 [x+d,x+d*1,x+d*2,...,x+d*k]，其中 x+d*k+k >= y
// 即将数据分为 k+1 份后，再分别进行排序（可选插入排序等稳定排序算法）
func partition(arr []int, l, r int) int {
	// 1.三数取中法
	// 2.随机法
	//pivot, counter := l+rand.Intn(r-l+1), l+1	// 随机数
	//arr[l], arr[pivot] = arr[pivot], arr[l]	// 或交换到 r
	//pivot = l
	pivot, counter := l, l+1
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[pivot] {
			arr[counter], arr[i] = arr[i], arr[counter]
			counter++
		}
	}
	counter-- // 否则可能 counter==r+1，导致无限循环
	arr[counter], arr[pivot] = arr[pivot], arr[counter]
	return counter
}
