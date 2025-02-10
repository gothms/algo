package main

import "fmt"

func main() {
	s := "xaxcd" // aawcd
	k := 4
	s = "zbbz" // aaaz
	k = 3
	t := getSmallestString(s, k)
	fmt.Println(t)
}

// leetcode submit region begin(Prohibit modification and deletion)
func getSmallestString(s string, k int) string {
	buf := []byte(s)
	for i, b := range buf {
		d := int(min('z'-b+1, b-'a'))
		if d > k {
			buf[i] -= byte(k)
			break
		}
		buf[i] = 'a'
		k -= d
	}
	return string(buf)
}

//leetcode submit region end(Prohibit modification and deletion)
