package maths

/*
哈希部分：
	insert 和 find 函数使用线性探测来解决哈希冲突，类似于 C++ 版本。
	使用数组 h 来存储哈希节点。
数学函数：
	实现了 log10、extendGcd、inv、fac 和 C 函数，用于计算对数、扩展欧几里得算法、模逆、阶乘和组合数。
Solution 类：
	包含了 calc 和 treeOfInfiniteSouls 方法，分别用于计算和解决问题。
	calc 方法通过回溯和动态规划来计算结果。
	treeOfInfiniteSouls 方法初始化必要的数据并调用 calc 方法。
这个 Go 版本的代码在逻辑上尽量保持与原 C++ 代码一致，并处理了语言差异，如位操作、数组和切片的使用等。请注意，Go 语言中的切片和数组的处理方式与 C++ 中的向量和数组不同，因此在实现时需要特别注意这些差异。

这段代码通过动态规划和组合数学的方法，解决了一个复杂的组合优化问题。
它使用了哈希表来计数和查找特定组合，并通过递归和动态规划来计算结果。整个过程涉及到大量的位运算和数学计算，非常高效且复杂。
1.哈希表的实现
	代码定义了一个简单的哈希表，用于计数特定的整数。该哈希表使用开放寻址法解决冲突
2.数学工具函数
	log10：计算整数的十进制对数。
	extend_gcd 和 inv：计算扩展欧几里得算法和模逆。
	fac 和 C：计算阶乘和组合数。
3.动态规划和回溯
	代码通过定义一个类 Solution 来实现组合优化问题的求解。该类包含了一个主方法 treeOfInfiniteSouls，以及一个辅助方法 calc。
4.详细逻辑
	treeOfInfiniteSouls 方法首先初始化各种变量和数组。
	通过位运算和动态规划计算特定组合的值，并使用哈希表来计数和查找这些组合。
	使用 calc 方法进行详细的计算，calc 方法通过递归和动态规划来解决问题。
	calc 方法依赖于三个函数 f1, f2, f3，它们定义了特定的条件和逻辑，用于在计算过程中进行不同的操作。
*/

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
//	len             [1 << N]int
//	n, ans, p, r, B int
//	T               = 1
//	h               [1<<S + 105]hnode
//)
//
//type hnode struct {
//	x, y, t int
//}
//
//func H(x int) int {
//	return (x * M) >> S1
//}
//
//func insert(x int) {
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
//func find(x int) int {
//	for p := &h[H(x)]; p.t == T; p = &h[H(x)+1] {
//		if p.x == x {
//			return p.y
//		}
//	}
//	return 0
//}
//
//type Node struct {
//	v1, v2, l int
//}
//
//func log10(n int) int {
//	if n < 10 {
//		return 1
//	}
//	return log10(n/10) + 1
//}
//
//func extendGCD(a, b int) (int, int, int) {
//	if b == 0 {
//		return a, 1, 0
//	}
//	d, x1, y1 := extendGCD(b, a%b)
//	x := y1
//	y := x1 - (a/b)*y1
//	return d, x, y
//}
//
//func inv(a, mod int) int {
//	_, x, _ := extendGCD(a, mod)
//	return (x%mod + mod) % mod
//}
//
//func fac(n int) *big.Int {
//	res := big.NewInt(1)
//	for n > 0 {
//		res.Mul(res, big.NewInt(int64(n)))
//		n--
//	}
//	return res
//}
//
//func C(n, m int) *big.Int {
//	res := big.NewInt(1)
//	for i := 1; i <= m; i++ {
//		res.Mul(res, big.NewInt(int64(n-i+1)))
//		res.Div(res, big.NewInt(int64(i)))
//	}
//	return res
//}
//
//func calc(t0 int, f1, f2, f3 func(int, int) bool) {
//	var d [1 << (N + 1)][]Node
//	for i := (1 << n) + 1; i < (1 << (n + 1)); i++ {
//		d[i] = nil
//	}
//	for i := 1 << n; i < (1 << (n + 1)); i++ {
//		if pop(i) <= t0 {
//			for j := (i - 1) & i; j > 0; j = (j - 1) & i {
//				if f2(i, j) || f1(i, j) {
//					for _, t := range d[j] {
//						l1 := len[j-(1<<n)] - t.l + 2*(j > (1<<n))
//						for _, v2 := range c[i-j] {
//							d[i] = append(d[i], Node{
//								v1: (t.v1 + pow10[l1]) % p,
//								v2: (t.v2*pow10[len[i-j]+1] + v2*10 + 9) % p,
//								l:  t.l + len[i-j] + 1,
//							})
//							if !f2(i, j) {
//								d[i] = append(d[i], Node{
//									v1: (t.v1 + pow10[l1+len[i-j]] + v2*pow10[l1]) % p,
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
//				insert(v)
//			}
//			for _, t := range d[i] {
//				ans += find((((r-t.v2-t.v1*pow10[len[j]+t.l])%p + p) * pinv[t.l]) % p)
//			}
//		}
//	}
//}
//
//var c [1 << N][]int
//
//func treeOfInfiniteSouls(a []int, P, R int) int {
//	n = len(a)
//	ans = 0
//	B = (n + 2) / 3
//	p = P
//	r = R
//	if p == 2 || p == 5 {
//		if p == 2 && r == 1 || p == 5 && r == 4 {
//			return int(C((n-1)*2, n-1).Div(C((n-1)*2, n-1), big.NewInt(int64(n))).Mul(fac(n), C((n-1)*2, n-1)).Int64())
//		}
//		return 0
//	}
//	pow10[0] = 1 % p
//	for i := 1; i < M0; i++ {
//		pow10[i] = (pow10[i-1] * 10) % p
//	}
//	for i := 0; i < M0; i++ {
//		pinv[i] = inv(pow10[i], p)
//	}
//	for i := 0; i < n; i++ {
//		l[i] = log10(a[i])
//	}
//	for i := 1; i < (1 << n); i++ {
//		len[i] = (pop(i)*2 - 1) * 2
//		for j := 0; j < n; j++ {
//			if i&(1<<j) != 0 {
//				len[i] += l[j]
//			}
//		}
//	}
//	for i := 0; i < n; i++ {
//		c[1<<i] = append(c[1<<i], (a[i]*10+pow10[l[i]+1]+9)%p)
//	}
//	for i := 1; i < (1 << n); i++ {
//		if pop(i) <= B*2 {
//			t := pow10[len[i]-1] + 9
//			for j := (i - 1) & i; j > 0; j = (j - 1) & i {
//				if n == 9 || pop(i) < max((n+1)/2, 2) || max(pop(j), pop(i-j)) <= min(B, (n-1)/2) {
//					for _, v1 := range c[j] {
//						t1 := v1*pow10[len[i-j]+1] + t
//						for _, v2 := range c[i-j] {
//							c[i] = append(c[i], (v2*10+t1)%p)
//						}
//					}
//				}
//			}
//		}
//	}
//	d[1<<n] = append(d[1<<n], Node{0, 0, 0})
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
//
//func main() {
//	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	p := 17
//	r := 10
//	fmt.Println(treeOfInfiniteSouls(a, p, r))
//}
