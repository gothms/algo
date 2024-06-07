package main

import (
	"fmt"
	"math/big"
	"math/bits"
)

func main() {
	gem := []int{2, 3}
	p, target := 100000007, 11391299 // 1
	//gem = []int{3, 21, 3}
	//p, target = 7, 5 // 4
	//gem = []int{988440132, 523615747, 753032918, 124332214, 865741351, 513616557, 65302524}
	//p, target = 48479, 20879 // 32
	//gem = []int{988440132, 523615747, 753032918, 124332214, 865741351, 513616557, 65302524, 1e9}
	//p, target = 48479, 20879 // 346
	//gem = []int{82, 78, 0, 44, 44, 70, 74, 94}
	//p, target = 37, 2 // 492600
	//gem = []int{82, 78, 1, 44, 44, 70, 74, 94}
	//p, target = 37, 2 // 487176
	//gem = []int{56824237, 125979399, 137221112, 127912002, 311549300, 701280270, 616099502, 954006405}
	//p, target = 513302941, 8375467 // 4
	//gem = []int{123456789, 123456789, 123456789, 123456789, 123456789, 123456789, 123456789, 123456789}
	//p, target = 3, 0 // 17297280
	//gem = []int{321113, 909148, 2108330, 853584, 1000839, 674651, 1585598, 38486, 42347102}
	//p, target = 2, 1 // 518918400 = 108533 -> 32127 -> 27470 Milliseconds
	//gem = []int{7, 75, 393, 5486, 25179, 906116, 7313534, 34808317, 125437907}
	//p, target = 100000007, 11479217 // 21 = 27319 Milliseconds

	//souls := treeOfInfiniteSouls(gem, p, target)
	//fmt.Println(souls)
	//fmt.Println(math.MaxInt64) // 19

	fmt.Println(treeOfInfiniteSouls(gem, p, target))
}

/*
Edge 结构体：表示图的边，包含起点 u、终点 v 和权重 w。

UnionFind 结构体：实现并查集，用于管理和合并节点集合。
	NewUnionFind(n int)：初始化并查集。
	Find(x int)：查找根节点。
	Union(x, y int)：合并两个集合。
Kruskal 算法：
	sort.Slice(edges, func(i, j int) bool { return edges[i].w < edges[j].w })：按边的权重从小到大排序。
	初始化并查集 uf。
	遍历排序后的边，使用 Union 函数将两个顶点合并，如果合并成功则将边的权重加到总权重中。
main 函数：
	读取输入的顶点数 n 和边数 m。
	读取每条边的信息并存储在 edges 切片中。
	调用 Kruskal 函数计算最小生成树的总权重并输出。
这个实现确保 Kruskal 算法的所有核心步骤都得到正确实现，包括边排序和使用并查集避免环的形成。
*/

// leetcode submit region begin(Prohibit modification and deletion)

//const lcp82 = 3
//
//var tree82 [][][][2]int
//
//func init() {
//	// dp
//	// 二叉树中序遍历的递归逻辑
//	tree82 = make([][][][2]int, lcp82)
//	tree82[0] = [][][2]int{{{1, 10}, {9, 10}}} // gem.length == 0
//	for i, c := 1, 1; i < lcp82; i++ {
//		c = c * (i<<2 - 2) / (i + 1) // 卡特兰数
//		tree82[i] = make([][][2]int, 0, c)
//	}
//	for i := 1; i < lcp82; i++ { // 最多 9 位数
//		for j := 0; j < i; j++ { // 左子树的节点数
//			for _, l := range tree82[j] { // 左子树
//				for _, r := range tree82[i-j-1] { // 右子树
//					// 状态转移方程：1 + 左子树 / 右子树 + 9
//					temp := make([][2]int, i+2)
//					k := len(l) - 1
//					copy(temp, l[:k])
//					copy(temp[k:], r)
//					temp[0][0] += temp[0][1]
//					temp[0][1] *= 10
//					temp[k][0] += l[k][0] * r[0][1]
//					temp[k][1] *= l[k][1]
//					temp[i+1][0] = temp[i+1][0]*10 + 9
//					temp[i+1][1] *= 10
//					tree82[i] = append(tree82[i], temp)
//				}
//			}
//		}
//	}
//	for i, t := range tree82 {
//		//	if i == 8 {
//		fmt.Println(i, t)
//		//	}
//	}
//	//memo := make(map[string]int)	// 检查是否重复
//	//for _, ts := range tree82[length(tree82)-1] {
//	//	var sb strings.Builder
//	//	for _, v := range ts {
//	//		sb.WriteString(strconv.Itoa(v[0]))
//	//	}
//	//	if i, ok := memo[sb.String()]; ok {
//	//		fmt.Println(i, sb.String())
//	//	}
//	//	memo[sb.String()]++
//	//}
//	//fmt.Println(memo)
//}

//const (
//	S  = 16
//	S1 = 32 - S
//	M  = 1996090921
//	N  = 9
//	M0 = 205
//)
//
//var (
//	pow10           [M0]int
//	pinv            [M0]int
//	l               [N]int
//	length          [1 << N]int
//	n, ans, p, r, B int
//	T               = 1
//	h               [1<<S + 105]hnode
//)
//
//type hnode struct { // 代码定义了一个简单的哈希表，用于计数特定的整数
//	x, y, t int
//}
//
//func H(x int) int {
//	return (x * M) >> S1
//}
//
//func insert82(x int) {
//	p := &h[H(x)]
//	for ; p.t == T; p = &h[H(x)+1] {
//		if p.x == x {
//			p.y++
//			return
//		}
//	}
//	p.t = T
//	p.x = x
//	p.y = 1
//}
//
//func find82(x int) int {
//	for p := &h[H(x)]; p.t == T; p = &h[H(x)+1] {
//		if p.x == x {
//			return p.y
//		}
//	}
//	return 0
//}
//
//type Node82 struct {
//	v1, v2, l int
//}
//
//func log10(n int) int { // 计算整数的十进制对数
//	if n < 10 {
//		return 1
//	}
//	return log10(n/10) + 1
//}
//
//func extendGCD(a, b int) (int, int, int) { // 计算扩展欧几里得算法和模逆
//	if b == 0 {
//		return a, 1, 0
//	}
//	d, x1, y1 := extendGCD(b, a%b)
//	x := y1
//	y := x1 - (a/b)*y1
//	return d, x, y
//}
//
//func inv(a, mod int) int { // 计算扩展欧几里得算法和模逆
//	_, x, _ := extendGCD(a, mod)
//	return (x%mod + mod) % mod
//}
//
//func fac(n int) *big.Int { // 计算阶乘
//	res := big.NewInt(1)
//	for n > 0 {
//		res.Mul(res, big.NewInt(int64(n)))
//		n--
//	}
//	return res
//}
//
//func C(n, m int) *big.Int { // 计算组合数
//	res := big.NewInt(1)
//	for i := 1; i <= m; i++ {
//		res.Mul(res, big.NewInt(int64(n-i+1)))
//		res.Div(res, big.NewInt(int64(i)))
//	}
//	return res
//}
//
//// 使用 calc 方法进行详细的计算，calc 方法通过递归和动态规划来解决问题
//// calc 方法依赖于三个函数 f1, f2, f3，它们定义了特定的条件和逻辑，用于在计算过程中进行不同的操作
//func calc(t0 int, f1, f2 func(int, int) bool, f3 func(int) bool) {
//	//var d [1 << (N + 1)][]Node82
//	for i := (1 << n) + 1; i < (1 << (n + 1)); i++ {
//		d[i] = nil
//	}
//	for i := 1 << n; i < (1 << (n + 1)); i++ {
//		if pop(i) <= t0 {
//			for j := (i - 1) & i; j > 0; j = (j - 1) & i {
//				if f2(i, j) || f1(i, j) {
//					for _, t := range d[j] {
//						//l1 := length[j-(1<<n)] - t.l + 2*(j > (1<<n))
//						l1 := length[j-(1<<n)] - t.l
//						if j > 1<<n {
//							l1 += 2
//						}
//						for _, v2 := range c[i-j] {
//							d[i] = append(d[i], Node82{
//								v1: (t.v1 + pow10[l1]) % p,
//								v2: (t.v2*pow10[length[i-j]+1] + v2*10 + 9) % p,
//								l:  t.l + length[i-j] + 1,
//							})
//							if !f2(i, j) {
//								d[i] = append(d[i], Node82{
//									v1: (t.v1 + pow10[l1+length[i-j]] + v2*pow10[l1]) % p,
//									v2: (t.v2*10 + 9) % p,
//									l:  t.l + 1,
//								})
//							}
//						}
//					}
//				}
//			}
//			j := (1 << (n + 1)) - 1 - i
//			T++
//			if f3(j) {
//				continue
//			}
//			for _, v := range c[j] {
//				insert82(v)
//			}
//			for _, t := range d[i] {
//				ans += find82((((r-t.v2-t.v1*pow10[length[j]+t.l])%p + p) * pinv[t.l]) % p)
//			}
//		}
//	}
//}
//
//var c [1 << N][]int
//var d [1 << (N + 1)][]Node82
//
//// 首先初始化各种变量和数组
//// 通过位运算和动态规划计算特定组合的值，并使用哈希表来计数和查找这些组合
//func treeOfInfiniteSouls(a []int, P, R int) int {
//	n = len(a)
//	ans = 0
//	B = (n + 2) / 3 // TODO
//	p = P
//	r = R
//	if p == 2 || p == 5 { // fast path
//		if p == 2 && r == 1 || p == 5 && r == 4 {
//			return int(C((n-1)*2, n-1).Div(C((n-1)*2, n-1), big.NewInt(int64(n))).Mul(fac(n), C((n-1)*2, n-1)).Int64())
//		}
//		return 0
//	}
//	pow10[0] = 1 % p
//	for i := 1; i < M0; i++ {
//		pow10[i] = (pow10[i-1] * 10) % p
//	}
//	//fmt.Println("pow10:", pow10)
//	for i := 0; i < M0; i++ {
//		pinv[i] = inv(pow10[i], p)
//	}
//	for i := 0; i < n; i++ {
//		l[i] = log10(a[i]) // 元素的长度
//	}
//	//fmt.Println(l)
//	for i := 1; i < (1 << n); i++ {
//		length[i] = (pop(i)*2 - 1) * 2 // TODO
//		for j := 0; j < n; j++ {
//			if i&(1<<j) != 0 {
//				length[i] += l[j] // 路径数值的总长度
//			}
//		}
//	}
//	//fmt.Println(length)
//	for i := 0; i < n; i++ {
//		c[1<<i] = append(c[1<<i], (a[i]*10+pow10[l[i]+1]+9)%p) // 1x9
//	}
//	//fmt.Println(c)
//	for i := 1; i < (1 << n); i++ {
//		if pop(i) <= B*2 {
//			t := pow10[length[i]-1] + 9
//			for j := (i - 1) & i; j > 0; j = (j - 1) & i { // 取出每一个二进制
//				if n == 9 || pop(i) < max((n+1)/2, 2) || max(pop(j), pop(i-j)) <= min(B, (n-1)/2) {
//					for _, v1 := range c[j] {
//						t1 := v1*pow10[length[i-j]+1] + t
//						for _, v2 := range c[i-j] {
//							c[i] = append(c[i], (v2*10+t1)%p)
//						}
//					}
//				}
//			}
//		}
//	}
//	fmt.Println(c)
//	d[1<<n] = append(d[1<<n], Node82{0, 0, 0})
//	yes := func(x, y int) bool { return true }
//	no := func(x, y int) bool { return false }
//	if n == 9 {
//		calc(4, yes, no, func(j int) bool { return pop(j) != 6 })
//		calc(5, func(i, j int) bool { return j != (1<<n) || pop(i-j) >= 2 }, no, func(j int) bool { return pop(j) != 5 })
//		calc(6, func(i, j int) bool { return j != (1<<n) || pop(i-j) == 3 }, func(i, j int) bool { return j == (1<<n) && pop(i-j) == 4 }, func(j int) bool { return pop(j) != 4 })
//	} else {
//		calc(n/2+1, yes, func(i, j int) bool { return n%2 == 0 && pop(j) == 1 && pop(i-j) == n/2 }, func(j int) bool { return pop(j) < (n+1)/2 || pop(j) > B*2 })
//	}
//	return ans
//}
//
//func pop(x int) int {
//	return bits.OnesCount(uint(x))
//}

const (
	N  = 9
	M0 = 205
)

type Node struct {
	v1, v2, l int
}
type Solution struct {
	pow10           [M0]int
	pinv            [M0]int
	l               [N]int
	len             [1 << N]int
	n, ans, p, r, B int
	//c               map[int][]int
	//d               map[int][]Node
	c    [1 << N][]int
	d    [1 << (N + 1)][]Node
	hash map[int]int
}

func NewSolution() *Solution {
	return &Solution{
		//c:    make(map[int][]int),
		//d:    make(map[int][]Node),
		c: [1 << N][]int{},
		d: [1 << (N + 1)][]Node{},
		//c:    make(map[int][]int),
		//d:    make(map[int][]Node),
		hash: make(map[int]int),
	}
}
func treeOfInfiniteSouls(a []int, p, r int) int {
	s := NewSolution()
	return s.treeOfInfiniteSouls(a, p, r)
}

// 首先初始化各种变量和数组
// 通过位运算和动态规划计算特定组合的值，并使用哈希表来计数和查找这些组合
func (s *Solution) treeOfInfiniteSouls(a []int, p, r int) int {
	s.n = len(a)
	s.ans = 0
	s.B = (s.n + 2) / 3
	s.p = p
	s.r = r
	if p == 2 || p == 5 {
		if p == 2 && r == 1 || p == 5 && r == 4 {
			return int(s.C((s.n-1)*2, s.n-1).Div(s.C((s.n-1)*2, s.n-1), big.NewInt(int64(s.n))).Mul(s.fac(s.n), s.C((s.n-1)*2, s.n-1)).Int64())
		}
		return 0
	}
	s.pow10[0] = 1 % p
	for i := 1; i < M0; i++ { // 计算整数的十进制对数
		s.pow10[i] = (s.pow10[i-1] * 10) % p
	}
	for i := 0; i < M0; i++ {
		s.pinv[i] = s.inv(s.pow10[i], p)
	}
	for i := 0; i < s.n; i++ { // 元素的长度
		s.l[i] = s.log10(a[i])
	}
	for i := 1; i < (1 << s.n); i++ {
		s.len[i] = (s.pop(i)*2 - 1) * 2 // TODO
		for j := 0; j < s.n; j++ {
			if i&(1<<j) != 0 {
				s.len[i] += s.l[j] // 路径数值的总长度
			}
		}
	}
	for i := 0; i < s.n; i++ {
		key := 1 << i
		s.c[key] = append(s.c[key], (a[i]*10+s.pow10[s.l[i]+1]+9)%p) // 1x9
	}
	for i := 1; i < (1 << s.n); i++ {
		if s.pop(i) <= s.B*2 {
			t := s.pow10[s.len[i]-1] + 9
			for j := (i - 1) & i; j != 0; j = (j - 1) & i {
				if s.n == 9 || s.pop(i) < max((s.n+1)/2, 2) || max(s.pop(j), s.pop(i-j)) <= min(s.B, (s.n-1)/2) {
					for _, v1 := range s.c[j] {
						t1 := v1*s.pow10[s.len[i-j]+1] + t
						for _, v2 := range s.c[i-j] {
							key := i
							s.c[key] = append(s.c[key], (v2*10+t1)%p)
						}
					}
				}
			}
		}
	}
	fmt.Println(s.c)
	s.d[1<<s.n] = append(s.d[1<<s.n], Node{0, 0, 0})
	yes := func(x, y int) bool { return true }
	//no := func(x, y int) bool { return false }
	//if s.n == 9 {
	//	s.calc(4, yes, no, func(j int) bool { return s.pop(j) != 6 })
	//	s.calc(5, func(i, j int) bool { return j != (1<<s.n) || s.pop(i-j) >= 2 }, no, func(j int) bool { return s.pop(j) != 5 })
	//	s.calc(6, func(i, j int) bool { return j != (1<<s.n) || s.pop(i-j) == 3 }, func(i, j int) bool { return j == (1<<s.n) && s.pop(i-j) == 4 }, func(j int) bool { return s.pop(j) != 4 })
	//} else {
	s.calc(s.n/2+1, yes, func(i, j int) bool { return s.n%2 == 0 && s.pop(j) == 1 && s.pop(i-j) == s.n/2 }, func(j int) bool { return s.pop(j) < (s.n+1)/2 || s.pop(j) > s.B*2 })
	//}
	return s.ans
}

// 使用 calc 方法进行详细的计算，calc 方法通过递归和动态规划来解决问题
// calc 方法依赖于三个函数 f1, f2, f3，它们定义了特定的条件和逻辑，用于在计算过程中进行不同的操作
func (s *Solution) calc(t0 int, f1, f2 func(int, int) bool, f3 func(int) bool) {
	for i := (1 << s.n) + 1; i < (1 << (s.n + 1)); i++ {
		s.d[i] = nil
	}
	for i := 1 << s.n; i < (1 << (s.n + 1)); i++ {
		if s.pop(i) <= t0 {
			for j := (i - 1) & i; (j >> s.n) > 0; j = (j - 1) & i {
				if f2(i, j) || f1(i, j) {
					for _, t := range s.d[j] {
						//l1 := s.len[j-(1<<s.n)] - t.l + 2*(j > (1<<s.n))
						l1 := s.len[j-(1<<s.n)] - t.l
						if j > (1 << s.n) {
							l1 += 2
						}
						for _, v2 := range s.c[i-j] {
							s.d[i] = append(s.d[i], Node{
								v1: (t.v1 + s.pow10[l1]) % s.p,
								v2: (t.v2*s.pow10[s.len[i-j]+1] + v2*10 + 9) % s.p,
								l:  t.l + s.len[i-j] + 1,
							})
							if !f2(i, j) {
								s.d[i] = append(s.d[i], Node{
									v1: (t.v1 + s.pow10[l1+s.len[i-j]] + v2*s.pow10[l1]) % s.p,
									v2: (t.v2*10 + 9) % s.p,
									l:  t.l + 1,
								})
							}
						}
					}
				}
			}
			j := (1 << (s.n + 1)) - 1 - i
			s.hash = make(map[int]int)
			if f3(j) {
				continue
			}
			for _, v := range s.c[j] {
				s.hash[v]++
			}
			for _, t := range s.d[i] {
				s.ans += s.hash[((s.r-t.v2-t.v1*s.pow10[s.len[j]+t.l])%s.p+s.p)*s.pinv[t.l]%s.p]
			}
		}
	}
}

func (s *Solution) log10(n int) int {
	if n < 10 {
		return 1
	}
	return s.log10(n/10) + 1
}

func (s *Solution) inv(a, mod int) int {
	_, x, _ := s.extendGCD(a, mod)
	return (x%mod + mod) % mod
}

func (s *Solution) extendGCD(a, b int) (int, int, int) { // 计算扩展欧几里得算法和模逆
	if b == 0 {
		return a, 1, 0
	}
	d, x1, y1 := s.extendGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return d, x, y
}

func (s *Solution) fac(n int) *big.Int { // 计算阶乘
	res := big.NewInt(1)
	for n > 0 {
		res.Mul(res, big.NewInt(int64(n)))
		n--
	}
	return res
}

func (s *Solution) C(n, m int) *big.Int { // 计算组合数
	res := big.NewInt(1)
	for i := 1; i <= m; i++ {
		res.Mul(res, big.NewInt(int64(n-i+1)))
		res.Div(res, big.NewInt(int64(i)))
	}
	return res
}

func (s *Solution) pop(x int) int {
	return bits.OnesCount(uint(x))
}

//leetcode submit region end(Prohibit modification and deletion)
