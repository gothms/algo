package tree

/*
lc
	493
*/
// ====================树状数组，从 0 开始====================

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

// ====================树状数组，从 1 开始====================

func BITOne(arr []int) []int {
	n := len(arr)
	b := make([]int, n+1)
	for i, j := 1, 0; i <= n; i++ {
		b[i] += arr[i-1]
		if j = i + i&-i; j <= n {
			b[j] += b[i]
		}
	}
	return b
}

func UpdateOne(b []int, i, delta int) {
	for ; i < len(b); i += i & -i {
		b[i] += delta
	}
}
func PrefixSumOne(b []int, i int) int {
	sum := 0
	for ; i >= 0; i &= i - 1 {
		sum += b[i]
	}
	return sum
}
func RangeSumOne(b []int, f, t int) int {
	return PrefixSumOne(b, t) - PrefixSumOne(b, f-1)
}
