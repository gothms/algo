package main

import "fmt"

func main() {
	words := []string{"acaeeccaaedabeddddae",
		"bedbbeaabbbaaeabbdea",
		"eeaadeddbebcdeaecccd",
		"edccaaeeaaceaaeedbce",
		"eadacedabddebdeebddb",
		"ecdeecbcdeceabacadac",
		"becbddbdbeededeecebb"}
	i := similarPairs(words)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func similarPairs(words []string) int {
	// 位运算
	memo := make(map[int]int)
	var ans int
	for _, w := range words {
		var v int
		for _, c := range w {
			v |= 1 << (c - 'a')
		}
		ans += memo[v]
		memo[v]++
	}
	return ans

	// hash 优化
	//memo := make(map[[26]bool]int)
	//var ans int
	//for _, w := range words {
	//	cache := [26]bool{}
	//	for _, c := range w {
	//		cache[c-'a'] = true
	//	}
	//	ans += memo[cache]
	//	memo[cache]++
	//}
	//return ans

	//const prime = 16777619
	//memo := make(map[uint32]int)
	//var ans int
	//for _, w := range words {
	//	cache := [26]bool{}
	//	for _, c := range w {
	//		cache[c-'a'] = true
	//	}
	//	var hash uint32
	//	for i, b := range cache {
	//		if b {
	//			hash = hash*prime + uint32(i+1) // i+1：不能出现 0
	//		}
	//	}
	//	ans += memo[hash]
	//	memo[hash]++
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
