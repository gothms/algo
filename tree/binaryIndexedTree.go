package tree

// BIT 构建 树状数组
func BIT(arr []int) []int {
	n := len(arr)
	b := make([]int, n)
	for i, j := 0, 0; i < n; i++ {
		b[i] += arr[i]
		if j = i | (i + 1); j < n {
			b[j] += b[i]
		}
	}
	return b
}

// Update 源数据的 arr[i] 值更新为 delta
func Update(b []int, i, delta int) {
	for ; i < len(b); i |= i + 1 {
		b[i] += delta
	}
}

// PrefixSum [0,i]区间和
func PrefixSum(b []int, i int) int {
	sum := 0
	for ; i >= 0; i -= ^i & (i + 1) {
		sum += b[i]
	}
	return sum
}

// RangeSum [f,t]区间和
func RangeSum(b []int, f, t int) int {
	return PrefixSum(b, t) - PrefixSum(b, f-1)
}
