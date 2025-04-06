package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 1, 2, 3} // [2,3,4,1]
	odd := sortEvenOdd(nums)
	fmt.Println(odd)
}

// leetcode submit region begin(Prohibit modification and deletion)
func sortEvenOdd(nums []int) []int {
	n := len(nums)
	arr := [2]sort.IntSlice{}
	arr[0], arr[1] = make(sort.IntSlice, 0, (n+1)>>1), make(sort.IntSlice, 0, n>>1)
	for i, v := range nums {
		arr[i&1] = append(arr[i&1], v)
	}
	sort.Sort(arr[0])
	sort.Sort(sort.Reverse(arr[1]))
	nums = nums[:0]
	for i, v := range arr[1] {
		nums = append(nums, arr[0][i], v)
	}
	if n&1 > 0 {
		nums = append(nums, arr[0][len(arr[0])-1])
	}
	return nums
}

//func MergeSort(arr []int, l, r int) {
//	if l+1 >= r {
//		return
//	}
//	mid := l + (r-l)>>1
//	MergeSort(arr, l, mid)
//	MergeSort(arr, mid+1, r)
//	merge(arr, l, mid, r)
//}
//
//func merge(arr []int, l, m, r int) {
//	temp := make([]int, r-l+1)
//	j, k := m+1, 0
//	for i := l; i <= m; {
//		if (i^j)&1 == 1 { // 不同奇偶
//			if i+1 <= m {
//
//			}
//			if j+1 <= r {
//
//			}
//		}
//		//if j <= r {
//		//	switch i & 1 {
//		//	case 0: // 同偶，升序
//		//		if arr[j] < arr[i] {
//		//			temp[k], k = arr[j], k+1
//		//		}
//		//	case 1: // 同奇，降序
//		//		if arr[j] > arr[i] {
//		//			temp[k], k = arr[j], k+1
//		//		}
//		//	}
//		//}
//		if j <= r && (i&1 == 0 && arr[j] < arr[i] || i&1 == 1 && arr[j] > arr[i]) {
//			temp[k], k = arr[j], k+1
//			j++
//		} else {
//			temp[k], k = arr[i], k+1
//			i++
//		}
//		//for ; j <= r && arr[j] < arr[i]; j++ {
//		//	temp[k], k = arr[j], k+1
//		//}
//		//temp[k], k = arr[i], k+1
//	}
//	copy(arr[l:], temp[:k])
//}

//leetcode submit region end(Prohibit modification and deletion)
