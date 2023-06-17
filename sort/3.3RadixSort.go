package sort

func RadixSort(arr []int) {
	max, n := 0, len(arr)
	for i := 0; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	digitCount := 1
	for ; max != 0; max /= 10 {
		digitCount++
	}
	radix := [10][]int{}
	for i, div := 0, 1; i < digitCount; i, div = i+1, div*10 {
		for j := 0; j < n; j++ {
			idx := arr[j] / div % 10
			radix[idx] = append(radix[idx], arr[j])
		}
		for j, k := 0, 0; j < 10; j++ {
			copy(arr[k:], radix[j])
			radix[j], k = radix[j][len(radix[j]):], k+len(radix[j])
		}
	}
}
