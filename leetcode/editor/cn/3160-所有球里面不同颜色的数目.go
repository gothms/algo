package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func queryResults(limit int, queries [][]int) []int {
	ball, color := make(map[int]int), make(map[int]int)
	ans := make([]int, len(queries))
	for i, q := range queries {
		b, c := q[0], q[1]
		if last := ball[b]; last > 0 {
			if color[last] == 1 {
				delete(color, last)
			} else {
				color[last]--
			}
		}
		ball[b] = c
		color[c]++
		ans[i] = len(color)
    }
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
