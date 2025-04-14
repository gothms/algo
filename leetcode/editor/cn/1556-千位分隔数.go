package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 1234
	n = 0
	separator := thousandSeparator(n)
	fmt.Println(separator)
}

// leetcode submit region begin(Prohibit modification and deletion)
func thousandSeparator(n int) string {
	buf := []byte(strconv.Itoa(n))
	m := len(buf)
	i := m % 3
	if i == 0 {
		i = min(3, m)
	}
	var sb strings.Builder
	sb.Write(buf[:i])
	for ; i < m; i += 3 {
		sb.WriteRune('.')
		sb.Write(buf[i : i+3])
	}
	return sb.String()

	//buf := []byte(strconv.Itoa(n))
	//slices.Reverse(buf)
	//m := len(buf)
	//var sb strings.Builder
	//sb.Write(buf[:min(3, m)])
	//for i := 3; i < m; i += 3 {
	//	sb.WriteRune('.')
	//	sb.Write(buf[i:min(i+3, m)])
	//}
	//buf = []byte(sb.String())
	//slices.Reverse(buf)
	//return string(buf)
}

//leetcode submit region end(Prohibit modification and deletion)
