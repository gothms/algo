package main

import (
	"fmt"
)

func main() {
	n := 13       // 5
	n = 100000000 // 31
	replacement := integerReplacement(n)
	fmt.Println(replacement)

	//fmt.Println(bits.TrailingZeros(12))
}

// leetcode submit region begin(Prohibit modification and deletion)
func integerReplacement(n int) int {
	// lc
	ans := 0
	for n != 1 {
		if n&1 == 0 {
			n >>= 1
			ans++
		} else if n&3 == 1 { // 末尾二进制 ...01
			n >>= 2
			ans += 3
		} else if n == 3 { // 特殊处理
			ans += 2
			break
		} else { // 末尾二进制 ...11
			n++
			ans++
		}
	}
	return ans

	// 个人
	//ans := 0
	//for n != 1 {
	//	if n&1 == 0 {
	//		c := bits.TrailingZeros(uint(n))
	//		ans += c
	//		n >>= c
	//	} else if n == 3 { // 重要
	//		ans += 2
	//		break
	//	} else {
	//		x, y := bits.OnesCount(uint(n+1)), bits.OnesCount(uint(n-1))	// 直觉：但严格证明呢
	//		if x <= y {
	//			n++
	//		} else {
	//			n--
	//		}
	//		ans++
	//	}
	//}
	//return ans

	// 个人：记忆化搜索
	//memo := map[int]int{0: -1, 1: 0, 2: 1}
	//var dfs func(int) int
	//dfs = func(i int) int {
	//	if v, ok := memo[i]; ok {
	//		return v
	//	}
	//	ret := bits.TrailingZeros(uint(i))
	//	i >>= ret
	//	ret += min(dfs(i-1), dfs(i+1)) + 1
	//	memo[i] = ret
	//	return ret
	//}
	//return dfs(n)
}

//leetcode submit region end(Prohibit modification and deletion)
