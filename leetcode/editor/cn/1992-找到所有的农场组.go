package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findFarmland(land [][]int) [][]int {
	// dfs

	// 模拟
	m, n := len(land), len(land[0])
	ans := make([][]int, 0)
	for i := range land {
		for j := 0; j < n; j++ {
			if land[i][j] == 1 && (i == 0 || land[i-1][j] == 0) && (j == 0 || land[i][j-1] == 0) {
				x := i
				for x++; x < m && land[x][j] == 1; {
					x++
				}
				x-- // 防越界
				y := j
				for y++; y < n && land[x][y] == 1; {
					y++
				}
				ans = append(ans, []int{i, j, x, y - 1})
				j = y
			}
		}
	}
	return ans

	// 个人
	//type p struct {
	//	x, y int
	//}
	//memo := make(map[[2]int][2]int)
	//row := make([]*p, len(land[0])+1)
	//for i, l := range land {
	//	for j, v := range l {
	//		if v == 1 {
	//			if i == 0 || land[i-1][j] == 0 {
	//				row[j+1] = nil
	//			}
	//			if row[j+1] == nil && j > 0 && land[i][j-1] == 1 {
	//				row[j+1] = row[j]
	//			}
	//			pair := row[j+1]
	//			if pair == nil {
	//				pair = &p{i, j}
	//				row[j+1] = pair
	//			}
	//			memo[[2]int{pair.x, pair.y}] = [2]int{i, j}
	//		}
	//	}
	//}
	//ans := make([][]int, 0, len(memo))
	//for k, v := range memo {
	//	ans = append(ans, []int{k[0], k[1], v[0], v[1]})
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
