//一个 2D 网格中的 峰值 是指那些 严格大于 其相邻格子(上、下、左、右)的元素。
//
// 给你一个 从 0 开始编号 的 m x n 矩阵 mat ，其中任意两个相邻格子的值都 不相同 。找出 任意一个 峰值 mat[i][j] 并 返回其位置
// [i,j] 。
//
// 你可以假设整个矩阵周边环绕着一圈值为 -1 的格子。
//
// 要求必须写出时间复杂度为 O(m log(n)) 或 O(n log(m)) 的算法
//
//
//
//
//
// 示例 1:
//
//
//
//
//输入: mat = [[1,4],[3,2]]
//输出: [0,1]
//解释: 3 和 4 都是峰值，所以[1,0]和[0,1]都是可接受的答案。
//
//
// 示例 2:
//
//
//
//
//输入: mat = [[10,20,15],[21,30,14],[7,16,32]]
//输出: [1,1]
//解释: 30 和 32 都是峰值，所以[1,1]和[2,2]都是可接受的答案。
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 500
// 1 <= mat[i][j] <= 10⁵
// 任意两个相邻元素均不相等.
//
//
// Related Topics 数组 二分查找 矩阵 👍 78 👎 0

package main

import "fmt"

func main() {
	mat := [][]int{{1, 4}, {3, 2}}
	mat = [][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}
	mat = [][]int{
		{25, 37, 23, 37, 19},
		{45, 19, 2, 43, 26},
		{18, 1, 37, 44, 50}}
	grid := findPeakGrid(mat)
	fmt.Println(grid)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findPeakGrid(mat [][]int) []int {
	// lc：二分
	idx := func(x int) int { // 某一行的峰值
		y := 0
		for i, v := range mat[x] {
			if v > mat[x][y] {
				y = i
			}
		}
		return y
	}
	l, r := 0, len(mat)-1
	for l < r { // 二分行
		x := (l + r) >> 1
		y := idx(x)
		if mat[x][y] > mat[x+1][y] {
			r = x
		} else {
			l = x + 1
		}
	}
	return []int{l, idx(l)}

	// 二分：个人
	//m, n := len(mat), len(mat[0])-1
	//l, r, u, d := 0, n, 0, m-1
	//for u < d {
	//	x := (u + d) >> 1
	//	l, r = 0, n
	//	for l < r {
	//		mid := (l + r) >> 1
	//		if mat[x][mid] > mat[x][mid+1] {
	//			r = mid
	//		} else {
	//			l = mid + 1
	//		}
	//	}
	//	fmt.Println(x, l, r)
	//	if mat[x][l] > mat[x+1][l] && (x == 0 || mat[x][l] > mat[x-1][l]) {
	//		u = x
	//		break
	//	} else if mat[x][l] > mat[x+1][l] {
	//		d = x
	//	} else {
	//		u = x + 1
	//	}
	//}
	//return []int{u, l}
}

//leetcode submit region end(Prohibit modification and deletion)
