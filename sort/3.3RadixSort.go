package sort

// RadixSort 稳定算法
func RadixSort(arr []int) {
	maxV, n := 0, len(arr)
	for i := 0; i < n; i++ {
		if arr[i] > maxV {
			maxV = arr[i]
		}
	}
	digitCount := 1
	for ; maxV != 0; maxV /= 10 {
		digitCount++
	}
	radix := [10][]int{}
	for i, div := 0, 1; i < digitCount; i, div = i+1, div*10 {
		for j := 0; j < n; j++ { // 根据每一位来排序，可以用桶排序或计数排序
			idx := arr[j] / div % 10 // 此处利用十进制为桶
			radix[idx] = append(radix[idx], arr[j])
		}
		for j, k := 0, 0; j < 10; j++ {
			copy(arr[k:], radix[j])
			radix[j], k = radix[j][len(radix[j]):], k+len(radix[j])
		}
	}
}
