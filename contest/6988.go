package main

import (
	"fmt"
)

func main() {
	arr := [4]int{8, 9, 4, 6}
	//for i := 0; i < 4; i++ {
	//	arr[i] = rand.Intn(100)
	//}
	//fmt.Println(arr)
	fmt.Println(arr[0] ^ arr[1])
	fmt.Println(arr[2] ^ arr[3])
	fmt.Println((arr[0] ^ arr[1]) + (arr[2] ^ arr[3]))
	fmt.Println((arr[0] ^ arr[2]) | (arr[1] ^ arr[3]))
	fmt.Println((arr[0] + arr[2]) ^ (arr[1] + arr[3]))

	//coordinates := [][]int{{1, 2}, {4, 2}, {1, 3}, {5, 2}}
	//k := 5
	//pairs := countPairs(coordinates, k)
	//fmt.Println(pairs)
}
func countPairs(coordinates [][]int, k int) int {
	cnt := 0
	memo := make(map[[2]int]int)
	for _, c := range coordinates {
		for i := 0; i <= k; i++ {
			x := c[0] ^ i
			y := c[1] ^ (k - i)
			cnt += memo[[2]int{x, y}]
		}
		memo[[2]int{c[0], c[1]}]++
	}
	return cnt
}
