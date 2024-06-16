package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
var f357 []int

func init() {
	const n = 9
	comb9 := [10]int{} // C(9,i)
	comb9[0] = 1
	for i, j := 1, 9; i < len(comb9); i++ {
		comb9[i] = comb9[i-1] * j / i
		j--
	}
	f357 = make([]int, n)
	f357[0], f357[1] = 1, 10
	for i, j, p, k, q := 2, 9, 10, 1, 1; i < n; i++ {
		p *= j // 所有的排列
		j--
		q *= k // 用于计算 0 开头的排列
		k++
		f357[i] = f357[i-1] + p - comb9[i-1]*q // 主要算法
	}
	//fmt.Println(f357)
}
func countNumbersWithUniqueDigits(n int) int {
	return f357[n]
}

//leetcode submit region end(Prohibit modification and deletion)
