package str

import (
	"fmt"
	"index/suffixarray"
	"math"
	"reflect"
	"sort"
	"unsafe"
)

/*
https://oi-wiki.org/string/sa/

后缀数组（Suffix Array）
	数组 sa
		sa[i] 表示将所有后缀排序后第 i 小的后缀的编号，也是所说的后缀数组，后文也称编号数组 sa
	数组 rk
		rk[i] 表示后缀 i 的排名，是重要的辅助数组，后文也称排名数组 rk
	这两个数组满足性质：sa[rk[i]]=rk[sa[i]]=i
后缀数组怎么求？	TODO
	O(n^2logn) 做法
	O(nlog^2n) 做法
		倍增思想
	O(nlogn) 做法
	O(n) 做法

height 数组
	LCP（最长公共前缀）
	height 数组的定义
		height[i]=lcp(sa[i],sa[i-1])，即第 i 名的后缀与它前一名的后缀的最长公共前缀
		height[1] 可以视作 0
	O(n) 求 height 数组需要的一个引理
		height[rank[i]] >= height[rank[i-1]]-1
	证明：
		当 height[rk[i-1]] <= 1 时，上式显然成立（右边小于等于 0）。
		当 height[rk[i-1]]>1 时：
		根据 height 定义，有 lcp(sa[rk[i-1]], sa[rk[i-1]-1]) = height[rk[i-1]] > 1。
		既然后缀 i-1 和后缀 sa[rk[i-1]-1] 有长度为 height[rk[i-1]] 的最长公共前缀，
			注意
				rk[i-1] 为后缀 i-1 的排名，rk[i-1]-1 即是比后缀 i-1 小的前一个后缀的排名
				那么 sa[rk[i-1]-1] 就是前一个后缀
		那么不妨用 aA 来表示这个最长公共前缀。（其中 a 是一个字符，A 是长度为 height[rk[i-1]]-1 的字符串，非空）
		那么后缀 i-1 可以表示为 aAD，后缀 sa[rk[i-1]-1] 可以表示为 aAB。（B < D，B 可能为空串，D 非空）
		进一步地，后缀 i 可以表示为 AD，存在后缀（sa[rk[i-1]-1]+1）AB。
			注意
				后缀 i-1 去掉首字符 a，即为后缀 i == AD
				同理后缀 sa[rk[i-1]-1]+1 == AB
		因为后缀 sa[rk[i]-1] 在大小关系的排名上仅比后缀 sa[rk[i]] 也就是后缀 i，小一位，而 AB < AD。
		所以 AB <= 后缀 sa[rk[i]-1] < AD，显然后缀 i 和后缀 sa[rk[i]-1] 有公共前缀 A。
			注意
				sa[rk[i]-1] 和 AD 是排名紧邻的，那么必然存在 AB <= sa[rk[i]-1]
				所以后缀 i 和后缀 sa[rk[i]-1] 至少存在公共前缀 A
		于是就可以得出 lcp(i,sa[rk[i]-1]) 至少是 height[rk[i-1]]-1，也即 height[rk[i]] >= height[rk[i-1]]-1。
			总结
				height[rank[i-1]] = lcp(sa[rk[i-1]], sa[rk[i-1]-1])
					存在最长公共前缀 aA
				height[rank[i]] = lcp(i,sa[rk[i]-1])
					至少存在最长公共前缀 A

lc
	1923
*/

// longestCommonSubpath lc 1923
func longestCommonSubpath(n int, paths [][]int) int {
	const t = math.MaxInt32
	vs := make([]int, 0)
	minL := t // 二分的右边界
	for _, path := range paths {
		vs = append(vs, t)       // 用一个不存在 paths 中的数拼接所有路径
		vs = append(vs, path...) // 合为一个集合
		minL = min(minL, len(path))
	}
	N, m := len(vs), len(paths)
	ids, id := make([]int, N), -1
	for i, v := range vs {
		if v == t {
			ids[i] = m // 集合分隔标记
			id++
		} else {
			ids[i] = id // 同一集合标记
		}
	}

	// 后缀数组 sa 和 height 数组
	s := make([]byte, 0, N<<2)
	for _, v := range vs {
		s = append(s, byte(v), byte(v>>8), byte(v>>16), byte(v>>24)) // int32 拆为 4 个 byte
	}
	// Go 自带后缀数组，由 SA-IS 算法实现，复杂度 O(n)
	_sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).
		Elem().FieldByName("sa").Field(0).UnsafeAddr())) // 转为 []int32 数组
	sa, rk, height := make([]int32, 0, N), make([]int, N), make([]int, N)
	for _, v := range _sa {
		if v&3 == 0 { // 重点：连续的 4 个 byte，只取第一个 byte 的排名（也可以是最后一个）
			sa = append(sa, v>>2) // [0,N) 的排名
		}
	}
	for i := range rk { // rank
		rk[sa[i]] = i
	}
	for i, k := 0, 0; i < N; i++ { // height 数组
		if rk[i] == 0 {
			continue
		}
		if k > 0 {
			k--
		}
		for j := int(sa[rk[i]-1]); i+k < N && j+k < N && vs[i+k] == vs[j+k]; {
			k++
		}
		height[rk[i]] = k
	}

	// 二分查找
	return sort.Search(minL, func(k int) bool {
		k++                   // 因为 sort.Search 返回的是 k+1
		vis := make([]int, m) // 重点：标记，一个集合只访问一次
		for i := 1; i < N; i++ {
			if height[i] < k {
				continue
			}
			cnt := 0
			// 检查 sa[i] 和 sa[i-1]
			if id := ids[sa[i-1]]; id < m { // i 和 i-1 的公共前缀长度 = height[i]，但只考虑第一次的 i-1
				vis[id] = i
				cnt++
			}
			for tag := i; i < N && height[i] >= k; i++ {
				if id := ids[sa[i]]; id < m && vis[id] != tag {
					vis[id] = tag
					cnt++
				}
				//if id := ids[sa[i]]; id < m && vis[id] != tag {
				//	vis[id] = tag
				//	cnt++
				//}
				//if id := ids[sa[i-1]]; id < m && vis[id] != tag {
				//	vis[id] = tag
				//	cnt++
				//}
			}
			if cnt == m { // 每个集合的公共子串长度 == k++
				return false
			}
		}
		return true
	})
}

// suffixArray O(n) 求 height 数组的代码实现
func suffixArray() {
	s := "aabaaaab"
	n := len(s)
	// Go 后缀数组，由 SA-IS 算法实现，复杂度 O(n)
	// TODO
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).
		Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rk := make([]int, n)
	for i := range rk { // rank 数组
		rk[sa[i]] = i
	}
	height := make([]int, n)
	k := 0
	for i := 0; i < n; i++ { // height 数组
		if rk[i] == 0 {
			continue
		}
		if k > 0 {
			k--
		}
		j := int(sa[rk[i]-1])
		for i+k < n && j+k < n && s[i+k] == s[j+k] {
			k++
		}
		height[rk[i]] = k
	}
	fmt.Println(sa)     // [3 4 5 0 6 1 7 2]
	fmt.Println(rk)     // [3 5 7 0 1 2 4 6]
	fmt.Println(height) // [0 3 2 3 1 2 0 1]
}
