package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
const (
	MOUSE_TURN = 0
	CAT_TURN   = 1
	UNKNOWN    = 0
	MOUSE_WIN  = 1
	CAT_WIN    = 2
	MAX_MOVES  = 1000
	WALL       = '#'
	CAT        = 'C'
	MOUSE      = 'M'
	FOOD       = 'F'
)

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

var (
	rows    int
	cols    int
	food    int
	degrees [][][]int
	results [][][][]int
)

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	rows = len(grid)
	cols = len(grid[0])
	mouseStart, catStart := -1, -1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			c := grid[i][j]
			if c == CAT {
				catStart = getPos(i, j, cols)
			} else if c == MOUSE {
				mouseStart = getPos(i, j, cols)
			} else if c == FOOD {
				food = getPos(i, j, cols)
			}
		}
	}
	total := rows * cols
	degrees = make([][][]int, total)
	results = make([][][][]int, total)
	for i := 0; i < total; i++ {
		degrees[i] = make([][]int, total)
		results[i] = make([][][]int, total)
		for j := 0; j < total; j++ {
			degrees[i][j] = make([]int, 2)
			results[i][j] = make([][]int, 2)
			for k := 0; k < 2; k++ {
				results[i][j][k] = make([]int, 2)
			}
		}
	}
	queue := [][]int{}
	for mouse := 0; mouse < total; mouse++ {
		mouseRow, mouseCol := mouse/cols, mouse%cols
		if grid[mouseRow][mouseCol] == WALL {
			continue
		}
		for cat := 0; cat < total; cat++ {
			catRow, catCol := cat/cols, cat%cols
			if grid[catRow][catCol] == WALL {
				continue
			}
			degrees[mouse][cat][MOUSE_TURN]++
			degrees[mouse][cat][CAT_TURN]++
			for _, dir := range dirs {
				for row, col, jump := mouseRow+dir[0], mouseCol+dir[1], 1; row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] != WALL && jump <= mouseJump; row, col, jump = row+dir[0], col+dir[1], jump+1 {
					nextMouse, nextCat := getPos(row, col, cols), getPos(catRow, catCol, cols)
					degrees[nextMouse][nextCat][MOUSE_TURN]++
				}
				for row, col, jump := catRow+dir[0], catCol+dir[1], 1; row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] != WALL && jump <= catJump; row, col, jump = row+dir[0], col+dir[1], jump+1 {
					nextMouse, nextCat := getPos(mouseRow, mouseCol, cols), getPos(row, col, cols)
					degrees[nextMouse][nextCat][CAT_TURN]++
				}
			}
		}
	}
	for pos := 0; pos < total; pos++ {
		if pos == food {
			continue
		}
		row, col := pos/cols, pos%cols
		if grid[row][col] == WALL {
			continue
		}
		results[pos][pos][MOUSE_TURN][0] = CAT_WIN
		results[pos][pos][MOUSE_TURN][1] = 0
		results[pos][pos][CAT_TURN][0] = CAT_WIN
		results[pos][pos][CAT_TURN][1] = 0
		queue = append(queue, []int{pos, pos, MOUSE_TURN})
		queue = append(queue, []int{pos, pos, CAT_TURN})
	}
	for mouse := 0; mouse < total; mouse++ {
		mouseRow, mouseCol := mouse/cols, mouse%cols
		if grid[mouseRow][mouseCol] == WALL || mouse == food {
			continue
		}
		results[mouse][food][MOUSE_TURN][0] = CAT_WIN
		results[mouse][food][MOUSE_TURN][1] = 0
		results[mouse][food][CAT_TURN][0] = CAT_WIN
		results[mouse][food][CAT_TURN][1] = 0
		queue = append(queue, []int{mouse, food, MOUSE_TURN})
		queue = append(queue, []int{mouse, food, CAT_TURN})
	}
	for cat := 0; cat < total; cat++ {
		catRow, catCol := cat/cols, cat%cols
		if grid[catRow][catCol] == WALL || cat == food {
			continue
		}
		results[food][cat][MOUSE_TURN][0] = MOUSE_WIN
		results[food][cat][MOUSE_TURN][1] = 0
		results[food][cat][CAT_TURN][0] = MOUSE_WIN
		results[food][cat][CAT_TURN][1] = 0
		queue = append(queue, []int{food, cat, MOUSE_TURN})
		queue = append(queue, []int{food, cat, CAT_TURN})
	}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		mouse, cat, turn := state[0], state[1], state[2]
		result := results[mouse][cat][turn][0]
		steps := results[mouse][cat][turn][1]
		prevStates := getPrevStates(mouse, cat, turn, grid, catJump, mouseJump)
		for _, prevState := range prevStates {
			prevMouse, prevCat, prevTurn := prevState[0], prevState[1], prevState[2]
			if results[prevMouse][prevCat][prevTurn][0] == UNKNOWN {
				winState := (result == MOUSE_WIN && prevTurn == MOUSE_TURN) || (result == CAT_WIN && prevTurn == CAT_TURN)
				if winState {
					results[prevMouse][prevCat][prevTurn][0] = result
					results[prevMouse][prevCat][prevTurn][1] = steps + 1
					queue = append(queue, []int{prevMouse, prevCat, prevTurn})
				} else {
					degrees[prevMouse][prevCat][prevTurn]--
					if degrees[prevMouse][prevCat][prevTurn] == 0 {
						results[prevMouse][prevCat][prevTurn][0] = result
						results[prevMouse][prevCat][prevTurn][1] = steps + 1
						queue = append(queue, []int{prevMouse, prevCat, prevTurn})
					}
				}
			}
		}
	}
	return results[mouseStart][catStart][MOUSE_TURN][0] == MOUSE_WIN && results[mouseStart][catStart][MOUSE_TURN][1] <= MAX_MOVES
}

func getPrevStates(mouse int, cat int, turn int, grid []string, catJump int, mouseJump int) [][]int {
	prevStates := [][]int{}
	mouseRow, mouseCol := mouse/cols, mouse%cols
	catRow, catCol := cat/cols, cat%cols
	var prevTurn int
	var maximumJump int
	var startRow int
	var startCol int
	if turn == MOUSE_TURN {
		prevTurn = CAT_TURN
	} else {
		prevTurn = MOUSE_TURN
	}
	if prevTurn == MOUSE_TURN {
		maximumJump = mouseJump
		startRow = mouseRow
		startCol = mouseCol
	} else {
		maximumJump = catJump
		startRow = catRow
		startCol = catCol
	}
	prevStates = append(prevStates, []int{mouse, cat, prevTurn})
	for _, dir := range dirs {
		for i, j, jump := startRow+dir[0], startCol+dir[1], 1; i >= 0 && i < rows && j >= 0 && j < cols && grid[i][j] != WALL && jump <= maximumJump; i, j, jump = i+dir[0], j+dir[1], jump+1 {
			var prevMouseRow int
			var prevMouseCol int
			var prevCatRow int
			var prevCatCol int
			if prevTurn == MOUSE_TURN {
				prevMouseRow = i
				prevMouseCol = j
				prevCatRow = catRow
				prevCatCol = catCol
			} else {
				prevMouseRow = mouseRow
				prevMouseCol = mouseCol
				prevCatRow = i
				prevCatCol = j
			}
			prevMouse := getPos(prevMouseRow, prevMouseCol, cols)
			prevCat := getPos(prevCatRow, prevCatCol, cols)
			prevStates = append(prevStates, []int{prevMouse, prevCat, prevTurn})
		}
	}
	return prevStates
}

func getPos(row int, col int, cols int) int {
	return row*cols + col
}

//leetcode submit region end(Prohibit modification and deletion)
