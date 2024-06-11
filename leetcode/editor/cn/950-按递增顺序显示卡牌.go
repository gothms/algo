package main

import (
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func deckRevealedIncreasing(deck []int) []int {
	// 模拟
	sort.Ints(deck) // 排序
	n := len(deck)
	if n <= 2 {
		return deck
	}
	m := n<<1 - 2 // 需要模拟 n-2 次操作，则需要 n<<1-2 的空间，因为后两位不用模拟
	temp := make([]int, m)
	temp[m-2], temp[m-1] = deck[n-2], deck[n-1] // 拷贝末尾两位
	for i, j, k := n-3, m-3, m-1; i >= 0; i-- { // 模拟：逆操作
		temp[j] = temp[k]
		j--
		k--
		temp[j] = deck[i]
		j--
	}
	return temp[:n]
}

//leetcode submit region end(Prohibit modification and deletion)
