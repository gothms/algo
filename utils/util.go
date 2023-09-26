package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	s = "[[1,5],[5,4],[4,6],[5,8],[8,9],[9,10],[5,2],[2,3],[3,7]]"
	s = "[[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]"
	s = "[[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]"
	s = "[[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]"
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
func temp(arr [][]int) {
	//v = lfu.Get(11)
	//lfu.Put(9, 12)
	for _, v := range arr {
		switch len(v) {
		case 1:
			fmt.Printf("v = lfu.Get(%d)\nfmt.Println(v)\n", v[0])
		case 2:
			fmt.Printf("lfu.Put(%d, %d)\n", v[0], v[1])
		}
	}
}
