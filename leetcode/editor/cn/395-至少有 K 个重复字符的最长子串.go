package main

import "fmt"

func main() {
	s := "ababbc" // 5
	k := 2
	s = "ababbacacdbcadd"
	k = 3 // 15
	s = "a"
	k = 1 // 1
	s = "bbaaacbd"
	k = 3 // 3
	substring := longestSubstring(s, k)
	fmt.Println(substring)
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestSubstring(s string, k int) int {
	// 滑动窗体
	//ans := 0
	//for i := 1; i <= 26; i++ {
	//	cache := [26]int{}
	//	charCnt, kCnt, l := 0, 0, 0
	//	for r, c := range s {
	//		c -= 'a'
	//		if cache[c] == 0 {
	//			charCnt++
	//			kCnt++
	//		}
	//		cache[c]++
	//		if cache[c] == k {
	//			kCnt--
	//		}
	//		if charCnt > i {
	//			for charCnt > i {
	//				o := s[l] - 'a'
	//				if cache[o] == k {
	//					kCnt++
	//				}
	//				cache[o]--
	//				if cache[o] == 0 {
	//					charCnt--
	//					kCnt--
	//				}
	//				l++
	//			}
	//		}
	//		if kCnt == 0 {
	//			ans = max(ans, r-l+1)
	//		}
	//	}
	//}
	//return ans

	// 分治
	var div func(int, int, [26]int) int
	div = func(l, r int, cache [26]int) int {
		if l > r {
			return 0
		}
		ans, idx := 0, uint8(0)
		for i, v := range cache {
			if v > 0 && v < k {
				idx = uint8(i + 'a') // 分割
				break
			}
		}
		if idx == 0 { // 不分割
			return r - l + 1
		}
		cache = [26]int{}
		i, p := l, l
		for ; i <= r; i++ {
			if s[i] == idx {
				if i-p > ans { // 剪枝
					ans = max(ans, div(p, i-1, cache))
				}
				if i > p { // 优化
					cache = [26]int{}
				}
				p = i + 1
				continue
			}
			cache[s[i]-'a']++
		}
		if i-p > ans { // 后缀
			ans = max(ans, div(p, i-1, cache))
		}
		return ans
	}
	cache := [26]int{}
	for _, c := range s {
		cache[c-'a']++
	}
	return div(0, len(s)-1, cache)

	// 个人
	//ans, n := 0, len(s)
	//if k == 1 {
	//	return n
	//}
	//for i := 0; i <= n-k; i++ {
	//	cache, cnt := [26]int{}, 0
	//	for j, idx := i, 0; j < n; j++ {
	//		idx = int(s[j] - 'a')
	//		cache[idx]++
	//		if cache[idx] == 1 {
	//			cnt++
	//			continue
	//		} else if cache[idx] == k {
	//			cnt--
	//		}
	//		if cnt == 0 {
	//			ans = max(ans, j-i+1)
	//		}
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
