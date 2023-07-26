//在 `n*m` 大小的棋盘中，有黑白两种棋子，黑棋记作字母 `"X"`, 白棋记作字母 `"O"`，空余位置记作 `"."`。当落下的棋子与其他相同颜色的棋
//子在行、列或对角线完全包围（中间不存在空白位置）另一种颜色的棋子，则可以翻转这些棋子的颜色。
//
//![1.gif](https://pic.leetcode-cn.com/1630396029-eTgzpN-6da662e67368466a96d203
//f67bb6e793.gif){:height=170px}![2.gif](https://pic.leetcode-cn.com/1630396240-
//nMvdcc-8e4261afe9f60e05a4f740694b439b6b.gif){:height=170px}![3.gif](https://pic.
//leetcode-cn.com/1630396291-kEtzLL-6fcb682daeecb5c3f56eb88b23c81d33.gif){:height=170
//px}
//
//「力扣挑战赛」黑白翻转棋项目中，将提供给选手一个未形成可翻转棋子的棋盘残局，其状态记作 `chessboard`。若下一步可放置一枚黑棋，请问选手最多能翻转
//多少枚白棋。
//
//**注意：**
//- 若翻转白棋成黑棋后，棋盘上仍存在可以翻转的白棋，将可以 **继续** 翻转白棋
//- 输入数据保证初始棋盘状态无可以翻转的棋子且存在空余位置
//
//**示例 1：**
//
//> 输入：`chessboard = ["....X.","....X.","XOOO..","......","......"]`
//>
//> 输出：`3`
//>
//> 解释：
//> 可以选择下在 `[2,4]` 处，能够翻转白方三枚棋子。
//
//**示例 2：**
//
//> 输入：`chessboard = [".X.",".O.","XO."]`
//>
//> 输出：`2`
//>
//> 解释：
//> 可以选择下在 `[2,2]` 处，能够翻转白方两枚棋子。
//> ![2126c1d21b1b9a9924c639d449cc6e65.gif](https://pic.leetcode-cn.com/16266832
//55-OBtBud-2126c1d21b1b9a9924c639d449cc6e65.gif)
//
//**示例 3：**
//
//> 输入：`chessboard = [".......",".......",".......","X......",".O.....","..O....
//","....OOX"]`
//>
//> 输出：`4`
//>
//> 解释：
//> 可以选择下在 `[6,3]` 处，能够翻转白方四枚棋子。
//> ![803f2f04098b6174397d6c696f54d709.gif](https://pic.leetcode-cn.com/16303937
//70-Puyked-803f2f04098b6174397d6c696f54d709.gif)
//
//**提示：**
//- `1 <= chessboard.length, chessboard[i].length <= 8`
//- `chessboard[i]` 仅包含 `"."、"O"` 和 `"X"`
//
// Related Topics 广度优先搜索 数组 矩阵 👍 19 👎 0

package main

import "fmt"

func main() {
	chessboard := []string{
		"....X.",
		"....X.",
		"XOOO..",
		"......",
		"......",
	}
	chessboard = []string{
		".X.",
		".O.",
		"XO.",
	}
	chess := flipChess(chessboard)
	fmt.Println(chess)
}

/*
思路：dfs
	1.碰到黑子就开始寻找可以翻转的白子
		8个方向上找到一条可以翻转的路径后，把路径上的 '.' 变为黑子
		获取8个方向可翻转的最大值
	2.存在问题：
		8个方向翻转的数据互相直接会混淆
思路：bfs
	1.不能以黑子为发起 bfs 的点，因为会一次bfs会放下 8 个新的黑子，各自数据混淆了
		否则每一个翻转的白子都要 defer
	2.bfs
		2.1.每一个 '.' 都尝试放下一个黑子
		2.2.如果有可翻转的白子，全部翻为白子，并记录为下一次 bfs 的起点
		2.3.每一次尝试，都需要全新的原数据来记录翻转情况，防止数据混淆
*/
// leetcode submit region begin(Prohibit modification and deletion)
func flipChess(chessboard []string) int {
	// bfs
	cnt, n, m := 0, len(chessboard), len(chessboard[0])
	ok := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < m
	}
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	di, dj := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}, [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	var bfs func([][2]int) int
	bfs = func(queue [][2]int) int {
		c, memo := 0, make([][]uint8, n) // memo 是全新的原数据，避免混淆
		for i := 0; i < n; i++ {
			memo[i] = make([]uint8, m)
			for j := range memo[i] {
				memo[i][j] = chessboard[i][j] // 否则每一个翻转的白子都要 defer
			}
		}
		memo[queue[0][0]][queue[0][1]] = 'X' // 开始bfs的索引
		for len(queue) > 0 {                 // bfs
			i, j := queue[0][0], queue[0][1]
			queue = queue[1:]
		out:
			for k := 0; k < 8; k++ { // 查询可翻转的白子
				ni, nj := i+di[k], j+dj[k]
			reverse:
				for ok(ni, nj) {
					switch memo[ni][nj] {
					case '.': // 不可翻转
						continue out
					case 'X': // 可翻转
						break reverse
					}
					ni, nj = ni+di[k], nj+dj[k]
				}
				if !ok(ni, nj) { // 防止越界
					continue
				}
				for xi, yi := i+di[k], j+dj[k]; memo[xi][yi] != 'X'; xi, yi = xi+di[k], yi+dj[k] {
					queue, memo[xi][yi] = append(queue, [2]int{xi, yi}), 'X' // 新的黑子
					c++
				} // xi,yi 要使用新的变量
			}
		}
		return c
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if chessboard[i][j] == '.' {
				cnt = maxVal(cnt, bfs([][2]int{{i, j}}))
			}
		}
	}
	return cnt
}
	// bfs
	//cnt, n, m := 0, len(chessboard), len(chessboard[0])
	//memo := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	memo[i] = make([]int, m)
	//}
	//ok := func(i, j int) bool {
	//	return i >= 0 && i < n && j >= 0 && j < m
	//}
	//abs := func(a int) int {
	//	if a < 0 {
	//		return -a
	//	}
	//	return a
	//}
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//di, dj := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}, [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	//var bfs func(int, int) int
	//bfs = func(i int, j int) int {
	//	//fmt.Println(i, j)
	//	memo[i][j] = -1 // 已访问
	//	max := 0
	//	for k := 0; k < 8; k++ {
	//		ni, nj := i+di[k], j+dj[k]
	//	out:
	//		for ok(ni, nj) {
	//			//fmt.Println("n", ni, nj)
	//			if memo[ni][nj] < 0 {
	//				break out
	//			}
	//			switch chessboard[ni][nj] {
	//			case 'O':
	//				ni, nj = ni+di[k], nj+dj[k]
	//			//case 'X':
	//			//	break out
	//			case '.': // 错误：要连锁反应
	//				fmt.Println(i, j, ni, nj, abs(ni-i)-1)
	//				if c := maxVal(abs(ni-i), abs(nj-j)) - 1; c > 0 {
	//					memo[ni][nj] += c
	//					max = maxVal(max, memo[ni][nj])
	//				}
	//				break out
	//			}
	//		}
	//	}
	//	return max
	//}
	//for i := 0; i < n; i++ {
	//	for j := 0; j < m; j++ {
	//		if chessboard[i][j] == 'X' && memo[i][j] == 0 {
	//			cnt = maxVal(cnt, bfs(i, j))
	//		}
	//	}
	//}
	//for i := 0; i < n; i++ {
	//	fmt.Println(memo[i])
	//}
	//return cnt

	// dfs
	//cnt, n, m := 0, len(chessboard), len(chessboard[0])
	//memo := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	memo[i] = make([]int, m)
	//}
	//ok := func(i, j int) bool {
	//	return i >= 0 && i < n && j >= 0 && j < m
	//}
	//abs := func(a int) int {
	//	if a < 0 {
	//		return -a
	//	}
	//	return a
	//}
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//di, dj := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}, [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	//var dfs func(int, int) int
	//dfs = func(i int, j int) int {
	//	fmt.Println(i, j)
	//	memo[i][j] = 1
	//	sum := 0
	//	for k := 0; k < 8; k++ {
	//		ni, nj := i+di[k], j+dj[k]
	//	out:
	//		for ok(ni, nj) {
	//			//fmt.Println("n", ni, nj)
	//			if memo[ni][nj] > 0 {
	//				break out
	//			}
	//			switch chessboard[ni][nj] {
	//			case 'O':
	//				ni, nj = ni+di[k], nj+dj[k]
	//			case 'X':
	//				memo[ni][nj] = 1
	//				sum += dfs(ni, nj)
	//			case '.':
	//				c := abs(ni-i) - 1
	//				if c <= 0 {
	//					break out
	//				}
	//				//fmt.Println(ni, nj, "=====", c)
	//				memo[ni][nj] = 1 // 而且只能放一个黑子，不能放多个
	//				for bi, bj := ni-di[k], nj-dj[k]; memo[bi][bj] != 1; bi, bj = bi-di[k], bj-dj[k] {
	//					memo[bi][bj] = 1
	//				}
	//				sum += c + dfs(ni, nj) // 存在白子重复计数的问题
	//			}
	//		}
	//	}
	//	return sum
	//}
	//for i := 0; i < n; i++ {
	//	for j := 0; j < m; j++ {
	//		if chessboard[i][j] == 'X' && memo[i][j] == 0 {
	//			cnt = maxVal(cnt, dfs(i, j))
	//		}
	//	}
	//}
	//for i := 0; i < n; i++ {
	//	fmt.Println(memo[i])
	//}
	//return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
