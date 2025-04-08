package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func merge(A []int, m int, B []int, n int) {
	for i, j, k := m-1, n-1, m+n-1; j >= 0; k-- {
		if i >= 0 && A[i] > B[j] {
			A[k] = A[i]
			i--
		} else {
			A[k] = B[j]
			j--
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
