package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 10
	n = 2  // 0
	n = 12 // 5
	i := countPrimes(n)
	fmt.Println(i)
}

//leetcode submit region begin(Prohibit modification and deletion)

const n204 = 5_000_001

var prime204 []int

func init() {
	prime := make([]bool, n204)
	for i := 2; i < n204; i++ {
		if prime[i] {
			continue
		}
		prime204 = append(prime204, i)
		for j := i << 1; j < n204; j += i {
			prime[j] = true
		}
	}
	//fmt.Println(prime204)
}
func countPrimes(n int) int {
	return sort.Search(len(prime204), func(i int) bool {
		return prime204[i] >= n
	})
}

//leetcode submit region end(Prohibit modification and deletion)
