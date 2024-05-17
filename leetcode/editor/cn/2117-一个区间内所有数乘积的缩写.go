package main

import (
    "fmt"
    "math/big"
    "strings"
)

func main() {
	left, right := 1, 4
	left, right = 2, 11
	left, right = 4838, 6186 // 36088...36896e337
	left, right = 44, 9556   // 10205...06688e2378
	left, right = 1269, 7292 // 20535...03584e1505
	product := abbreviateProduct(left, right)
	fmt.Println(product)

	//fmt.Println(math.Log(10))
}

// leetcode submit region begin(Prohibit modification and deletion)
func abbreviateProduct(left int, right int) string {
	// 力扣将数据范围由 10^6 下调至 10^4，大大降低了本题的难度，暴力模拟即可通过此题
	s := new(big.Int).MulRange(int64(left), int64(right)).String()
	tz := len(s)
	s = strings.TrimRight(s, "0")
	tz -= len(s)
	if len(s) > 10 {
		return fmt.Sprintf("%s...%se%d", s[:5], s[len(s)-5:], tz)
	}
	return fmt.Sprintf("%se%d", s, tz)

	// 误差分析：https://leetcode.cn/problems/abbreviating-the-product-of-a-range/solutions/1177259/yi-ge-shu-ju-tuan-mie-jue-da-bu-fen-dai-234yd/
	// https://leetcode.cn/problems/abbreviating-the-product-of-a-range/solutions/1176795/chai-fen-cheng-si-ge-wen-ti-ji-suan-dui-6karq/
	//var lim, _ = new(big.Int).SetString(strings.Repeat("9", 200), 10)
	//var div, _ = new(big.Int).SetString("1"+strings.Repeat("0", 100), 10)
	//fmt.Println(lim)
	//fmt.Println(div)
    //
	//cnt2, cnt5, suf, mul := 0, 0, 1, 1
	//update := func(x int) {
	//	suf = suf * x % 1e5
	//	if mul != -1 {
	//		mul *= x
	//		if mul >= 1e10 { // 长度超过 10
	//			mul = -1
	//		}
	//	}
	//}
    //
	//pre := big.NewInt(1)
	//for i := left; i <= right; i++ {
	//	pre.Mul(pre, big.NewInt(int64(i)))
	//	if pre.Cmp(lim) > 0 { // 超过一定位数就截断
	//		pre.Quo(pre, div) // pre/div
	//	}
	//	x := i
	//	tz := bits.TrailingZeros(uint(x)) // 因子 2 的个数
	//	cnt2 += tz
	//	x >>= tz
	//	for ; x%5 == 0; x /= 5 {
	//		cnt5++ // 因子 5 的个数
	//	}
	//	update(x)
	//}
	//cnt10 := min(cnt2, cnt5)
	//for i := cnt10; i < cnt2; i++ {
	//	update(2) // 补上多拆出来的 2
	//}
	//for i := cnt10; i < cnt5; i++ {
	//	update(5) // 补上多拆出来的 5
	//}
    //
	//if mul != -1 { // 不需要缩写
	//	return fmt.Sprintf("%de%d", mul, cnt10)
	//}
	//return fmt.Sprintf("%s...%05de%d", pre.String()[:5], suf, cnt10)

	// lc 官方题解：并不能 AC
	//logAns := 0.0
	//cnt2, cnt5 := 0, 0
	//for i := left; i <= right; i++ {
	//	x := i
	//	c2 := bits.TrailingZeros(uint(x))
	//	x >>= c2
	//	cnt2 += c2
	//	for ; x%5 == 0; x /= 5 {
	//		cnt5++
	//	}
	//	logAns += math.Log10(float64(i))
	//}
	//cnt, k := min(cnt2, cnt5), int(logAns)
	//if k+1-cnt > 10 {
	//	ansD := logAns - float64(k)
	//	first := math.Floor(math.Pow(10, ansD)*10_000 + 0.5)
	//	fmt.Println(first)
	//	fmt.Println(math.Pow(10, ansD) * 10_000)
	//	need2, need5 := cnt, cnt
	//	last := 1
	//	for i := left; i <= right; i++ {
	//		x := i
	//		for need2 > 0 && x&1 == 0 {
	//			x >>= 1
	//			need2--
	//		}
	//		for need5 > 0 && x%5 == 0 {
	//			x /= 5
	//			need5--
	//		}
	//		last = last * x % 100_000
	//	}
	//	sLast := strconv.Itoa(last + 100_000)
	//	return strconv.Itoa(int(first)) + "..." + sLast[1:] + "e" + strconv.Itoa(cnt)
	//} else {
	//	need2, need5 := cnt, cnt
	//	ans := 1
	//	for i := left; i <= right; i++ {
	//		x := i
	//		for need2 > 0 && x&1 == 0 {
	//			x >>= 1
	//			need2--
	//		}
	//		for need5 > 0 && x%5 == 0 {
	//			x /= 5
	//			need5--
	//		}
	//		ans *= x
	//	}
	//	return strconv.Itoa(ans) + "e" + strconv.Itoa(cnt)
	//}
}

//leetcode submit region end(Prohibit modification and deletion)
