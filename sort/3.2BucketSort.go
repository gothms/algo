package sort

import "sort"

func BucketSort(arr []int, size int) {
	min, max, n := arr[0], arr[0], len(arr)
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		} else if arr[i] > max {
			max = arr[i]
		}
	}
	bucketCount := (max-min+1)/size + 1
	buckets := make([][]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		buckets[i] = make([]int, 0)
	}
	for i := 0; i < n; i++ {
		idx := (arr[i] - min) / size
		buckets[idx] = append(buckets[idx], arr[i])
	}
	for i, j := 0, 0; i < bucketCount; i, j = i+1, j+len(buckets[i]) {
		sort.Ints(buckets[i])
		copy(arr[j:], buckets[i])
	}
}
