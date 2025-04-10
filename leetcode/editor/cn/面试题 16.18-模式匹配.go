package main

import (
	"fmt"
	"strings"
)

func main() {
	p := "abba" // true
	v := "dogcatcatdog"
	p = "ab" // false
	v = ""
	//p = "baabab"
	//v = "eimmieimmieeimmiee"
	//p = "bbbbbbbbabbbbbbbbbbbabbbbbbba" // true
	//v = "zezezezezezezezezkxzezezezezezezezezezezezkxzezezezezezezezkx"
	matching := patternMatching(p, v)
	fmt.Println(matching)
}

// leetcode submit region begin(Prohibit modification and deletion)
func patternMatching(pattern string, value string) bool {
	// 由测试数据可知，a/b 可以匹配空串（但 a b 不能同时匹配空串，如 ab 匹配 "" 为 false，因为此时 a==b）
	// 一旦 a 确定了，那么根据 pattern 中 a b 的个数和 value 的总长度，b 也就确定了
	// 注意点：1.pattern 中 a b 的个数可能 =0，注意除 0 error。2.pattern[0] 可能是 b
	// 个人：枚举 a
	n := len(value)
	var (
		ac, bc int // 可能 =0
		a, b   string
	)
	if pattern[0] == 'b' { // 将 pattern 换为 a 开头，或者其他写法
		buf := make([]rune, 0, len(pattern))
		for _, c := range pattern {
			buf = append(buf, 'a'+c&1)
		}
		pattern = string(buf)
	}

	for _, c := range pattern { // 统计频率
		if c == 'a' {
			ac++
		} else {
			bc++
		}
	}
	if ac == 0 || bc == 0 { // 频率其中之一为 0
		//if n == 0 { // 防止 ac+bc = 0
		//	return true
		//}
		c := ac + bc
		return n == 0 || n%c == 0 && strings.Repeat(value[:n/c], c) == value
	}
	//if n == 0 {	// 也可以不判断
	//	return false
	//}
out:
	for i := 0; i <= n/ac; i++ { // 枚举 a 的长度：a = value[:i]，不含 i
		if (n-ac*i)%bc != 0 { // 整分为 bc 个 b，长度为 (n-ac*i)/bc
			continue
		}
		j := (n - ac*i) / bc // b 的长度
		bNil := true         // b 尚未赋值
		a, b = value[:i], ""
		var sb strings.Builder
		for _, c := range pattern {
			switch c {
			case 'a':
				sb.WriteString(a)
			case 'b':
				if bNil { // 赋值
					b = value[sb.Len() : sb.Len()+j]
					if a == b {
						continue out // 终止本次循环
					}
					bNil = false
				}
				sb.WriteString(b)
			}
		}
		if sb.String() == value {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
