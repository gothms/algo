package main

import (
	"fmt"
	"sort"
)

func main() {
	s, a, b := "isawsquirrelnearmysquirrelhouseohmy", "my", "squirrel"
	k := 15 // [16,33]
	s, a, b = "wl", "xjigt", "wl"
	k = 2 // []
	indices := beautifulIndices(s, a, b, k)
	fmt.Println(indices)
}

// leetcode submit region begin(Prohibit modification and deletion)
func beautifulIndices(s string, a string, b string, k int) []int {
	// 升级版：lc-3008
	// BoyerMoore

	// Rabin-Karp
	n := len(s)
	if n < len(a) || n < len(b) {
		return nil
	}
	const prime = 16777619
	rkHelper := func(p string) (uint32, uint32) {
		var hash, pow, base uint32 = 0, 1, prime
		for _, c := range p {
			hash = hash*prime + uint32(c)
		}
		for n := len(p); n != 0; n >>= 1 {
			if n&1 != 0 {
				pow *= base
			}
			base *= base
		}
		return hash, pow
	}
	rk := func(s, p string) []int {
		hash, pow := rkHelper(p)
		ans := make([]int, 0)
		m := len(p)
		var h uint32
		for i := 0; i < m; i++ {
			h = h*prime + uint32(s[i])
		}
		if h == hash {
			ans = append(ans, 0)
		}
		for i := m; i < len(s); i++ {
			h = h*prime + uint32(s[i]) - pow*(uint32(s[i-m]))
			if h == hash {
				ans = append(ans, i-m+1)
			}
		}
		return ans
	}
	idsA, idsB := rk(s, a), rk(s, b)
	idx, left := 0, 0
	for _, i := range idsB {
		if left == len(idsA) {
			break
		}
		l, r := max(left, i-k), min(n-1, i+k)
		//	if idxA[left] > r {                // 已截取
		//		continue
		//	}
		if l > idsA[left] {
			left = sort.SearchInts(idsA[idx:], l) + idx
		}
		right := sort.SearchInts(idsA[idx:], r+1) + idx
		copy(idsA[idx:], idsA[left:right])
		idx += right - left
		left = right
	}
	return idsA[:idx]

	// KMP
	//kmp := func(p string) []int {
	//	n := len(p)
	//	prefix := make([]int, n)
	//	for i, k := 1, 0; i < n; i++ {
	//		for k > 0 && p[i] != p[k] {
	//			k = prefix[k-1]
	//		}
	//		if p[i] == p[k] {
	//			k++
	//		}
	//		prefix[i] = k
	//	}
	//	return prefix
	//}
	//match := func(p string, prefix []int) []int {
	//	n := len(p)
	//	ret := make([]int, 0)
	//	for i, j := 0, 0; i < len(s); i++ {
	//		for j > 0 && s[i] != p[j] {
	//			j = prefix[j-1]
	//		}
	//		if s[i] == p[j] {
	//			j++
	//		}
	//		if j == n {
	//			ret = append(ret, i-n+1)
	//			j = prefix[j-1]
	//		}
	//	}
	//	return ret
	//}
	//pa, pb := kmp(a), kmp(b)                 // prefix 数组
	//idxA, idxB := match(a, pa), match(b, pb) // 所有匹配的索引
	//
	//ans := make([]int, 0)
	//m, left := len(s), 0 // 标记 left 左边的 idxA 的索引均已截取
	//for _, i := range idxB {
	//	if left == len(idxA) { // 已截取完
	//		break
	//	}
	//	l, r := max(0, i-k), min(m-1, i+k) // 左右边界
	//	if idxA[left] > r {                // 已截取
	//		continue
	//	}
	//	if l > idxA[left] { // 需要更新 left
	//		left = sort.SearchInts(idxA, l)
	//	}
	//	right := sort.SearchInts(idxA, r+1) // 查询 right
	//	ans = append(ans, idxA[left:right]...)
	//	left = right // 新的 left
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
