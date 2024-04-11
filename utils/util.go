package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	s = "[[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]"
	s = "[[5,2,4],[3,0,5],[0,7,2]]"
	s = "[[0,2],[0,5],[2,4],[1,6],[5,4]]"
	s = "[[17,0],[30,17],[41,30],[10,30],[13,10],[7,13],[6,7],[45,10],[2,10],[14,2],[40,14],[28,40],[29,40],[8,29],[15,29],[26,15],[23,40],[19,23],[34,19],[18,23],[42,18],[5,42],[32,5],[16,32],[35,14],[25,35],[43,25],[3,43],[36,25],[38,36],[27,38],[24,36],[31,24],[11,31],[39,24],[12,39],[20,12],[22,12],[21,39],[1,21],[33,1],[37,1],[44,37],[9,44],[46,2],[4,46]]"
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
