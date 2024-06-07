package maths

import (
	"fmt"
	"math/big"
	"math/bits"
	"strings"
)

/*
	big.Int 的比较
		x.Cmp(y)
	big.Int 转为 int64
		x.Int64()

lc
	2117
*/

func abbreviateProduct(left int, right int) string {
	// 力扣将数据范围由 10^6 下调至 10^4，大大降低了本题的难度，暴力模拟即可通过此题
	//s := new(big.Int).MulRange(int64(left), int64(right)).String()
	//tz := len(s)
	//s = strings.TrimRight(s, "0")
	//tz -= len(s)
	//if len(s) > 10 {
	//	return fmt.Sprintf("%s...%se%d", s[:5], s[len(s)-5:], tz)
	//}
	//return fmt.Sprintf("%se%d", s, tz)

	// 误差分析：https://leetcode.cn/problems/abbreviating-the-product-of-a-range/solutions/1177259/yi-ge-shu-ju-tuan-mie-jue-da-bu-fen-dai-234yd/
	// https://leetcode.cn/problems/abbreviating-the-product-of-a-range/solutions/1176795/chai-fen-cheng-si-ge-wen-ti-ji-suan-dui-6karq/
	var lim, _ = new(big.Int).SetString(strings.Repeat("9", 200), 10)
	var div, _ = new(big.Int).SetString("1"+strings.Repeat("0", 100), 10)
	cnt2, cnt5, suf, mul := 0, 0, 1, 1
	update := func(x int) {
		suf = suf * x % 1e5
		if mul != -1 {
			mul *= x
			if mul >= 1e10 { // 长度超过 10
				mul = -1
			}
		}
	}

	pre := big.NewInt(1)
	for i := left; i <= right; i++ {
		pre.Mul(pre, big.NewInt(int64(i)))
		if pre.Cmp(lim) > 0 { // 超过一定位数就截断
			pre.Quo(pre, div)
		}
		x := i
		tz := bits.TrailingZeros(uint(x)) // 因子 2 的个数
		cnt2 += tz
		x >>= tz
		for ; x%5 == 0; x /= 5 {
			cnt5++ // 因子 5 的个数
		}
		update(x)
	}
	cnt10 := min(cnt2, cnt5)
	for i := cnt10; i < cnt2; i++ {
		update(2) // 补上多拆出来的 2
	}
	for i := cnt10; i < cnt5; i++ {
		update(5) // 补上多拆出来的 5
	}

	if mul != -1 { // 不需要缩写
		return fmt.Sprintf("%de%d", mul, cnt10)
	}
	return fmt.Sprintf("%s...%05de%d", pre.String()[:5], suf, cnt10)
}
