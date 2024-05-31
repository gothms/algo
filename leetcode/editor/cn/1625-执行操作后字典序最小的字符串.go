package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "5525"
	a, b := 9, 2 // 2050
	s = "31"
	a, b = 4, 1 // 11
	s = "593290172167"
	a, b = 7, 4 // 206658319916
	s = "84688926"
	a, b = 2, 5 // 00148068
	s = "43081271531750"
	a, b = 2, 5 // 00147355195245
	s = "5562438547"
	a, b = 1, 3 // 0014305132
	smallestString := findLexSmallestString(s, a, b)
	fmt.Println(smallestString)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findLexSmallestString(s string, a int, b int) string {
	// bfs：个人
	// 由于 s.len 为偶数，则每次一操作都是 s 中一半的元素序列
	// b 为偶数：只有 s 初始的偶数序列可以被一操作
	// b 为奇数：奇/偶序列都可以被操作，即所有元素都可以被一操作
	// 扩展：如果 s.length 可以为奇数，则结果为 00..00x，那么 x 为多少呢？
	var (
		n   int
		buf []int32
	)
	rebuild := func(dif int) int32 { // 根据 a，一操作能得到的最小序列
		var d int32
		switch a {
		case 5:
			if buf[dif] >= 5 {
				d = 5
			}
		case 1, 3, 7, 9:
			d = 10 - buf[dif]
		case 2, 4, 6, 8:
			d = 10 + buf[dif]&1 - buf[dif]
		}
		return d
	}
	getD := func(dif int, ds *[2]int32) {
		if b&1 != 0 { // b 为偶数，则偶数序列不能被一操作
			ds[0] = rebuild(dif)
		}
		ds[1] = rebuild((dif + 1) % n)
	}

	n = len(s)
	buf = make([]int32, n)
	for i, c := range s {
		buf[i] = c - '0'
	}
	k := func(x, y int) int { // 最大公约数
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}(n, b)
	k = b / k * (n - 1) // 最小公倍数 - b（因为 0 占了第一个位置，所以 -b）
	var dif int
	ds, ts := [2]int32{}, [2]int32{}
	getD(dif, &ds)

	for c := b; c <= k; c += b { // 寻找字典序最小字符串
		i := c % n // 要 check 的序列的起始索引
		getD(i, &ts)
		for j := 0; j < n; j++ { // check 序列 i
			x, y := (buf[(dif+j)%n]+ds[j&1])%10, (buf[(i+j)%n]+ts[j&1])%10
			if x < y { // check 失败
				break
			} else if x > y { // check 成功，更新最小字符串
				dif, ds = i, ts
				break
			}
		}
	}
	var sb strings.Builder
	for j := 0; j < n; j++ { // 结果集
		sb.WriteRune((buf[(dif+j)%n]+ds[j&1])%10 + '0')
	}
	return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)
