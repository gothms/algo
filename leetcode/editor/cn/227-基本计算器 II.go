package main

import (
	"fmt"
)

func main() {
	s := "3+2*2"
	s = " 3/2 "
	s = " 3+5 / 2 "
	s = " (3+5) / 2 "
	s = " 42 "
	s = "1+1+1"
	i := calculate(s)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) int {
	// 栈

	// 递归
	//n := len(s)
	//operation := func(vs *[]int, op *[]uint8, all bool) {
	//	if len(*op) == 0 {
	//		return
	//	}
	//	v := (*vs)[len(*vs)-1]
	//	*vs = (*vs)[:len(*vs)-1]
	//out:
	//	for len(*op) > 0 {
	//		switch (*op)[len(*op)-1] {
	//		case '+':
	//			if !all {
	//				break out
	//			}
	//			v += (*vs)[len(*vs)-1]
	//		case '-':
	//			if !all {
	//				break out
	//			}
	//			v = (*vs)[len(*vs)-1] - v
	//		case '*':
	//			v *= (*vs)[len(*vs)-1]
	//		case '/':
	//			v = (*vs)[len(*vs)-1] / v
	//		}
	//		*vs, *op = (*vs)[:len(*vs)-1], (*op)[:len(*op)-1]
	//	}
	//	*vs = append(*vs, v)
	//}
	//var dfs func(int) (int, int)
	//dfs = func(i int) (int, int) {
	//	stVs, stOp := make([]int, 0), make([]uint8, 0)
	//out:
	//	for i < n {
	//		switch s[i] {
	//		case '+', '-':
	//			operation(&stVs, &stOp, true)
	//			stOp = append(stOp, s[i])
	//			i++
	//		case '*', '/':
	//			if len(stOp) > 0 && (stOp[len(stOp)-1] == '*' || stOp[len(stOp)-1] == '/') {
	//				operation(&stVs, &stOp, false)
	//			}
	//			stOp = append(stOp, s[i])
	//			i++
	//		case ' ':
	//			i++
	//		case '(':
	//			j, v := dfs(i + 1)
	//			i = j
	//			stVs = append(stVs, v)
	//		case ')':
	//			i++
	//			break out
	//		default:
	//			v := int(s[i] - '0')
	//			for i++; i < n && unicode.IsDigit(rune(s[i])); i++ {
	//				v = v*10 + int(s[i]-'0')
	//			}
	//			stVs = append(stVs, v)
	//		}
	//	}
	//	operation(&stVs, &stOp, true)
	//	return i, stVs[0]
	//}
	//_, ans := dfs(0)
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
