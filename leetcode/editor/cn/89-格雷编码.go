package main

import "fmt"

func main() {
	n := 3
	code := grayCode(n)
	fmt.Println(code)
}

/*
   思路：
   	n = 0	1	2	3	...
   		0	0	00	000
   			1	01	001
   				11	011
   				10	010
   					110
   					111
   					101
   					100
*/
// leetcode submit region begin(Prohibit modification and deletion)

var gray89 []int

func init() {
	const n = 16
	gray89 = make([]int, 1, 1<<n)
	for i := 0; i < n; i++ {
		k := 1 << i
		for j := len(gray89) - 1; j >= 0; j-- {
			gray89 = append(gray89, k|gray89[j])
		}
	}
	//fmt.Println(gray89)
}

func grayCode(n int) []int {
	return gray89[:1<<n]
}

//leetcode submit region end(Prohibit modification and deletion)

func grayCode_(n int) []int {
	ret := make([]int, 1, 1<<(n+1))    // n = 0
	for i, head := 1, 1; i <= n; i++ { // 1-n
		for j := head - 1; j >= 0; j-- { // 逆序：长度为 head - 1
			ret = append(ret, head|ret[j]) // 头部添 1
		}
		head <<= 1
	}
	return ret
}
