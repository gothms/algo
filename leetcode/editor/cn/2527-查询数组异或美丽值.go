package main

import "fmt"

func main() {
	nums := []int{15, 45, 20, 2, 34, 35, 5, 44, 32, 30}
	beauty := xorBeauty(nums)
	fmt.Println(beauty)
}

/*
位运算经典技巧：由于每个比特位互不相干，所以拆分成每个比特位分别计算。
	由于只有 0 和 1，这样就好算了。
	对异或有影响的是 1，所以只需要统计 (a∣b)&c=1 的情况。
	那么 c 必须是 1，a 和 b 不能都是 0。

	设有 y 个 1，那么就有 x=n−y 个 0。
	那么 c 有 y 个，a∣b 有 n^2−x^2 个（任意选是 n^2，减去两个都是 0 的 x^2 个）。
	根据乘法原理，一共可以产生
	ones = (n^2−x^2)y = (n^2−(n−y)^2)y = (2ny−y^2)y 个 1。
	由于异或只在乎 ones 的奇偶性，所以 2ny 可以去掉，那么就变成看 y^3 的奇偶性，也就是 y 的奇偶性。
	如果 y 是奇数，那么这个比特位的异或值就是 1。
	这实际上就是看每个比特位的异或值是否为 1。
	那么把 nums 的每个数异或起来，就是答案。
*/

// leetcode submit region begin(Prohibit modification and deletion)
func xorBeauty(nums []int) int {
	var xor int
	for _, v := range nums {
		xor ^= v
	}
	return xor
}

//leetcode submit region end(Prohibit modification and deletion)
