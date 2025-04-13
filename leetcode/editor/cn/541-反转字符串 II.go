package main

import "slices"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	// lc
	buf := []byte(s)
	n := len(buf)
	for i := 0; i < n; i += k << 1 {
		slices.Reverse(buf[i:min(i+k, n)])
	}
	return string(buf)

	// 个人
	//n := len(s)
	//var sb strings.Builder
	//for i := 0; i < n; {
	//	j := min(n, i+k)
	//	buf := []byte(s[i:j])
	//	slices.Reverse(buf)
	//	sb.Write(buf)
	//	i = min(n, j+k)
	//	sb.WriteString(s[j:i])
	//}
	//return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)
