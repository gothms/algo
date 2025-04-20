package main

import "fmt"

func main() {
	num := 58
	k := 9 // 2
	numbers := minimumNumbers(num, k)
	fmt.Println(numbers)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumNumbers(num int, k int) int {
	// 个人
	if num == 0 {
		return 0
	}
	//if k&1 == 0 && num&1 == 1 {
	//	return -1
	//}
	for i, v, t := 1, k, num%10; i <= 10 && v <= num; i++ {
		if v%10 == t {
			return i
		}
		v += k
	}
	return -1

	// 记忆化搜索：枚举多重集的最小值
	// Error：读错题意，“返回该多重集的最小大小”。而此方法求的是有多少种不重复的多重集
	//if num < k || k&1 == 0 && num&1 == 1 {
	//	return -1
	//}
	//memo := make([]int, num+1)
	//memo[k] = 1
	//for i := k + 1; i <= num; i++ {
	//	memo[i] = -1
	//}
	//var dfs func(int, int) int
	//dfs = func(i, f int) int {
	//	v := &memo[i]
	//	if *v >= 0 {
	//		return *v
	//	}
	//	ret := 0
	//	if i%10 == k {
	//		ret = 1
	//	}
	//	for j := f; j <= i>>1; j += 10 { // 枚举最小值（最多枚举 i>>1），增量为 10
	//		ret += dfs(i-j, j)
	//	}
	//	*v = ret
	//	return ret
	//}
	//return dfs(num, k)
}

//leetcode submit region end(Prohibit modification and deletion)
