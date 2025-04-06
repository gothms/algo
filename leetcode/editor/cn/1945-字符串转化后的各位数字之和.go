package main

import "fmt"

func main() {
	s := "fleyctuuajsr" // 8
	k := 5
	lucky := getLucky(s, k)
	fmt.Println(lucky)
}

// leetcode submit region begin(Prohibit modification and deletion)
func getLucky(s string, k int) int {
	// 1 <= s.length <= 100
	buf := make([]int, len(s)) // 100 位的 string 不能用 int 来存
	for i, c := range s {
		buf[i] = int(c - 'a' + 1)
	}
	v := 0
	for _, b := range buf { // 第一次加法 ，将切片转为 int
		v += b/10 + b%10
	}
	for i := 1; i < k; i++ { // 后面的加法
		w := 0
		for ; v != 0; v /= 10 {
			w += v % 10
		}
		v = w
	}
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
