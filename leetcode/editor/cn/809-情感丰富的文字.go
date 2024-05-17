package main

import "fmt"

func main() {
	s := "heeellooo"
	words := []string{"hello", "hi", "helo"}
	s = "abcd"
	words = []string{"abc"}
	i := expressiveWords(s, words)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func expressiveWords(s string, words []string) int {
	// 优化
	ans, n := 0, len(s)
	cache := make([]int, n)
	for i := 1; i < n; i++ { // 正序：统计出现次数
		if s[i] == s[i-1] {
			cache[i] = cache[i-1] + 1
		}
	}
	for i := n - 1; i >= 0; {
		if cache[i] >= 2 {
			for j := i + 1; i >= 0 && cache[i] != 0; i-- { // 倒序：编排“下一个索引”
				cache[i] = j
			}
		} else {
			cache[i] = n + 1 // 区别于 n
			i--
		}
	}
	for _, w := range words {
		i, j := 0, 0
		for j < len(w) && i < n {
			if w[j] == s[i] {
				i++
				j++
				continue
			} else {
				i = cache[i] // 匹配下一个
			}
		}
		if j == len(w) && (i == n || cache[i] == n) { // 区别于 cache[i] == n+1
			ans++
		}
	}
	return ans

	// lc
	//n := len(s)
	//check := func(w string) bool {
	//	x, y := 0, 0
	//	for y < len(w) && x < n {
	//		if w[y] != s[x] {
	//			return false
	//		}
	//		lx, ly := x, y
	//		for x++; x < n && s[x] == s[x-1]; {
	//			x++
	//		}
	//		for y++; y < len(w) && w[y] == w[y-1]; {
	//			y++
	//		}
	//		if dx, dy := x-lx, y-ly; dx < dy || dx > dy && dx < 3 {
	//			return false
	//		}
	//	}
	//	return x == n && y == len(w)
	//}
	//var ans int
	//for _, w := range words {
	//	if check(w) {
	//		ans++
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
