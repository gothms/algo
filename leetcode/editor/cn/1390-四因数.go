package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 45
	divisors := sumFourDivisors(nums)
	fmt.Println(divisors)
}

// leetcode submit region begin(Prohibit modification and deletion)
var prime1309 []int // 质数集

func init() {
	prime1309 = make([]int, 0, 9592)
	n := int(1e5 + 1)
	temp := make([]bool, n)
	for i := 2; i < n; i++ {
		if !temp[i] {
			prime1309 = append(prime1309, i)
			for j := i << 1; j < n; j += i {
				temp[j] = true
			}
		}
	}
}
func sumFourDivisors(nums []int) int {
	ans := 0
	for _, v := range nums {
		m := int(math.Pow(float64(v), 0.5)) // 开平方
		if m*m == v {                       // 如 4：1 2 4
			continue
		}
		for i := 0; prime1309[i] <= m; i++ {
			p := prime1309[i]
			if v%p != 0 {
				continue
			}
			if w := v / p; prime1309[sort.SearchInts(prime1309, w)] == w || // 另一个因子也是质数
				p*p == w { // p^3 = v，如 8：1 2 4 8
				ans += 1 + v + p + w
				break
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
