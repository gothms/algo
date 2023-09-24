package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	s = "[[1,5],[5,4],[4,6],[5,8],[8,9],[9,10],[5,2],[2,3],[3,7]]"
	array := LCArray2GoArray(s)
	fmt.Println(array)
}

func LCArray2GoArray(s string) string {
	var sb strings.Builder
	for i, c := range s {
		switch c {
		case '[':
			sb.WriteByte('{')
		case ']':
			sb.WriteByte('}')
		default:
			sb.WriteRune(c)
			if c == ',' && i-1 >= 0 && s[i-1] == ']' {
				sb.WriteByte('\n')
			}
		}
	}
	return sb.String()
}
