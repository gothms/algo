package stack

import (
	"math"
	"unicode"
)

// StackOperation 表达式求值，模拟 + - * /（没括号）
// 规则：
// 如果比运算符栈顶元素的优先级高，就将当前运算符压入栈
// 如果比运算符栈顶元素的优先级低或者相同，从运算符栈中取栈顶运算符，从操作数栈的栈顶取 2 个操作数，然后进行计算
// 再把计算完的结果压入操作数栈，继续比较
func StackOperation(s string) int {
	stVal, stOpt := make([]int, 0), make([]uint8, 0)
	n := len(s)
	op := func(add bool) {
		i := len(stOpt) - 1
		if i < 0 { // fast fail
			return
		}
		v := stVal[len(stVal)-1] // 栈顶第一个元素
		stVal = stVal[:len(stVal)-1]
	out:
		for ; i >= 0; i-- {
			o, cur := stOpt[i], stVal[len(stVal)-1] // 运算符 & 栈顶第二个元素
			switch o {                              // + - * /
			case '+':
				if !add {
					break out
				}
				v += cur
			case '-':
				if !add {
					break out
				}
				v = cur - v
			case '*':
				v *= cur
			case '/':
				v = cur / v
			}
			stVal = stVal[:len(stVal)-1]
		}
		stOpt = stOpt[:i+1]      // 移除已计算的运算符
		stVal = append(stVal, v) // 把计算结果压入栈
	}
	for i := 0; i < n; {
		c := s[i]
		switch c {
		case '+', '-': // 优先级最低，将计算所有运算符
			op(true)
			stOpt = append(stOpt, c)
			i++
		case '*', '/':
			if last := len(stOpt) - 1; last >= 0 && (stOpt[last] == '*' || stOpt[last] == '/') {
				op(false) // 优先级相同，才计算
			}
			stOpt = append(stOpt, c)
			i++
		case ' ':
			i++
		default:
			v := int(s[i] - '0')
			for i++; i < n && unicode.IsDigit(rune(s[i])); i++ {
				v = v*10 + int(s[i]-'0')
			}
			//v, _ := strconv.Atoi(s[f:i]) // 取得操作数
			stVal = append(stVal, v)
		}
	}
	op(true) // 收尾
	if len(stVal) == 1 {
		return stVal[0]
	}
	return math.MinInt // error
}
