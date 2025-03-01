package main

import "fmt"

func main() {
	s, p := "baa", "aa"
	anagrams := findAnagrams(s, p)
	fmt.Println(anagrams)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	// lc：滑动窗体
	curr, cache := [26]int{}, [26]int{}
	for _, c := range p {
		cache[c-'a']++
	}
	n := len(p)
	ans := make([]int, 0)
	for l, r := 0, 0; r < len(s); r++ {
		i := s[r] - 'a'
		curr[i]++
		for ; curr[i] > cache[i]; l++ { // 1.控制和比较 curr[i] 的数量
			curr[s[l]-'a']--
		}
		if r-l+1 == n { // 2.控制和比较滑动窗体的长度
			ans = append(ans, l)
		}
	}
	return ans

	// 个人
	//const C = 26
	//m, t := [C]int{}, [C]int{}
	//for _, c := range p {
	//	t[c-'a']++
	//}
	//cnt, n := C, len(p)
	//for _, v := range t {
	//	if v == 0 {
	//		cnt--
	//	}
	//}
	//ans := make([]int, 0)
	//for i, c := range s {
	//	in := c - 'a'
	//	if m[in] == t[in] {
	//		cnt++
	//	}
	//	m[in]++
	//	if m[in] == t[in] {
	//		cnt--
	//	}
	//	if i >= n {
	//		out := s[i-n] - 'a'
	//		if m[out] == t[out] {
	//			cnt++
	//		}
	//		m[out]--
	//		if m[out] == t[out] {
	//			cnt--
	//		}
	//	}
	//	if cnt == 0 {
	//		ans = append(ans, i-n+1)
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
