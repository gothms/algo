package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numEquivDominoPairs(dominoes [][]int) int {
	memo := make(map[int]int)
	ans := 0
	for _, d := range dominoes {
		x, y := d[0], d[1]
		v := min(x, y)*10 + max(x, y)
		ans += memo[v]
		memo[v]++
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
