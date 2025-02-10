package main

import "fmt"

func main() {
	graph := [][]int{{2, 5}, {3}, {0, 4, 5}, {1, 4, 5}, {2, 3}, {0, 2, 3}}
	game := catMouseGame(graph)
	fmt.Println(game)
}

// leetcode submit region begin(Prohibit modification and deletion)
const (
	HOLE        = 0
	MOUSE_START = 1
	CAT_START   = 2
	MOUSE_TURN  = 0
	CAT_TURN    = 1
	UNKNOWN     = 0
	MOUSE_WIN   = 1
	CAT_WIN     = 2
)

func catMouseGame(graph [][]int) int {
	var (
		n       int
		degrees [][][]int
		results [][][]int
	)

	n = len(graph)
	degrees = make([][][]int, n)
	results = make([][][]int, n)
	for i := 0; i < n; i++ {
		degrees[i] = make([][]int, n)
		results[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			degrees[i][j] = make([]int, 2)
			results[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			degrees[i][j][MOUSE_TURN] = len(graph[i])
			degrees[i][j][CAT_TURN] = len(graph[j])
		}
	}
	for i := 0; i < n; i++ {
		for _, j := range graph[HOLE] {
			degrees[i][j][CAT_TURN]--
		}
	}
	queue := [][]int{}
	for i := 1; i < n; i++ {
		results[i][i][MOUSE_TURN] = CAT_WIN
		results[i][i][CAT_TURN] = CAT_WIN
		queue = append(queue, []int{i, i, MOUSE_TURN})
		queue = append(queue, []int{i, i, CAT_TURN})
	}
	for j := 1; j < n; j++ {
		results[HOLE][j][MOUSE_TURN] = MOUSE_WIN
		results[HOLE][j][CAT_TURN] = MOUSE_WIN
		queue = append(queue, []int{HOLE, j, MOUSE_TURN})
		queue = append(queue, []int{HOLE, j, CAT_TURN})
	}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		mouse, cat, turn := state[0], state[1], state[2]
		result := results[mouse][cat][turn]
		prevStates := getPrevStates(mouse, cat, turn, graph)
		for _, prevState := range prevStates {
			prevMouse, prevCat, prevTurn := prevState[0], prevState[1], prevState[2]
			if results[prevMouse][prevCat][prevTurn] == UNKNOWN {
				winState := (result == MOUSE_WIN && prevTurn == MOUSE_TURN) || (result == CAT_WIN && prevTurn == CAT_TURN)
				if winState {
					results[prevMouse][prevCat][prevTurn] = result
					queue = append(queue, []int{prevMouse, prevCat, prevTurn})
				} else {
					degrees[prevMouse][prevCat][prevTurn]--
					if degrees[prevMouse][prevCat][prevTurn] == 0 {
						results[prevMouse][prevCat][prevTurn] = result
						queue = append(queue, []int{prevMouse, prevCat, prevTurn})
					}
				}
			}
		}
	}
	return results[MOUSE_START][CAT_START][MOUSE_TURN]
}

func getPrevStates(mouse int, cat int, turn int, graph [][]int) [][]int {
	prevStates := [][]int{}
	var prevTurn int
	if turn == MOUSE_TURN {
		prevTurn = CAT_TURN
	} else {
		prevTurn = MOUSE_TURN
	}
	if prevTurn == CAT_TURN {
		for _, prevCat := range graph[cat] {
			if prevCat != HOLE {
				prevStates = append(prevStates, []int{mouse, prevCat, prevTurn})
			}
		}
	} else {
		for _, prevMouse := range graph[mouse] {
			prevStates = append(prevStates, []int{prevMouse, cat, prevTurn})
		}
	}
	return prevStates
}

//leetcode submit region end(Prohibit modification and deletion)
