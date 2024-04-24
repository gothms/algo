package sort

import "sort"

func BucketSort(arr []int, size int) {
	minV, maxV, n := arr[0], arr[0], len(arr)
	for i := 1; i < n; i++ {
		if arr[i] < minV {
			minV = arr[i]
		} else if arr[i] > maxV {
			maxV = arr[i]
		}
	}
	bucketCount := (maxV-minV+1)/size + 1
	buckets := make([][]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		buckets[i] = make([]int, 0)
	}
	for i := 0; i < n; i++ {
		idx := (arr[i] - minV) / size
		buckets[idx] = append(buckets[idx], arr[i])
	}
	for i, j := 0, 0; i < bucketCount; i, j = i+1, j+len(buckets[i]) {
		sort.Ints(buckets[i]) // 采用稳定算法，则桶排序结果也稳定
		copy(arr[j:], buckets[i])
	}
}
