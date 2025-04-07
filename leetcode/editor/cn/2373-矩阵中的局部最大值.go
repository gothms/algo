package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func largestLocal(grid [][]int) [][]int {
	// 二维单调队列：https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/monotone_queue.go#L214

	// 暴力

	// 个人
	//n := len(grid)
	//temp := make([][]int, n)
	//for i, g := range grid {
	//	temp[i] = make([]int, n-2)
	//	stack := make([]int, 0)
	//	for j, v := range g {
	//		for len(stack) > 0 && grid[i][stack[len(stack)-1]] <= v {
	//			stack = stack[:len(stack)-1]
	//		}
	//		stack = append(stack, j)
	//		if j >= 2 {
	//			if stack[0]+3 <= j {
	//				stack = stack[1:]
	//			}
	//			temp[i][j-2] = grid[i][stack[0]]
	//		}
	//	}
	//}
	//ans := make([][]int, n-2)
	//for i := range ans {
	//	ans[i] = make([]int, n-2)
	//}
	//for j := 0; j < n-2; j++ {
	//	stack := make([]int, 0)
	//	for i := 0; i < n; i++ {
	//		v := temp[i][j]
	//		for len(stack) > 0 && temp[stack[len(stack)-1]][j] <= v {
	//			stack = stack[:len(stack)-1]
	//		}
	//		stack = append(stack, i)
	//		if i >= 2 {
	//			if stack[0]+3 <= i {
	//				stack = stack[1:]
	//			}
	//			ans[i-2][j] = temp[stack[0]][j]
	//		}
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
