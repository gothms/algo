package main

import (
	"fmt"
	"strings"
)

func main() {
	a, b, c := "zab", "aba", "baz"   // "zabaz"
	a, b, c = "a", "c", "b"          // "abc"
	a, b, c = "xyyyz", "xzyz", "zzz" // "xyyyzxzyzzz"
	s := minimumString(a, b, c)
	fmt.Println("s", s)
}
func minimumString(a string, b string, c string) string {
	res := merge(merge(a, b), c) // 初始化
	getMinString := func(s string) {
		if len(s) < len(res) || len(s) == len(res) && s < res {
			res = s // 更新
		}
	}
	s := merge(merge(a, c), b)
	getMinString(s)
	s = merge(merge(b, a), c)
	getMinString(s)
	s = merge(merge(b, c), a)
	getMinString(s)
	s = merge(merge(c, a), b)
	getMinString(s)
	s = merge(merge(c, b), a)
	getMinString(s)
	return res
}
func merge(a, b string) string { // 合并两个字符串
	i, j := len(a), len(b)
	if i > j {
		if strings.Contains(a, b) { // 子串
			return a
		}
	} else {
		if strings.Contains(b, a) {
			return b
		}
	}
	for k := minVal(i, j); ; k-- { // 从最长开始比较
		if a[i-k:] == b[:k] {
			return a + b[k:]
		}
	}
}
