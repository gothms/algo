package main

import "fmt"

func main() {
	arr := []int{0}
	arr[0], arr[0] = arr[0]+1, arr[0]-1
	fmt.Println(arr) // -1
}

// leetcode submit region begin(Prohibit modification and deletion)
type Bank struct {
	b []int64
}

func Constructor(balance []int64) Bank {
	return Bank{append([]int64{0}, balance...)}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	b := this.b
	if account1 < len(b) && account2 < len(b) && b[account1] >= money {
		b[account1] -= money
		b[account2] += money
		return true
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	b := this.b
	if account < len(b) {
		b[account] += money
		return true
	}
	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	b := this.b
	if account < len(b) && b[account] >= money {
		b[account] -= money
		return true
	}
	return false
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
//leetcode submit region end(Prohibit modification and deletion)
