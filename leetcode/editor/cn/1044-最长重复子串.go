package main

import (
	"fmt"
	"index/suffixarray"
	"reflect"
	"sort"
	"strings"
	"unsafe"
)

func main() {
	s := "banana"
	s = "havbc"
	s = "aeeebaabd"
	s = "aaaaaaa"
	s = "babab"
	substring := longestDupSubstring(s)
	fmt.Println(substring, len(s))

	fmt.Println(len(strings.Split(s, "havbc")))
	fmt.Println(len(strings.Split(s, "rvlmms")))

	//fmt.Println('a', '`')
	//fmt.Printf("%c\n", '`')
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestDupSubstring(s string) string {
	// 后缀数组
	ans, idx, n := 0, 0, len(s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rk := make([]int, n)
	for i, v := range sa {
		rk[v] = i
	}
	height := make([]int, n)
	for i, k := 0, 0; i < n; i++ {
		if rk[i] == 0 {
			continue
		}
		if k > 0 {
			k--
		}
		for j := int(sa[rk[i]-1]); i+k < n && j+k < n && s[i+k] == s[j+k]; {
			k++
		}
		if k > ans {
			ans, idx = k, i
		}
		height[rk[i]] = k
	}
	fmt.Println(height)
	return s[idx : idx+ans]
}

//leetcode submit region end(Prohibit modification and deletion)

func longestDupSubstring_(s string) string {
	// 优化
	const p1, p2, B = 16777619, 2099999999, '`'
	check := func(i int) int {
		var hash1, hash2, pow1, pow2 uint32 = 0, 0, 1, 1
		for _, c := range s[:i] {
			v := uint32(c - B)
			hash1, hash2 = hash1*p1+v, hash2*p2+v
			pow1, pow2 = pow1*p1, pow2*p2
		}
		memo := make(map[[2]uint32]struct{}, len(s)-i+1)
		memo[[2]uint32{hash1, hash2}] = struct{}{}
		for j, c := range s[i:] {
			v, last := uint32(c-B), uint32(s[j]-B)
			hash1, hash2 = hash1*p1+v-last*pow1, hash2*p2+v-last*pow2
			key := [2]uint32{hash1, hash2}
			if _, ok := memo[key]; ok {
				return j + 1
			}
			memo[key] = struct{}{}
		}
		return -1
	}
	var idx, maxL int
	sort.Search(len(s), func(i int) bool {
		if i <= maxL {
			return false
		}
		if j := check(i); j != -1 { // j > 0
			idx, maxL = j, i
			return false
		}
		return true
	})
	return s[idx : idx+maxL]

	// lc
	//randInt := func(a, b int) int {
	//	return a + rand.Intn(b-a)
	//}
	//pow := func(x, n, mod int) int {
	//	res := 1
	//	for ; n > 0; n >>= 1 {
	//		if n&1 > 0 {
	//			res = res * x % mod
	//		}
	//		x = x * x % mod
	//	}
	//	return res
	//}
	//check := func(arr []byte, m, a1, a2, mod1, mod2 int) int {
	//	aL1, aL2 := pow(a1, m, mod1), pow(a2, m, mod2)
	//	h1, h2 := 0, 0
	//	for _, c := range arr[:m] {
	//		h1 = (h1*a1 + int(c)) % mod1
	//		h2 = (h2*a2 + int(c)) % mod2
	//	}
	//	// 存储一个编码组合是否出现过
	//	seen := map[[2]int]bool{{h1, h2}: true}
	//	for start := 1; start <= len(arr)-m; start++ {
	//		h1 = (h1*a1 - int(arr[start-1])*aL1 + int(arr[start+m-1])) % mod1
	//		if h1 < 0 {
	//			h1 += mod1
	//		}
	//		h2 = (h2*a2 - int(arr[start-1])*aL2 + int(arr[start+m-1])) % mod2
	//		if h2 < 0 {
	//			h2 += mod2
	//		}
	//		// 如果重复，则返回重复串的起点
	//		if seen[[2]int{h1, h2}] {
	//			return start
	//		}
	//		seen[[2]int{h1, h2}] = true
	//	}
	//	// 没有重复，则返回 -1
	//	return -1
	//}
	//
	//rand.Seed(time.Now().UnixNano())
	//// 生成两个进制
	//a1, a2 := randInt(26, 100), randInt(26, 100)
	//// 生成两个模
	//mod1, mod2 := randInt(1e9+7, math.MaxInt32), randInt(1e9+7, math.MaxInt32)
	//// 先对所有字符进行编码
	//arr := []byte(s)
	//for i := range arr {
	//	arr[i] -= 'a'
	//}
	//
	//l, r := 1, len(s)-1
	//length, start := 0, -1
	//for l <= r {
	//	m := l + (r-l+1)/2
	//	idx := check(arr, m, a1, a2, mod1, mod2)
	//	if idx != -1 { // 有重复子串，移动左边界
	//		l = m + 1
	//		length = m
	//		start = idx
	//	} else { // 无重复子串，移动右边界
	//		r = m - 1
	//	}
	//}
	//if start == -1 {
	//	return ""
	//}
	//return s[start : start+length]
}
