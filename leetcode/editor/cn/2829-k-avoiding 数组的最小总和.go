package main

import "fmt"

func main() {
	n, k := 5, 4
	//n, k = 2, 6
	i := minimumSum(n, k)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumSum(n int, k int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func minimumSum_(n int, k int) int {
	m := (k - 1) >> 1
	if k-m > n {
		return (n + 1) * n >> 1
	}
	return (n+1)*n>>1 + (n-k+m+1)*m
}
