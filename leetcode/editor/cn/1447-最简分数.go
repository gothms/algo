package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 3
	fractions := simplifiedFractions(n)
	fmt.Println(fractions)
}

//leetcode submit region begin(Prohibit modification and deletion)

var factor1447 []map[int]struct{}

func init() {
	const n = 101
	factor1447 = make([]map[int]struct{}, n)
	prime := make(map[int][]int)
	for i := 2; i < n; i++ {
		temp := make([]int, 0, n/i)
		for j := i; j < n; j += i {
			temp = append(temp, j) // 有公因数的数，都在同一个集合
		}
		prime[i] = temp
	}
	for _, p := range prime {
		for i, v := range p {
			if factor1447[v] == nil {
				factor1447[v] = make(map[int]struct{})
			}
			for _, val := range p[:i] {
				factor1447[v][val] = struct{}{}
			}
		}
	}
	//for _, f := range factor1447 {
	//	fmt.Println(f)
	//}
}

func simplifiedFractions(n int) []string {
	//gcd := func(a, b int) bool {
	//	for b != 0 {
	//		a, b = b, a%b
	//	}
	//	return a == 1
	//}
	//ans := make([]string, 0)
	//for i := 2; i <= n; i++ {
	//	ans = append(ans, fmt.Sprintf("%d/%d", 1, i))
	//	for j := 2; j < i; j++ {
	//		if gcd(i, j) {
	//			ans = append(ans, fmt.Sprintf("%d/%d", j, i))
	//		}
	//	}
	//}
	//return ans

	// 个人
	ans := make([]string, 0)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if _, ok := factor1447[i][j]; !ok {
				ans = append(ans, strconv.Itoa(j)+"/"+strconv.Itoa(i))
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
