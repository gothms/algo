package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	paths := [][]int{{0, 1, 2, 3, 4}, {2, 3, 4}, {4, 0, 1, 2, 3}}
	n := 5 // 2
	paths = [][]int{{1, 2, 3, 4}, {4, 1, 2, 3}, {4}}
	n = 5 // 1
	//paths = [][]int{{1, 7, 0, 6, 9, 0, 7, 4, 3, 9, 1, 5, 0, 8, 0, 6, 3, 6, 0, 8, 3, 7, 8, 3, 5, 3, 7, 4, 0, 6, 8, 1, 4},
	//	{1, 7, 0, 6, 9, 0, 7, 4, 3, 9, 1, 5, 0, 8, 0, 6, 3, 6, 0, 8, 3, 7, 8, 3, 5, 3, 7, 4, 0, 6, 8, 1, 5},
	//	{8, 1, 7, 0, 6, 9, 0, 7, 4, 3, 9, 1, 5, 0, 8, 0, 6, 3, 6, 0, 8, 3, 7, 8, 3, 5, 3, 7, 4, 0, 6, 8, 1}}
	//n = 10 // 32
	//paths = [][]int{
	//	{0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1},
	//	{1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2}}
	//n = 3 // 1023
	//paths = [][]int{{0, 1, 0, 1, 0, 1, 0, 1, 0}, {0, 1, 3, 0, 1, 4, 0, 1, 0}}
	//n = 5 // 3

	const N = 1024
	paths = make([][]int, 2)
	p := make([]int, 1, N)
	p[0] = 1
	for i := 1; i < N; i <<= 1 {
		for j := 0; j < i; j++ {
			p = append(p, p[j]^3)
		}
		paths[0] = p
	}
	paths[1] = append(paths[0][1:], paths[0][0])
	//for i, path := range paths {
	//	fmt.Println(i, path)
	//}
	subpath := longestCommonSubpath(n, paths)
	fmt.Println(subpath)
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonSubpath(n int, paths [][]int) int {
	//const t = math.MaxInt32
	//vs := make([]int, 0)
	//minL := t // 二分的右边界
	//for _, path := range paths {
	//	vs = append(vs, t)       // 用一个不存在 paths 中的数拼接所有路径
	//	vs = append(vs, path...) // 合为一个集合
	//	minL = min(minL, len(path))
	//}
	//N, m := len(vs), len(paths)
	//ids, id := make([]int, N), -1
	//for i, v := range vs {
	//	if v == t {
	//		ids[i] = m // 集合分隔标记
	//		id++
	//	} else {
	//		ids[i] = id // 同一集合标记
	//	}
	//}
	//// 后缀数组 sa 和 height 数组
	//s := make([]byte, 0, N<<2)
	//for _, v := range vs {
	//	s = append(s, byte(v), byte(v>>8), byte(v>>16), byte(v>>24)) // int32 拆为 4 个 byte
	//}
	//// Go 自带后缀数组，由 SA-IS 算法实现，复杂度 O(n)
	//_sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).
	//	Elem().FieldByName("sa").Field(0).UnsafeAddr())) // 转为 []int32 数组
	//sa, rk, height := make([]int32, 0, N), make([]int, N), make([]int, N)
	//for _, v := range _sa {
	//	if v&3 == 0 { // 重点：连续的 4 个 byte，只取第一个 byte 的排名（也可以是最后一个）
	//		sa = append(sa, v>>2) // [0,N) 的排名
	//	}
	//}
	//for i := range rk { // rank
	//	rk[sa[i]] = i
	//}
	//for i, k := 0, 0; i < N; i++ { // height 数组
	//	if rk[i] == 0 {
	//		continue
	//	}
	//	if k > 0 {
	//		k--
	//	}
	//	for j := int(sa[rk[i]-1]); i+k < N && j+k < N && vs[i+k] == vs[j+k]; {
	//		k++
	//	}
	//	height[rk[i]] = k
	//}
	//// 二分查找
	//return sort.Search(minL, func(k int) bool {
	//	k++                   // 因为 sort.Search 返回的是 k+1
	//	vis := make([]int, m) // 重点：标记，一个集合只访问一次
	//	for i := 1; i < N; i++ {
	//		if height[i] < k {
	//			continue
	//		}
	//		cnt := 0
	//		// 检查 sa[i] 和 sa[i-1]
	//		if id := ids[sa[i-1]]; id < m { // i 和 i-1 的公共前缀长度 = height[i]，但只考虑第一次的 i-1
	//			vis[id] = i
	//			cnt++
	//		}
	//		for tag := i; i < N && height[i] >= k; i++ {
	//			if id := ids[sa[i]]; id < m && vis[id] != tag {
	//				vis[id] = tag
	//				cnt++
	//			}
	//			//if id := ids[sa[i]]; id < m && vis[id] != tag {
	//			//	vis[id] = tag
	//			//	cnt++
	//			//}
	//			//if id := ids[sa[i-1]]; id < m && vis[id] != tag {
	//			//	vis[id] = tag
	//			//	cnt++
	//			//}
	//		}
	//		if cnt == m { // 每个集合的公共子串长度 == k++
	//			return false
	//		}
	//	}
	//	return true
	//})

	// 简化：但是更耗时，因为 slices.Compare
	//const prime, X, M = 16777619, 1, 100_000
	//binaryCheck := func(i int) bool {
	//	if len(paths[0]) < i {
	//		return false
	//	}
	//	var hash, pow uint32 = 0, 1
	//	for _, v := range paths[0][:i] {
	//		hash = hash*prime + uint32(v+X)
	//		pow *= prime
	//	}
	//	memo := make(map[uint32]int, len(paths[0])-i+1)
	//	memo[hash] = 0
	//	for k, v := range paths[0][i:] {
	//		hash = hash*prime + uint32(v+X) - uint32(paths[0][k]+X)*pow
	//		memo[hash] = k + 1
	//	}
	//	for _, path := range paths[1:] {
	//		if len(path) < i {
	//			return false
	//		}
	//		hash = 0
	//		temp := make(map[uint32]int, len(memo))
	//		for _, v := range path[:i] {
	//			hash = hash*prime + uint32(v+X)
	//		}
	//		if idx, ok := memo[hash]; ok &&
	//			slices.Compare(paths[0][idx:idx+i], path[0:i]) == 0 {
	//			temp[hash] = idx
	//		}
	//		for k, v := range path[i:] {
	//			hash = hash*prime + uint32(v+X) - uint32(path[k]+X)*pow
	//			if idx, ok := memo[hash]; ok &&
	//				slices.Compare(paths[0][idx:idx+i], path[k+1:i+k+1]) == 0 {
	//				temp[hash] = idx
	//			}
	//		}
	//		if len(temp) == 0 {
	//			return false
	//		}
	//		memo, temp = temp, nil
	//	}
	//	return true
	//}
	//var ans int
	//sort.Search(M, func(i int) bool {
	//	if i <= ans {
	//		return false
	//	}
	//	if binaryCheck(i) {
	//		ans = i
	//		return false
	//	}
	//	return true
	//})
	//return ans

	// 个人
	//const prime, X, M = 16777619, 1, 100_000
	//same := func(i, j, k, m int) bool {
	//	for c := 0; c < m; c++ {
	//		if paths[0][i+c] != paths[j][k+c] {
	//			return false
	//		}
	//	}
	//	return true
	//}
	//binaryCheck := func(i int) bool {
	//	if len(paths[0]) < i {
	//		return false
	//	}
	//	var hash, pow uint32 = 0, 1
	//	for _, v := range paths[0][:i] {
	//		hash = hash*prime + uint32(v+X)
	//		pow *= prime
	//	}
	//	memo := make(map[uint32]int, len(paths[0])-i+1)
	//	memo[hash] = 0
	//	for k, v := range paths[0][i:] {
	//		hash = hash*prime + uint32(v+X) - uint32(paths[0][k]+X)*pow
	//		memo[hash] = k + 1
	//	}
	//	for j, path := range paths[1:] {
	//		if len(path) < i {
	//			return false
	//		}
	//		hash = 0
	//		temp := make(map[uint32]int, len(memo))
	//		for _, v := range path[:i] {
	//			hash = hash*prime + uint32(v+X)
	//		}
	//		if idx, ok := memo[hash]; ok && same(idx, j+1, 0, i) {
	//			temp[hash] = idx
	//		}
	//		for k, v := range path[i:] {
	//			hash = hash*prime + uint32(v+X) - uint32(path[k]+X)*pow
	//			if idx, ok := memo[hash]; ok && same(idx, j+1, k+1, i) {
	//				temp[hash] = idx
	//			}
	//		}
	//		if len(temp) == 0 {
	//			return false
	//		}
	//		memo, temp = temp, nil
	//	}
	//	return true
	//}
	//var ans int
	//sort.Search(M, func(i int) bool {
	//	if i <= ans {
	//		return false
	//	}
	//	if binaryCheck(i) {
	//		ans = i
	//		return false
	//	}
	//	return true
	//})
	//return ans

	// 错误：唯独其中一个 prime==2 通过
	minL := math.MaxInt32
	for _, path := range paths {
		minL = min(minL, len(path))
	}
	const prime1, prime2, X = 16777619, 2099999999, 1
	//const prime1, prime2, X = 100030001, 2, 1
	binaryCheck := func(i int) bool {
		var pow1, pow2 uint32 = 1, 1
		var key [2]uint32
		for _, v := range paths[0][:i] {
			val := uint32(v + X)
			key[0], key[1] = key[0]*prime1+val, key[1]*prime2+val
			//key[0], key[1] = (key[0]*prime1+val)%mod, (key[1]*prime2+val)%mod
			pow1, pow2 = pow1*prime1, pow2*prime2
		}
		memo := make(map[[2]uint32]struct{}, len(paths[0])-i+1)
		memo[key] = struct{}{}
		for j, v := range paths[0][i:] {
			val, out := uint32(v+X), uint32(paths[0][j]+X)
			key[0], key[1] = key[0]*prime1+val-out*pow1, key[1]*prime2+val-out*pow2
			//key[0], key[1] = (key[0]*prime1+val-out*pow1)%mod, (key[1]*prime2+val-out*pow2)%mod
			memo[key] = struct{}{}
		}
		for _, path := range paths[1:] {
			key[0], key[1] = 0, 0
			temp := make(map[[2]uint32]struct{}, len(memo))
			for _, v := range path[:i] {
				val := uint32(v + X)
				key[0], key[1] = key[0]*prime1+val, key[1]*prime2+val
				//key[0], key[1] = (key[0]*prime1+val)%mod, (key[1]*prime2+val)%mod
			}
			if _, ok := memo[key]; ok {
				temp[key] = struct{}{}
			}
			for j, v := range path[i:] {
				val, out := uint32(v+X), uint32(path[j]+X)
				key[0], key[1] = key[0]*prime1+val-out*pow1, key[1]*prime2+val-out*pow2
				//key[0], key[1] = (key[0]*prime1+val-out*pow1)%mod, (key[1]*prime2+val-out*pow2)%mod
				if _, ok := memo[key]; ok {
					temp[key] = struct{}{}
				}
			}
			if len(temp) == 0 {
				return false
			}
			memo, temp = temp, nil
		}
		return true
	}
	var ans int
	sort.Search(minL+1, func(i int) bool {
		if i <= ans {
			return false
		}
		if binaryCheck(i) {
			ans = i
			return false
		}
		return true
	})
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
