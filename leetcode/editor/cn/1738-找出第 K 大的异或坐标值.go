package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 6, 1, 8, 4, 2, 5, 9, 7}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)

	matrix := [][]int{{8, 10, 5, 8, 5, 7, 6, 0, 1, 4, 10, 6, 4, 3, 6, 8, 7, 9, 4, 2}}
	k := 2
	value := kthLargestValue(matrix, k)
	fmt.Println(value)
}

// leetcode submit region begin(Prohibit modification and deletion)
func kthLargestValue(matrix [][]int, k int) int {
	// heap
	//n := len(matrix[0])
	//arr := make(sort.IntSlice, 1, k)
	//dp, h := make([]int, n+1), &hp1738{arr}
	//for _, mt := range matrix {
	//	last := 0
	//	for i, v := range mt {
	//		dp[i+1], last = dp[i+1]^dp[i]^last^v, dp[i+1]
	//		if dp[i+1] > h.IntSlice[0] {
	//			if h.Len() == k {
	//				h.IntSlice[0] = dp[i+1]
	//				heap.Fix(h, 0)
	//			} else {
	//				heap.Push(h, dp[i+1])
	//			}
	//		}
	//	}
	//}
	//return h.IntSlice[0]

	// lc：快速选择算法
	m, n := len(matrix), len(matrix[0])
	val, dp := make([]int, 0, m*n), make([]int, n+1)
	for _, mt := range matrix {
		last := 0
		for i, v := range mt {
			dp[i+1], last = dp[i+1]^dp[i]^last^v, dp[i+1]
			val = append(val, dp[i+1])
		}
	}
	k = len(val) - k // 反过来找第 k 小
	//rand.Shuffle(len(val), func(i, j int) { val[i], val[j] = val[j], val[i] })	// 可以不要
	for l, r := 0, len(val)-1; l < r; {
		v := val[l]
		i, j := l, r+1
		for {
			for i++; i < r && val[i] < v; { // i++ 包含 val[i] == v
				i++
			}
			for j--; j > l && val[j] > v; {
				j--
			}
			if i >= j {
				break
			}
			val[i], val[j] = val[j], val[i]
		}
		val[l], val[j] = val[j], v
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}
	return val[k]
}

type hp1738 struct {
	sort.IntSlice
}

func (h *hp1738) Push(x any) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h hp1738) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
