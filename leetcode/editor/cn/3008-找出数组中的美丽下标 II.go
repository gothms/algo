package main

import (
	"fmt"
)

func main() {
	s, a, b := "isawsquirrelnearmysquirrelhouseohmy", "my", "squirrel"
	k := 15 // [16,33]
	//s, a, b = "wl", "xjigt", "wl"
	//k = 2 // []
	//s, a, b = "abcd", "a", "a"
	//k = 4 // [0]
	indices := beautifulIndices(s, a, b, k)
	fmt.Println(indices)
}

// leetcode submit region begin(Prohibit modification and deletion)
func beautifulIndices(s string, a string, b string, k int) []int {
	// BoyerMoore

	// RabinKarp

	// kmp

}

//leetcode submit region end(Prohibit modification and deletion)
