package main

import "fmt"

func main() {
	base := []int{2, 3}
	topping := []int{4, 5, 100}
	target := 18
	base = []int{10}
	topping = []int{1}
	target = 1
	cost := closestCost(base, topping, target)
	fmt.Println(cost)
}

// leetcode submit region begin(Prohibit modification and deletion)
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	// lc
	x := baseCosts[0]
	for _, c := range baseCosts {
		x = min(x, c)
	}
	if x > target {
		return x
	}
	can := make([]bool, target+1)
	ans := 2*target - x
	for _, c := range baseCosts {
		if c <= target {
			can[c] = true
		} else {
			ans = min(ans, c) // 超过 target 的最小值
		}
	}
	for _, c := range toppingCosts {
		for count := 0; count < 2; count++ { // 0-2 份
			for i := target; i > 0; i-- {
				if can[i] && i+c > target { // 超过 target 的最小值
					ans = min(ans, i+c)
				}
				if i-c > 0 {
					can[i] = can[i] || can[i-c] // 递推
				}
			}
		}
	}
	for i := 0; i <= ans-target; i++ { // 小于 target 的最大值
		if can[target-i] {
			return target - i
		}
	}
	return ans

	// 个人
	//abs := func(x int) int {
	//	if x < 0 {
	//		return -x
	//	}
	//	return x
	//}
	//ans, d := baseCosts[0], abs(target-baseCosts[0])
	//check := func(v int) {
	//	if cur := abs(target - v); cur < d || cur == d && v < ans {
	//		ans, d = v, cur
	//	}
	//}
	//for _, v := range baseCosts {
	//	memo := make(map[int]struct{})
	//	memo[v] = struct{}{}
	//	check(v)
	//	for _, t := range toppingCosts {
	//		temp := make(map[int]struct{})
	//		for k := range memo {
	//			temp[k+t] = struct{}{}
	//			temp[k+t<<1] = struct{}{}
	//			check(k + t)
	//			check(k + t<<1)
	//		}
	//		for k := range temp {
	//			memo[k] = struct{}{}
	//		}
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
