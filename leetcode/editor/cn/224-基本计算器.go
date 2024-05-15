package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	s := "1 + 1"
	s = "2-1 + 2" // 3
	//s = "(1+(4+5+2)-3)+(6+8)"   // 23
	//s = "1-(     -2)"           // 3
	//s = "- (3 - (- (4 + 5) ) )" // -12
	i := calculate(s)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func calculate_(s string) int {
	// 栈：迭代
	ans, n := 0, len(s)
	op := []bool{true} // op := []int{1}
	sign := true       // 标记括号范围内的符号
	for i := 0; i < n; i++ {
		switch s[i] {
		case '+':
			sign = op[len(op)-1] // 括号范围的最初标记
		case '-':
			sign = !op[len(op)-1] // 符号反转
		case '(':
			op = append(op, sign) // 入栈
		case ')':
			op = op[:len(op)-1] // 出栈
		case ' ':
		default:
			j := i
			for i++; i < n && unicode.IsDigit(rune(s[i])); {
				i++
			}
			v, _ := strconv.Atoi(s[j:i])
			if !sign {
				ans += ^v + 1
			} else {
				ans += v
			}
			i--
		}
	}
	return ans

	// 递归
	//n := len(s)
	//var dfs func(i int) (int, int)
	//dfs = func(i int) (int, int) {
	//	v, op := 0, true
	//out:
	//	for i < n {
	//		switch s[i] {
	//		case '(':
	//			val, j := dfs(i + 1)
	//			i = j
	//			if op {
	//				v += val
	//			} else {
	//				v -= val
	//			}
	//		case ')':
	//			i++
	//			break out
	//		case ' ':
	//			i++
	//		case '+':
	//			op = true
	//			i++
	//		case '-':
	//			op = false
	//			i++
	//		default:
	//			val := int(s[i] - '0')
	//			for i++; i < n && unicode.IsDigit(rune(s[i])); i++ {
	//				val = val*10 + int(s[i]-'0')
	//			}
	//			if op {
	//				v += val
	//			} else {
	//				v -= val
	//			}
	//		}
	//	}
	//	return v, i
	//}
	//v, _ := dfs(0)
	//return v
}
