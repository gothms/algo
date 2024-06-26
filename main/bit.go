package main

import (
	"fmt"
	"math/bits"
)

func main() {
	for x := 1; x < 20; x++ {
		fmt.Print(x, x|(x+1), " ") // 最小位的0，变1
		fmt.Print(^x&(x+1), " ")   // 取出最小位的0，用1标识
		fmt.Println((x + 1) &^ x)  // 取出最小位的0，用1标识
	}

	var x uint = 4
	y := bits.LeadingZeros(x)
	fmt.Println(y)
	y = bits.TrailingZeros(x)
	fmt.Println(y)
	y = bits.Len(x)
	fmt.Println(y)

	n := 20
	for i := 1; i < n; i++ {
		v := i ^ (i - 1)
		fmt.Println(i, v)
	}

	//a, b := 2*3*7*11*17*2*2*3, 3*17*13*5*2*5*23
	//for a != 0 {
	//	b, a = a, b%a
	//}
	//fmt.Println(a, b)
}
