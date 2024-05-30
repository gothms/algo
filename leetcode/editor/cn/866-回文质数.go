package main

import (
	"fmt"
	"sort"
)

func main() {
	//for i := 0; i < 99; i++ {
	//	x := bits.Len(uint(i))
	//	fmt.Println(i, bits.Reverse32(uint32(i))>>(32-x))
	//}

	n := 13
	n = 9989900
	n = 16777619 + 100
	palindrome := primePalindrome(n)
	fmt.Println(palindrome)
}

// leetcode submit region begin(Prohibit modification and deletion)
const n866 = 1e8 + 30002 // 100030001

var prime866 []int

func init() {
	prime866 = make([]int, 0, 782)
	//p, b := [n866]bool{}, make([]int, 9)
	p := [n866]bool{}
	for i := 2; i < n866; i++ {
		if p[i] {
			continue
		}
		for j := i << 1; j < n866; j += i {
			p[j] = true
		}
		val := 0
		for v := i; v > 0; v /= 10 {
			val = val*10 + v%10
		}
		if val == i {
			prime866 = append(prime866, i)
		}
	}
	fmt.Println(prime866, len(prime866))
}
func primePalindrome(n int) int {
	return prime866[sort.SearchInts(prime866, n)]
}

//leetcode submit region end(Prohibit modification and deletion)
