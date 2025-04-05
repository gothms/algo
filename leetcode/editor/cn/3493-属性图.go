package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfComponents(properties [][]int, k int) int {
	n := len(properties)
	ans := n
	uni := make([]int, n)
	for i := range uni {
		uni[i] = i
	}
	find := func(x int) int {
		for x != uni[x] {
			x, uni[x] = uni[x], uni[uni[x]]
		}
		//if x != uni[x] {
		//	uni[x] = find(uni[x])
		//}
		return uni[x]
	}
	union := func(x, y int) bool {
		rx, ry := find(x), find(y)
		if rx != ry {
			uni[rx] = ry
			ans--
			return true
		}
		return false
	}

	memo := make([]map[int]struct{}, n)
	for i, p := range properties {
		memo[i] = make(map[int]struct{})
		for _, j := range p {
			memo[i][j] = struct{}{}
		}
	}
	for i, cur := range memo {
		for j, pre := range memo[:i] {
			cnt := 0
			for x := range pre {
				if _, ok := cur[x]; ok {
					cnt++
				}
			}
			if cnt >= k {
				union(j, i)
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
