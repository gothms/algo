package main

import (
	"fmt"
	"math"
)

func main() {
	s := "hello"
	ofString := scoreOfString(s)
	fmt.Println(ofString)
}

// leetcode submit region begin(Prohibit modification and deletion)
func scoreOfString(s string) int {
	var ans float64
	for i := 1; i < len(s); i++ {
		ans += math.Abs(float64(s[i]) - float64(s[i-1]))
	}
	return int(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
