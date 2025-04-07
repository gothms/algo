package main

import "strings"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func printBin(num float64) string {
	var sb strings.Builder
	sb.WriteString("0.")
	for i := 0; i < 6; i++ {
		num *= 2
		if num >= 1 {
			sb.WriteRune('1')
			num--
			if num == 0 {
				return sb.String()
			}
		} else {
			sb.WriteRune('0')
		}
	}
	return "ERROR"
}

//leetcode submit region end(Prohibit modification and deletion)
