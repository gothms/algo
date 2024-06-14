package main

import "fmt"

func main() {
	numBottles, numExchange := 15, 4
	numBottles, numExchange = 9, 3 // 13
	bottles := numWaterBottles(numBottles, numExchange)
	fmt.Println(bottles) // 19
}

// leetcode submit region begin(Prohibit modification and deletion)
func numWaterBottles(numBottles int, numExchange int) int {
	// math
	if numBottles < numExchange {
		return numBottles
	}
	return numBottles + (numBottles-numExchange)/(numExchange-1) + 1

	// 模拟
	//ans, empty := numBottles, numBottles
	//for ; empty >= numExchange; empty = empty%numExchange + empty/numExchange {
	//	ans += empty / numExchange
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
