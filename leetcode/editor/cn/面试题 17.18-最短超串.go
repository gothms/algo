package main

import "fmt"

func main() {
	big := []int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7}
	small := []int{1, 5, 9} // [7,10]
	seq := shortestSeq(big, small)
	fmt.Println(seq)
}

// leetcode submit region begin(Prohibit modification and deletion)
func shortestSeq(big []int, small []int) []int {
	memo := make(map[int]int)
	for _, v := range small {
		memo[v] = -1
	}
	idx, mn := -1, len(big) // mn 为窗口长度 -1
	j, c := 0, len(memo)    // 左指针 & 窗口中与 small 相差的个数
	for i, v := range big {
		cur, ok := memo[v]
		if !ok { // 右：数字在 small 中不存在
			continue
		}
		if memo[v]++; cur+1 == 0 {
			c--
		}
		if c != 0 { // 右：窗体与 small 不匹配
			continue
		}
		for ; c == 0; j++ {
			if _, ok = memo[big[j]]; !ok { // 左：数字在 small 中不存在
				continue
			}
			if memo[big[j]]--; memo[big[j]] < 0 { // 左：窗体与 small 不匹配
				c++
				if i-j < mn { // 更新
					idx, mn = j, i-j
				}
			}
		}
	}
	if idx >= 0 {
		return []int{idx, idx + mn}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
