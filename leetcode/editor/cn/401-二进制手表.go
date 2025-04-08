package main

import (
	"fmt"
)

func main() {
	combinations := generateCombinations(4, 2)
	fmt.Println(combinations)

	turnedOn := 8
	watch := readBinaryWatch(turnedOn)
	fmt.Println(watch)
}

// 穷举的方式
func generateCombinations(n, k int) [][]int {
	var result [][]int
	var backtrack func(start int, path []int)

	backtrack = func(start int, path []int) {
		if len(path) == k {
			comb := make([]int, k)
			copy(comb, path)
			result = append(result, comb)
			return
		}
		for i := start; i <= n; i++ {
			backtrack(i+1, append(path, i))
		}
	}

	backtrack(0, []int{})
	return result
}

//leetcode submit region begin(Prohibit modification and deletion)

var h401, m401 [][]int

func init() {
	// 参考 lc-77
	bitN := func(n, mx int) [][]int {
		ans := make([][]int, n)
		ans[0] = []int{0}
		for k := 1; k < n; k++ { // 穷举出二进制位含 k 个 1 的数字，且数字 < mx
			ans[k] = make([]int, 0)
			v := 1<<k - 1 // 如 k=3，则初始为 111
			for v < mx {
				ans[k] = append(ans[k], v)
				// 将 1011100 变为下一个数 1100011
				// 方法一
				//c := bits.TrailingZeros(uint(v))
				//v = v>>c + 1     // 11000
				//w := v & -v      // 1000
				//v <<= c          // 1100000
				//v = v | (w-1)>>1 // 1100011

				// 方法二

				// 方法三
				//c := bits.TrailingZeros(uint(v))              // 2
				//v += v & -v                                   // 1100000
				//v |= 1<<(bits.TrailingZeros(uint(v))-c-1) - 1 // v | 11
			}
		}
		return ans
	}
	h, m := 4, 6
	h401, m401 = bitN(h, 12), bitN(m, 60)
}

func readBinaryWatch(turnedOn int) []string {
	// 个人
	// 参考：面试题 05.04-下一个数
	if turnedOn > 8 {
		return nil
	}
	const H, M = 3, 5
	ans := make([]string, 0)
	for h := 0; h <= min(turnedOn, H); h++ {
		if m := turnedOn - h; m <= M {
			for _, v := range h401[h] {
				for _, w := range m401[m] {
					ans = append(ans, fmt.Sprintf("%d:%02d", v, w))
				}
			}
		}
	}
	return ans

	// lc
	//ans := make([]string, 0)
	//for i := 0; i < 1<<10; i++ { // 1024
	//	h, m := i>>6, i&63
	//	if h < 12 && m < 60 && bits.OnesCount(uint(h))+bits.OnesCount(uint(m)) == turnedOn {
	//		ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
