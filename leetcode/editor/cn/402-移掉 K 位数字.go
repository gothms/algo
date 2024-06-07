package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	var a int32 = 'a'
	var b uint8 = 'b'
	sizeofA := unsafe.Sizeof(a)
	sizeofB := unsafe.Sizeof(b)
	fmt.Println(sizeofA)
	fmt.Println(sizeofB)
	fmt.Printf("%b,%b\n", a, b)

	num := "1432219"
	k := 3
	//num = "10200"
	//k = 1
	//num = "10"
	//k = 2
	//num = "9"
	//k = 1
	//num = "10001"
	//k = 1 // "1"
	kdigits := removeKdigits(num, k)
	fmt.Println(kdigits)
}

// leetcode submit region begin(Prohibit modification and deletion)
func removeKdigits(num string, k int) string {
	// 优化
	st := []uint8{'0'}
	for i := range num {
		for k > 0 && num[i] < st[len(st)-1] {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, num[i])
	}
	st = st[1 : len(st)-k]
	ans := strings.TrimLeft(string(st), "0")
	if ans == "" {
		return "0"
	}
	return ans

	// 栈
	//	st := []uint8{'0'}
	//	var i int
	//out:
	//	for i = range num {
	//		for num[i] < st[len(st)-1] {
	//			st = st[:len(st)-1]
	//			k--
	//			if k == 0 {
	//				st = append(st, num[i])
	//				break out
	//			}
	//		}
	//		st = append(st, num[i])
	//	}
	//	st = st[:len(st)-k] // k 还有剩余
	//	st = append(st, num[i+1:]...)
	//	j := 1
	//	for j < len(st) && st[j] == '0' {
	//		j++
	//	}
	//	if j == len(st) {
	//		return "0"
	//	}
	//	return string(st[j:])
}

//leetcode submit region end(Prohibit modification and deletion)
