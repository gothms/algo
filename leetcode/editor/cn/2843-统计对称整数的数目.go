package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)

var symmetric []int

func init() {
	symmetric = make([]int, 0, 10000)
	for i := 1; i < 10; i++ {
		symmetric = append(symmetric, i*10+i)
	}
	for i := 1000; i < 10000; i++ {
		l, r := i/100, i%100
		if l/10+l%10 == r/10+r%10 {
			symmetric = append(symmetric, i)
		}
	}
}

func countSymmetricIntegers(low int, high int) int {
	return sort.SearchInts(symmetric, high+1) - sort.SearchInts(symmetric, low)
}

//leetcode submit region end(Prohibit modification and deletion)
