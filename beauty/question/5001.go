package main

import (
	"fmt"
)

func main() {
	s := "abcdefg"
	k := 13
	s2 := rotate(s, k)
	fmt.Println(s2)
}
func rotate(password string, target int) string {
	//n := len(password)
	//k := target % n
	//var sb strings.Builder
	//for i := 0; i < n; i++ {
	//	sb.WriteByte(password[(n-k+i)%n])
	//}
	//return sb.String()

	buf := []byte(password)
	n := len(buf)
	for i := 0; i < target%n; i++ {
		temp := buf[n-1]
		for j := n - 1; j > 0; j-- {
			buf[j] = buf[j-1]
		}
		buf[0] = temp
	}
	return string(buf)
}
