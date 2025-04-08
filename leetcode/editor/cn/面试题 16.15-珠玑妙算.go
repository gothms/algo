package main

import (
	"fmt"
	"math"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func masterMind(solution string, guess string) []int {
	memo := make(map[byte]int, 4)
	a0, a1 := 0, 8
	for i := range solution {
		if solution[i] == guess[i] {
			a0++
			a1 -= 2
		} else {
			memo[solution[i]]++
			memo[guess[i]]--
		}
	}
	for _, v := range memo {
		a1 -= int(math.Abs(float64(v)))
	}
	return []int{a0, a1 >> 1}
}

//leetcode submit region end(Prohibit modification and deletion)
