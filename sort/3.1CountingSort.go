package sort

// CountingSort 非稳定排序
func CountingSort(arr []int, maxVal int) {
	bucket, n := make([]int, maxVal+1), len(arr)
	for i := 0; i < n; i++ {
		bucket[arr[i]]++
	}
	for i, k := 0, 0; i <= maxVal; i++ {
		for j := 0; j < bucket[i]; j++ {
			arr[k], k = i, k+1
		}
	}
}

// CountingSortGeekBang 稳定排序
func CountingSortGeekBang(arr []int, maxVal int) {
	bucket, n := make([]int, maxVal+1), len(arr)
	for i := 0; i < n; i++ {
		bucket[arr[i]]++
	}
	// 尝试 bucket[0]--，则省去下一个for循环 arr[i]]-1 的 -1 操作
	bucket[0]--
	for i := 1; i <= maxVal; i++ {
		bucket[i] += bucket[i-1]
	}
	cache := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		//cache[bucket[arr[i]]-1] = arr[i] // 从1开始，第 bucket[arr[i]] 小
		cache[bucket[arr[i]]] = arr[i] // bucket[0]--
		bucket[arr[i]]--
	}
	copy(arr, cache)
}
