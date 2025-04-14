package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 443
	//num = 63
	//num = 10
	//num = 181
	//num = 13     // false
	//num = 100000 // false
	num = 187   // true
	num = 190   // false
	num = 211   // false
	num = 20442 // false
	num = 99998 // false
	num = 1291  // true
	//num = 111_1111_1111 // false
	reverse := sumOfNumberAndReverse(num)
	fmt.Println(reverse)

	//fmt.Println('0'&1, '1'&1, '2'&1)
}

// leetcode submit region begin(Prohibit modification and deletion)
func sumOfNumberAndReverse(num int) bool {
	// 其他：https://leetcode.cn/problems/sum-of-number-and-its-reverse/solutions/1896433/ji-yi-hua-sou-suo-geng-da-shu-ju-fan-wei-xyer/
	s := strconv.Itoa(num) // 将数字转为字符串
	n := len(s)

	// 定义递归闭包函数
	var f func(int, int, int, int) bool
	f = func(i, j, pre, suf int) bool {
		if i < j { // 未到达中间位置
			// 尝试可能的进位值（0或1）
			for m := 0; m <= 1; m++ {
				// 计算当前位的总和
				sum := pre*10 + int(s[i]-'0') - m

				// 检查合法性
				minVal := 0
				if i == 0 { // 最高位不能为0
					minVal = 1
				}

				// 验证总和是否合法，且末位匹配
				if sum >= minVal && sum <= 18 && (sum+suf)%10 == int(s[j]-'0') {
					// 递归处理下一对数字
					if f(i+1, j-1, m, (sum+suf)/10) {
						return true
					}
				}
			}
			return false
		}

		// 处理中间情况
		if i == j { // 奇数长度，中间位
			return (int(s[i]-'0')-suf)%2 == 0 // 必须为偶数
		}
		return pre == suf // 偶数长度，进位必须相等
	}
	// 主逻辑：尝试两种情况
	if f(0, n-1, 0, 0) {
		return true
	}
	// 处理最高位进位的情况（例如 999 + 1 = 1000）
	if n > 0 && s[0] == '1' {
		return f(1, n-1, 1, 0)
	}
	return false

	//// 个人：dfs
	////cnt := 0
	//var dfs func([]byte, bool) bool
	//dfs = func(buf []byte, adv bool) bool { // adv 标记“是否进位”
	//	//cnt++
	//	fmt.Println(string(buf))
	//	n := len(buf)
	//	switch n {
	//	case 0: // 因为判断了 case 2，所以没有 case 0
	//		return true
	//	case 1:
	//		return buf[0]&1 == 0 // 必须是偶数
	//	case 2: // 只可能是 xx 或者 1y 且 y 是偶数，才是 true
	//		return !adv && buf[0] == buf[1] || buf[0] == '1' && buf[1]&1 == 0 // 比如 211 时，必须判断 adv
	//	}
	//
	//	l, r := buf[0], buf[n-1]
	//	switch l {
	//	case '1':
	//		if l == r && dfs(buf[1:n-1], false) { // 不用进位
	//			return true
	//		}
	//		// 必须进位，判断次高位与个位的关系
	//		if buf[1] != r && buf[1] != r+1 || buf[n-1] == '9' { // 两个个位数相加，不可能是 19
	//			return false
	//		}
	//		j := n - 2
	//		for j > 0 && buf[j] == '0' { // 向前借位
	//			buf[j] = '9'
	//			j--
	//		}
	//		if j == 0 || j == 1 && buf[1] == r { // 不够借
	//			return false
	//		}
	//		buf[j]--
	//		// 写法一
	//		//switch buf[1] {
	//		//case r: // 下一轮不进位
	//		//	return dfs(buf[2:n-1], false)
	//		//case r + 1: // 下一轮进位
	//		//	buf[1] = '1'
	//		//	return dfs(buf[1:n-1], true)
	//		//}
	//		// 写法二
	//		buf[1] -= r - '0'
	//		return dfs(buf[1:n-1], buf[1] == '1')
	//	// 写法一
	//	//case r:
	//	//	return dfs(buf[1:n-1], false)
	//	//case r + 1:
	//	//	buf[0] = '1'
	//	//	return dfs(buf[:n-1], true)
	//	// 写法二
	//	case r, r + 1:
	//		buf[0] -= r - '0'
	//		return dfs(buf[:n-1], buf[0] == '1')
	//	}
	//	return false
	//}
	////defer func() { fmt.Println(cnt) }()
	//return dfs([]byte(strconv.Itoa(num)), false)

	// 暴力：根据题意，可以前导 0
	//	n := len(strconv.Itoa(num))
	//out:
	//	for i := 0; i <= num>>1; i++ {
	//		a, b := fmt.Sprintf("%0*d", n, i), fmt.Sprintf("%0*d", n, num-i)
	//		l, r := 0, n-1
	//		if a[0] == '0' && b[0] == '0' { // 特殊情况
	//			l = 1
	//		}
	//		for ; l < n; l, r = l+1, r-1 {
	//			if a[l] != b[r] {
	//				continue out
	//			}
	//		}
	//		return true
	//	}
	//	return false
}

//leetcode submit region end(Prohibit modification and deletion)
