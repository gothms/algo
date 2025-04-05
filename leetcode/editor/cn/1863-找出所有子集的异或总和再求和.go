package main

func main() {

}

/*
	https://leetcode.cn/problems/sum-of-all-subset-xor-totals/solutions/3614974/on-shu-xue-zuo-fa-pythonjavaccgojsrust-b-9txy/
*/

// leetcode submit region begin(Prohibit modification and deletion)
func subsetXORSum(nums []int) int {
	// 1.子集的第 i 位的异或值，只和第 i 位的 1 的个数有关
	// 2.m 个元素，子集的元素个数为奇数的总子集数为 2^(m-1)
	// 假设第 i 位有 x 个 1，则子集中含奇数个 1 的子集总数为 2^(x-1)
	// 又只含 0 的子集总数为 2^(n-x)
	// 则异或值 =1 的子集总数为 2^(x-1)*2^(n-x) = 2^(n-1)
	// 那么第 i 位的异或总和为 2^i * 2^(n-1)
	xor := 0
	for _, v := range nums {
		xor |= v
	}
	return xor << (len(nums) - 1)
}

//leetcode submit region end(Prohibit modification and deletion)
