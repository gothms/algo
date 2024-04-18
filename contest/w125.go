package main

import "fmt"

func main() {
	n := 1
	judge := findJudge(n, nil)
	fmt.Println(judge)

	n = 5
	lamps := [][]int{{0, 0}, {4, 4}}
	queries := [][]int{{1, 1}, {1, 0}}
	queries = [][]int{{1, 1}, {1, 1}} // [1 1]
	lamps = [][]int{{0, 0}, {0, 4}}
	queries = [][]int{{0, 4}, {0, 1}, {1, 4}} // [1 1 0]
	illumination := gridIllumination(n, lamps, queries)
	fmt.Println(illumination)
}
func findJudge(n int, trust [][]int) int {
	cnt, j := make([]int, n+1), make([]bool, n+1)
	for _, t := range trust {
		cnt[t[1]]++
		j[t[0]] = true
	}
	for i := 1; i <= n; i++ {
		if cnt[i] == n-1 && !j[i] {
			return i
		}
	}
	return -1
}
func numRookCaptures(board [][]byte) int {
	var (
		n    = len(board)
		m    = len(board[0])
		x, y int
		ret  int
	)
out:
	for i, b := range board {
		for j, v := range b {
			if v == 'R' {
				x, y = i, j
				break out
			}
		}
	}
up:
	for i := x - 1; i >= 0; i-- {
		switch board[i][y] {
		case 'p':
			ret++
			fallthrough
		case 'B':
			break up
		}
	}
down:
	for i := x + 1; i < n; i++ {
		switch board[i][y] {
		case 'p':
			ret++
			fallthrough
		case 'B':
			break down
		}
	}
left:
	for i := y - 1; i >= 0; i-- {
		switch board[x][i] {
		case 'p':
			ret++
			fallthrough
		case 'B':
			break left
		}
	}
right:
	for i := y + 1; i < m; i++ {
		switch board[x][i] {
		case 'p':
			ret++
			fallthrough
		case 'B':
			break right
		}
	}
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	tn := &TreeNode{val, nil, nil}
	if root.Val < val { // fast path
		tn.Left = root
		return tn
	}
	var dfs func(*TreeNode, *TreeNode)
	dfs = func(p, c *TreeNode) {
		if c == nil || c.Val < val {
			p.Right = tn
			tn.Left = c
			return
		}
		dfs(c, c.Right) // 只会是右子树
	}
	dfs(root, root.Right)
	return root
}
func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	var (
		//r = make([]int, n)
		//c = make([]int, n)
		//pie  = make([]int, n<<1-1)
		//na   = make([]int, n<<1-1)
		r    = make(map[int]int) // 使用 map，降低空间复杂度
		c    = make(map[int]int)
		pie  = make(map[int]int) // 使用 map，防止 int 溢出：1 <= n <= 1e9
		na   = make(map[int]int)
		memo = make(map[[2]int]struct{}) // 去重
		//dx   = [9]int{0, 0, -1, -1, -1, 0, 1, 1, 1}
		//dy   = [9]int{0, -1, -1, 0, 1, 1, 1, 0, -1}
		ret = make([]int, len(queries))
	)
	for _, l := range lamps {
		x, y := l[0], l[1]
		if _, ok := memo[[2]int{x, y}]; !ok {
			memo[[2]int{x, y}] = struct{}{}
			r[x]++
			c[y]++
			//pie[x+y]++    // 撇的算法 x+y
			//na[x-y+n-1]++ // 捺的算法 x-y+n-1
			pie[x-n+y]++ // 撇的算法 x-n+y：防止溢出（或直接使用 x+y）
			na[x-y]++    // 捺的算法 x-y
		}
	}
	for idx, q := range queries {
		x, y := q[0], q[1]
		//if r[x]+c[y]+pie[x-n+y]+na[x-y] > 0 {
		//	ret[idx] = 1
		//}
		//for d := 0; d < 9; d++ {
		//	if nx, ny := x+dx[d], y+dy[d]; 0 <= nx && nx < n && 0 <= ny && ny < n {
		//		if _, ok := memo[[2]int{nx, ny}]; ok {
		//			r[nx]--
		//			c[ny]--
		//			pie[nx-n+ny]--
		//			na[nx-ny]--
		//			delete(memo, [2]int{nx, ny})
		//		}
		//	}
		//}

		// 简化
		if r[x] > 0 || c[y] > 0 || pie[x-n+y] > 0 || na[x-y] > 0 {
			ret[idx] = 1
		}
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				if 0 <= i && i < n && 0 <= j && j < n {
					if _, ok := memo[[2]int{i, j}]; ok {
						r[i]--
						c[j]--
						pie[i-n+j]--
						na[i-j]--
						delete(memo, [2]int{i, j})
					}
				}
			}
		}
	}
	return ret
}
