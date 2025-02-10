package main

import "fmt"

func main() {
	s, t := "rwohu", "rwuoh"
	difference := findPermutationDifference(s, t)
	fmt.Println(difference)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findPermutationDifference(s string, t string) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	idx := [26]int{}
	for i, c := range s {
		idx[c-'a'] = i
	}
	ans := 0
	for i, c := range t {
		ans += abs(i - idx[c-'a'])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
