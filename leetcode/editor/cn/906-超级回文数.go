package main

import (
	"fmt"
	"math"
	"math/bits"
	"sort"
	"strconv"
)

func main() {
	s := strconv.Itoa(math.MaxInt64)
	fmt.Println(s, len(s))
	fmt.Println(math.MaxInt64, math.MaxInt)

	v := 1234321
	k := bits.Len(uint(v))
	reverse := bits.Reverse(uint(v))
	fmt.Printf("%b,%b\n", v, reverse>>(64-k))

	c := 0
	for i := 4000000008000000004; i != 0; i /= 10 {
		c++
	}
	fmt.Println(c)

	inRange := superpalindromesInRange(strconv.Itoa(1), strconv.Itoa(math.MaxInt64>>1))
	fmt.Println(inRange)
}

// leetcode submit region begin(Prohibit modification and deletion)
//var sir906 []int
//
//func init() {
//	isPalindrome := func(x int) bool {
//		v := x
//		y := 0
//		for v != 0 {
//			y = y*10 + v%10
//			v /= 10
//		}
//		return x == y
//	}
//	const n, m, R = 70, 100_000, 1e18
//	sir906 = make([]int, 0, n)
//	for i := 1; i < m; i++ {
//		v := i
//		even := i
//		for v != 0 {
//			even = even*10 + v%10
//			v /= 10
//		}
//		even *= even
//		if even < R && isPalindrome(even) {
//			sir906 = append(sir906, even)
//		}
//		v = i / 10
//		odd := i
//		for v != 0 {
//			odd = odd*10 + v%10
//			v /= 10
//		}
//		odd *= odd
//		if isPalindrome(odd) {
//			sir906 = append(sir906, odd)
//		}
//	}
//	sort.Ints(sir906)
//	//for _, v := range sir906 {
//	//	fmt.Print(v, ",")
//	//}
//	//fmt.Println()
//	//fmt.Println(sir906, len(sir906))
//}

func superpalindromesInRange(left string, right string) int {
	// 设 P^2=R，有 R<1e18，则 P<1e9，则 p 的“左半部分” < 1e5
	// 考虑偶数和奇数，并且 P^2<=R
	sir906 := []int{
		1, 4, 9, 121, 484, 10201, 12321, 14641, 40804, 44944,
		1002001, 1234321, 4008004, 100020001, 102030201, 104060401, 121242121, 123454321, 125686521, 400080004,
		404090404, 10000200001, 10221412201, 12102420121, 12345654321, 40000800004, 1000002000001, 1002003002001, 1004006004001, 1020304030201,
		1022325232201, 1024348434201, 1210024200121, 1212225222121, 1214428244121, 1232346432321, 1234567654321, 4000008000004, 4004009004004, 100000020000001,
		100220141022001, 102012040210201, 102234363432201, 121000242000121, 121242363242121, 123212464212321, 123456787654321, 400000080000004, 10000000200000001, 10002000300020001,
		10004000600040001, 10020210401202001, 10022212521222001, 10024214841242001, 10201020402010201, 10203040504030201, 10205060806050201, 10221432623412201, 10223454745432201, 12100002420000121,
		12102202520220121, 12104402820440121, 12122232623222121, 12124434743442121, 12321024642012321, 12323244744232321, 12343456865434321, 12345678987654321, 40000000800000004, 40004000900040004}
	l, _ := strconv.Atoi(left)
	r, _ := strconv.Atoi(right)
	return sort.SearchInts(sir906, r) - sort.SearchInts(sir906, l-1)
}

//leetcode submit region end(Prohibit modification and deletion)
