package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumOperationsToWriteY(grid [][]int) int {
	n := len(grid)
	m := n >> 1
	memo := [2][3]int{}
	for i, g := range grid {
		for j, v := range g {
			if i <= m && (i == j || i+j == n-1) || i > m && j == m {
				memo[0][v]++
			} else {
				memo[1][v]++
			}
		}
	}
	maxV := 0
	for i, y := range memo[0] {
		for j, ny := range memo[1] {
			if i != j {
				maxV = max(maxV, y+ny)
			}
		}
	}
	return n*n - maxV
}

//leetcode submit region end(Prohibit modification and deletion)
