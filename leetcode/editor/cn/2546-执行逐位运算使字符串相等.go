package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			fmt.Println(i, j, i|j, i^j)
		}
	}
	/*
	   0 0 0 0
	   0 1 1 1
	   1 0 1 1
	   1 1 1 0
	*/
}

// leetcode submit region begin(Prohibit modification and deletion)
func makeStringsEqual(s string, target string) bool {
	// lc
	return strings.Contains(s, "1") == strings.Contains(target, "1")

	// 个人
	//var sCnt, tCnt [2]int
	//for _, c := range s {
	//	sCnt[c-'0']++
	//}
	//for _, c := range target {
	//	tCnt[c-'0']++
	//}
	//// 观察 main 测试中的规律
	//// 1.1 可以转为 0，但最终总会剩 1 个 1
	//// 2.没有 1，则不能转为 1
	//return !(tCnt[1] == 0 && sCnt[1] > 0 || tCnt[1] > 0 && sCnt[1] == 0)
}

//leetcode submit region end(Prohibit modification and deletion)
