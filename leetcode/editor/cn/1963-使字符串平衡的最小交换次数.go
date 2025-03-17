package main

import "fmt"

func main() {
	s := "]]][[["            // 2
	s = "][[][[][[][][[]]]]" // 1
	swaps := minSwaps(s)
	fmt.Println(swaps)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minSwaps(s string) int {
	// lc：将多余的 ] 换到最后面
	bs := []byte(s)
	ans, cnt, i := 0, 0, len(bs)-1
	for _, c := range bs {
		if c == '[' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			for bs[i] == ']' {
				i--
			}
			bs[i] = ']'
			ans++
			cnt++
		}
	}
	return ans

	// 个人
	//ans, l, r := 0, 0, len(s)-1
	//lc, rc := 0, 0
	//for l < r {
	//	for l < r && lc >= 0 {
	//		if s[l] == '[' {
	//			lc++
	//		} else {
	//			lc--
	//		}
	//		l++
	//	}
	//	for l < r && rc >= 0 {
	//		if s[r] == ']' {
	//			rc++
	//		} else {
	//			rc--
	//		}
	//		r--
	//	}
	//	if lc < 0 || rc < 0 {
	//		ans++
	//		lc += 2 // +2
	//		rc += 2
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
