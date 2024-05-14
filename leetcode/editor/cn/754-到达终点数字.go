package main

import (
	"fmt"
	"math"
)

func main() {
	target := 2 // 3
	target = 3  // 2
	target = 4  // 3
	//target = -1000000000 // 44723
	//target = 8           // 4
	number := reachNumber(target)
	fmt.Println(number)
}

// leetcode submit region begin(Prohibit modification and deletion)
func reachNumber(target int) int {
	// lc
	if target < 0 {
		target = -target
	}
	n := int(math.Ceil((-1 + math.Sqrt(float64(8*target+1))) / 2))
	if (n*(n+1)/2-target)%2 == 0 {
		return n
	}
	return n + 1 + n%2

	//if target < 0 {
	//	target = -target
	//}
	//i := int(math.Sqrt(float64(target << 1)))
	//d := target - i*(i+1)>>1
	//if d == 0 { // 刚好
	//	return i
	//} else if d > 0 { // 超过 target，方便计算
	//	i++
	//	d -= i
	//}
	//if d&1 == 0 { // 超过了偶数，则第 -d/2 步往回走
	//	return i
	//}
	//if i&1 == 0 { // 当前步数是偶数，则下步是奇数。再往前一步，d 变成偶数
	//	return i + 1
	//}
	//return i + 2 // 当前步数是奇数，则下步是偶数。再往前一步，d 仍然是奇数，继续再走一步，d 变成偶数

	// 超时
	//if target < 0 {
	//	target = -target
	//}
	//memo := make(map[int]bool)
	//memo[0] = true
	//var dfs func(int, int) int
	//dfs = func(x, i int) int {
	//	if x == target {
	//		return i
	//	}
	//	i++
	//	l, r := math.MaxInt32, math.MaxInt32
	//	if !memo[x+i] && x+i <= target {
	//		memo[x+i] = true
	//		l = dfs(x+i, i)
	//	}
	//	if !memo[x-i] && x-i > -target {
	//		memo[x-i] = true
	//		r = dfs(x-i, i)
	//	}
	//	return min(l, r)
	//}
	//return dfs(0, 0)
}

//leetcode submit region end(Prohibit modification and deletion)
