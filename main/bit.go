package main

import "fmt"

func main() {
	for x := 1; x < 20; x++ {
		fmt.Print(x, x|(x+1), " ") // 最小位的0，变1
		fmt.Print(^x&(x+1), " ")   // 取出最小位的0，用1标识
		fmt.Println((x + 1) &^ x)  // 取出最小位的0，用1标识
	}
}
