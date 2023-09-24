package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "010"
	number := maximumOddBinaryNumber(s)
	fmt.Println(number)
}
func maximumOddBinaryNumber(s string) string {
	cnt := -1
	for _, c := range s {
		if c == '1' {
			cnt++
		}
	}
	var sb strings.Builder
	sb.Grow(len(s))
	for i := 0; i < cnt; i++ {
		sb.WriteByte('1')
	}
	for i := cnt; i < len(s)-1; i++ {
		sb.WriteByte('0')
	}
	sb.WriteByte('1')
	return sb.String()
}
