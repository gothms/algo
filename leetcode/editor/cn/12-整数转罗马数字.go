package main

import (
	"bytes"
	"fmt"
	"slices"
)

func main() {
	num := 10
	roman := intToRoman(num)
	fmt.Println(roman)
}

// leetcode submit region begin(Prohibit modification and deletion)
func intToRoman(num int) string {
	o, f := []byte{'I', 'X', 'C', 'M'}, []byte{'V', 'L', 'D'}
	//var sb strings.Builder
	var sb bytes.Buffer
	for i := 0; num > 0; i++ {
		v := num % 10
		num /= 10
		switch v {
		case 0: // do nothing
		case 4:
			sb.WriteByte(f[i])
			sb.WriteByte(o[i])
		case 9:
			sb.WriteByte(o[i+1])
			sb.WriteByte(o[i])
		case 1, 2, 3:
			for j := 0; j < v; j++ {
				sb.WriteByte(o[i])
			}
		default:
			for j := 0; j < v-5; j++ {
				sb.WriteByte(o[i])
			}
			sb.WriteByte(f[i])
		}
	}
	slices.Reverse(sb.Bytes())
	return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)
