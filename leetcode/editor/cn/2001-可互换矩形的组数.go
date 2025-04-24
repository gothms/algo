package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func interchangeableRectangles(rectangles [][]int) int64 {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	memo := make(map[[2]int]int)
	ans := 0
	for _, r := range rectangles {
		x, y := r[0], r[1]
		g := gcd(x, y)
		k := [2]int{x / g, y / g}
		ans += memo[k]
		memo[k]++
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
