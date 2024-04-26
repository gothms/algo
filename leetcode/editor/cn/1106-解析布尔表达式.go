//布尔表达式 是计算结果不是 true 就是 false 的表达式。有效的表达式需遵循以下约定：
//
//
// 't'，运算结果为 true
// 'f'，运算结果为 false
// '!(subExpr)'，运算过程为对内部表达式 subExpr 进行 逻辑非（NOT）运算
// '&(subExpr1, subExpr2, ..., subExprn)'，运算过程为对 2 个或以上内部表达式 subExpr1, subExpr2,
// ..., subExprn 进行 逻辑与（AND）运算
// '|(subExpr1, subExpr2, ..., subExprn)'，运算过程为对 2 个或以上内部表达式 subExpr1, subExpr2,
// ..., subExprn 进行 逻辑或（OR）运算
//
//
// 给你一个以字符串形式表述的 布尔表达式 expression，返回该式的运算结果。
//
// 题目测试用例所给出的表达式均为有效的布尔表达式，遵循上述约定。
//
//
//
// 示例 1：
//
//
//输入：expression = "&(|(f))"
//输出：false
//解释：
//首先，计算 |(f) --> f ，表达式变为 "&(f)" 。
//接着，计算 &(f) --> f ，表达式变为 "f" 。
//最后，返回 false 。
//
//
// 示例 2：
//
//
//输入：expression = "|(f,f,f,t)"
//输出：true
//解释：计算 (false OR false OR false OR true) ，结果为 true 。
//
//
// 示例 3：
//
//
//输入：expression = "!(&(f,t))"
//输出：true
//解释：
//首先，计算 &(f,t) --> (false AND true) --> false --> f ，表达式变为 "!(f)" 。
//接着，计算 !(f) --> NOT false --> true ，返回 true 。
//
//
//
//
// 提示：
//
//
// 1 <= expression.length <= 2 * 10⁴
// expression[i] 为 '('、')'、'&'、'|'、'!'、't'、'f' 和 ',' 之一
//
//
// Related Topics 栈 递归 字符串 👍 205 👎 0

package main

import "fmt"

func main() {
	expression := "&(|(f))"
	//expression = "|(f,f,f,t)"
	//expression = "!(&(f,t))"
	//expression = "!(f)"
	//expression = "!(&(&(f),&(!(t),&(f),|(f)),&(!(&(f)),&(t),|(f,f,t))))" // true，len=53
	expr := parseBoolExpr(expression)
	fmt.Println(expr)

	//fmt.Println('t', 'f') // 116 102
}

/*
思路
	基本思路仍然是采用双栈思路，简化点：
	1.stAll 栈存储所有的运算符 ! & | 和所有的运算值 t f，stOpt 栈存储所有的运算符在 stAll 栈中的位置（即索引，以便高效的定位）
	2.运算符操作的简化：
		!：只可能运算一个值，所以当遇到此运算符时，直接计算结果
		&：当出现第一个 f 时，结果肯定为 f，而不用考虑该运算符有效的范围内（即它所运算的括号范围）后续的值和运算
		|：同 &，当出现第一个 t 时，结果肯定为 t
	3.为了达成第 2. 的简化，用 cnt 记录左括号 '(' 出现的次数，以跳过无用的值和运算
	4.因为已经达成了第 2. 的简化，所以当出现右括号 ')' 时：
		&：运算结果肯定是 t，因为如果出现 f，已经被提前计算了结果，所以此时 & 运算的值必定全部是 t
		|：同 &，运算结果肯定是 f
*/
// leetcode submit region begin(Prohibit modification and deletion)
func parseBoolExpr(expression string) bool {
	// 栈：简化
	cnt, n := 0, len(expression)
	if n == 1 { // fast path
		return expression[0] == 't'
	}
	stAll, stOpt := []uint8{0}, []int{0} // 哨兵
	check := func(i int, v uint8) int {
		idx := stOpt[len(stOpt)-1]
		// 可再优化：分别考虑 v==t/f，如 v==t 时只有 last==|/!，才会继续 for 循环
		for last := stAll[idx]; last == '&' && v == 'f' || // &：一个 false，结果 false
			last == '|' && v == 't' || // |：一个 true，结果 true
			last == '!'; last = stAll[idx] { // !：只会操作一个值
			switch last {
			case '!':
				if v == 't' { // 取反
					v = 'f'
				} else {
					v = 't'
				}
				i++ // 跳过 ')'
				cnt--
			case '&', '|': // “结果已知”，即为 v
				for c := cnt - 1; cnt != c; { // 跳过所有“无用”的操作和值
					i++
					switch expression[i] {
					case '(':
						cnt++
					case ')':
						cnt--
					}
				}
			}
			stAll = stAll[:idx] // 结束本次运算后，处理 栈
			stOpt = stOpt[:len(stOpt)-1]
			idx = stOpt[len(stOpt)-1]
		}
		stAll = append(stAll, v)
		return i
	}
	for i := 0; i < n; i++ {
		v := expression[i]
		switch v {
		case '&', '|', '!':
			stOpt = append(stOpt, len(stAll)) // 操作符在 stAll 栈中对应的位置，方便快捷定位
			stAll = append(stAll, v)          // 操作符 + t, f 值
		case 't', 'f': // 检查是否“结果已知”
			i = check(i, v) // i 值更新
		case ')': // 先计算一次结果，再检查是否“结果已知”
			cnt--
			idx := stOpt[len(stOpt)-1]
			last := stAll[idx]
			switch last {
			case '!':
				if v == 't' {
					v = 'f'
				} else {
					v = 't'
				}
			case '&': // 入栈的肯定全部是 t
				v = 't'
			case '|': // 入栈的肯定全部是 f
				v = 'f'
			}
			stAll = stAll[:idx] // 本次计算的收尾工作
			stOpt = stOpt[:len(stOpt)-1]
			i = check(i, v)
		case '(':
			cnt++ // 统计 '(' 已“入栈”的次数，用于 check 函数中
		}
	}
	return stAll[1] == 't'
}

// leetcode submit region end(Prohibit modification and deletion)

//func parseBoolExpr(expression string) bool {
//	// 栈：个人
//	if len(expression) == 1 { // fast path
//		return expression[0] == 't'
//	}
//	stAll, stOpt := []int32{-1}, []int32{-1} // 哨兵
//	addNew := func(v bool) {
//		var nb int32
//		if v { // 存回 t, f
//			nb = 't'
//		} else {
//			nb = 'f'
//		}
//		stAll = append(stAll[:len(stAll)-1], nb) // 移除 stAll 的操作符
//	}
//	for _, e := range expression {
//		switch e {
//		case '&', '|', '!':
//			stAll = append(stAll, e) // 操作符 + t, f
//			stOpt = append(stOpt, e) // 操作符
//		case 't', 'f':
//			stAll = append(stAll, e)
//		case ')':
//			last, b := stOpt[len(stOpt)-1], stAll[len(stAll)-1] == 't'
//			stOpt, stAll = stOpt[:len(stOpt)-1], stAll[:len(stAll)-1]
//			if last == '!' { // ! 只操作一个值
//				addNew(!b)
//				break
//			}
//			for cur := stAll[len(stAll)-1]; cur != last; cur = stAll[len(stAll)-1] { // & | 操作多个值
//				switch last {
//				case '&':
//					b = b && cur == 't'
//				case '|':
//					b = b || cur == 't'
//				}
//				stAll = stAll[:len(stAll)-1]
//			}
//			addNew(b)
//		}
//	}
//	return stAll[1] == 't'
//}
