package main

import "fmt"

func main() {
	nums := []int{2, 5, 1, 4}
	beautifulPairs := countBeautifulPairs(nums)
	fmt.Println(beautifulPairs)
}

// leetcode submit region begin(Prohibit modification and deletion)
var primePair2748 map[[2]int]struct{}

func init() {
	gcd := func(a, b int) bool {
		for b != 0 {
			a, b = b, a%b
		}
		return a == 1
	}
	primePair2748 = make(map[[2]int]struct{})
	for i := 1; i <= 9; i++ {
		for j := i; j <= 9; j++ {
			if gcd(i, j) {
				primePair2748[[2]int{i, j}] = struct{}{}
				primePair2748[[2]int{j, i}] = struct{}{}
			}
		}
	}
}
func countBeautifulPairs(nums []int) int {
	memo := make(map[int]int) // 可换为 [10]int
	ans := 0
	for _, v := range nums {
		t := v % 10
		for v > 9 {
			v /= 10
		}
		for k, c := range memo {
			if _, ok := primePair2748[[2]int{t, k}]; ok {
				ans += c
			}
		}
		memo[v]++
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
