package main

import "fmt"

func main() {
	s := "1^0|0|1"
	result := 0
	s = "0&0&0&1^1|0"
	result = 1
	eval := countEval(s, result)
	fmt.Println(eval)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countEval(s string, result int) int {
	// dp

	// 记忆化搜索
	n := len(s)
	memo := make([][][2]int, n)
	for i := range memo {
		memo[i] = make([][2]int, n)
		if i&1 == 0 { // 偶数位为 0/1
			v := int(s[i] - '0')
			memo[i][i][v], memo[i][i][v^1] = 1, 0
		}
		for j := i + 2; j < n; j += 2 { // 默认为 -1
			memo[i][j][0], memo[i][j][1] = -1, -1
		}
	}
	var dfs func(int, int, int) int
	dfs = func(l, r, t int) int {
		v := &memo[l][r][t]
		if *v >= 0 {
			return *v
		}
		ret := 0
		for i := l + 1; i < r; i += 2 { // 奇数位为运算符
			switch t {
			case 0:
				l0, r0 := dfs(l, i-1, 0), dfs(i+1, r, 0)
				ret += l0 * r0
				switch s[i] {
				case '&': // 0 0 / 1 0 / 0 1
					l1, r1 := dfs(l, i-1, 1), dfs(i+1, r, 1)
					ret += l1*r0 + l0*r1
				case '|': // 0 0
				case '^': // 0 0 / 1 1
					l1, r1 := dfs(l, i-1, 1), dfs(i+1, r, 1)
					ret += l1 * r1
				}
			case 1:
				l1, r1 := dfs(l, i-1, 1), dfs(i+1, r, 1)
				switch s[i] {
				case '|': // 1 0 / 0 1 / 1 1
					ret += l1 * r1
					fallthrough
				case '^': // 1 0 / 0 1
					l0, r0 := dfs(l, i-1, 0), dfs(i+1, r, 0)
					ret += l1*r0 + l0*r1
				case '&': // 1 1
					ret += l1 * r1
				}
			}
		}
		*v = ret // 记忆化
		return ret
	}
	return dfs(0, n-1, result)
}

//leetcode submit region end(Prohibit modification and deletion)
