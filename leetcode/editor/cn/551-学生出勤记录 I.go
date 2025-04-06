package main

import "strings"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func checkRecord(s string) bool {
	// lc
	return strings.Count(s, "A") < 2 && !strings.Contains(s, "LLL")

	// 个人
	//ac, ll := 0, 0
	//for _, c := range s {
	//	switch c {
	//	case 'A':
	//		ac++
	//		if ac >= 2 {
	//			return false
	//		}
	//		ll = 0
	//	case 'L':
	//		ll++
	//		if ll >= 3 {
	//			return false
	//		}
	//	case 'P':
	//		ll = 0
	//	}
	//}
	//return true
}

//leetcode submit region end(Prohibit modification and deletion)
