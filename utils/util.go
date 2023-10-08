package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	s = "[[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]"
	s = "[[1,0,0,0,0,1,0],[0,1,1,0,0,0,0],[0,1,1,0,0,0,0],[0,0,0,1,0,0,0],[0,0,0,0,1,0,0],[1,0,0,0,0,1,0],[0,0,0,0,0,0,1]]"
	s = "[[1,2,3],[4,5,6],[7,8,9]]"
	array := LCArray2GoArray(s)
	fmt.Println(array)

}

func LCArray2GoArray(s string) string {
	var sb strings.Builder
	sb.WriteString("[][]int")
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
