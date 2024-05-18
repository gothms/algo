package main

import "fmt"

func main() {
	nums := []int{73, 13, 20, 6}
	divisors := []int{56, 75, 83, 26, 24, 53, 56, 61}
	score := maxDivScore(nums, divisors)
	fmt.Println(score)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxDivScore(nums []int, divisors []int) int {
	memo := make(map[int]struct{})
	ans, maxV := divisors[0], 0
	for _, d := range divisors {
		if _, ok := memo[d]; ok {
			continue
		}
		memo[d] = struct{}{}
		cnt := 0
		for _, v := range nums {
			if v%d == 0 {
				cnt++
			}
		}
		if cnt > maxV || cnt == maxV && d < ans {
			ans, maxV = d, cnt
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
