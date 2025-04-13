package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

func main() {
	n := 3
	k := 5 // 27
	n = 2
	k = 1
	n = 3
	k = 1 // 243
	integers := countGoodIntegers(n, k)
	fmt.Println(integers)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countGoodIntegers(n int, k int) int64 {
	// 个人思路
	// 1.对于 n 位数，只需要讨论前 n>>1 位数（如果是 n 是奇数，则需要讨论第 n/2 位的中间数），然后将前 n>>1 位数反转并拼接
	// 2.将 n>>1 位数符合条件的数放入 map 中去重
	// 3.检验 map 中的数列能组成的不同排列数，如果是奇数，则还需要乘以第 n/2 位的中间数可以取值的总数（根据拼接后的 mod k 的值确定）
	factorial := make([]int, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	start := int(math.Pow10((n - 1) >> 1))
	end := start * 10
	memo := make(map[string]struct{})
	ans := 0
	for i := start; i < end; i++ {
		// 第一步：检查 v%k==0
		v, l := i, i
		if n&1 == 1 {
			l /= 10
		}
		for ; l != 0; l /= 10 {
			v = v*10 + l%10
		}
		if v%k != 0 {
			continue
		}
		// 第二步：去重
		buf := []byte(strconv.Itoa(v))
		slices.Sort(buf)
		s := string(buf)
		if _, ok := memo[s]; ok {
			continue
		}
		memo[s] = struct{}{}
		// 第三步：计算排列数
		cnt := [10]int{}
		for _, b := range buf {
			cnt[b-'0']++
		}
		pmt := (n - cnt[0]) * factorial[n-1]
		for _, c := range cnt {
			pmt /= factorial[c]
		}
		ans += pmt
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
