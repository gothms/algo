package main

import "fmt"

func main() {
	s := "aaaa"
	//s = "abcdef" // -1
	length := maximumLength(s)
	fmt.Println(length)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumLength(s string) int {
	memo := [27][3]int{}
	var (
		cnt  int
		last = 'a' + 26 // 哨兵
	)
	update := func() {
		if i := last - 'a'; cnt > memo[i][2] { // 插入 cnt
			if cnt > memo[i][1] {
				memo[i][2] = memo[i][1]
				if cnt > memo[i][0] {
					memo[i][0], memo[i][1] = cnt, memo[i][0]
				} else {
					memo[i][1] = cnt
				}
			} else {
				memo[i][2] = cnt
			}
		}
	}
	for _, c := range s {
		if c != last { // 更新 memo
			update()
			cnt = 1
			last = c // 记录“上一个”字符
		} else {
			cnt++
		}
	}
	update()
	ans := 0
	for _, m := range memo {
		ans = max(ans, m[2], min(m[1], m[0]-1), m[0]-2) // 至少三次的特殊子字符串
	}
	if ans == 0 {
		return -1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
