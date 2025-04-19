package main

import "fmt"

func main() {
	s := "ihroyeeb"
	k := 5
	substrings := beautifulSubstrings(s, k)
	fmt.Println(substrings)
}

// leetcode submit region begin(Prohibit modification and deletion)
func beautifulSubstrings(s string, k int) int {
	// lc

	// 个人
	m := map[int][]int{0: {-1}}
	ans, cnt := 0, 0
	for i, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			cnt++
		default:
			cnt--
		}
		for _, j := range m[cnt] {
			n := (i - j) >> 1
			if n*n%k == 0 {
				ans++
			}
		}
		m[cnt] = append(m[cnt], i)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
