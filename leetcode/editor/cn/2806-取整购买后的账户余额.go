package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func accountBalanceAfterPurchase(purchaseAmount int) int {
	m := purchaseAmount % 10
	if m >= 5 {
		return 100 - purchaseAmount - 10 + m
	}
	return 100 - purchaseAmount + m
}

//leetcode submit region end(Prohibit modification and deletion)
