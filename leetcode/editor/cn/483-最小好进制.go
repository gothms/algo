package main

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
)

func main() {
	n := "13"
	n = "4681"
	//for i := bits.Len(uint(n)) - 1; i > 1; i-- {
	//	k := int(math.Pow(float64(n), 1/float64(i)))
	//	fmt.Println(i, k)
	//}
	base := smallestGoodBase(n)
	fmt.Println(base)
}

// leetcode submit region begin(Prohibit modification and deletion)
func smallestGoodBase(n string) string {

}

//leetcode submit region end(Prohibit modification and deletion)

func smallestGoodBase_(n string) string {
	v, _ := strconv.Atoi(n)
	mx := bits.Len(uint(v)) - 1
	pre := 0
	for i := mx; i > 1; i-- {
		k := int(math.Pow(float64(v), 1/float64(i)))
		if k == pre {
			continue
		}
		pow, sum := 1, 1
		for j := 1; j <= i; j++ {
			pow *= k
			sum += pow
		}
		if sum == v {
			return strconv.Itoa(k)
		}
		pre = k
	}
	return strconv.Itoa(v - 1)
}
