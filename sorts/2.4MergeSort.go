package sorts

import (
	"math"
)

func MergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)>>1
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	//merge(arr, l, mid, r)
	mergeSentinel(arr, l, mid, r)
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

// mergeSentinel 请利用哨兵，简化编程（和 GPT 答案相同）
func mergeSentinel(arr []int, l, m, r int) {
	lSize, rSize := m-l+1, r-m
	left, right := make([]int, lSize+1), make([]int, rSize+1)
	copy(left[:lSize], arr[l:]) // copy(left[:lSize], arr[l:]) + 哨兵
	copy(right[:rSize], arr[m+1:])
	left[lSize], right[rSize] = math.MaxInt, math.MaxInt // 哨兵
	for i, j, k := 0, 0, l; k <= r; k++ {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
}
