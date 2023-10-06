//设计一个数据结构，有效地找到给定子数组的 多数元素 。
//
// 子数组的 多数元素 是在子数组中出现 threshold 次数或次数以上的元素。
//
// 实现 MajorityChecker 类:
//
//
// MajorityChecker(int[] arr) 会用给定的数组 arr 对 MajorityChecker 初始化。
// int query(int left, int right, int threshold) 返回子数组中的元素 arr[left...right] 至少出
//现 threshold 次数，如果不存在这样的元素则返回 -1。
//
//
//
//
// 示例 1：
//
//
//输入:
//["MajorityChecker", "query", "query", "query"]
//[[[1, 1, 2, 2, 1, 1]], [0, 5, 4], [0, 3, 3], [2, 3, 2]]
//输出：
//[null, 1, -1, 2]
//
//解释：
//MajorityChecker majorityChecker = new MajorityChecker([1,1,2,2,1,1]);
//majorityChecker.query(0,5,4); // 返回 1
//majorityChecker.query(0,3,3); // 返回 -1
//majorityChecker.query(2,3,2); // 返回 2
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 2 * 10⁴
// 1 <= arr[i] <= 2 * 10⁴
// 0 <= left <= right < arr.length
// threshold <= right - left + 1
// 2 * threshold > right - left + 1
// 调用 query 的次数最多为 10⁴
//
//
// Related Topics 设计 树状数组 线段树 数组 二分查找 👍 160 👎 0

package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	arr := []int{1, 1, 2, 2, 1, 1}
	mc := Constructor(arr)
	fmt.Println(mc.cache)
	fmt.Println(mc.vPre)
	fmt.Println(mc.st)
	query := mc.Query(0, 5, 4)
	fmt.Println(query)
	query = mc.Query(0, 3, 3)
	fmt.Println(query)
	query = mc.Query(2, 3, 2)
	fmt.Println(query)
}

/*
题意：2 * threshold > right - left + 1

暴力：摩尔投票求出众数，遍历求出次数
	同分块方法中 Query，if right-left <= this.s
线段树 + 摩尔投票 + hash + 二分

分块 / sqrt 分解？
	https://leetcode.cn/problems/online-majority-element-in-subarray/solutions/19976/san-chong-fang-fa-bao-li-fen-kuai-xian-duan-shu-by/

线段树 + 分块
	摩尔投票
		此方案在遇到 n 很大，而元素重复次数很低时，不可行
		比如 20000 个元素，出现次数最多的元素是 3
	线段树
*/

// leetcode submit region begin(Prohibit modification and deletion)
type MajorityChecker struct {
	st    [][2]int    // 用于查绝对众数
	arr   []int       // 原数据
	cache map[int]int // val[i] 记录“可能元素”
	vPre  [][]int     // 前缀和：vPre[i][j] 记录元素 val[i] 在 arr[j] 前出现的次数
	vN    int         // “可能元素”的数量
	s     int         // 分块大小
	n     int
}

func Constructor(arr []int) MajorityChecker {
	n := len(arr)
	stLen := 1 << (bits.Len(uint(n-1)) + 1)
	st := make([][2]int, stLen)
	stBuild(arr, 0, n-1, 1, st) // 线段树

	s := int(math.Sqrt(float64(n << 1))) // 分块大小：根号 2*n
	vN := 0                              // 初始“可能元素”的数量为 0
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++ // 统计元素出现的次数
	}
	cache := make(map[int]int, s) // 或写死为 142，因为 1 <= arr.length <= 2 * 104，所以 vN 最大为 142
	vPre := make([][]int, 0, s)   // val := make([]int, 0, s+1) & vPre := make([][]int, 0, s+1)
	for i, v := range m {
		if v > s>>1 { // 元素出现次数大于分块长度的一半，即为“可能元素”
			cache[i] = vN
			vPre = append(vPre, make([]int, n+1))
			for j := 0; j < n; j++ { // 前缀和：针对元素值为 i
				if arr[j] == i {
					vPre[vN][j+1] = vPre[vN][j] + 1
				} else {
					vPre[vN][j+1] = vPre[vN][j] // vPre[i][j+1] 记录元素 val[vN] 出现的次数
				}
			} // 也可使用 差分 进行统计
			vN++
		}
	}
	//fmt.Println(cache)
	return MajorityChecker{
		st:    st,
		arr:   arr,
		cache: cache,
		vPre:  vPre,
		vN:    vN,
		s:     s,
		n:     n,
	}
}
func stBuild(arr []int, l, r, i int, st [][2]int) { // 构建线段树
	if l == r {
		st[i] = [2]int{arr[l], 1}
		return
	}
	m, idx := int(uint(l+r)>>1), i<<1
	stBuild(arr, l, m, idx, st)
	stBuild(arr, m+1, r, idx+1, st)
	st[i] = add(st[idx], st[idx+1])
}
func add(l, r [2]int) [2]int { // 摩尔投票：维护线段树
	if l[0] == r[0] {
		l[1] += r[1]
		return l
	} else if l[1] > r[1] {
		l[1] -= r[1]
		return l
	} else {
		r[1] -= l[1]
		return r
	}
}
func stQuery(f, t int, l, r int, i int, st [][2]int) (ret [2]int) {
	if f <= l && r <= t {
		return st[i]
	}
	m, idx := int(uint(l+r)>>1), i<<1
	if f <= m {
		ret = add(ret, stQuery(f, t, l, m, idx, st))
	}
	if t > m {
		ret = add(ret, stQuery(f, t, m+1, r, idx+1, st))
	}
	return
}
func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	var i, v, cnt int
	if right-left <= this.s { // 长度小于等于分块，直接用摩尔投票法
		for i = left; i <= right; i++ { // 摩尔投票
			if this.arr[i] == v {
				cnt++
			} else if cnt > 0 {
				cnt--
			} else {
				cnt = 1
				v = this.arr[i]
			}
		}
		for i, cnt = left, 0; i <= right; i++ { // 统计次数
			if this.arr[i] == v {
				cnt++
			}
		}
		if cnt < threshold {
			return -1
		}
		return v
	}
	v = stQuery(left, right, 0, this.n-1, 1, this.st)[0] // 众数
	i, ok := this.cache[v]                               // 众数的前缀和数组的位置
	if ok && this.vPre[i][right+1]-this.vPre[i][left] >= threshold {
		return v
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

// ====================分块 / sqrt 分解？====================

//type MajorityChecker struct {
//	arr  []int   // 原数据
//	val  []int   // val[i] 记录“可能元素”
//	vPre [][]int // 前缀和：vPre[i][j] 记录元素 val[i] 在 arr[j] 前出现的次数
//	vN   int     // “可能元素”的数量
//	s    int     // 分块大小
//}
//
//func Constructor(arr []int) MajorityChecker {
//	n := len(arr)
//	s := int(math.Sqrt(float64(n << 1))) // 分块大小：根号 2*n
//	vN := 0                              // 初始“可能元素”的数量为 0
//	m := make(map[int]int)
//	for _, v := range arr {
//		m[v]++ // 统计元素出现的次数
//	}
//	val := make([]int, 0, s)    // 或写死为 142，因为 1 <= arr.length <= 2 * 104，所以 vN 最大为 142
//	vPre := make([][]int, 0, s) // val := make([]int, 0, s+1) & vPre := make([][]int, 0, s+1)
//	for i, v := range m {
//		if v > s>>1 { // 元素出现次数大于分块长度的一半，即为“可能元素”
//			val = append(val, i) // 记录元素值
//			vPre = append(vPre, make([]int, n+1))
//			for j := 0; j < n; j++ { // 前缀和：针对元素值为 i
//				vPre[vN][j+1] = vPre[vN][j] // vPre[i][j+1] 记录元素 val[vN] 出现的次数
//				if arr[j] == i {
//					vPre[vN][j+1]++
//				}
//			} // 也可使用 差分 进行统计
//			vN++
//		}
//	}
//	return MajorityChecker{
//		arr:  arr,
//		val:  val,
//		vPre: vPre,
//		vN:   vN,
//		s:    s,
//	}
//}
//func (this *MajorityChecker) Query(left int, right int, threshold int) int {
//	var i, v, cnt int
//	if right-left <= this.s { // 长度小于等于分块，直接用摩尔投票法
//		for i = left; i <= right; i++ { // 摩尔投票
//			if this.arr[i] == v {
//				cnt++
//			} else if cnt > 0 {
//				cnt--
//			} else {
//				cnt = 1
//				v = this.arr[i]
//			}
//		}
//		for i, cnt = left, 0; i <= right; i++ { // 统计次数
//			if this.arr[i] == v {
//				cnt++
//			}
//		}
//		if cnt < threshold {
//			return -1
//		}
//		return v
//	}
//	for i = 0; i < this.vN; i++ { // 长度大于分块长度，在“可能元素”中查找
//		if this.vPre[i][right+1]-this.vPre[i][left] >= threshold {
//			return this.val[i]
//		}
//	}
//	return -1
//}

// ====================线段树====================

//type MajorityChecker struct {
//	st [][2]int      // 用于查绝对众数
//	m  map[int][]int // 用于查绝对众数出现次数
//	n  int           // 原数据长度
//}
//
//func stBuild(arr []int, l, r, i int, st [][2]int) { // 构建线段树
//	if l == r {
//		st[i] = [2]int{arr[l], 1}
//		return
//	}
//	m, idx := int(uint(l+r)>>1), i<<1
//	stBuild(arr, l, m, idx, st)
//	stBuild(arr, m+1, r, idx+1, st)
//	st[i] = add(st[idx], st[idx+1])
//}
//func add(l, r [2]int) [2]int { // 摩尔投票：维护线段树
//	if l[0] == r[0] {
//		l[1] += r[1]
//		return l
//	} else if l[1] > r[1] {
//		l[1] -= r[1]
//		return l
//	} else {
//		r[1] -= l[1]
//		return r
//	}
//}
//func Constructor(arr []int) MajorityChecker {
//	n := len(arr)
//	stLen := 1 << (bits.Len(uint(n-1)) + 1)
//	st := make([][2]int, stLen)
//	stBuild(arr, 0, n-1, 1, st) // 线段树
//	m := make(map[int][]int)
//	for i, v := range arr {
//		m[v] = append(m[v], i) // 元素:索引集
//	}
//	return MajorityChecker{st: st, m: m, n: n}
//}
//func (this *MajorityChecker) Query(left int, right int, threshold int) int {
//	ret := stQuery(left, right, 0, this.n-1, 1, this.st)            // 查出众数
//	mv := this.m[ret[0]]                                            // 元素结合
//	cnt := sort.SearchInts(mv, right+1) - sort.SearchInts(mv, left) // 查 right +1
//	if cnt < threshold {                                            // 不满足查询条件
//		return -1
//	}
//	return ret[0]
//}
//func stQuery(f, t int, l, r int, i int, st [][2]int) (ret [2]int) {
//	if f <= l && r <= t {
//		return st[i]
//	}
//	m, idx := int(uint(l+r)>>1), i<<1
//	if f <= m {
//		ret = add(ret, stQuery(f, t, l, m, idx, st))
//	}
//	if t > m {
//		ret = add(ret, stQuery(f, t, m+1, r, idx+1, st))
//	}
//	return
//}
