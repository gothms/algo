package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	s := "1 + 1"
	s = "2-1 + 2"
	s = "(1+(4+5+2)-3)+(6+8)"   // 23
	s = "1-(     -2)"           // 3
	s = "- (3 - (- (4 + 5) ) )" // -12
	i := calculate(s)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) int {
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

	//vs, op := make([]int, 0), make([]uint8, 0)
	//n := len(s)
	//val := func(i int) (int, int) {
	//	j := i
	//	for i < n && unicode.IsDigit(rune(s[i])) {
	//		i++
	//	}
	//	v, _ := strconv.Atoi(s[j:i])
	//	return i, v
	//}
	//operation := func(b int) {
	//out:
	//	for len(op) > 0 {
	//		o := op[len(op)-1]
	//		switch o {
	//		case '+':
	//			a := vs[len(vs)-1]
	//			b += a
	//			vs = vs[:len(vs)-1]
	//		case '-':
	//			if len(vs) > 0 {
	//				a := vs[len(vs)-1]
	//				b = a - b
	//				vs = vs[:len(vs)-1]
	//			} else {
	//				b = -b
	//			}
	//		case '(':
	//			break out
	//		}
	//		op = op[:len(op)-1]
	//	}
	//	vs = append(vs, b)
	//}
	//for i, last := 0, false; i < n; i++ {
	//	//if i == 9 {
	//	//	fmt.Println()
	//	//}
	//	switch s[i] {
	//	case '+':
	//		op = append(op, s[i])
	//		last = true
	//	case '-':
	//		if last {
	//			j, v := val(i + 1)
	//			i, v = j-1, -v
	//			operation(v)
	//			last = false
	//		} else {
	//			op = append(op, s[i])
	//			last = true
	//		}
	//	case '(':
	//		op = append(op, s[i])
	//		last = true
	//	case ')':
	//		v := vs[len(vs)-1]
	//		op = op[:len(op)-1]
	//		vs = vs[:len(vs)-1]
	//		operation(v)
	//		last = false
	//	case ' ':
	//	default:
	//		j, v := val(i)
	//		i = j - 1
	//		operation(v)
	//		last = false
	//	}
	//	fmt.Println(i, vs)
	//}
	////fmt.Println(vs)
	//return vs[0]
}

//leetcode submit region end(Prohibit modification and deletion)
