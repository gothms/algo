package main

import "unicode"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func diffWaysToCompute(expression string) []int {
	// dp

	// 个人
	n := len(expression)
	value, operation := make([]int, 0), make([]uint8, 0)
	for i := 0; i < n; {
		c := expression[i]
		switch c {
		case '+', '-', '*':
			operation = append(operation, c)
			i++
		default:
			v, j := int(c-'0'), i+1
			for ; j < n && unicode.IsDigit(rune(expression[j])); j++ {
				v = v*10 + int(expression[j]-'0')
			}
			value = append(value, v)
			i = j
		}
	}
	m := len(value)
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, m)
		memo[i][i] = []int{value[i]}
	}
	var dfs func(int, int) []int
	dfs = func(l, r int) []int {
		if vs := memo[l][r]; len(vs) > 0 {
			return vs
		}
		vs := make([]int, 0)
		for i := l; i < r; i++ {
			lv, rv := dfs(l, i), dfs(i+1, r)
			for _, x := range lv {
				for _, y := range rv {
					op := operation[i]
					switch op {
					case '+':
						vs = append(vs, x+y)
					case '-':
						vs = append(vs, x-y)
					case '*':
						vs = append(vs, x*y)
					}
				}
			}
		}
		memo[l][r] = vs
		return vs
	}
	return dfs(0, m-1)
}

//leetcode submit region end(Prohibit modification and deletion)
