package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	stamp := "abca"
	target := "aabcaca"
	stamp = "h"
	target = "hhhhh"
	stamp = "oz"
	target = "ooozz" // [0,3,1,2]
	stamp = "ozz"
	target = "ooozzzzzzzzz" // [0 1 4 3 2]
	toStamp := movesToStamp(stamp, target)
	fmt.Println(toStamp)
}

// leetcode submit region begin(Prohibit modification and deletion)
func movesToStamp(stamp string, target string) []int {
	// 最优解思路：找最长公共子串，先找最长的，再找次长的，直到最后找长为 1 的
	// 但是本题并不要求最优解，而是 10 * target.length，则有充足的戳印次数
	// 关键点：逆向操作，将字符全换为 ?
	// 1.拓扑排序：考虑相邻的戳印，一种方案是先戳印换为 ? 的操作，则应该相对较后操作
	// 2.从左往右 和 从右往左 各操作一遍
	n, m := len(stamp), len(target)
	buf := []byte(target)
	ans := make([]int, 0)
	cnt := 0
	for i, j := strings.Index(target, stamp), 0; i != -1; i = strings.Index(target[j:], stamp) {
		j += i
		ans = append(ans, j)
		for k := j; k < j+n; k++ {
			buf[k] = '?'
		}
		j += n
		cnt += n
	}
	check := func(i int) {
		k := 0
		for j := 0; j < n; j++ {
			if buf[i+j] == stamp[j] {
				k++
				continue
			}
			if buf[i+j] != '?' {
				return
			}
		}
		if k == 0 {
			return
		}
		for j := i; j < i+n; j++ {
			if buf[j] != '?' {
				buf[j] = '?'
				cnt++
			}
		}
		ans = append(ans, i)
	}
	for i := 0; i <= m-n; i++ {
		check(i)
	}
	for i := m - n; i >= 0; i-- {
		check(i)
	}
	if cnt < m {
		return nil
	}
	slices.Reverse(ans)
	return ans

	// lc
	//n := len(stamp)
	//s := []byte(target)
	//var res []int
	//var check = func(start int) {
	//	for i := 0; i < len(stamp); i++ {
	//		j := i + start
	//		if stamp[i] != s[j] && s[j] != '#' {
	//			return
	//		}
	//	}
	//	flag := false
	//	for i := 0; i < len(stamp); i++ {
	//		j := i + start
	//		if s[j] != '#' {
	//			flag = true
	//		}
	//		s[j] = '#'
	//	}
	//	if flag {
	//		res = append(res, start)
	//	}
	//}
	//for i := 0; i+n-1 < len(s); i++ {
	//	check(i)
	//}
	//for i := len(s) - n; i >= 0; i-- {
	//	check(i)
	//}
	//for i := range s {
	//	if s[i] != '#' {
	//		return []int{}
	//	}
	//
	//}
	//return res
}

//leetcode submit region end(Prohibit modification and deletion)
