package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"tars", "rats", "arts", "star"}
	//strs = []string{"omv", "ovm"}
	//strs = []string{"aaaaabdc", "aaaaacdb", "aaaaacbd", "aaaaadbc"}
	//strs = []string{"abc", "abc"}                                // 1
	//strs = []string{"jvhpg", "jhvpg", "hpvgj", "hvpgj", "vhgjp"} // 3
	//strs = []string{"dzqsu", "udzqs", "qzsud", "sudzq", "zsduq", "duszq", "sduqz", "suqzd", "szqdu", "qzuds", "dzqsu", "uqdzs", "zsduq", "quzds"}
	groups := numSimilarGroups(strs)
	fmt.Println(groups)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numSimilarGroups(strs []string) int {
	// 并查集
	//n, m := len(strs), len(strs[0])
	//uni, size := make([]int, n), make([]int, n)
	//setCount := n
	//for i := range uni {
	//	uni[i], size[i] = i, 1
	//}
	//find := func(p int) int {
	//	if uni[p] != p { // 启发式策略，不用 for
	//		p, uni[p] = uni[p], uni[uni[p]]
	//	}
	//	return uni[p]
	//}
	//union := func(p, q int) bool {
	//	rp, rq := find(p), find(q)
	//	if rp == rq {
	//		return false
	//	}
	//	if size[rp] < size[rq] {
	//		rp, rq = rq, rp
	//	}
	//	uni[rq] = rp
	//	size[rp] += size[rq] // 启发式策略
	//	setCount--
	//	return true
	//}
	//for i, s1 := range strs[:n-1] {
	//out:
	//	for j, s2 := range strs[i+1:] {
	//		var diff int
	//		for k := 0; k < m; k++ {
	//			if s1[k] != s2[k] {
	//				diff++
	//				if diff > 2 {
	//					continue out
	//				}
	//			}
	//		}
	//		union(i, j+i+1)
	//	}
	//}
	//return setCount

	// bfs
	cnt, m := 0, len(strs[0])
	memo := make(map[string]struct{})
	for _, s := range strs {
		memo[s] = struct{}{}
	}
	for _, s := range strs {
		if _, ok := memo[s]; !ok {
			continue
		}
		cnt++
		delete(memo, s)
		q := [][]byte{[]byte(s)}
		for ; len(q) > 0; q = q[1:] { // bfs
			str := q[0]
			for l := 0; l < m; l++ {
				for r := l + 1; r < m; r++ {
					if str[l] == str[r] {
						continue
					}
					str[l], str[r] = str[r], str[l] // 仅交换一对字符
					if _, ok := memo[string(str)]; ok {
						q = append(q, slices.Clone(str))
						delete(memo, string(str))
					}
					str[l], str[r] = str[r], str[l] // 回溯
				}
			}
		}
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
